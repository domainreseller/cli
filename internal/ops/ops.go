// Package ops is the operation registry: one entry per API operation, with
// the documentation, safety class and CLI ergonomics the OpenAPI document
// lacks. Everything user-facing is derived from it:
//
//   - CLI commands and flags (internal/cli)
//   - MCP tools and their input schemas (internal/mcp)
//   - `dna schema`, `dna agent guide` and the generated docs/reference
//
// Field names, types and constraints come from the embedded OpenAPI document,
// so the registry only says what the spec cannot.
package ops

import (
	"encoding/json"
	"fmt"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"unicode"

	"github.com/domainreseller/cli/internal/dna"
	"github.com/domainreseller/cli/internal/openapi"
)

// Safety classifies what an operation does to the account.
type Safety string

const (
	// Read operations have no side effects.
	Read Safety = "read"
	// Write operations change state but cost nothing and can be reverted.
	Write Safety = "write"
	// Destructive operations delete data or are hard to undo.
	Destructive Safety = "destructive"
	// Billable operations charge the reseller balance.
	Billable Safety = "billable"
)

// NeedsConfirmation reports whether the CLI asks before running it.
func (s Safety) NeedsConfirmation() bool { return s == Destructive || s == Billable }

// Mutates reports whether the operation changes anything.
func (s Safety) Mutates() bool { return s != Read }

// Field overrides how one input is presented. The key in Operation.Fields is
// the input path: a top-level key ("domainName") or a flattened nested key
// ("contact.firstName").
type Field struct {
	Flag     string // flag name, without dashes (default: kebab-case of the key)
	Short    string // one-letter shorthand
	Help     string // flag/tool description
	Required bool   // required even though the spec says optional
	Optional bool   // optional even though the spec says required
	Enum     string // enum (see dna.Enums) the value must belong to
	EnumList bool   // value is a comma-separated list of Enum values
	Split    bool   // split []string values on commas
	Hidden   bool   // no flag (still accepted via --body / MCP)
	File     bool   // a value of the form @path is read from that file
}

// Virtual is an input that is not sent as is: Prepare turns it into the real
// request (used when the API body is a bare array).
type Virtual struct {
	Key      string
	Type     string // "string", "strings"
	Help     string
	Required bool
}

// Example is one usage example.
type Example struct {
	Command     string
	Description string
}

// Converter is an extra repeatable CLI flag that builds a complex field
// (an array of objects) from simple strings, e.g. --ip 192.0.2.1.
type Converter struct {
	Flag   string
	Help   string
	Target string // input key it fills
	Build  func(values []string) (any, error)
}

// Operation is one API operation.
type Operation struct {
	// ID is the stable identifier, e.g. "domain.register". The CLI command is
	// the ID split on dots (`dna domain register`) and the MCP tool name is the
	// ID with dots and dashes turned into underscores (domain_register).
	ID      string
	Method  string
	Path    string
	Summary string // one line, starts with a verb
	// Description explains behaviour, prerequisites and pitfalls (markdown).
	Description string
	Safety      Safety
	// Idempotent marks a non-GET operation as safe to repeat.
	Idempotent bool
	// Args are positional arguments, mapped to input keys in order. A key
	// ending in "..." takes all remaining arguments.
	Args []string
	// Flatten lists body properties holding a single object whose fields
	// become top-level flags. The value is the flag prefix ("" for none).
	Flatten    map[string]string
	Fields     map[string]Field
	Virtual    []Virtual
	Converters []Converter
	// Prepare adjusts the input before it is validated and sent.
	Prepare  func(in Input) error
	Examples []Example
	// Columns are the default table columns for list output.
	Columns []string
	// TableRow adds derived, human-friendly columns (e.g. prices) to each
	// list row in table output. JSON output is never changed.
	TableRow func(row map[string]any)
	// Warnings inspects a successful response and returns notes for stderr.
	// stdout is never changed.
	Warnings func(in Input, result []byte) []string
	Aliases  []string

	spec *openapi.Operation
}

// Input is the canonical, JSON-shaped input of an operation: path and query
// parameters by camelCase name plus the body's top-level properties.
type Input map[string]any

// bodyKey is set by Prepare hooks to send a precomputed body.
const bodyKey = "$body"

// SetBody makes the request send body verbatim (used by Prepare hooks).
func (in Input) SetBody(body any) { in[bodyKey] = body }

// CommandPath is the CLI command path, e.g. ["domain", "register"].
func (o *Operation) CommandPath() []string { return strings.Split(o.ID, ".") }

// ToolName is the MCP tool name.
func (o *Operation) ToolName() string {
	return strings.NewReplacer(".", "_", "-", "_").Replace(o.ID)
}

// Spec returns the OpenAPI view of the operation.
func (o *Operation) Spec() *openapi.Operation { return o.spec }

// Paged reports whether the operation supports SkipCount/MaxResultCount.
func (o *Operation) Paged() bool {
	var skip, max bool
	for _, p := range o.spec.Params {
		switch p.Name {
		case "SkipCount":
			skip = true
		case "MaxResultCount":
			max = true
		}
	}
	return skip && max
}

// ResultType names the response schema ("" when the API returns no body).
func (o *Operation) ResultType() string {
	if o.spec.Result == nil {
		return ""
	}
	return doc().TypeName(o.spec.Result)
}

// BodyType names the request body schema ("" when there is none).
func (o *Operation) BodyType() string {
	if o.spec.Body == nil {
		return ""
	}
	return doc().TypeName(o.spec.Body)
}

func doc() *openapi.Document { return openapi.MustLoad() }

// ---------------------------------------------------------------------------
// Input fields (the CLI view: one entry per flag)

// Kind is how an input is entered on the command line.
type Kind string

const (
	KindString  Kind = "string"
	KindInt     Kind = "int"
	KindNumber  Kind = "number"
	KindBool    Kind = "bool"
	KindStrings Kind = "strings" // repeatable string flag
	KindMap     Kind = "map"     // repeatable KEY=VALUE flag
	KindJSON    Kind = "json"    // JSON value, @file or - for stdin
)

// InputField is one flag-level input.
type InputField struct {
	Key      string   // "domainName" or "contact.firstName"
	Path     []string // Key split on dots
	In       string   // "path", "query", "body", "virtual"
	Wire     string   // parameter name on the wire (path/query only)
	Kind     Kind
	Type     string // schema type name for docs ("string", "integer", "ContactHandleDto[]")
	Required bool
	Enum     *dna.EnumInfo
	EnumList bool
	Help     string
	Flag     string
	Short    string
	Split    bool
	Hidden   bool
	File     bool
	Schema   *openapi.Schema
}

// Inputs returns the operation's inputs in display order: positional
// arguments first, then required, then optional.
func (o *Operation) Inputs() []*InputField {
	d := doc()
	var out []*InputField
	add := func(f *InputField) {
		ov := o.Fields[f.Key]
		if ov.Flag != "" {
			f.Flag = ov.Flag
		}
		if f.Flag == "" {
			f.Flag = kebab(f.Path[len(f.Path)-1])
		}
		f.Short = ov.Short
		if ov.Help != "" {
			f.Help = ov.Help
		}
		if ov.Required {
			f.Required = true
		}
		if ov.Optional {
			f.Required = false
		}
		if ov.Enum != "" {
			f.Enum, _ = dna.LookupEnum(ov.Enum)
			f.EnumList = ov.EnumList
		}
		f.Split = ov.Split
		f.Hidden = ov.Hidden
		f.File = ov.File
		if f.Enum != nil && (f.Kind == KindInt || f.Kind == KindNumber) {
			f.Kind = KindString // accept names as well as numbers
		}
		out = append(out, f)
	}

	for _, p := range o.spec.Params {
		key := lowerFirst(p.Name)
		f := &InputField{Key: key, Path: []string{key}, In: p.In, Wire: p.Name, Required: p.Required || p.In == "path", Schema: p.Schema}
		f.Kind, f.Type = kindOf(d, p.Schema)
		add(f)
	}
	for _, v := range o.Virtual {
		f := &InputField{Key: v.Key, Path: []string{v.Key}, In: "virtual", Required: v.Required, Help: v.Help, Kind: KindString, Type: "string"}
		if v.Type == "strings" {
			f.Kind, f.Type = KindStrings, "string[]"
		}
		add(f)
	}
	if body := d.Resolve(o.spec.Body); body != nil && body.Type != "array" {
		owner := openapi.ShortName(openapi.RefName(o.spec.Body))
		for _, prop := range body.Properties {
			ps := prop.Schema
			if prefix, ok := o.Flatten[prop.Name]; ok {
				inner := d.Resolve(ps)
				innerOwner := openapi.ShortName(openapi.RefName(ps))
				for _, ip := range inner.Properties {
					key := prop.Name + "." + ip.Name
					f := &InputField{Key: key, Path: []string{prop.Name, ip.Name}, In: "body", Required: inner.IsRequired(ip.Name), Schema: ip.Schema}
					f.Kind, f.Type = kindOf(d, ip.Schema)
					f.Flag = prefix + kebab(ip.Name)
					f.Enum, _ = dna.EnumFor(innerOwner, ip.Name)
					f.Help = describeConstraints(ip.Schema)
					if f.Enum != nil && f.Kind != KindJSON {
						f.Kind = KindString
					}
					add(f)
				}
				continue
			}
			f := &InputField{Key: prop.Name, Path: []string{prop.Name}, In: "body", Required: body.IsRequired(prop.Name), Schema: ps}
			f.Kind, f.Type = kindOf(d, ps)
			f.Enum, _ = dna.EnumFor(owner, prop.Name)
			f.Help = describeConstraints(ps)
			if f.Enum != nil && f.Kind != KindJSON {
				f.Kind = KindString
			}
			add(f)
		}
	}

	pos := map[string]int{}
	for i, a := range o.Args {
		pos[strings.TrimSuffix(a, "...")] = i + 1
	}
	sort.SliceStable(out, func(i, j int) bool {
		pi, pj := pos[out[i].Key], pos[out[j].Key]
		if (pi > 0) != (pj > 0) {
			return pi > 0
		}
		if pi > 0 {
			return pi < pj
		}
		return out[i].Required && !out[j].Required
	})
	return out
}

func kindOf(d *openapi.Document, s *openapi.Schema) (Kind, string) {
	name := d.TypeName(s)
	r := d.Resolve(s)
	if r == nil {
		return KindString, name
	}
	switch r.Type {
	case "string":
		return KindString, name
	case "integer":
		return KindInt, name
	case "number":
		return KindNumber, name
	case "boolean":
		return KindBool, name
	case "array":
		it := d.Resolve(r.Items)
		if it != nil && (it.Type == "string" || it.Type == "integer" || it.Type == "number") {
			return KindStrings, name
		}
		return KindJSON, name
	}
	if r.HasAdditional && len(r.Properties) == 0 {
		return KindMap, name
	}
	return KindJSON, name
}

func describeConstraints(s *openapi.Schema) string {
	r := doc().Resolve(s)
	if r == nil {
		return ""
	}
	var parts []string
	if r.MaxLength != nil {
		parts = append(parts, fmt.Sprintf("max %d chars", *r.MaxLength))
	}
	if r.Minimum != nil && r.Maximum != nil {
		parts = append(parts, fmt.Sprintf("%g-%g", *r.Minimum, *r.Maximum))
	}
	if r.Format == "email" {
		parts = append(parts, "e-mail")
	}
	if len(parts) == 0 {
		return ""
	}
	return "(" + strings.Join(parts, ", ") + ")"
}

// ---------------------------------------------------------------------------
// Validation and request building

// FieldError is one problem with the input.
type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// InputError is returned when the input is invalid. It is reported before
// anything is sent to the API.
type InputError struct {
	Operation string       `json:"operation"`
	Errors    []FieldError `json:"errors"`
}

func (e *InputError) Error() string {
	parts := make([]string, len(e.Errors))
	for i, fe := range e.Errors {
		parts[i] = fe.Field + ": " + fe.Message
	}
	return "invalid input: " + strings.Join(parts, "; ")
}

// BuildRequest validates and normalizes in, then turns it into an API
// request. in is not modified.
func (o *Operation) BuildRequest(in Input) (*dna.Request, error) {
	in = cloneInput(in)
	if o.Prepare != nil {
		if err := o.Prepare(in); err != nil {
			return nil, &InputError{Operation: o.ID, Errors: []FieldError{{Field: "input", Message: err.Error()}}}
		}
	}
	if errs := o.normalize(in); len(errs) > 0 {
		return nil, &InputError{Operation: o.ID, Errors: errs}
	}
	req := &dna.Request{Method: o.Method, Path: o.Path, Query: url.Values{}, Idempotent: o.Idempotent || !o.Safety.Mutates()}
	for _, p := range o.spec.Params {
		key := lowerFirst(p.Name)
		val, ok := in[key]
		delete(in, key)
		if !ok || isEmpty(val) {
			continue
		}
		switch p.In {
		case "path":
			req.Path = strings.ReplaceAll(req.Path, "{"+p.Name+"}", url.PathEscape(formatScalar(val)))
		case "query":
			for _, s := range formatQuery(val) {
				req.Query.Add(p.Name, s)
			}
		}
	}
	if b, ok := in[bodyKey]; ok {
		req.Body = b
		delete(in, bodyKey)
	} else if o.spec.Body != nil {
		body := map[string]any{}
		schema := doc().Resolve(o.spec.Body)
		for _, prop := range schema.Properties {
			if v, ok := in[prop.Name]; ok {
				if !isEmpty(v) {
					body[prop.Name] = v
				}
				delete(in, prop.Name)
			}
		}
		req.Body = body
	}
	for _, v := range o.Virtual {
		delete(in, v.Key)
	}
	if len(in) > 0 { // normalize already rejects unknown keys; keep as a guard
		return nil, &InputError{Operation: o.ID, Errors: []FieldError{{Field: strings.Join(sortedKeys(in), ","), Message: "unknown input"}}}
	}
	return req, nil
}

// InputKeys lists the accepted top-level input keys.
func (o *Operation) InputKeys() []string {
	var keys []string
	for _, p := range o.spec.Params {
		keys = append(keys, lowerFirst(p.Name))
	}
	for _, v := range o.Virtual {
		keys = append(keys, v.Key)
	}
	if body := doc().Resolve(o.spec.Body); body != nil && body.Type != "array" {
		for _, prop := range body.Properties {
			keys = append(keys, prop.Name)
		}
	}
	return keys
}

// normalize checks in against the spec (and registry overrides), coercing
// values to their wire types in place: enum names become their values,
// numeric strings become numbers, and so on.
func (o *Operation) normalize(in Input) []FieldError {
	d := doc()
	var errs []FieldError
	known := map[string]bool{bodyKey: true}
	for _, k := range o.InputKeys() {
		known[k] = true
	}
	for _, k := range sortedKeys(in) {
		if !known[k] {
			msg := "unknown input; accepted: " + strings.Join(o.InputKeys(), ", ")
			if s := suggest(k, o.InputKeys()); s != "" {
				msg = fmt.Sprintf("unknown input (did you mean %q?); accepted: %s", s, strings.Join(o.InputKeys(), ", "))
			}
			errs = append(errs, FieldError{k, msg})
		}
	}

	required := func(key string, specRequired bool) bool {
		ov := o.Fields[key]
		return (specRequired || ov.Required) && !ov.Optional
	}

	for _, p := range o.spec.Params {
		key := lowerFirst(p.Name)
		val, ok := in[key]
		if !ok || isEmpty(val) {
			if required(key, p.Required || p.In == "path") {
				errs = append(errs, FieldError{key, "is required"})
			}
			continue
		}
		ov := o.Fields[key]
		var enum *dna.EnumInfo
		if ov.Enum != "" {
			enum, _ = dna.LookupEnum(ov.Enum)
		}
		nv, fe := coerce(d, key, p.Schema, enum, ov.EnumList, val)
		errs = append(errs, fe...)
		in[key] = nv
	}
	_, raw := in[bodyKey]
	for _, v := range o.Virtual {
		if val, ok := in[v.Key]; v.Required && !raw && (!ok || isEmpty(val)) {
			errs = append(errs, FieldError{v.Key, "is required"})
		}
	}
	if raw || o.spec.Body == nil {
		return errs
	}
	body := d.Resolve(o.spec.Body)
	if body.Type == "array" {
		return errs
	}
	owner := openapi.ShortName(openapi.RefName(o.spec.Body))
	for _, prop := range body.Properties {
		val, ok := in[prop.Name]
		if !ok || isEmpty(val) {
			if required(prop.Name, body.IsRequired(prop.Name)) {
				errs = append(errs, FieldError{prop.Name, "is required"})
			}
			continue
		}
		var enum *dna.EnumInfo
		if ov := o.Fields[prop.Name]; ov.Enum != "" {
			enum, _ = dna.LookupEnum(ov.Enum)
		} else {
			enum, _ = dna.EnumFor(owner, prop.Name)
		}
		nv, fe := coerce(d, prop.Name, prop.Schema, enum, o.Fields[prop.Name].EnumList, val)
		errs = append(errs, fe...)
		in[prop.Name] = nv
	}
	return errs
}

// coerce validates v against schema s and converts it to its wire form.
func coerce(d *openapi.Document, path string, s *openapi.Schema, enum *dna.EnumInfo, enumList bool, v any) (any, []FieldError) {
	owner := openapi.ShortName(openapi.RefName(s))
	r := d.Resolve(s)
	if r == nil {
		return v, nil
	}
	if enum != nil {
		if enumList {
			str := fmt.Sprint(v)
			if list, ok := toSlice(v); ok {
				parts := make([]string, len(list))
				for i, x := range list {
					parts[i] = fmt.Sprint(x)
				}
				str = strings.Join(parts, ",")
			}
			var out []string
			for _, part := range strings.Split(str, ",") {
				part = strings.TrimSpace(part)
				if part == "" {
					continue
				}
				nv, ok := enum.Normalize(part)
				if !ok {
					return v, []FieldError{{path, fmt.Sprintf("%q is not a valid %s (allowed: %s)", part, enum.Name, enum.Summary())}}
				}
				out = append(out, fmt.Sprint(nv))
			}
			return strings.Join(out, ","), nil
		}
		if r.Type != "array" {
			nv, ok := enum.Normalize(fmt.Sprint(v))
			if !ok {
				return v, []FieldError{{path, fmt.Sprintf("%v is not a valid %s (allowed: %s)", v, enum.Name, enum.Summary())}}
			}
			return nv, nil
		}
	}
	switch r.Type {
	case "string":
		str, ok := v.(string)
		if !ok {
			switch v.(type) {
			case float64, int, int64, bool, json.Number:
				str = fmt.Sprint(v)
			default:
				return v, []FieldError{{path, "must be a string"}}
			}
		}
		if r.MaxLength != nil && len([]rune(str)) > *r.MaxLength {
			return str, []FieldError{{path, fmt.Sprintf("is longer than %d characters", *r.MaxLength)}}
		}
		return str, nil
	case "integer":
		n, ok := toInt(v)
		if !ok {
			return v, []FieldError{{path, fmt.Sprintf("must be an integer, got %v", v)}}
		}
		if r.Minimum != nil && float64(n) < *r.Minimum || r.Maximum != nil && float64(n) > *r.Maximum {
			return n, []FieldError{{path, fmt.Sprintf("must be between %g and %g", deref(r.Minimum, math.Inf(-1)), deref(r.Maximum, math.Inf(1)))}}
		}
		return n, nil
	case "number":
		switch x := v.(type) {
		case float64:
			return x, nil
		case int:
			return float64(x), nil
		case string:
			f, err := strconv.ParseFloat(strings.TrimSpace(x), 64)
			if err != nil {
				return v, []FieldError{{path, "must be a number"}}
			}
			return f, nil
		}
		return v, []FieldError{{path, "must be a number"}}
	case "boolean":
		switch x := v.(type) {
		case bool:
			return x, nil
		case string:
			if b, ok := ParseBool(x); ok {
				return b, nil
			}
		}
		return v, []FieldError{{path, "must be true or false"}}
	case "array":
		list, ok := toSlice(v)
		if !ok {
			list = []any{v} // a single value where a list is expected
		}
		var errs []FieldError
		out := make([]any, 0, len(list))
		for i, item := range list {
			nv, fe := coerce(d, fmt.Sprintf("%s[%d]", path, i), r.Items, enum, false, item)
			errs = append(errs, fe...)
			out = append(out, nv)
		}
		return out, errs
	}
	// objects
	if r.HasAdditional && len(r.Properties) == 0 {
		m, ok := toMap(v)
		if !ok {
			return v, []FieldError{{path, "must be an object of string values"}}
		}
		return m, nil
	}
	m, ok := toMap(v)
	if !ok {
		return v, []FieldError{{path, "must be an object"}}
	}
	var errs []FieldError
	out := map[string]any{}
	for _, k := range sortedKeys(m) {
		if _, ok := r.Properties.Get(k); !ok {
			names := make([]string, len(r.Properties))
			for i, p := range r.Properties {
				names[i] = p.Name
			}
			errs = append(errs, FieldError{path + "." + k, "unknown field; accepted: " + strings.Join(names, ", ")})
		}
	}
	for _, prop := range r.Properties {
		val, ok := m[prop.Name]
		if !ok || isEmpty(val) {
			if r.IsRequired(prop.Name) {
				errs = append(errs, FieldError{path + "." + prop.Name, "is required"})
			}
			continue
		}
		pe, _ := dna.EnumFor(owner, prop.Name)
		nv, fe := coerce(d, path+"."+prop.Name, prop.Schema, pe, false, val)
		errs = append(errs, fe...)
		out[prop.Name] = nv
	}
	return out, errs
}

func deref(p *float64, def float64) float64 {
	if p == nil {
		return def
	}
	return *p
}

// ParseBool accepts true/false, yes/no, on/off, 1/0, enable(d)/disable(d).
func ParseBool(s string) (bool, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "yes", "y", "on", "1", "enable", "enabled":
		return true, true
	case "false", "no", "n", "off", "0", "disable", "disabled":
		return false, true
	}
	return false, false
}

func toInt(v any) (int, bool) {
	switch x := v.(type) {
	case int:
		return x, true
	case int64:
		return int(x), true
	case float64:
		if x == math.Trunc(x) {
			return int(x), true
		}
	case json.Number:
		n, err := x.Int64()
		return int(n), err == nil
	case string:
		n, err := strconv.Atoi(strings.TrimSpace(x))
		return n, err == nil
	}
	return 0, false
}

func toSlice(v any) ([]any, bool) {
	switch x := v.(type) {
	case []any:
		return x, true
	case []string:
		out := make([]any, len(x))
		for i, s := range x {
			out[i] = s
		}
		return out, true
	case []int:
		out := make([]any, len(x))
		for i, s := range x {
			out[i] = s
		}
		return out, true
	case []map[string]any:
		out := make([]any, len(x))
		for i, s := range x {
			out[i] = s
		}
		return out, true
	}
	return nil, false
}

func toMap(v any) (map[string]any, bool) {
	switch x := v.(type) {
	case map[string]any:
		return x, true
	case map[string]string:
		out := make(map[string]any, len(x))
		for k, s := range x {
			out[k] = s
		}
		return out, true
	case Input:
		return map[string]any(x), true
	}
	return nil, false
}

func isEmpty(v any) bool {
	switch x := v.(type) {
	case nil:
		return true
	case string:
		return x == ""
	case []any:
		return len(x) == 0
	case []string:
		return len(x) == 0
	case map[string]any:
		return len(x) == 0
	case map[string]string:
		return len(x) == 0
	}
	return false
}

func formatScalar(v any) string {
	switch x := v.(type) {
	case float64:
		if x == math.Trunc(x) {
			return strconv.FormatInt(int64(x), 10)
		}
		return strconv.FormatFloat(x, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(x)
	}
	return fmt.Sprint(v)
}

func formatQuery(v any) []string {
	if list, ok := toSlice(v); ok {
		out := make([]string, 0, len(list))
		for _, x := range list {
			out = append(out, formatScalar(x))
		}
		return out
	}
	return []string{formatScalar(v)}
}

func cloneInput(in Input) Input {
	out := make(Input, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
}

func sortedKeys[M ~map[string]V, V any](m M) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// suggest returns the closest accepted key, for typo hints.
func suggest(k string, accepted []string) string {
	norm := func(s string) string { return strings.ToLower(strings.NewReplacer("-", "", "_", "").Replace(s)) }
	best, bestDist := "", 3
	for _, a := range accepted {
		if d := levenshtein(norm(k), norm(a)); d < bestDist {
			best, bestDist = a, d
		}
	}
	return best
}

func levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	prev := make([]int, len(rb)+1)
	for j := range prev {
		prev[j] = j
	}
	for i := 1; i <= len(ra); i++ {
		cur := make([]int, len(rb)+1)
		cur[0] = i
		for j := 1; j <= len(rb); j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			cur[j] = min(prev[j]+1, cur[j-1]+1, prev[j-1]+cost)
		}
		prev = cur
	}
	return prev[len(rb)]
}

func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}

// kebab turns camelCase into kebab-case: nameServers -> name-servers.
func kebab(s string) string {
	switch s {
	case "eMail":
		return "email"
	}
	var sb strings.Builder
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			if i > 0 && !unicode.IsUpper(runes[i-1]) {
				sb.WriteByte('-')
			}
			sb.WriteRune(unicode.ToLower(r))
			continue
		}
		sb.WriteRune(r)
	}
	return sb.String()
}
