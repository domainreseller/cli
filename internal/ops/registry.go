package ops

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/domainreseller/cli/internal/openapi"
)

var (
	initOnce sync.Once
	byID     map[string]*Operation
	byTool   map[string]*Operation
)

func ensure() {
	initOnce.Do(func() {
		d := openapi.MustLoad()
		byID = map[string]*Operation{}
		byTool = map[string]*Operation{}
		for _, op := range registry {
			spec, ok := d.Operation(op.Method, op.Path)
			if !ok {
				panic(fmt.Sprintf("ops: %s refers to %s %s, which is not in the OpenAPI document", op.ID, op.Method, op.Path))
			}
			op.spec = spec
			if _, dup := byID[op.ID]; dup {
				panic("ops: duplicate id " + op.ID)
			}
			byID[op.ID] = op
			byTool[op.ToolName()] = op
		}
	})
}

// All returns every operation sorted by ID.
func All() []*Operation {
	ensure()
	out := make([]*Operation, len(registry))
	copy(out, registry)
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	return out
}

// Lookup finds an operation by ID ("domain.search"), tool name
// ("domain_search") or CLI path ("domain search").
func Lookup(name string) (*Operation, bool) {
	ensure()
	name = strings.TrimSpace(name)
	if op, ok := byID[name]; ok {
		return op, true
	}
	if op, ok := byTool[name]; ok {
		return op, true
	}
	if op, ok := byID[strings.Join(strings.Fields(name), ".")]; ok {
		return op, true
	}
	return nil, false
}

// Groups returns the top-level command groups with a one-line description.
func Groups() map[string]string {
	return map[string]string{
		"account": "Reseller account: balance and transactions",
		"contact": "Contact handles (registrant, admin, tech, billing)",
		"domain":  "Search, register, renew, transfer and manage domains",
		"dns":     "DNS zone records of domains on the platform nameservers",
		"dnssec":  "DNSSEC (DS records) of domains",
		"ssl":     "SSL certificates: buy, configure, issue, reissue, revoke",
		"hosting": "Web hosting packages and services",
		"order":   "Orders",
		"product": "Product catalogue and prices (TLDs, SSL, products)",
	}
}

var domainArg = []string{"domainName"}

// contactFields mirrors the API's contact validation (create and update):
// every field except company name, fax and disclose flag is required.
var contactFields = map[string]Field{
	"contact.firstName":        {Required: true, Help: "First name (at least 2 letters)"},
	"contact.lastName":         {Required: true, Help: "Last name (at least 2 letters)"},
	"contact.companyName":      {Help: "Company name (defaults to first + last name)"},
	"contact.eMail":            {Required: true, Help: "E-mail address"},
	"contact.address":          {Required: true, Help: "Street address"},
	"contact.city":             {Required: true, Help: "City"},
	"contact.state":            {Required: true, Help: "State or province"},
	"contact.postalCode":       {Required: true, Help: "Postal code"},
	"contact.country":          {Required: true, Help: "2-letter ISO country code, e.g. TR, US"},
	"contact.phoneCountryCode": {Required: true, Help: "Country calling code, digits only without +, e.g. 90"},
	"contact.phone":            {Required: true, Help: "Phone number, digits only"},
	"contact.discloseFlag":     {Help: "Publish the contact in WHOIS (unset keeps the registry default)"},
}

var registry = []*Operation{
	// ------------------------------------------------------------------ account
	{
		ID: "account.get", Method: "GET", Path: "/api/v1/deposit/accounts/me", Safety: Read,
		Aliases: []string{"balance", "me"},
		Summary: "Show the reseller account and its TRY/USD balances",
		Description: "Returns the reseller ID, name, group and the current TRY and USD balances. " +
			"This is also the cheapest way to verify that credentials work.",
		Examples: []Example{{"dna account get", "Show balances"}, {"dna account get --jq .usdBalance", "Print only the USD balance"}},
	},
	{
		ID: "account.transactions", Method: "GET", Path: "/api/v1/deposit/transactions", Safety: Read,
		Summary:     "List balance transactions (charges, refunds, deposits)",
		Description: "Pages through the balance ledger. Filter by date range, transaction type, product type and charge direction.",
		Fields: map[string]Field{
			"transactionTypes": {Enum: "TransactionType", Help: "Only these transaction types (name or number, repeatable)"},
			"productTypes":     {Enum: "ProductType", Help: "Only these product types (name or number, repeatable)"},
			"chargeType":       {Enum: "ChargeType", Help: "Add (credits) or Subtract (debits)"},
			"startDate":        {Help: "From this time (RFC 3339, e.g. 2026-01-01T00:00:00Z)"},
			"endDate":          {Help: "Until this time (RFC 3339)"},
			"resellerId":       {Hidden: true},
			"resellerNames":    {Hidden: true},
			"accountType":      {Hidden: true},
		},
		Columns:  []string{"creationTime", "type", "status", "change", "balance", "period", "description"},
		TableRow: transactionTableRow,
		Examples: []Example{{"dna account transactions --start-date 2026-11-01T00:00:00Z --transaction-types Register", "Registrations this month"}},
	},

	// ------------------------------------------------------------------ contact
	{
		ID: "contact.get", Method: "GET", Path: "/api/v1/contacts/{handleCode}", Safety: Read, Args: []string{"handleCode"},
		Summary:  "Show a contact handle and its verification state",
		Examples: []Example{{"dna contact get DNA-12345", ""}},
	},
	{
		ID: "contact.create", Method: "POST", Path: "/api/v1/contacts", Safety: Write,
		Summary: "Create a reusable contact handle",
		Description: "Creates a contact and returns its handle, which `dna domain register --handle` accepts. " +
			"country is a 2-letter ISO code; phone and phoneCountryCode are digits only (e.g. 90 and 5551234567).",
		Flatten:  map[string]string{"contact": ""},
		Fields:   contactFields,
		Examples: []Example{{`dna contact create --first-name Ada --last-name Lovelace --email ada@example.com --address "1 Example St" --city Istanbul --state Istanbul --country TR --postal-code 34000 --phone-country-code 90 --phone 5551234567`, ""}},
	},
	{
		ID: "contact.update", Method: "PUT", Path: "/api/v1/contacts/{handleCode}", Safety: Write, Args: []string{"handleCode"},
		Summary: "Update a contact handle (and the domains that use it)",
		Description: "Changes the contact details. Every domain using the handle is updated; the result lists per-domain outcomes " +
			"(some registries require documents or split the handle). Send the full contact: the API validates it like a new one.",
		Flatten:  map[string]string{"contact": ""},
		Fields:   contactFields,
		Examples: []Example{{"dna contact update DNA-12345 --email new@example.com --body @contact.json", "Start from a JSON file, override one field"}},
	},
	{
		ID: "contact.verification", Method: "GET", Path: "/api/v1/contacts/verification/check", Safety: Read, Args: []string{"handleCode"},
		Summary: "Show e-mail/phone verification status of a contact",
	},
	{
		ID: "contact.resend-verification", Method: "POST", Path: "/api/v1/contacts/verification/request", Safety: Write, Args: []string{"handleCode"},
		Summary: "Send the contact verification notice again",
		Fields:  map[string]Field{"handleCode": {Required: true}},
	},

	// ------------------------------------------------------------------ domain
	{
		ID: "domain.search", Method: "POST", Path: "/api/v1/domains/search", Safety: Read, Args: domainArg,
		Warnings:    currencyWarnings,
		Summary:     "Check if a domain is available and get its price",
		Description: "Returns availability (info.status), price, currency, premium flag and whether registry documents are required. Rate limit: 2 searches/second.",
		Fields:      map[string]Field{"domainName": {Required: true}, "currency": {Help: "Price currency, e.g. USD or TRY. The API may answer in the TLD's own currency instead (check currency in the result)"}},
		Examples:    []Example{{"dna domain search example.com", ""}, {"dna domain search example.com.tr --currency TRY -o json", ""}},
	},
	{
		ID: "domain.bulk-search", Method: "POST", Path: "/api/v1/domains/bulk-search", Safety: Read, Args: []string{"domains..."},
		Summary: "Check availability and prices of several domains at once",
		Virtual: []Virtual{
			{Key: "domains", Type: "strings", Required: true, Help: "Domain names to check"},
			{Key: "currency", Type: "string", Help: "Price currency, e.g. USD or TRY. The API may answer in the TLD's own currency instead (check currency in the result)"},
		},
		Prepare:  prepareBulkSearch,
		Warnings: currencyWarnings,
		Columns:  []string{"domainName", "status", "price", "currency", "period", "isPremium", "isDocumentRequired", "reason"},
		Examples: []Example{{"dna domain bulk-search example.com example.net example.org", ""}},
	},
	{
		ID: "domain.list", Method: "GET", Path: "/api/v1/domains", Safety: Read,
		Summary:     "List your domains",
		Description: "Pages through the reseller's domains. Use --all to fetch every page.",
		Fields: map[string]Field{
			"domainName": {Help: "Filter by (part of) the domain name"},
			"statusIds":  {Enum: "DomainStatus", EnumList: true, Help: "Filter by status, comma separated names or numbers (e.g. Active,PendingDelete)"},
		},
		Columns:  []string{"domainName", "status", "expirationDate", "remainingDay", "lockStatus", "privacyProtectionStatus", "renewalMode"},
		Examples: []Example{{"dna domain list --all -o json", "Every domain as JSON"}, {"dna domain list --jq '.items[] | select(.remainingDay < 30) | .domainName'", "Domains expiring within 30 days"}},
	},
	{
		ID: "domain.info", Method: "GET", Path: "/api/v1/domains/info", Safety: Read, Args: domainArg,
		Summary: "Show full details of a domain",
		Description: "Dates, status, lock and privacy state, nameservers, child hosts, contacts, auth code and registry attributes. " +
			"Not rate limited and costs 1 quota point.",
		Examples: []Example{{"dna domain info example.com", ""}, {"dna domain info example.com --jq .authCode", "Print the transfer auth code"}},
	},
	{
		ID: "domain.register", Method: "POST", Path: "/api/v1/domains/register", Safety: Billable, Args: domainArg,
		Summary: "Register a domain using existing contact handles (charges balance)",
		Description: "Registers the domain for --period years with the given nameservers and contact handles. " +
			"Check availability and price first with `dna domain search`. Some TLDs need --tld-attributes (see `dna product tlds --filter-text <tld>`).",
		Fields: map[string]Field{
			"period":        {Help: "Registration period in years (1-10)"},
			"nameServers":   {Flag: "ns", Split: true, Help: "Nameserver host names (repeat or comma separate; at least 2 usually required)"},
			"handles":       {Help: `Contact handles as JSON: [{"contactType":"Registrant","handle":"DNA-1"},...]`},
			"tldAttributes": {Flag: "tld-attribute", Help: "Registry-specific attribute"},
		},
		Converters: []Converter{handleConverter("handles")},
		Examples: []Example{
			{"dna domain register example.com --period 1 --ns ns1.example.net --ns ns2.example.net --handle DNA-12345 --yes", "Same handle for all four roles"},
			{"dna domain register example.com --ns ns1.example.net,ns2.example.net --handle Registrant=DNA-1 --handle Admin=DNA-2 --handle Tech=DNA-2 --handle Billing=DNA-2 --dry-run", "Preview the request"},
		},
	},
	{
		ID: "domain.register-with-contacts", Method: "POST", Path: "/api/v1/domains/register-with-contacts", Safety: Billable, Args: domainArg,
		Summary: "Register a domain with inline contact details (charges balance)",
		Description: "Like `domain register`, but contacts are given inline. All four roles (Registrant, Admin, Tech, Billing) are required " +
			"unless --use-trustee-contact is set. Pass them with --contacts @contacts.json, or one contact for every role with --contact @contact.json.",
		Fields: map[string]Field{
			"period":            {Help: "Registration period in years (1-10)"},
			"nameServers":       {Flag: "ns", Split: true, Help: "Nameserver host names (repeat or comma separate)"},
			"contacts":          {Help: "Contacts as a JSON array (or @file); each needs contactType"},
			"tldAttributes":     {Flag: "tld-attribute", Help: "Registry-specific attribute"},
			"useTrusteeContact": {Help: "Use the registrar's trustee contact (TLDs with isSupportTrustee)"},
		},
		Converters: []Converter{allRolesContactConverter("contacts")},
		Examples:   []Example{{"dna domain register-with-contacts example.com --ns ns1.example.net,ns2.example.net --contact @contact.json --dry-run", ""}},
	},
	{
		ID: "domain.renew", Method: "POST", Path: "/api/v1/domains/renew", Safety: Billable, Args: domainArg,
		Summary:     "Renew a domain (charges balance)",
		Description: "Extends the registration by --period years. Run `dna domain renew-check` first to see renewability and premium pricing.",
		Fields:      map[string]Field{"period": {Required: true, Help: "Years to add (1-10)"}},
		Examples:    []Example{{"dna domain renew example.com --period 1 --yes", ""}},
	},
	{
		ID: "domain.renew-check", Method: "POST", Path: "/api/v1/domains/renew/check", Safety: Read, Args: domainArg,
		Summary: "Check whether a domain can be renewed and at what price",
		Fields:  map[string]Field{"period": {Help: "Years to add (1-10)"}},
	},
	{
		ID: "domain.transfer", Method: "POST", Path: "/api/v1/domains/transfer", Safety: Billable, Args: domainArg,
		Summary:     "Transfer a domain in from another registrar (charges balance)",
		Description: "Starts an incoming transfer with the auth (EPP) code. Validate first with `dna domain transfer-check`; follow progress with `dna domain transfer-status`.",
		Fields: map[string]Field{
			"authCode": {Help: "Transfer authorization (EPP) code from the current registrar"},
			"period":   {Help: "Years added on transfer (1-10)"},
			"contacts": {Help: "Optional contacts as a JSON array (or @file)"},
		},
		Examples: []Example{{"dna domain transfer example.com --auth-code 'Xy7!kP2q' --yes", ""}},
	},
	{
		ID: "domain.transfer-check", Method: "POST", Path: "/api/v1/domains/transfers/check", Safety: Read, Args: []string{"domainName", "authCode"},
		Summary: "Check whether a domain can be transferred in with an auth code",
		Description: "Reports transfer availability, lock state, whether the auth code is valid or required, registry statuses and premium price. " +
			"Run before `domain transfer`.",
		Examples: []Example{{"dna domain transfer-check example.com 'Xy7!kP2q'", ""}},
	},
	{
		ID: "domain.transfer-status", Method: "POST", Path: "/api/v1/domains/transfers/query", Safety: Read, Args: domainArg,
		Summary: "Show the state of a pending transfer",
	},
	{
		ID: "domain.restore", Method: "POST", Path: "/api/v1/domains/restore", Safety: Billable, Args: domainArg,
		Summary:     "Restore a deleted domain from redemption (charges balance)",
		Description: "Only possible during the registry's redemption period (status PendingDelete). Restore fees are usually high; check `dna product tlds`.",
	},
	{
		ID: "domain.lock", Method: "POST", Path: "/api/v1/domains/lock", Safety: Write, Idempotent: true, Args: domainArg,
		Summary: "Enable the transfer lock",
		Fields:  map[string]Field{"domainName": {Required: true}},
	},
	{
		ID: "domain.unlock", Method: "POST", Path: "/api/v1/domains/unlock", Safety: Write, Idempotent: true, Args: domainArg,
		Summary: "Disable the transfer lock (needed before transferring out)",
		Fields:  map[string]Field{"domainName": {Required: true}},
	},
	{
		ID: "domain.privacy", Method: "POST", Path: "/api/v1/domains/privacy", Safety: Write, Idempotent: true, Args: []string{"domainName", "privacyStatus"},
		Summary: "Turn WHOIS privacy protection on or off",
		Fields: map[string]Field{
			"domainName":    {Required: true},
			"privacyStatus": {Flag: "enabled", Required: true, Help: "on/off (true/false)"},
		},
		Examples: []Example{{"dna domain privacy example.com on", ""}, {"dna domain privacy example.com off", ""}},
	},
	{
		ID: "domain.update-auth-code", Method: "PUT", Path: "/api/v1/domains/auth", Safety: Write, Args: domainArg,
		Summary:     "Update the transfer auth (EPP) code and return it",
		Description: "Asks the registry to update the domain's transfer authorization code and returns the resulting code. To only read the current code use `dna domain info --jq .authCode`.",
	},
	{
		ID: "domain.contacts.update", Method: "PUT", Path: "/api/v1/domains/contacts/update", Safety: Write, Args: domainArg,
		Summary:     "Replace a domain's contacts with inline contact details",
		Description: "Each contact needs contactType (Registrant, Admin, Tech, Billing). Registrant changes may trigger an owner change process and documents on some TLDs.",
		Fields: map[string]Field{
			"domainName": {Required: true},
			"contacts":   {Required: true, Help: "Contacts as a JSON array (or @file); each needs contactType"},
		},
		Converters: []Converter{allRolesContactConverter("contacts")},
	},
	{
		ID: "domain.contacts.set-handles", Method: "PUT", Path: "/api/v1/domains/contacts/update-with-handle", Safety: Write, Args: domainArg,
		Summary: "Point a domain's contact roles at existing contact handles",
		Fields: map[string]Field{
			"domainName":  {Required: true},
			"handleItems": {Required: true, Help: `Handles as JSON: [{"contactType":"Admin","handle":"DNA-2"}]`},
		},
		Converters: []Converter{handleConverter("handleItems")},
		Examples:   []Example{{"dna domain contacts set-handles example.com --handle Admin=DNA-2 --handle Tech=DNA-2", ""}},
	},
	{
		ID: "domain.documents.list", Method: "GET", Path: "/api/v1/domains/documents", Safety: Read, Args: domainArg,
		Summary: "List registry documents of a domain (mostly .tr)",
		Fields: map[string]Field{
			"documentTypes": {Enum: "DocumentType", EnumList: true, Help: "Filter by document type, comma separated names or numbers"},
		},
		Columns: []string{"name", "type", "status", "operationType", "uploadedDate"},
	},
	{
		ID: "domain.documents.upload", Method: "POST", Path: "/api/v1/domains/documents", Safety: Write, Args: domainArg,
		Summary: "Upload registry documents for a domain",
		Description: "Upload the documents a registry asked for (domain status WaitingForDocument). Use --document TYPE=PATH; the file is base64-encoded for you. " +
			"Accepted files: PDF, JPG/JPEG, PNG, DOC, DOCX, ODT; at most 10 MB per file and 50 MB per request.",
		Fields:     map[string]Field{"documents": {Help: "Documents as JSON: [{documentType, fileName, fileContent(base64)}]"}},
		Converters: []Converter{documentConverter("documents")},
		Examples:   []Example{{"dna domain documents upload example.com.tr --document IdentityCard=./id.pdf", ""}},
	},
	{
		ID: "domain.forward.get", Method: "GET", Path: "/api/v1/domains/forwards", Safety: Read, Args: domainArg,
		Summary: "Show the URL forwarding of a domain",
		Fields:  map[string]Field{"domainName": {Required: true}},
	},
	{
		ID: "domain.forward.set", Method: "POST", Path: "/api/v1/domains/forwards", Safety: Write, Args: []string{"domainName", "redirectAddress"},
		Summary: "Forward a domain to a URL",
		Fields: map[string]Field{
			"domainName":      {Required: true},
			"redirectAddress": {Flag: "to", Required: true, Help: "Target URL, e.g. https://www.example.com"},
			"forwardType":     {Flag: "type", Help: "Standard (301 redirect, default) or Frame"},
		},
		Prepare:  func(in Input) error { setDefault(in, "forwardType", "Standard"); return nil },
		Examples: []Example{{"dna domain forward set example.com https://www.example.org", ""}},
	},
	{
		ID: "domain.forward.delete", Method: "DELETE", Path: "/api/v1/domains/forwards", Safety: Destructive, Args: domainArg,
		Summary: "Remove the URL forwarding of a domain",
	},
	{
		ID: "domain.nameservers.set", Method: "PUT", Path: "/api/v1/domains/dns/name-server", Safety: Write, Idempotent: true, Args: []string{"domainName", "nameServers..."},
		Summary: "Replace the nameservers of a domain",
		Description: "Sets the complete nameserver list. On a DNSSEC-signed domain the change is blocked unless --dnssec-mode is " +
			"KeepDnssec (new provider serves a matching DNSKEY) or DisableDnssec (remove DS records first).",
		Fields: map[string]Field{
			"nameServers": {Flag: "ns", Split: true, Required: true, Help: "Nameserver host names (repeat or comma separate)"},
			"dnssecMode":  {Help: "Block (default), KeepDnssec or DisableDnssec"},
		},
		Examples: []Example{{"dna domain nameservers set example.com ns1.example.net ns2.example.net", ""}},
	},
	{
		ID: "domain.host.add", Method: "POST", Path: "/api/v1/domains/dns/host", Safety: Write, Args: []string{"domainName", "hostName"},
		Summary:     "Create a child nameserver (glue record) such as ns1.example.com",
		Description: "Registers a host object under the domain with one or more IP addresses. The IP version is detected for you when you use --ip.",
		Fields:      map[string]Field{"ipAddresses": {Help: `IPs as JSON: [{"ipAddress":"192.0.2.1","ipVersion":"v4"}]`}},
		Converters:  []Converter{ipConverter("ipAddresses")},
		Examples:    []Example{{"dna domain host add example.com ns1.example.com --ip 192.0.2.10 --ip 2001:db8::10", ""}},
	},
	{
		ID: "domain.host.update", Method: "PUT", Path: "/api/v1/domains/dns/host", Safety: Write, Args: []string{"domainName", "hostName"},
		Summary:    "Change a child nameserver's name or IP addresses",
		Fields:     map[string]Field{"newHostName": {Help: "Rename the host"}, "ipAddresses": {Help: "IPs as JSON"}},
		Converters: []Converter{ipConverter("ipAddresses")},
		Examples:   []Example{{"dna domain host update example.com ns1.example.com --ip 192.0.2.20", ""}},
	},
	{
		ID: "domain.host.delete", Method: "DELETE", Path: "/api/v1/domains/dns/host", Safety: Destructive, Args: []string{"domainName", "hostName"},
		Summary: "Delete a child nameserver",
	},

	// ------------------------------------------------------------------ dns (zone)
	{
		ID: "dns.list", Method: "GET", Path: "/api/v1/domains/zones", Safety: Read, Args: domainArg,
		Summary: "List the DNS records of a domain",
		Fields:  map[string]Field{"domainName": {Required: true}},
		Columns: []string{"name", "type", "ttl", "records", "changeType"},
	},
	{
		ID: "dns.add", Method: "POST", Path: "/api/v1/domains/zones", Safety: Write, Args: domainArg,
		Summary: "Add values to a DNS record set (creates the set if needed)",
		Description: "Adds one or more values to the record set identified by --name and --type; existing values of that set are kept. " +
			"Names may be relative (www), the apex (@) or fully qualified. Record types are case-sensitive in the API; the CLI fixes the case. " +
			"TXT content is quoted automatically; CAA content needs exactly 3 parts (`0 issue \"letsencrypt.org\"`); CNAME needs a name. " +
			"Records only resolve publicly once the domain uses the platform nameservers (see `dna dns apply`).",
		Flatten: map[string]string{"zoneStruct": ""},
		Fields: map[string]Field{
			"domainName":          {Required: true},
			"zoneStruct.name":     {Help: "Record name, e.g. www (CNAME requires a name)"},
			"zoneStruct.ttl":      {Optional: true, Help: "TTL in seconds (1-86400, default 3600)"},
			"zoneStruct.contents": {Flag: "content", Help: "Record value; repeat for several values"},
			"zoneStruct.type":     {Help: "Record type (A, AAAA, CNAME, MX, TXT, NS, SRV, CAA, ...)"},
		},
		Prepare:  func(in Input) error { setNestedDefault(in, "zoneStruct", "ttl", 3600); return nil },
		Examples: []Example{{"dna dns add example.com --name www --type A --content 192.0.2.10 --ttl 3600", ""}, {`dna dns add example.com --name @ --type TXT --content "v=spf1 include:_spf.example.net ~all"`, ""}},
	},
	{
		ID: "dns.update", Method: "PUT", Path: "/api/v1/domains/zones", Safety: Write, Idempotent: true, Args: domainArg,
		Summary:     "Replace a DNS record set with new values",
		Description: "Replaces every value of the set --record-name + --type with the given --content values (and TTL); --name can rename it. Use `dns add` to append a value instead.",
		Flatten:     map[string]string{"zoneStruct": ""},
		Fields: map[string]Field{
			"domainName":          {Required: true},
			"recordName":          {Required: true, Help: "Name of the existing record set to replace"},
			"zoneStruct.name":     {Help: "New record name"},
			"zoneStruct.contents": {Flag: "content", Help: "Record value; repeat for several values"},
		},
		Examples: []Example{{"dna dns update example.com --record-name www --name www --type A --ttl 600 --content 192.0.2.20", ""}},
	},
	{
		ID: "dns.delete", Method: "DELETE", Path: "/api/v1/domains/zones", Safety: Destructive, Args: domainArg,
		Summary:     "Remove one value from a DNS record set",
		Description: "Removes the value --content from the set --name + --type. Other values of the set stay.",
		Fields: map[string]Field{
			"domainName": {Required: true},
			"name":       {Help: "Record name"},
			"recordType": {Flag: "type", Enum: "RecordType", Help: "Record type"},
			"record":     {Flag: "content", Help: "Exact record value to delete"},
		},
		Examples: []Example{{"dna dns delete example.com --name www --type A --content 192.0.2.10 --yes", ""}},
	},
	{
		ID: "dns.apply", Method: "POST", Path: "/api/v1/domains/zones/apply", Safety: Write, Idempotent: true, Args: domainArg,
		Summary: "Switch the domain to the platform nameservers so its DNS zone is served",
		Description: "If the domain does not already use the platform's nameservers, its nameservers are changed to them, so the records managed with `dna dns` become live. " +
			"This replaces any external DNS provider (e.g. Cloudflare) for the domain. Reports previous, current and system nameservers; does nothing when they already match.",
		Fields: map[string]Field{"domainName": {Required: true}},
	},

	// ------------------------------------------------------------------ dnssec
	{
		ID: "dnssec.get", Method: "GET", Path: "/api/v1/domains/dns/dnssec", Safety: Read, Args: domainArg,
		Summary: "Show DNSSEC status, DS records and allowed algorithms",
	},
	{
		ID: "dnssec.add", Method: "POST", Path: "/api/v1/domains/dns/dnssec", Safety: Write, Args: domainArg,
		Summary:     "Publish DS records at the registry (enable DNSSEC)",
		Description: "Adds DS records from your DNS provider. Wrong DS records make the domain unresolvable for validating resolvers - copy them exactly.",
		Fields:      map[string]Field{"records": {Help: `DS records as JSON: [{"keyTag":2371,"algorithm":13,"digestType":2,"digest":"..."}]`}},
		Converters:  []Converter{dsConverter("records")},
		Examples:    []Example{{`dna dnssec add example.com --ds "2371 13 2 1F987CC6583E92DF0890718C42..."`, "KEYTAG ALGORITHM DIGESTTYPE DIGEST"}},
	},
	{
		ID: "dnssec.remove", Method: "DELETE", Path: "/api/v1/domains/dns/dnssec", Safety: Destructive, Args: domainArg,
		Summary: "Remove one DS record",
		Fields: map[string]Field{
			"algorithm":  {Enum: "DNSSECAlgorithm", Help: "IANA algorithm number or name"},
			"digestType": {Enum: "DSDigestType", Help: "IANA digest type number or name"},
		},
	},
	{
		ID: "dnssec.disable", Method: "POST", Path: "/api/v1/domains/dns/dnssec/disable", Safety: Destructive, Args: domainArg,
		Summary: "Remove every DS record (disable DNSSEC)",
	},

	// ------------------------------------------------------------------ ssl
	{
		ID: "ssl.generate-csr", Method: "POST", Path: "/api/v1/ssls/generate-csr", Safety: Read, Args: domainArg,
		Summary:     "Generate a CSR and private key",
		Description: "Returns a CSR and its private key. Store the private key safely: it is needed to install the certificate.",
		Fields:      map[string]Field{"alternativeNames": {Flag: "san", Split: true, Help: "Extra names (SAN), repeat or comma separate"}},
		Examples:    []Example{{"dna ssl generate-csr example.com --country TR --state Istanbul --locality Istanbul --organization 'Example Ltd' --organization-unit IT -o json > csr.json", ""}},
	},
	{
		ID: "ssl.check-csr", Method: "POST", Path: "/api/v1/ssls/check-csr", Safety: Read,
		Summary:  "Decode a CSR and show its subject",
		Fields:   map[string]Field{"csr": {File: true, Help: "PEM-encoded CSR"}},
		Examples: []Example{{"dna ssl check-csr --csr @example.csr", ""}},
	},
	{
		ID: "ssl.initialize", Method: "POST", Path: "/api/v1/ssls/initialize", Safety: Billable, Args: domainArg,
		Summary:     "Buy an SSL certificate package for a domain (charges balance)",
		Description: "Find package codes with `dna product ssls`. Then run `dna ssl configure`.",
		Fields:      map[string]Field{"domainName": {Required: true}, "packageCode": {Required: true, Help: "SSL product code from `dna product ssls`"}},
	},
	{
		ID: "ssl.configure", Method: "PUT", Path: "/api/v1/ssls/configure", Safety: Write, Args: domainArg,
		Summary: "Submit CSR, validation method and contact for issuance",
		Description: "approverMail, authenticationType and webServerType are required for SSL packages. List valid approver addresses with " +
			"`dna ssl approvers`. Then poll `dna ssl status` until the certificate is issued.",
		Flatten: map[string]string{"contact": "contact-"},
		Fields: map[string]Field{
			"csr":                {File: true, Help: "PEM-encoded CSR"},
			"privateKey":         {File: true, Help: "PEM-encoded private key (optional when the CSR was generated elsewhere)"},
			"authenticationType": {Help: "Domain validation method"},
			"webServerType":      {Help: "Web server software"},
			"sanList":            {Help: "Additional names, comma separated"},
			"wildcardList":       {Help: "Wildcard names, comma separated"},
		},
		Examples: []Example{{"dna ssl configure example.com --csr @example.csr --approver-mail admin@example.com --authentication-type Email --web-server-type Apache --body @ssl-contact.json", ""}},
	},
	{
		ID: "ssl.status", Method: "PUT", Path: "/api/v1/ssls/check-status", Safety: Read, Idempotent: true, Args: domainArg,
		Summary: "Refresh and show certificate status",
	},
	{
		ID: "ssl.get", Method: "GET", Path: "/api/v1/ssls/certificate", Safety: Read, Args: domainArg,
		Summary: "Show certificate details (subscription, configuration, contacts)",
	},
	{
		ID: "ssl.content", Method: "GET", Path: "/api/v1/ssls/certificate-content", Safety: Read, Args: domainArg,
		Summary:  "Download the issued certificate and chain (PEM)",
		Examples: []Example{{"dna ssl content example.com --jq .certificate > example.crt", ""}},
	},
	{
		ID: "ssl.reissue", Method: "POST", Path: "/api/v1/ssls/reissue", Safety: Write, Args: domainArg,
		Summary: "Reissue a certificate with a new CSR/key or approver",
		Fields:  map[string]Field{"csr": {File: true, Help: "PEM-encoded CSR"}, "privateKey": {File: true, Help: "PEM-encoded private key"}},
	},
	{
		ID: "ssl.approvers", Method: "GET", Path: "/api/v1/ssls/approvers", Safety: Read, Args: domainArg,
		Summary: "List e-mail addresses allowed for e-mail validation",
		Columns: []string{"mails"},
	},
	{
		ID: "ssl.revoke", Method: "DELETE", Path: "/api/v1/ssls/revoke", Safety: Destructive, Args: domainArg,
		Summary: "Revoke a certificate (irreversible)",
	},

	// ------------------------------------------------------------------ hosting
	{
		ID: "hosting.packages", Method: "GET", Path: "/api/v1/hostings/packages", Safety: Read,
		Summary: "List hosting packages and their service plans",
		Columns: []string{"name", "requiredMinLeasing", "permittedMaxLeasing", "servicePlans"},
	},
	{
		ID: "hosting.plans", Method: "GET", Path: "/api/v1/hostings/plans", Safety: Read,
		Summary: "List hosting service plans with their limits",
		Columns: []string{"name", "diskQuotaMb", "bandwidthLimitMb", "maxDomains", "maxDatabases", "maxEmailAccounts"},
	},
	{
		ID: "hosting.create", Method: "POST", Path: "/api/v1/hostings/create", Safety: Billable, Args: domainArg,
		Summary:     "Order a hosting for a domain (charges balance)",
		Description: "Provisioning is asynchronous (HTTP 202). Poll `dna hosting status` until it is ready; the response then includes panel credentials.",
		Fields:      map[string]Field{"domainName": {Required: true}, "servicePlan": {Required: true, Help: "Service plan from `dna hosting packages`"}},
	},
	{
		ID: "hosting.get", Method: "GET", Path: "/api/v1/hostings", Safety: Read, Args: domainArg,
		Summary: "Show a hosting by domain name or service ID",
	},
	{
		ID: "hosting.get-by-token", Method: "GET", Path: "/api/v1/hostings/by-token", Safety: Read, Args: []string{"token"},
		Summary: "Show a hosting by the token returned when it was created",
	},
	{
		ID: "hosting.status", Method: "GET", Path: "/api/v1/hostings/check-status", Safety: Read, Args: domainArg,
		Summary: "Show hosting provisioning status (and panel access once ready)",
	},

	// ------------------------------------------------------------------ order
	{
		ID: "order.get", Method: "GET", Path: "/api/v1/order/{orderCode}", Safety: Read, Args: []string{"orderCode"},
		Summary: "Show an order and its lines",
	},

	// ------------------------------------------------------------------ product
	{
		ID: "product.list", Method: "GET", Path: "/api/v1/products", Safety: Read,
		Summary: "List products",
		Fields:  map[string]Field{"productType": {Enum: "ProductType", Help: "Domain, SSL, Deposit, DomainTrustee, Hosting or OwnerChange"}},
		Columns: []string{"productName", "productCode", "productType", "entityId"},
	},
	{
		ID: "product.info", Method: "GET", Path: "/api/v1/products/info", Safety: Read, Args: []string{"productName"},
		Summary:     "Show the price of a product for an order type and period",
		Description: "productName is a TLD without the leading dot (com, com.tr) or another product name. Example: price of a 2-year .com renewal.",
		Fields: map[string]Field{
			"productName": {Help: "TLD without dot (com, net, com.tr) or product name"},
			"orderType":   {Enum: "OrderType", Help: "Register, Renew, Transfer, Restore, ..."},
			"period":      {Help: "Years (1-10)"},
		},
		Examples: []Example{{"dna product info com --order-type Renew --period 2", ""}},
	},
	{
		ID: "product.tlds", Method: "GET", Path: "/api/v1/products/tlds", Safety: Read,
		Aliases:     []string{"tld"},
		Summary:     "List TLDs with prices, periods, rules and required attributes",
		Description: "The catalogue of TLDs you can sell: register/renew/transfer/restore prices per period, grace periods, capabilities (privacy, DNSSEC, IDN, trustee) and registry attributes. Not rate limited.",
		Fields:      map[string]Field{"filterText": {Help: "Filter by TLD, e.g. com.tr"}, "currency": {Help: "Price currency, e.g. USD or TRY"}},
		Columns:     []string{"name", "register", "renew", "transfer", "minRegistrationPeriod", "maxRegistrationPeriod", "isSaleAllowed", "isRequireDocument", "requiredAttributes"},
		TableRow:    tldTableRow,
		Examples:    []Example{{"dna product tlds --filter-text com.tr -o json", ""}, {"dna product tlds --all --jq '.items[] | select(.isDnssecAvailable) | .name'", "TLDs that support DNSSEC"}},
	},
	{
		ID: "product.ssls", Method: "GET", Path: "/api/v1/products/ssls", Safety: Read,
		Summary:  "List SSL certificate products with prices",
		Columns:  []string{"name", "productCode", "price", "vendorName", "validationType", "isWildcard", "isMultiDomain", "issueTimeMinutes"},
		TableRow: sslTableRow,
	},
}
