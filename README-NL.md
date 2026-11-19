<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa.svg" height="32" alt="العربية"></a>
  <a href="README-NL.md"><img src="./assets/lang/nl-active.svg" height="32" alt="Nederlands"></a>
  <a href="README-AZ.md"><img src="./assets/lang/az.svg" height="32" alt="Azərbaycanca"></a>
  <a href="README-CN.md"><img src="./assets/lang/cn.svg" height="32" alt="中文"></a>
  <a href="README-FR.md"><img src="./assets/lang/fr.svg" height="32" alt="Français"></a>
  <a href="README-IT.md"><img src="./assets/lang/it.svg" height="32" alt="Italiano"></a>
</p>

<!--
  dna: the official Domain Name API command line client and MCP server, built for AI agents and developers.
  Every animation on this page is a hand-made SVG in ./assets. No external widgets, no trackers.
-->

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/nl/hero.svg" width="1200" alt="dna, de Domain Name API CLI aangestuurd door AI-agents: een agent zoekt een domeinnaam op, bekijkt de factureerbare registratie vooraf met --dry-run, vraagt om goedkeuring en registreert het domein."></a>

<h1 align="center">dna: de Domain Name API CLI voor AI-agents</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="Nieuwste release"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b> bestuurt het resellerplatform van <a href="https://www.domainnameapi.com">Domain Name API</a> vanaf de opdrachtregel en vanuit AI-agents.
  Het dekt alle 64 API-operaties: domeinen zoeken, registratie, verlenging, verhuizing, DNS, DNSSEC, contacten, SSL en hosting.
  Koppel het aan <b>Claude Code, Claude Desktop, Cursor, Codex</b> of een andere MCP-client, en elke operatie wordt een veilige, goed beschreven tool.
</p>

<img src="./assets/nl/stats.svg" width="1200" alt="64 API-operaties, 65 MCP-tools voor AI-agents, 4 afgedwongen veiligheidsklassen, 6 platforms vanuit één binary.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> Gebouwd voor AI-agents

|  | Wat een agent krijgt |
|---|---|
| 🧰 **Eén tool per operatie** | `dna mcp serve` biedt 65 MCP-tools met JSON Schemas die uit de API-specificatie zijn opgebouwd, zodat modellen de invoer meteen de eerste keer correct invullen. |
| 🛡️ **Een veiligheidspoort** | Factureerbare en destructieve aanroepen worden nooit vanzelf uitgevoerd. De CLI vereist `--yes`, MCP vereist `confirm=true`, en zonder die bevestiging krijgt de agent een voorbeeld om aan jou te laten zien. |
| 🧾 **Alles machineleesbaar** | JSON-uitvoer bij piping, JSON-fouten met stabiele exitcodes, `--dry-run` voor het exacte verzoek en `dna schema` voor elke invoer. |
| 🧪 **Een sandbox** | `--env ote` wijst naar de testomgeving: echte API, geen echt geld. `--read-only` blokkeert elke wijziging. |
| 📘 **Zelfdocumenterend** | `dna agent guide` leert een agent de hele CLI in één aanroep. `dna agent skill --install` voegt een Claude Code-skill toe. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> Installatie

> [!IMPORTANT]
> **Installeer altijd vanaf de _nieuwste release_.** Release-binaries hebben een versienummer, zijn getest en worden gepubliceerd met SHA-256-checksums.

**Linux / macOS**

```bash
curl -fsSL https://raw.githubusercontent.com/domainreseller/cli/main/scripts/install.sh | sh
```

**Fedora / RHEL · Debian / Ubuntu**

```bash
# Fedora / RHEL
sudo dnf install https://github.com/domainreseller/cli/releases/latest/download/dna_linux_amd64.rpm
# Debian / Ubuntu
curl -fsSLO https://github.com/domainreseller/cli/releases/latest/download/dna_linux_amd64.deb && sudo apt install ./dna_linux_amd64.deb
```

**Windows (PowerShell)**

```powershell
irm https://raw.githubusercontent.com/domainreseller/cli/main/scripts/install.ps1 | iex
```

**Met Go 1.26+**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>Handmatige download</b></summary>
<br>

Kies je archief op de [nieuwste release](https://github.com/domainreseller/cli/releases/latest), controleer het met `checksums.txt`, pak het uit en zet `dna` (of `dna.exe`) in je `PATH`.

| OS | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

Controleer de installatie met `dna version` en `dna doctor`.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> Snel aan de slag

```bash
# 1. Log eerst in op de testomgeving (die heeft een eigen API-sleutel)
dna auth login --env ote

# 2. Kijk rond: niets hiervan wijzigt iets
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. Bekijk een aankoop vooraf en voer hem daarna uit (er wordt om bevestiging gevraagd)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. Productie
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> Typ gewoon `dna`: de interactieve shell

<img src="./assets/nl/shell.svg" width="1200" alt="De interactieve shell van dna: bij het typen van acme verschijnen de vrije extensies met prijzen, de gebruiker kiest acme.net, vult een kort registratieformulier in, bevestigt de betaalde stap in een bevestigingsvenster en het domein wordt geregistreerd.">

Start `dna` zonder argumenten in een terminal. Je hoeft geen enkel commando te kennen: alles kies je met ↑↓ en Enter, en Esc gaat altijd één stap terug.

|  | Wat je krijgt |
|---|---|
| 🔎 **Live zoeken** | Typ een naam zoals `acme`: vrije extensies en hun prijzen verschijnen terwijl je typt. Met ↓ kies je er een en registreer je die via een kort formulier. |
| ⚡ **Snelle acties** | ↓ op een lege regel: je domeinen, een domein registreren, DNS-records, saldo, wisselen tussen OTE en productie. |
| 🗂️ **Domeinmenu’s** | Zoek in je domeinen en druk op Enter: details, DNS, verlengen, vergrendeling, WHOIS-privacy, nameservers, autorisatiecode. |
| 🛡️ **Niets per ongeluk** | Elke betaalde of destructieve stap eindigt in een bevestigingsvenster met het exacte verzoek en de prijs. Annuleren blijft geselecteerd tot je iets anders kiest. |
| 🤖 **Een AI-agent koppelen** | `/agent` toont meer dan 20 agents: Claude Code, ChatGPT, Cursor, Claude Desktop, Codex en meer. dna richt hem in, controleert de inloggegevens en geeft je kant-en-klare prompts. |

Scripts en AI-agents zien de shell nooit: via een pipe of gestart door een agent geeft `dna` net als voorheen JSON en de help.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> Verbind je AI-agent

<img src="./assets/nl/mcp.svg" width="1200" alt="Model Context Protocol: dna mcp serve biedt 65 tools met veiligheidsannotaties. Een factureerbare tool-aanroep zonder confirm geeft een voorbeeld terug; met confirm=true na goedkeuring wordt hij uitgevoerd.">

> [!TIP]
> **De snelste weg:** start `dna`, druk op ↓ en kies **Connect an AI agent** (of typ `/agent`). Kies een van de meer dan 20 agents, OTE of productie en hoeveel hij mag doen. dna voert de installatie uit of voegt zichzelf toe aan het configuratiebestand van de agent, controleert de inloggegevens en kopieert kant-en-klare prompts voor je.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # snippets voor elke andere client
```

<details>
<summary><b>Claude Desktop · Cursor · Windsurf · Gemini CLI</b></summary>
<br>

```json
{
  "mcpServers": {
    "dna": {
      "command": "dna",
      "args": ["mcp", "serve", "--env", "ote"]
    }
  }
}
```

Claude Desktop: `claude_desktop_config.json` · Cursor: `~/.cursor/mcp.json` · Windsurf: `~/.codeium/windsurf/mcp_config.json` · Gemini CLI: `~/.gemini/settings.json`
</details>

<details>
<summary><b>VS Code</b></summary>
<br>

`.vscode/mcp.json`:

```json
{
  "servers": {
    "dna": { "type": "stdio", "command": "dna", "args": ["mcp", "serve", "--env", "ote"] }
  }
}
```
</details>

<details>
<summary><b>OpenAI Codex CLI</b></summary>
<br>

`~/.codex/config.toml`:

```toml
[mcp_servers.dna]
command = "dna"
args = ["mcp", "serve", "--env", "ote"]
```
</details>

<details>
<summary><b>Agents met een shell (skill en guide)</b></summary>
<br>

```bash
dna agent skill --install                       # Claude Code-skill: ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # projectskill, commit deze mee in je repo
dna agent guide                                 # compacte gebruiksgids voor elke agent
```
</details>

Jij bepaalt hoe ver de agent mag gaan:

```bash
dna mcp serve --read-only            # kijken, niet aankomen
dna mcp serve --no-billable          # alles behalve aankopen
dna mcp serve --groups domain,dns    # alleen deze commandogroepen
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> Kant-en-klare prompts voor Claude en ChatGPT

Kopieer een prompt, plak die in uw AI-assistent en laat de assistent het werk doen. Elke prompt begint in de OTE-testomgeving en wacht op uw goedkeuring voordat er iets betaalds gebeurt.

**dna instellen in Claude Code**

```text
Zet de dna CLI voor me op (https://github.com/domainreseller/cli).
1. Installeer hem met het officiële installatiescript voor mijn besturingssysteem en
   voer daarna `dna version` en `dna doctor` uit.
2. Installeer de skill voor Claude Code met `dna agent skill --install` en voeg de
   MCP-server toe met `claude mcp add dna -- dna mcp serve --env ote`.
3. Vraag me niet om mijn API-key in deze chat te plakken. Zeg me dat ik
   `dna auth login --env ote` zelf moet uitvoeren en controleer daarna de verbinding
   met `dna --ote account get`.
Werk alleen in de OTE-testomgeving en gebruik nooit `--yes` zonder mijn goedkeuring.
```

**dna instellen in ChatGPT Codex**

```text
Zet de dna CLI voor me op (https://github.com/domainreseller/cli).
1. Installeer hem met het officiële installatiescript voor mijn besturingssysteem en
   voer daarna `dna version` en `dna doctor` uit.
2. Voeg hem toe aan mijn Codex-configuratie (`~/.codex/config.toml`) als MCP-server
   met de naam `dna`, het commando `dna` en de argumenten `mcp serve --env ote`.
3. Vraag me niet om mijn API-key in deze chat te plakken. Zeg me dat ik
   `dna auth login --env ote` zelf moet uitvoeren en controleer daarna de verbinding
   met `dna --ote account get`.
Werk alleen in de OTE-testomgeving en gebruik nooit `--yes` zonder mijn goedkeuring.
```

**Geen terminal? Vraag het aan ChatGPT of Claude in de chat**

```text
Lees https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt en de
dna-documentatie waarnaar daar wordt gelinkt. Geef me daarna stap voor stap de exacte
dna-commando's om een domeinnaam te registreren in de OTE-testomgeving: installeren,
inloggen, zoeken, als voorbeeld tonen met `--dry-run` en registreren. Leg uit wat elk
commando doet.
```

<details>
<summary><b>Meer prompts voor dagelijks werk</b></summary>
<br>

**Domeinnaam vinden en registreren**

```text
Controleer met dna of brandname.com, brandname.net en brandname.io beschikbaar zijn.
Zet de resultaten met prijzen in een tabel en laat daarna met `--dry-run` een
voorbeeld zien van de registratie van de goedkoopste beschikbare naam. Registreer
niets voordat ik ja zeg.
```

**Aankomende verlengingen**

```text
Geef met dna een lijst van mijn domeinnamen die in de komende 60 dagen verlopen, met
hun verlengingsprijzen, gesorteerd op vervaldatum. Verleng niets.
```

**Domeinnaam naar een server laten verwijzen**

```text
Laat example.com met dna naar 203.0.113.10 verwijzen: voeg A-records toe voor `@` en
`www` en laat me daarna de uiteindelijke DNS-zone zien. Vraag het me voordat je een
bestaand record wijzigt of verwijdert.
```

**Domeinnaam naar uw account verhuizen**

```text
Controleer met dna of example.org naar mijn account verhuisd kan worden, vertel me wat
ik nog nodig heb (authcode, ontgrendeling) en laat met `--dry-run` een voorbeeld van
de verhuizing zien.
```

**SSL-certificaat bestellen**

```text
Zoek met dna het goedkoopste DV-SSL-certificaat voor example.com, maak een CSR aan en
laat me een voorbeeld van de bestelling zien. Wacht op mijn goedkeuring voordat je het
koopt.
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> Hoe het werkt

<img src="./assets/nl/flow.svg" width="1200" alt="Een AI-agent of jijzelf roept dna aan via MCP of de CLI; dna valideert de invoer, handhaaft de veiligheidspoort en communiceert met de Domain Name API (productie of de OTE-sandbox), die de registries via EPP bereikt.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> Veiligheidspoort

<img src="./assets/nl/safety.svg" width="1200" alt="Lees- en schrijfopdrachten worden direct uitgevoerd; destructieve en factureerbare opdrachten stoppen totdat ze zijn bevestigd met --yes in de CLI of confirm=true via MCP.">

| klasse | voorbeelden | gedrag |
|---|---|---|
| `read` | zoeken, lijsten, info, prijzen | wordt direct uitgevoerd |
| `write` | DNS-records, vergrendeling, privacy, contacten | wordt direct uitgevoerd (gratis, omkeerbaar) |
| `destructive` | een record verwijderen, SSL intrekken, DNSSEC uitschakelen | stopt tot er bevestiging is |
| `billable` | registreren, verlengen, verhuizen, herstellen, SSL/hosting kopen | stopt tot er bevestiging is |

- **In een script of agent-shell:** zonder `--yes` stopt een opdracht achter de veiligheidspoort met exitcode **9** en wordt er niets verstuurd.
- **Via MCP:** zonder `confirm=true` geeft de tool een voorbeeld van het exacte verzoek terug.
- **Vooraf bekijken:** `--dry-run` toont het HTTP-verzoek met de API-sleutel gemaskeerd.
- **Alles op slot:** `--read-only`, `DNA_READ_ONLY=1` of een alleen-lezenprofiel blokkeert elke wijziging (exitcode **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> Uitvoer, fouten en exitcodes

```bash
dna domain info example.com                   # tabel in een terminal
dna domain info example.com -o json           # JSON (de standaard bij piping)
dna domain list --all -o jsonl                # één domein per regel
dna domain info example.com --jq .authCode    # ingebouwde jq
dna dns add --body @record.json               # volledige invoer uit een bestand, of --body - voor stdin
dna schema domain.register                    # JSON Schema van elke invoer
dna enums DomainStatus                        # wat status 7 betekent
```

Fouten worden als JSON naar stderr geschreven: `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| code | betekenis | code | betekenis |
|---:|---|---:|---|
| 0 | gelukt | 6 | rate limit bereikt |
| 1 | onverwachte fout | 7 | serverfout |
| 2 | ongeldige invoer, niets verstuurd | 8 | netwerk |
| 3 | authenticatie | 9 | bevestiging vereist |
| 4 | niet gevonden | 10 | alleen-lezenmodus |
| 5 | geweigerd door de API | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> Opdrachten

| groep | wat het doet |
|---|---|
| `dna` | de interactieve shell: live zoeken, menu’s, begeleide formulieren, AI-agents koppelen |
| `dna domain` | zoeken, bulkzoeken, lijst, info, registreren, verlengen, verhuizen, herstellen, vergrendelen, privacy, authcode, contacten, documenten, doorsturen, nameservers, child hosts |
| `dna dns` · `dna dnssec` | zonerecords, DS-records |
| `dna contact` | contacthandles en verificatie |
| `dna ssl` | CSR, kopen, configureren, status, downloaden, opnieuw uitgeven, intrekken |
| `dna hosting` · `dna product` · `dna order` · `dna account` | hosting, catalogus en prijzen, bestellingen, saldo en transacties |
| `dna api` | ruw verzoek naar elk endpoint |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | voor agents en tools |
| `dna auth` · `dna config` · `dna doctor` | inloggegevens, profielen, diagnose |

De `--help` van elke opdracht toont de veiligheidsklasse, de API-route en voorbeelden. Zie de volledige [opdrachtreferentie](docs/reference/README.md).

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> API-inloggegevens

`dna` heeft dezelfde twee waarden nodig als elke Domain Name API-module:

| waarde | wat het is |
|---|---|
| **Reseller ID** | een UUID zoals `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` |
| **API Key** | je API-sleutel; productie en OTE hebben **verschillende** sleutels |

> [!TIP]
> Je vindt beide in je Domain Name API-paneel onder **API Settings**. Sla ze op met `dna auth login`, of gebruik in CI de omgevingsvariabelen `DNA_RESELLER_ID` en `DNA_API_KEY`.

> [!CAUTION]
> Dit zijn **API-inloggegevens**. Het e-mailadres en wachtwoord waarmee je inlogt op het paneel werken hier **niet**.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> Veelgestelde vragen

<details>
<summary><b>Is het veilig om een AI-agent dna te laten gebruiken?</b></summary>
<br>
Factureerbare en destructieve operaties zitten achter een veiligheidspoort: de CLI weigert ze zonder <code>--yes</code>, en de MCP-server geeft een voorbeeld terug tenzij de aanroep <code>confirm=true</code> bevat, wat agents volgens hun instructies pas mogen instellen nadat jij toestemming hebt gegeven. Laat agents beginnen op <code>--env ote</code>, en gebruik <code>--read-only</code> of <code>--no-billable</code> om ze verder te beperken.
</details>

<details>
<summary><b>Kan ik het uitproberen zonder geld uit te geven?</b></summary>
<br>
Ja. <code>--env ote</code> gebruikt de OTE-testomgeving met een eigen API-sleutel en een eigen saldo. Niets daar raakt live domeinen of je echte saldo.
</details>

<details>
<summary><b>Welke AI-tools werken ermee?</b></summary>
<br>
Elke MCP-client: Claude Code, Claude Desktop, Cursor, VS Code, Windsurf, OpenAI Codex CLI, Gemini CLI en andere. Agents die alleen een shell hebben, kunnen de CLI rechtstreeks aanroepen; <code>dna agent guide</code> en de Claude Code-skill leggen hun uit hoe dat werkt.
</details>

<details>
<summary><b>Waar worden mijn inloggegevens opgeslagen?</b></summary>
<br>
<code>dna auth login</code> bewaart ze in een profielbestand in je gebruikersconfiguratiemap (<code>dna config path</code> laat zien waar), op Linux en macOS weggeschreven met rechten die alleen de eigenaar toegang geven. Omgevingsvariabelen gaan altijd voor op het bestand, en de API-sleutel wordt in alle uitvoer gemaskeerd.
</details>

<details>
<summary><b>Heb ik Go nodig?</b></summary>
<br>
Nee. De installatiescripts en de release-archieven bevatten één enkele binary. Go 1.26+ is alleen nodig voor <code>go install</code> of om vanuit de broncode te bouwen.
</details>

<details>
<summary><b>Is er een bibliotheek voor mijn eigen code?</b></summary>
<br>
<code>dna</code> is een CLI en MCP-server. Om de API vanuit code aan te roepen, gebruik je de officiële bibliotheken voor <a href="https://github.com/domainreseller/php-dna">PHP</a>, <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>, <a href="https://github.com/domainreseller/python-dna">Python</a>, <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> of <a href="https://github.com/domainreseller/go-dna">Go</a>.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> Documentatie

| handleiding | inhoud |
|---|---|
| [Aan de slag](docs/getting-started.md) | installatie, OTE-login, eerste registratie en DNS-wijziging |
| [Configuratie](docs/configuration.md) | inloggegevens, omgevingen, profielen, omgevingsvariabelen, alleen-lezenmodus |
| [Gebruik](docs/usage.md) | invoer, `--body`, uitvoerformaten, `--jq`, paginering, rate limits, recepten |
| [AI-agents en MCP](docs/ai-agents.md) | MCP-configuratie per client, bevestigingsprotocol, skill, agent guide |
| [Fouten en exitcodes](docs/errors.md) | fout-JSON, foutcodes, probleemoplossing |
| [Opdrachtreferentie](docs/reference/README.md) · [Enums](docs/enums.md) | elke opdracht, flag en gecodeerde waarde (gegenereerd) |
| [Architectuur](docs/architecture.md) · [Bijdragen](CONTRIBUTING.md) | hoe de code is georganiseerd |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> Contact

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API: resellerprogramma voor domeinnamen"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="Word domeinreseller"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="API en modules voor domeinresellers"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="Ondersteuning"></a>
</p>

<p align="right">
  <sub>Volg ons</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · gemaakt voor mensen en de agents die met hen samenwerken</sub></p>
