package ops

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	"github.com/domainreseller/cli/internal/dna"
	"github.com/domainreseller/cli/internal/openapi"
)

// Every operation in the OpenAPI document must be in the registry, and every
// registry entry must exist in the document. This keeps the CLI, MCP server
// and docs complete when the spec changes.
func TestRegistryCoversSpec(t *testing.T) {
	d := openapi.MustLoad()
	have := map[string]string{}
	for _, op := range All() {
		have[op.Method+" "+op.Path] = op.ID
	}
	for _, sop := range d.Operations() {
		if _, ok := have[sop.Key()]; !ok {
			t.Errorf("spec operation %s (%s) has no registry entry", sop.Key(), sop.OperationID)
		}
	}
	if len(have) != len(d.Operations()) {
		t.Errorf("registry has %d operations, spec has %d", len(have), len(d.Operations()))
	}
}

func TestRegistryEntriesAreDocumented(t *testing.T) {
	tools := map[string]bool{}
	for _, op := range All() {
		if op.Summary == "" {
			t.Errorf("%s: missing summary", op.ID)
		}
		if op.Safety == "" {
			t.Errorf("%s: missing safety class", op.ID)
		}
		if n := op.ToolName(); len(n) > 64 || tools[n] {
			t.Errorf("%s: bad or duplicate tool name %q", op.ID, n)
		}
		tools[op.ToolName()] = true
		keys := map[string]bool{}
		for _, k := range op.InputKeys() {
			if keys[k] {
				t.Errorf("%s: input key %q collides (query parameter vs body property)", op.ID, k)
			}
			keys[k] = true
		}
		for _, a := range op.Args {
			if !keys[strings.TrimSuffix(a, "...")] {
				t.Errorf("%s: positional arg %q is not an input", op.ID, a)
			}
		}
		for k := range op.Fields {
			root := strings.SplitN(k, ".", 2)[0]
			if !keys[root] {
				t.Errorf("%s: field override %q does not match an input", op.ID, k)
			}
		}
		flags := map[string]bool{}
		for _, f := range op.Inputs() {
			if flags[f.Flag] {
				t.Errorf("%s: duplicate flag --%s", op.ID, f.Flag)
			}
			flags[f.Flag] = true
		}
		for _, c := range op.Converters {
			if flags[c.Flag] {
				t.Errorf("%s: converter flag --%s collides with a field flag", op.ID, c.Flag)
			}
		}
		if _, err := json.Marshal(op.InputSchema()); err != nil {
			t.Errorf("%s: input schema: %v", op.ID, err)
		}
	}
}

func mustOp(t *testing.T, id string) *Operation {
	t.Helper()
	op, ok := Lookup(id)
	if !ok {
		t.Fatalf("no operation %s", id)
	}
	return op
}

func TestBuildRequestQueryAndBody(t *testing.T) {
	op := mustOp(t, "dns.add")
	req, err := op.BuildRequest(Input{
		"domainName": "example.com",
		"zoneStruct": map[string]any{"name": "www", "type": "cname", "contents": []string{"target.example.net."}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if req.Method != "POST" || req.Path != "/api/v1/domains/zones" || req.Query.Get("domainName") != "example.com" {
		t.Fatalf("unexpected request: %+v", req)
	}
	body := req.Body.(map[string]any)["zoneStruct"].(map[string]any)
	if body["type"] != "CNAME" {
		t.Errorf("record type not normalized: %v", body["type"])
	}
	if body["ttl"] != 3600 {
		t.Errorf("default ttl not applied: %v", body["ttl"])
	}
}

func TestBuildRequestPathParam(t *testing.T) {
	req, err := mustOp(t, "contact.get").BuildRequest(Input{"handleCode": "DNA 1/2"})
	if err != nil {
		t.Fatal(err)
	}
	if req.Path != "/api/v1/contacts/DNA%201%2F2" {
		t.Errorf("path = %s", req.Path)
	}
}

func TestValidation(t *testing.T) {
	cases := []struct {
		id    string
		in    Input
		field string
	}{
		{"domain.renew", Input{"domainName": "example.com"}, "period"},
		{"domain.transfer", Input{"domainName": "example.com", "period": 11}, "period"},
		{"domain.search", Input{"domainName": "example.com", "domian": "x"}, "domian"},
		{"domain.contacts.set-handles", Input{"domainName": "example.com", "handleItems": []any{map[string]any{"contactType": "Owner", "handle": "H"}}}, "handleItems[0].contactType"},
		{"domain.register", Input{"domainName": "example.com", "nameServers": []string{"ns1.x"}}, "handles"},
	}
	for _, c := range cases {
		_, err := mustOp(t, c.id).BuildRequest(c.in)
		ie, ok := err.(*InputError)
		if !ok {
			t.Errorf("%s %v: expected InputError, got %v", c.id, c.in, err)
			continue
		}
		found := false
		for _, fe := range ie.Errors {
			if fe.Field == c.field {
				found = true
			}
		}
		if !found {
			t.Errorf("%s: expected error on %q, got %v", c.id, c.field, ie.Errors)
		}
	}
}

func TestEnumNamesAndLists(t *testing.T) {
	req, err := mustOp(t, "domain.list").BuildRequest(Input{"statusIds": "active, PendingDelete"})
	if err != nil {
		t.Fatal(err)
	}
	if got := req.Query.Get("StatusIds"); got != "0,7" {
		t.Errorf("StatusIds = %q", got)
	}
	req, err = mustOp(t, "account.transactions").BuildRequest(Input{"transactionTypes": []string{"Register", "renew"}, "chargeType": "Subtract"})
	if err != nil {
		t.Fatal(err)
	}
	if got := req.Query["TransactionTypes"]; !reflect.DeepEqual(got, []string{"1", "2"}) {
		t.Errorf("TransactionTypes = %v", got)
	}
	if req.Query.Get("ChargeType") != "1" {
		t.Errorf("ChargeType = %q", req.Query.Get("ChargeType"))
	}
	req, err = mustOp(t, "domain.nameservers.set").BuildRequest(Input{"domainName": "example.com", "nameServers": []string{"a.ns", "b.ns"}, "dnssecMode": "Nope"})
	if err == nil {
		t.Errorf("expected invalid dnssecMode error, got %+v", req.Body)
	}
	req, err = mustOp(t, "domain.nameservers.set").BuildRequest(Input{"domainName": "example.com", "nameServers": []string{"a.ns", "b.ns"}, "dnssecMode": "DisableDnssec"})
	if err != nil {
		t.Fatal(err)
	}
	if req.Body.(map[string]any)["dnssecMode"] != 2 {
		t.Errorf("dnssecMode = %v", req.Body.(map[string]any)["dnssecMode"])
	}
}

func TestBulkSearchBody(t *testing.T) {
	req, err := mustOp(t, "domain.bulk-search").BuildRequest(Input{"domains": []string{"a.com", "b.net"}, "currency": "USD"})
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(req.Body)
	if string(b) != `[{"currency":"USD","domainName":"a.com"},{"currency":"USD","domainName":"b.net"}]` {
		t.Errorf("body = %s", b)
	}
}

func TestConverters(t *testing.T) {
	v, err := handleConverter("handles").Build([]string{"DNA-1", "admin=DNA-2"})
	if err != nil {
		t.Fatal(err)
	}
	b, _ := json.Marshal(v)
	want := `[{"contactType":"Registrant","handle":"DNA-1"},{"contactType":"Admin","handle":"DNA-2"},{"contactType":"Tech","handle":"DNA-1"},{"contactType":"Billing","handle":"DNA-1"}]`
	if string(b) != want {
		t.Errorf("handles = %s", b)
	}
	v, err = ipConverter("ipAddresses").Build([]string{"192.0.2.1", "2001:db8::1"})
	if err != nil {
		t.Fatal(err)
	}
	b, _ = json.Marshal(v)
	if string(b) != `[{"ipAddress":"192.0.2.1","ipVersion":"v4"},{"ipAddress":"2001:db8::1","ipVersion":"v6"}]` {
		t.Errorf("ips = %s", b)
	}
	if _, err := ipConverter("x").Build([]string{"not-an-ip"}); err == nil {
		t.Error("expected error for invalid IP")
	}
	v, err = dsConverter("records").Build([]string{"2371 13 2 AB CD"})
	if err != nil {
		t.Fatal(err)
	}
	b, _ = json.Marshal(v)
	if string(b) != `[{"algorithm":13,"digest":"ABCD","digestType":2,"keyTag":2371}]` {
		t.Errorf("ds = %s", b)
	}
}

func TestInputSchemaHasEnums(t *testing.T) {
	s := mustOp(t, "dns.add").InputSchema()
	zs := s.Properties.Get("zoneStruct")
	if zs == nil || zs.Properties.Get("type") == nil || len(zs.Properties.Get("type").Enum) == 0 {
		t.Fatalf("zoneStruct.type has no enum: %+v", zs)
	}
	if !contains(s.Required, "domainName") {
		t.Errorf("domainName should be required: %v", s.Required)
	}
}

// Every integer enum in the spec must be documented in internal/dna/enums.go,
// otherwise tables, `dna enums` and MCP schemas would show bare numbers.
func TestEveryIntegerEnumIsDocumented(t *testing.T) {
	d := openapi.MustLoad()
	for name, s := range d.Components.Schemas {
		owner := openapi.ShortName(name)
		for _, prop := range s.Properties {
			ps := prop.Schema
			if ps.Type == "array" && ps.Items != nil {
				ps = ps.Items
			}
			if ps.Type == "integer" && len(ps.Enum) > 0 {
				if _, ok := dna.EnumFor(owner, prop.Name); !ok {
					t.Errorf("%s.%s has enum values %v but no entry in enumFields (internal/dna/enums.go)", owner, prop.Name, ps.Enum)
				}
			}
		}
	}
	for _, op := range All() {
		for _, p := range op.Spec().Params {
			ps := p.Schema
			if ps != nil && ps.Type == "array" && ps.Items != nil {
				ps = ps.Items
			}
			if ps != nil && len(ps.Enum) > 0 && op.Fields[lowerFirst(p.Name)].Enum == "" {
				t.Errorf("%s: query parameter %s has enum values %v but no Field.Enum in the registry", op.ID, p.Name, ps.Enum)
			}
		}
	}
}

func TestCurrencyWarnings(t *testing.T) {
	single := []byte(`{"info":{"domainName":"a.com","currency":"USD"}}`)
	bulk := []byte(`{"infos":[{"domainName":"a.com","currency":"USD"},{"domainName":"a.com.tr","currency":"TRY"},{"domainName":"b.io","currency":"USD"}]}`)
	cases := []struct {
		name  string
		in    Input
		body  []byte
		want  int
		has   []string
		lacks []string
	}{
		{"mismatch", Input{"currency": "TRY"}, single, 1, []string{"TRY", "USD", "a.com"}, nil},
		{"match ignoring case", Input{"currency": "usd"}, single, 0, nil, nil},
		{"no currency asked", Input{}, single, 0, nil, nil},
		{"no currency returned", Input{"currency": "TRY"}, []byte(`{"info":{"domainName":"a.com"}}`), 0, nil, nil},
		{"null currency", Input{"currency": "TRY"}, []byte(`{"info":{"domainName":"a.com","currency":null}}`), 0, nil, nil},
		{"bulk mixed", Input{"currency": "TRY"}, bulk, 1, []string{"a.com,", "b.io", "USD"}, []string{"a.com.tr"}},
		{"not json", Input{"currency": "TRY"}, []byte(`oops`), 0, nil, nil},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := currencyWarnings(c.in, c.body)
			if len(got) != c.want {
				t.Fatalf("got %d warnings %q, want %d", len(got), got, c.want)
			}
			for _, s := range c.has {
				if !strings.Contains(got[0], s) {
					t.Errorf("warning %q lacks %q", got[0], s)
				}
			}
			for _, s := range c.lacks {
				if strings.Contains(got[0], s) {
					t.Errorf("warning %q should not mention %q", got[0], s)
				}
			}
		})
	}
}
