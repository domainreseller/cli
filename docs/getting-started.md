# Getting started

This walks through installing `dna`, connecting it to the **OTE test environment**, and doing a first registration
and DNS change without touching anything live.

## 1. Install

```bash
go install github.com/domainreseller/cli/cmd/dna@latest   # Go 1.26+
```

…or download a binary or a `.rpm` / `.deb` package from the [releases page](https://github.com/domainreseller/cli/releases/latest), or run
`scripts/install.sh` / `scripts/install.ps1` (see the [README](../README.md)). Then:

```bash
dna version
dna doctor        # checks config, network and credentials
```

Enable shell completion if you like: `dna completion --help` (bash, zsh, fish, PowerShell).

## 2. Get credentials

In the DomainNameAPI reseller panel open **API settings**. You need:

| value | looks like | note |
|---|---|---|
| Reseller ID | `1b2c3d4e-....-....-....-............` | a UUID, the same for both environments' panels |
| API key (production) | random string | used against `api.domainresellerapi.com` |
| API key (OTE) | random string | a **different** key, used against `ote.domainresellerapi.com` |

Panel logins (user name / password) do not work with the API.

## 3. Log in to OTE

```bash
dna auth login --env ote
```

`dna` asks for the reseller ID and the API key (hidden), verifies them with a real API call and stores them in the
profile `ote`. Check what is in effect at any time:

```bash
dna --ote auth status
```

`--ote` is short for `--env ote`; it also selects the `ote` profile.

## 4. Explore (read-only)

```bash
dna --ote account get                          # balances
dna --ote product tlds --filter-text com       # .com rules and prices
dna --ote domain search my-test-domain.com     # available? price?
dna --ote domain list
```

Nothing here changes anything. If you want a guarantee, add `--read-only`.

## 5. Register a domain

Registering needs contact handles. Create one (it can be used for all four roles):

```bash
dna --ote contact create --first-name Ada --last-name Lovelace --email ada@example.com \
  --address "1 Example Street" --city Istanbul --state Istanbul --country TR --postal-code 34000 \
  --phone-country-code 90 --phone 5551234567
# -> {"handle":"DNA-12345", ...}
```

Preview, then register:

```bash
dna --ote domain register my-test-domain.com --period 1 \
  --ns ns1.example.net --ns ns2.example.net --handle DNA-12345 --dry-run

dna --ote domain register my-test-domain.com --period 1 \
  --ns ns1.example.net --ns ns2.example.net --handle DNA-12345
# BILLABLE: Register a domain ... on OTE
# Continue? [y/N]
```

Registration is **billable**. On a terminal `dna` asks for confirmation; in scripts pass `--yes`.

## 6. Manage DNS

```bash
dna --ote dns add my-test-domain.com --name www --type A --content 192.0.2.10
dna --ote dns add my-test-domain.com --name @ --type MX --content "10 mail.example.net."
dna --ote dns list my-test-domain.com
dna --ote dns apply my-test-domain.com        # switch to the platform nameservers so the zone is served
```

## 7. Go live

```bash
dna auth login                 # production key, stored in profile "production"
dna domain list                # production is the default
```

## Next

- [Configuration](configuration.md): profiles, environment variables, read-only profiles
- [Usage](usage.md): every way to pass input, output formats, jq, pagination
- [AI agents](ai-agents.md): MCP server, Claude Code skill, safe automation
