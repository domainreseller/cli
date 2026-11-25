package dna

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

const testReseller = "11111111-2222-3333-4444-555555555555"

func get(c *Client, path string) (*Response, error) {
	return c.Do(context.Background(), &Request{Method: "GET", Path: path}, nil)
}

func post(c *Client, path string, body any) (*Response, error) {
	return c.Do(context.Background(), &Request{Method: "POST", Path: path, Body: body}, nil)
}

func newTestClient(t *testing.T, h http.HandlerFunc, opts ...Option) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	c, err := NewClient(testReseller, "secret-key", append([]Option{WithBaseURL(srv.URL), WithPacing(false)}, opts...)...)
	if err != nil {
		t.Fatal(err)
	}
	return c
}

func TestHeadersAndDecode(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-API-KEY") != "secret-key" || r.Header.Get("__reseller") != testReseller {
			t.Errorf("auth headers missing: %v", r.Header)
		}
		if !strings.HasPrefix(r.Header.Get("User-Agent"), "DomainNameAPI-CLI") {
			t.Errorf("user agent = %q", r.Header.Get("User-Agent"))
		}
		if r.Method != "POST" || r.URL.Path != "/api/v1/domains/search" || r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
		var in map[string]string
		_ = json.NewDecoder(r.Body).Decode(&in)
		if in["domainName"] != "example.com" {
			t.Errorf("body = %+v", in)
		}
		w.Header().Set("X-Quota-Scope", "domain")
		w.Header().Set("X-Quota-Used", "12.4")
		w.Header().Set("X-Quota-Limit", "200")
		w.Header().Set("X-Quota-Cost", "1")
		w.Header().Set("X-Correlation-Id", "abc")
		io.WriteString(w, `{"success":true,"info":{"domainName":"example.com","status":"available","price":9.99}}`)
	})
	var res struct {
		Success bool `json:"success"`
		Info    struct {
			Status string  `json:"status"`
			Price  float64 `json:"price"`
		} `json:"info"`
	}
	resp, err := c.Do(context.Background(), &Request{Method: "POST", Path: "/api/v1/domains/search", Body: map[string]string{"domainName": "example.com"}}, &res)
	if err != nil {
		t.Fatal(err)
	}
	if !res.Success || res.Info.Status != "available" || res.Info.Price != 9.99 {
		t.Errorf("decoded %+v", res.Info)
	}
	var raw json.RawMessage
	if _, err := c.Do(context.Background(), &Request{Method: "POST", Path: "/api/v1/domains/search", Body: map[string]string{"domainName": "example.com"}}, &raw); err != nil || !strings.HasPrefix(string(raw), `{"success":true`) {
		t.Errorf("raw passthrough = %s, %v", raw, err)
	}
	if resp.CorrelationID != "abc" || resp.Quota == nil || resp.Quota.Used != 12.4 || resp.Quota.Scope != "domain" {
		t.Errorf("meta %+v %+v", resp, resp.Quota)
	}
}

func TestErrors(t *testing.T) {
	cases := []struct {
		name   string
		status int
		ctype  string
		body   string
		code   string
		kind   ErrorKind
		msg    string
	}{
		{"suspended reseller", 302, "", "", "INVALID_RESELLER", KindAuth, "reseller ID is unknown"},
		{"bad key", 401, "text/plain", "Unauthorized! Invalid API Key.", "INVALID_API_KEY", KindAuth, "Invalid API Key"},
		{"missing key", 401, "text/plain", "Unauthorized! API Key is missing.", "MISSING_API_KEY", KindAuth, "missing"},
		{"abp validation", 400, "application/json", `{"error":{"code":"Dna:1018","message":"Invalid forward type","validationErrors":[{"message":"bad","members":["forwardType"]}]}}`, "Dna:1018", KindValidation, "Invalid forward type"},
		{"abp boilerplate", 500, "application/json", `{"error":{"code":"Dna.X:9","message":"Exception of type 'Volo.Abp.BusinessException' was thrown.","details":"Dna Domain is locked at registry"}}`, "Dna.X:9", KindServer, "Domain is locked at registry"},
		{"not found", 404, "", "", "HTTP_404", KindNotFound, "not found"},
		{"domain not found", 404, "application/json", `{"error":{"code":"Dna.DomainService:Domain:10007","message":"The domain name you requested could not be found."}}`, "Dna.DomainService:Domain:10007", KindNotFound, "could not be found"},
		{"internal error sent as 404", 404, "application/json", `{"error":{"code":"Dna.DomainService:General:10001","message":"An internal error has occurred. Please contact the system administrator."}}`, "Dna.DomainService:General:10001", KindServer, "internal error"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
				if tc.status == 302 {
					w.Header().Set("Location", "/Error/Suspend?reseller=&user=")
				}
				if tc.ctype != "" {
					w.Header().Set("Content-Type", tc.ctype)
				}
				w.WriteHeader(tc.status)
				io.WriteString(w, tc.body)
			}, WithRetries(0))
			_, err := get(c, "/api/v1/deposit/accounts/me")
			var ae *APIError
			if !errors.As(err, &ae) {
				t.Fatalf("want APIError, got %v", err)
			}
			if ae.Code != tc.code || ae.Kind() != tc.kind || !strings.Contains(ae.Message, tc.msg) {
				t.Errorf("got code=%s kind=%s msg=%q", ae.Code, ae.Kind(), ae.Message)
			}
		})
	}
}

func TestRetries(t *testing.T) {
	var calls atomic.Int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(429)
			return
		}
		io.WriteString(w, `{"success":true,"domainId":"00000000-0000-0000-0000-000000000000"}`)
	})
	// 429 is rejected before the backend runs, so even a billable POST is retried.
	if _, err := post(c, "/api/v1/domains/renew", map[string]any{"domainName": "example.com", "period": 1}); err != nil {
		t.Fatal(err)
	}
	if calls.Load() != 2 {
		t.Errorf("calls = %d, want 2", calls.Load())
	}

	calls.Store(0)
	c = newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		w.WriteHeader(503)
	}, WithRetries(1))
	if _, err := post(c, "/api/v1/domains/renew", map[string]any{"domainName": "example.com", "period": 1}); err == nil {
		t.Fatal("expected error")
	}
	if calls.Load() != 1 {
		t.Errorf("POST retried on 503: calls = %d", calls.Load())
	}
	calls.Store(0)
	if _, err := get(c, "/api/v1/domains/info?DomainName=example.com"); err == nil {
		t.Fatal("expected error")
	}
	if calls.Load() != 2 {
		t.Errorf("GET not retried on 503: calls = %d", calls.Load())
	}
}

func TestNoCredentials(t *testing.T) {
	c, _ := NewClient("", "")
	if _, err := get(c, "/api/v1/deposit/accounts/me"); !errors.Is(err, ErrNoCredentials) {
		t.Errorf("err = %v", err)
	}
}

func TestParseEnvironment(t *testing.T) {
	for in, want := range map[string]Environment{"": Production, "prod": Production, "OTE": OTE, "test": OTE, "true": OTE, "0": Production} {
		got, err := ParseEnvironment(in)
		if err != nil || got != want {
			t.Errorf("ParseEnvironment(%q) = %v, %v", in, got, err)
		}
	}
	if _, err := ParseEnvironment("staging2"); err == nil {
		t.Error("expected error")
	}
}

func TestEnums(t *testing.T) {
	if e, _ := LookupEnum("DomainStatus"); e.NameOf(7) != "PendingDelete" || e.NameOf(42) != "" {
		t.Error("DomainStatus.NameOf")
	}
	e, _ := LookupEnum("recordtype")
	if v, ok := e.Normalize("txt"); !ok || v != "TXT" {
		t.Errorf("Normalize txt = %v", v)
	}
	if e, ok := EnumFor("DnsEditNameServerInput", "dnssecMode"); !ok || e.Name != "DNSSECMode" {
		t.Error("EnumFor dnssecMode")
	}
}

func TestPathEscaping(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.EscapedPath() != "/api/v1/contacts/DNA%201%2F2" {
			t.Errorf("path = %s", r.URL.EscapedPath())
		}
	})
	if _, err := get(c, "/api/v1/contacts/DNA%201%2F2"); err != nil {
		t.Fatal(err)
	}
}

func TestPacerSpacesZoneCalls(t *testing.T) {
	p := newPacer()
	ctx := context.Background()
	start := time.Now()
	for i := 0; i < 2; i++ {
		if err := p.wait(ctx, "/api/v1/domains/zones"); err != nil {
			t.Fatal(err)
		}
	}
	if d := time.Since(start); d < 900*time.Millisecond {
		t.Errorf("zone calls not spaced: %s", d)
	}
}
