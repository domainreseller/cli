package cli

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/domainreseller/cli/internal/ops"
)

const testReseller = "11111111-2222-3333-4444-555555555555"

type recorded struct {
	Method string
	Path   string
	Query  string
	Body   string
}

type fakeAPI struct {
	t        *testing.T
	srv      *httptest.Server
	mu       sync.Mutex
	requests []recorded
	routes   map[string]func(w http.ResponseWriter, r *http.Request)
}

func newFakeAPI(t *testing.T) *fakeAPI {
	f := &fakeAPI{t: t, routes: map[string]func(http.ResponseWriter, *http.Request){}}
	f.srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		f.mu.Lock()
		f.requests = append(f.requests, recorded{r.Method, r.URL.Path, r.URL.RawQuery, string(b)})
		f.mu.Unlock()
		if r.Header.Get("X-API-KEY") != "k" || r.Header.Get("__reseller") != testReseller {
			w.WriteHeader(401)
			io.WriteString(w, "Unauthorized! Invalid API Key.")
			return
		}
		if h, ok := f.routes[r.Method+" "+r.URL.Path]; ok {
			h(w, r)
			return
		}
		w.WriteHeader(404)
		io.WriteString(w, `{"error":{"code":"NotFound","message":"no route"}}`)
	}))
	t.Cleanup(f.srv.Close)
	return f
}

func (f *fakeAPI) json(route, body string) {
	f.routes[route] = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, body)
	}
}

func (f *fakeAPI) all() []recorded {
	f.mu.Lock()
	defer f.mu.Unlock()
	return append([]recorded(nil), f.requests...)
}

func (f *fakeAPI) last() recorded {
	f.mu.Lock()
	defer f.mu.Unlock()
	if len(f.requests) == 0 {
		f.t.Fatal("no request recorded")
	}
	return f.requests[len(f.requests)-1]
}

type result struct {
	code   int
	stdout string
	stderr string
}

func run(t *testing.T, api *fakeAPI, env map[string]string, stdin string, args ...string) result {
	t.Helper()
	var out, errb bytes.Buffer
	base := map[string]string{
		"DNA_CONFIG":      filepath.Join(t.TempDir(), "config.json"),
		"DNA_RESELLER_ID": testReseller,
		"DNA_API_KEY":     "k",
	}
	if api != nil {
		base["DNA_BASE_URL"] = api.srv.URL
	}
	for k, v := range env {
		base[k] = v
	}
	app := &App{
		In: strings.NewReader(stdin), Out: &out, Err: &errb,
		Getenv: func(k string) string { return base[k] },
		Build:  BuildInfo{Version: "test"},
	}
	code := app.Run(context.Background(), args)
	return result{code, out.String(), errb.String()}
}

func decode(t *testing.T, s string) map[string]any {
	t.Helper()
	var m map[string]any
	if err := json.Unmarshal([]byte(s), &m); err != nil {
		t.Fatalf("not JSON: %v\n%s", err, s)
	}
	return m
}

func TestSearchPassesJSONThrough(t *testing.T) {
	api := newFakeAPI(t)
	api.json("POST /api/v1/domains/search", `{"success":true,"info":{"domainName":"example.com","status":"available"}}`)
	r := run(t, api, nil, "", "domain", "search", "example.com", "--currency", "USD")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	if got := decode(t, r.stdout)["info"].(map[string]any)["status"]; got != "available" {
		t.Errorf("status = %v", got)
	}
	if body := api.last().Body; body != `{"currency":"USD","domainName":"example.com"}` {
		t.Errorf("body = %s", body)
	}
}

func TestJQAndTable(t *testing.T) {
	api := newFakeAPI(t)
	api.json("GET /api/v1/domains", `{"totalCount":2,"items":[{"domainName":"a.com","status":0,"remainingDay":10},{"domainName":"b.com","status":7,"remainingDay":400}]}`)
	r := run(t, api, nil, "", "domain", "list", "--jq", ".items[] | select(.remainingDay < 30) | .domainName")
	if r.code != 0 || strings.TrimSpace(r.stdout) != "a.com" {
		t.Errorf("jq: exit %d out %q err %q", r.code, r.stdout, r.stderr)
	}
	r = run(t, api, nil, "", "domain", "list", "-o", "table")
	if !strings.Contains(r.stdout, "PendingDelete") || !strings.Contains(r.stdout, "DOMAIN") || strings.Contains(r.stdout, "\x1b[") {
		t.Errorf("table should show enum names:\n%s", r.stdout)
	}
}

func TestPaginationAll(t *testing.T) {
	api := newFakeAPI(t)
	api.routes["GET /api/v1/products/tlds"] = func(w http.ResponseWriter, r *http.Request) {
		skip := r.URL.Query().Get("SkipCount")
		switch skip {
		case "0":
			io.WriteString(w, `{"totalCount":3,"items":[{"name":"com"},{"name":"net"}]}`)
		default:
			io.WriteString(w, `{"totalCount":3,"items":[{"name":"org"}]}`)
		}
	}
	r := run(t, api, nil, "", "product", "tlds", "--all", "--max-result-count", "2", "--jq", "[.items[].name] | join(\",\")")
	if r.code != 0 || strings.TrimSpace(r.stdout) != "com,net,org" {
		t.Errorf("exit %d out %q err %s", r.code, r.stdout, r.stderr)
	}
}

func TestBillableNeedsYes(t *testing.T) {
	api := newFakeAPI(t)
	api.json("POST /api/v1/domains/renew", `{"success":true,"message":"ok"}`)
	r := run(t, api, nil, "", "domain", "renew", "example.com", "--period", "1")
	if r.code != ExitConfirmation {
		t.Fatalf("exit %d, want %d", r.code, ExitConfirmation)
	}
	if len(api.all()) != 0 {
		t.Fatal("request was sent without --yes")
	}
	if decode(t, r.stderr)["error"].(map[string]any)["kind"] != "confirmation_required" {
		t.Errorf("stderr = %s", r.stderr)
	}
	r = run(t, api, nil, "", "domain", "renew", "example.com", "--period", "1", "--yes")
	if r.code != 0 || api.last().Body != `{"domainName":"example.com","period":1}` {
		t.Errorf("exit %d body %s", r.code, api.last().Body)
	}
}

func TestDryRunSendsNothing(t *testing.T) {
	api := newFakeAPI(t)
	r := run(t, api, nil, "", "domain", "register", "example.com", "--ns", "ns1.x.net,ns2.x.net", "--handle", "DNA-1", "--dry-run")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	if len(api.all()) != 0 {
		t.Fatal("dry run sent a request")
	}
	m := decode(t, r.stdout)
	req := m["request"].(map[string]any)
	if req["method"] != "POST" || !strings.HasSuffix(req["url"].(string), "/api/v1/domains/register") {
		t.Errorf("request = %v", req)
	}
	if hdr := req["headers"].(map[string]any)["X-Api-Key"]; hdr != "****" {
		t.Errorf("API key not masked: %v", hdr)
	}
	body := req["body"].(map[string]any)
	if len(body["handles"].([]any)) != 4 || len(body["nameServers"].([]any)) != 2 {
		t.Errorf("body = %v", body)
	}
}

func TestReadOnlyBlocks(t *testing.T) {
	api := newFakeAPI(t)
	r := run(t, api, map[string]string{"DNA_READ_ONLY": "1"}, "", "domain", "lock", "example.com")
	if r.code != ExitPolicy || len(api.all()) != 0 {
		t.Errorf("exit %d, requests %d", r.code, len(api.all()))
	}
}

func TestErrorExitCodes(t *testing.T) {
	api := newFakeAPI(t)
	api.routes["GET /api/v1/domains/info"] = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		io.WriteString(w, `{"error":{"code":"Dna:Domain:1001","message":"Domain not in your account","validationErrors":[{"message":"invalid","members":["DomainName"]}]}}`)
	}
	r := run(t, api, nil, "", "domain", "info", "example.com")
	if r.code != ExitRejected {
		t.Fatalf("exit %d", r.code)
	}
	e := decode(t, r.stderr)["error"].(map[string]any)
	if e["code"] != "Dna:Domain:1001" || e["status"].(float64) != 400 || len(e["fields"].([]any)) != 1 {
		t.Errorf("error = %v", e)
	}
	r = run(t, api, map[string]string{"DNA_API_KEY": "wrong"}, "", "account", "get")
	if r.code != ExitAuth || decode(t, r.stderr)["error"].(map[string]any)["code"] != "INVALID_API_KEY" {
		t.Errorf("exit %d stderr %s", r.code, r.stderr)
	}
	r = run(t, api, nil, "", "domain", "renew", "example.com", "--period", "x")
	if r.code != ExitUsage {
		t.Errorf("bad flag: exit %d", r.code)
	}
	r = run(t, api, nil, "", "domain", "nosuch")
	if r.code != ExitUsage {
		t.Errorf("unknown command: exit %d (%s)", r.code, r.stderr)
	}
}

func TestBodyFromStdinAndFlagsOverride(t *testing.T) {
	api := newFakeAPI(t)
	api.json("POST /api/v1/domains/zones", ``)
	stdin := `{"domainName":"example.com","zoneStruct":{"name":"www","type":"A","ttl":300,"contents":["192.0.2.1"]}}`
	r := run(t, api, nil, stdin, "dns", "add", "--body", "-", "--ttl", "600", "--type", "aaaa", "--content", "2001:db8::1")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	last := api.last()
	if last.Query != "domainName=example.com" {
		t.Errorf("query = %s", last.Query)
	}
	if last.Body != `{"zoneStruct":{"contents":["2001:db8::1"],"name":"www","ttl":600,"type":"AAAA"}}` {
		t.Errorf("body = %s", last.Body)
	}
	if strings.TrimSpace(r.stdout) != `{"success":true}` {
		t.Errorf("empty response should print success: %q", r.stdout)
	}
}

func TestAuthLoginWritesProfile(t *testing.T) {
	api := newFakeAPI(t)
	api.json("GET /api/v1/deposit/accounts/me", `{"resellerName":"Acme","usdBalance":12.5}`)
	cfg := filepath.Join(t.TempDir(), "config.json")
	env := map[string]string{"DNA_CONFIG": cfg, "DNA_RESELLER_ID": "", "DNA_API_KEY": ""}
	r := run(t, nil, env, "k\n", "auth", "login", "--env", "ote", "--base-url", api.srv.URL, "--reseller-id", testReseller, "--api-key-stdin", "-o", "json")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	m := decode(t, r.stdout)
	if m["profile"] != "ote" || m["verified"] != true {
		t.Errorf("login result = %v", m)
	}
	b, err := os.ReadFile(cfg)
	if err != nil {
		t.Fatal(err)
	}
	var saved struct {
		Current  string                       `json:"current_profile"`
		Profiles map[string]map[string]string `json:"profiles"`
	}
	_ = json.Unmarshal(b, &saved)
	if saved.Current != "ote" || saved.Profiles["ote"]["api_key"] != "k" || saved.Profiles["ote"]["environment"] != "ote" {
		t.Errorf("saved config = %s", b)
	}
	// The stored profile is picked up by --env ote.
	r = run(t, nil, map[string]string{"DNA_CONFIG": cfg, "DNA_RESELLER_ID": "", "DNA_API_KEY": "", "DNA_BASE_URL": api.srv.URL}, "", "--env", "ote", "auth", "status", "-o", "json")
	if r.code != 0 || decode(t, r.stdout)["authenticated"] != true {
		t.Errorf("status: exit %d out %s err %s", r.code, r.stdout, r.stderr)
	}
}

func TestSchemaAndEnums(t *testing.T) {
	r := run(t, nil, nil, "", "schema")
	var list []map[string]any
	if err := json.Unmarshal([]byte(r.stdout), &list); err != nil || len(list) != len(ops.All()) {
		t.Fatalf("schema list: %v (%d items)", err, len(list))
	}
	for _, op := range ops.All() {
		r := run(t, nil, nil, "", "schema", op.ID, "--result")
		if r.code != 0 {
			t.Errorf("schema %s: exit %d %s", op.ID, r.code, r.stderr)
			continue
		}
		m := decode(t, r.stdout)
		if m["inputSchema"].(map[string]any)["type"] != "object" {
			t.Errorf("schema %s: bad input schema", op.ID)
		}
	}
	r = run(t, nil, nil, "", "enums", "DomainStatus")
	if decode(t, r.stdout)["name"] != "DomainStatus" {
		t.Errorf("enums: %s", r.stdout)
	}
}

func TestEveryCommandHasHelp(t *testing.T) {
	for _, op := range ops.All() {
		args := append(op.CommandPath(), "--help")
		r := run(t, nil, nil, "", args...)
		if r.code != 0 || !strings.Contains(r.stdout, op.Path) {
			t.Errorf("%s --help: exit %d", op.ID, r.code)
		}
	}
}

func TestMCPOverStdio(t *testing.T) {
	api := newFakeAPI(t)
	api.json("GET /api/v1/domains/info", `{"domainName":"example.com","status":"Active"}`)
	in := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{"protocolVersion":"2025-06-18","capabilities":{},"clientInfo":{"name":"test","version":"1"}}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized"}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list"}`,
		`{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"domain_info","arguments":{"domainName":"example.com"}}}`,
		`{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"domain_renew","arguments":{"domainName":"example.com","period":1}}}`,
	}, "\n") + "\n"
	r := run(t, api, nil, in, "mcp", "serve")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	byID := map[float64]map[string]any{}
	for _, line := range strings.Split(strings.TrimSpace(r.stdout), "\n") {
		m := decode(t, line)
		byID[m["id"].(float64)] = m
	}
	if v := byID[1]["result"].(map[string]any)["protocolVersion"]; v != "2025-06-18" {
		t.Errorf("protocol = %v", v)
	}
	tools := byID[2]["result"].(map[string]any)["tools"].([]any)
	if len(tools) != len(ops.All())+1 {
		t.Errorf("tools = %d", len(tools))
	}
	info := byID[3]["result"].(map[string]any)
	if info["isError"] != false || info["structuredContent"].(map[string]any)["status"] != "Active" {
		t.Errorf("domain_info = %v", info)
	}
	renew := byID[4]["result"].(map[string]any)
	if renew["isError"] != true || !strings.Contains(renew["content"].([]any)[0].(map[string]any)["text"].(string), "confirmation_required") {
		t.Errorf("domain_renew without confirm = %v", renew)
	}
	for _, req := range api.all() {
		if req.Path == "/api/v1/domains/renew" {
			t.Error("renew was executed without confirm")
		}
	}
}

func TestAgentGuideListsEveryCommand(t *testing.T) {
	g := AgentGuide()
	for _, op := range ops.All() {
		if !strings.Contains(g, "dna "+strings.Join(op.CommandPath(), " ")) {
			t.Errorf("guide is missing %s", op.ID)
		}
	}
	if !strings.HasPrefix(SkillFile(), "---\nname: dna\n") {
		t.Error("skill frontmatter")
	}
}

func TestInteractiveConfirmation(t *testing.T) {
	api := newFakeAPI(t)
	api.json("DELETE /api/v1/domains/forwards", ``)
	for _, tc := range []struct {
		answer string
		code   int
		sent   bool
	}{{"y\n", ExitOK, true}, {"n\n", ExitConfirmation, false}, {"\n", ExitConfirmation, false}} {
		before := len(api.all())
		var out, errb bytes.Buffer
		env := map[string]string{"DNA_CONFIG": filepath.Join(t.TempDir(), "c.json"), "DNA_RESELLER_ID": testReseller, "DNA_API_KEY": "k", "DNA_BASE_URL": api.srv.URL}
		app := &App{In: strings.NewReader(tc.answer), Out: &out, Err: &errb, Getenv: func(k string) string { return env[k] },
			Build: BuildInfo{Version: "test"}, StdinTTY: true, StderrTTY: true}
		code := app.Run(context.Background(), []string{"domain", "forward", "delete", "example.com"})
		if code != tc.code || (len(api.all()) > before) != tc.sent {
			t.Errorf("answer %q: exit %d, sent %v\n%s", tc.answer, code, len(api.all()) > before, errb.String())
		}
		if !strings.Contains(errb.String(), "DESTRUCTIVE") {
			t.Errorf("prompt missing: %s", errb.String())
		}
	}
}

func TestSearchWarnsOnCurrencyMismatch(t *testing.T) {
	api := newFakeAPI(t)
	api.json("POST /api/v1/domains/search", `{"success":true,"info":{"domainName":"a.com","currency":"USD","status":"AVAILABLE"}}`)
	r := run(t, api, nil, "", "domain", "search", "a.com", "--currency", "TRY", "-o", "json")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	if got := decode(t, r.stdout)["info"].(map[string]any)["currency"]; got != "USD" {
		t.Errorf("stdout currency = %v, want the API's USD untouched", got)
	}
	lines := strings.Split(strings.TrimSpace(r.stderr), "\n")
	if len(lines) != 1 {
		t.Fatalf("stderr = %q, want one warning line", r.stderr)
	}
	w, _ := decode(t, lines[0])["warning"].(string)
	if !strings.Contains(w, "TRY") || !strings.Contains(w, "USD") {
		t.Errorf("warning = %q", w)
	}
}

func TestBulkSearchWarnsOnCurrencyMismatch(t *testing.T) {
	api := newFakeAPI(t)
	api.json("POST /api/v1/domains/bulk-search", `{"success":true,"infos":[{"domainName":"a.com.tr","currency":"TRY"},{"domainName":"b.com","currency":"USD"}]}`)
	r := run(t, api, nil, "", "domain", "bulk-search", "a.com.tr", "b.com", "--currency", "TRY", "-o", "json")
	if r.code != 0 {
		t.Fatalf("exit %d: %s", r.code, r.stderr)
	}
	w, _ := decode(t, strings.TrimSpace(r.stderr))["warning"].(string)
	if !strings.Contains(w, "b.com") || strings.Contains(w, "a.com.tr") {
		t.Errorf("warning = %q, want b.com only", w)
	}
	if body := api.last().Body; !strings.Contains(body, `"currency":"TRY"`) {
		t.Errorf("body = %s, want currency still sent", body)
	}
}
