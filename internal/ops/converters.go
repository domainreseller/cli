package ops

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/domainreseller/cli/internal/dna"
)

var contactRoles = []string{dna.ContactTypeRegistrant, dna.ContactTypeAdmin, dna.ContactTypeTech, dna.ContactTypeBilling}

func normalizeRole(s string) (string, error) {
	e, _ := dna.LookupEnum("ContactType")
	v, ok := e.Normalize(s)
	if !ok {
		return "", fmt.Errorf("unknown contact role %q (use Registrant, Admin, Tech or Billing)", s)
	}
	return v.(string), nil
}

// handleConverter builds [{contactType, handle}] from --handle ROLE=HANDLE.
// A bare --handle HANDLE applies to every role that is not set explicitly.
func handleConverter(target string) Converter {
	return Converter{
		Flag:   "handle",
		Target: target,
		Help:   "Contact handle per role as ROLE=HANDLE (Registrant, Admin, Tech, Billing), or a bare HANDLE for all four roles; repeatable",
		Build: func(values []string) (any, error) {
			byRole := map[string]string{}
			var all string
			for _, v := range values {
				role, handle, ok := strings.Cut(v, "=")
				if !ok {
					if all != "" {
						return nil, fmt.Errorf("--handle %q: only one bare handle (for all roles) is allowed", v)
					}
					all = strings.TrimSpace(v)
					continue
				}
				r, err := normalizeRole(role)
				if err != nil {
					return nil, err
				}
				byRole[r] = strings.TrimSpace(handle)
			}
			var out []any
			for _, r := range contactRoles {
				h := byRole[r]
				if h == "" {
					h = all
				}
				if h != "" {
					out = append(out, map[string]any{"contactType": r, "handle": h})
				}
			}
			return out, nil
		},
	}
}

// allRolesContactConverter copies one contact (JSON or @file) into all four roles.
func allRolesContactConverter(target string) Converter {
	return Converter{
		Flag:   "contact",
		Target: target,
		Help:   "One contact as JSON or @file, used for all four roles (Registrant, Admin, Tech, Billing)",
		Build: func(values []string) (any, error) {
			if len(values) != 1 {
				return nil, fmt.Errorf("--contact can be given once")
			}
			raw, err := ReadValue(values[0])
			if err != nil {
				return nil, err
			}
			var contact map[string]any
			if err := json.Unmarshal([]byte(raw), &contact); err != nil {
				return nil, fmt.Errorf("--contact: expected a JSON object: %w", err)
			}
			out := make([]any, 0, len(contactRoles))
			for _, r := range contactRoles {
				c := make(map[string]any, len(contact)+1)
				for k, v := range contact {
					c[k] = v
				}
				c["contactType"] = r
				out = append(out, c)
			}
			return out, nil
		},
	}
}

// ipConverter builds [{ipAddress, ipVersion}] from --ip ADDRESS.
func ipConverter(target string) Converter {
	return Converter{
		Flag:   "ip",
		Target: target,
		Help:   "IP address of the host, IPv4 or IPv6 (version detected); repeatable",
		Build: func(values []string) (any, error) {
			var out []any
			for _, v := range values {
				for _, part := range strings.Split(v, ",") {
					part = strings.TrimSpace(part)
					if part == "" {
						continue
					}
					ip := net.ParseIP(part)
					if ip == nil {
						return nil, fmt.Errorf("--ip %q is not an IP address", part)
					}
					ver := dna.IPVersion6
					if ip.To4() != nil {
						ver = dna.IPVersion4
					}
					out = append(out, map[string]any{"ipAddress": part, "ipVersion": ver})
				}
			}
			return out, nil
		},
	}
}

// dsConverter builds DS records from --ds "KEYTAG ALGORITHM DIGESTTYPE DIGEST".
func dsConverter(target string) Converter {
	return Converter{
		Flag:   "ds",
		Target: target,
		Help:   `DS record as "KEYTAG ALGORITHM DIGESTTYPE DIGEST" (as shown by your DNS provider); repeatable`,
		Build: func(values []string) (any, error) {
			var out []any
			for _, v := range values {
				f := strings.Fields(strings.ReplaceAll(v, ",", " "))
				if len(f) < 4 {
					return nil, fmt.Errorf("--ds %q: expected KEYTAG ALGORITHM DIGESTTYPE DIGEST", v)
				}
				keyTag, err1 := strconv.Atoi(f[0])
				alg, err2 := strconv.Atoi(f[1])
				dt, err3 := strconv.Atoi(f[2])
				if err1 != nil || err2 != nil || err3 != nil {
					return nil, fmt.Errorf("--ds %q: key tag, algorithm and digest type must be numbers", v)
				}
				out = append(out, map[string]any{"keyTag": keyTag, "algorithm": alg, "digestType": dt, "digest": strings.Join(f[3:], "")})
			}
			return out, nil
		},
	}
}

// documentConverter reads --document TYPE=PATH files and base64-encodes them.
func documentConverter(target string) Converter {
	return Converter{
		Flag:   "document",
		Target: target,
		Help:   "Document to upload as TYPE=PATH, e.g. IdentityCard=./id.pdf (see `dna enums DocumentType`); repeatable",
		Build: func(values []string) (any, error) {
			e, _ := dna.LookupEnum("DocumentType")
			var out []any
			for _, v := range values {
				typ, path, ok := strings.Cut(v, "=")
				if !ok {
					return nil, fmt.Errorf("--document %q: expected TYPE=PATH", v)
				}
				dt, ok := e.Normalize(typ)
				if !ok {
					return nil, fmt.Errorf("--document %q: unknown document type (allowed: %s)", typ, e.Summary())
				}
				if !documentExts[strings.ToLower(filepath.Ext(path))] {
					return nil, fmt.Errorf("--document %q: unsupported file type (use PDF, JPG, PNG, DOC, DOCX or ODT)", path)
				}
				b, err := os.ReadFile(path)
				if err != nil {
					return nil, fmt.Errorf("--document: %w", err)
				}
				if len(b) > maxDocumentBytes {
					return nil, fmt.Errorf("--document %q: %d bytes is over the 10 MB limit", path, len(b))
				}
				out = append(out, map[string]any{
					"documentType": dt,
					"fileName":     filepath.Base(path),
					"fileContent":  base64.StdEncoding.EncodeToString(b),
				})
			}
			return out, nil
		},
	}
}

// Document upload limits enforced by the API.
var documentExts = map[string]bool{".pdf": true, ".jpg": true, ".jpeg": true, ".png": true, ".doc": true, ".docx": true, ".odt": true}

const maxDocumentBytes = 10 << 20

// ReadValue returns s, or the content of the file when s is "@path".
func ReadValue(s string) (string, error) {
	if strings.HasPrefix(s, "@") && len(s) > 1 {
		b, err := os.ReadFile(s[1:])
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
	return s, nil
}

func prepareBulkSearch(in Input) error {
	currency, _ := in["currency"].(string)
	list, ok := toSlice(in["domains"])
	if !ok {
		if s, isStr := in["domains"].(string); isStr && s != "" {
			list = []any{s}
		}
	}
	var body []map[string]any
	for _, item := range list {
		switch x := item.(type) {
		case string:
			for _, name := range strings.FieldsFunc(x, func(r rune) bool { return r == ',' || r == ' ' || r == '\n' }) {
				entry := map[string]any{"domainName": name}
				if currency != "" {
					entry["currency"] = currency
				}
				body = append(body, entry)
			}
		case map[string]any:
			entry := map[string]any{}
			for k, v := range x {
				entry[k] = v
			}
			if _, has := entry["currency"]; !has && currency != "" {
				entry["currency"] = currency
			}
			body = append(body, entry)
		default:
			return fmt.Errorf("domains: unsupported value %v", item)
		}
	}
	if len(body) == 0 {
		return fmt.Errorf("domains: at least one domain name is required")
	}
	in.SetBody(body)
	delete(in, "domains")
	delete(in, "currency")
	return nil
}

// currencyWarnings notes the domains a search priced in another currency than
// the one asked for: the search endpoints answer in each TLD's own currency
// (USD for .com, TRY for .com.tr) whatever "currency" says.
func currencyWarnings(in Input, result []byte) []string {
	want, _ := in["currency"].(string)
	if want == "" {
		return nil
	}
	type info struct {
		DomainName string  `json:"domainName"`
		Currency   *string `json:"currency"`
	}
	var body struct {
		Info  *info  `json:"info"`
		Infos []info `json:"infos"`
	}
	if json.Unmarshal(result, &body) != nil {
		return nil
	}
	infos := body.Infos
	if body.Info != nil {
		infos = append(infos, *body.Info)
	}
	byCurrency := map[string][]string{}
	for _, i := range infos {
		if i.Currency == nil || *i.Currency == "" || strings.EqualFold(*i.Currency, want) {
			continue
		}
		byCurrency[*i.Currency] = append(byCurrency[*i.Currency], i.DomainName)
	}
	if len(byCurrency) == 0 {
		return nil
	}
	currencies := make([]string, 0, len(byCurrency))
	for c := range byCurrency {
		currencies = append(currencies, c)
	}
	sort.Strings(currencies)
	parts := make([]string, len(currencies))
	for i, c := range currencies {
		parts[i] = strings.Join(byCurrency[c], ", ") + " in " + c
	}
	return []string{fmt.Sprintf("requested %s but the API priced %s; search prices use each TLD's own currency (dna product tlds --currency converts)", want, strings.Join(parts, " and "))}
}

func setDefault(in Input, key string, v any) {
	if cur, ok := in[key]; !ok || isEmpty(cur) {
		in[key] = v
	}
}

func setNestedDefault(in Input, obj, key string, v any) {
	m, ok := toMap(in[obj])
	if !ok {
		return // let validation report the missing object
	}
	if cur, ok := m[key]; !ok || isEmpty(cur) {
		m[key] = v
	}
	in[obj] = m
}

// firstPrice formats the 1-year (or shortest) price of the first price group,
// e.g. "11.99 USD", for table output.
func firstPrice(row map[string]any, list string) string {
	groups, _ := row["prices"].([]any)
	for _, g := range groups {
		gm, _ := g.(map[string]any)
		items, _ := gm[list].([]any)
		var best map[string]any
		for _, it := range items {
			im, _ := it.(map[string]any)
			if p, _ := im["period"].(float64); best == nil || p == 1 {
				best = im
				if p == 1 {
					break
				}
			}
		}
		if best != nil {
			cur, _ := best["currency"].(string)
			return strings.TrimSpace(fmt.Sprintf("%v %s", best["price"], cur))
		}
	}
	return ""
}

func tldTableRow(row map[string]any) {
	row["register"] = firstPrice(row, "register")
	row["renew"] = firstPrice(row, "renew")
	row["transfer"] = firstPrice(row, "transfer")
	var req []any
	attrs, _ := row["attributes"].([]any)
	for _, a := range attrs {
		am, _ := a.(map[string]any)
		if r, _ := am["isRequired"].(bool); r {
			if k, _ := am["key"].(string); k != "" {
				req = append(req, k)
			}
		}
	}
	row["requiredAttributes"] = req
}

func sslTableRow(row map[string]any) {
	row["price"] = firstPrice(row, "create")
}

// transactionTableRow adds "change" (balance minus previous balance), which
// is reliable even when the API leaves amount at 0.
func transactionTableRow(row map[string]any) {
	bal, ok1 := row["balance"].(float64)
	old, ok2 := row["oldBalance"].(float64)
	if !ok1 || !ok2 {
		return
	}
	row["change"] = strconv.FormatFloat(math.Round((bal-old)*100)/100, 'f', 2, 64)
	if bal > old {
		row["change"] = "+" + row["change"].(string)
	}
}
