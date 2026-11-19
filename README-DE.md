<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de-active.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa.svg" height="32" alt="العربية"></a>
  <a href="README-NL.md"><img src="./assets/lang/nl.svg" height="32" alt="Nederlands"></a>
  <a href="README-AZ.md"><img src="./assets/lang/az.svg" height="32" alt="Azərbaycanca"></a>
  <a href="README-CN.md"><img src="./assets/lang/cn.svg" height="32" alt="中文"></a>
  <a href="README-FR.md"><img src="./assets/lang/fr.svg" height="32" alt="Français"></a>
  <a href="README-IT.md"><img src="./assets/lang/it.svg" height="32" alt="Italiano"></a>
</p>

<!--
  dna: the official Domain Name API command line client and MCP server, built for AI agents and developers.
  Every animation on this page is a hand-made SVG in ./assets. No external widgets, no trackers.
-->

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/de/hero.svg" width="1200" alt="dna, die Domain Name API CLI, gesteuert von KI-Agenten: Ein Agent sucht eine Domain, prüft die kostenpflichtige Registrierung vorab mit --dry-run, bittet um Freigabe und registriert sie."></a>

<h1 align="center">dna: die Domain Name API CLI für KI-Agenten</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="Neueste Version"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b> steuert die Reseller-Plattform von <a href="https://www.domainnameapi.com">Domain Name API</a> über die Kommandozeile und über KI-Agenten.
  Es deckt alle 64 API-Operationen ab: Domainsuche, Registrierung, Verlängerung, Transfer, DNS, DNSSEC, Kontakte, SSL und Hosting.
  Binden Sie es in <b>Claude Code, Claude Desktop, Cursor, Codex</b> oder einen beliebigen MCP-Client ein, und jede Operation wird zu einem sicheren, gut beschriebenen Tool.
</p>

<img src="./assets/de/stats.svg" width="1200" alt="64 API-Operationen, 65 MCP-Tools für KI-Agenten, 4 durchgesetzte Sicherheitsklassen, 6 Plattformen aus einer einzigen Binärdatei.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> Für KI-Agenten entwickelt

|  | Was ein Agent erhält |
|---|---|
| 🧰 **Ein Tool pro Operation** | `dna mcp serve` stellt 65 MCP-Tools mit JSON-Schemas bereit, die aus der API-Spezifikation erzeugt werden, sodass Modelle die Eingaben auf Anhieb korrekt ausfüllen. |
| 🛡️ **Eine Sicherheitsschranke** | Kostenpflichtige und destruktive Aufrufe werden nie eigenständig ausgeführt. Die CLI benötigt `--yes`, MCP benötigt `confirm=true`, und ohne diese Angabe erhält der Agent eine Vorschau, die er Ihnen zeigen kann. |
| 🧾 **Alles maschinenlesbar** | JSON-Ausgabe bei Weiterleitung per Pipe, JSON-Fehler mit stabilen Exit-Codes, `--dry-run` für die exakte Anfrage und `dna schema` für jede Eingabe. |
| 🧪 **Eine Sandbox** | `--env ote` verweist auf die Testumgebung: echte API, kein echtes Geld. `--read-only` blockiert jede Änderung. |
| 📘 **Selbstdokumentierend** | `dna agent guide` vermittelt einem Agenten die gesamte CLI mit einem einzigen Aufruf. `dna agent skill --install` fügt einen Skill für Claude Code hinzu. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> Installation

> [!IMPORTANT]
> **Installieren Sie immer aus dem _neuesten Release_.** Release-Binärdateien sind versioniert, getestet und werden mit SHA-256-Prüfsummen veröffentlicht.

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

**Mit Go 1.26+**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>Manueller Download</b></summary>
<br>

Wählen Sie Ihr Archiv im [neuesten Release](https://github.com/domainreseller/cli/releases/latest) aus, prüfen Sie es anhand von `checksums.txt`, entpacken Sie es und legen Sie `dna` (bzw. `dna.exe`) in Ihren `PATH`.

| Betriebssystem | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

Überprüfen Sie die Installation mit `dna version` und `dna doctor`.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> Schnellstart

```bash
# 1. Zuerst in der Testumgebung anmelden (sie hat einen eigenen API-Key)
dna auth login --env ote

# 2. Umsehen: Nichts davon ändert etwas
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. Einen Kauf in der Vorschau prüfen und dann ausführen (mit Bestätigungsabfrage)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. Produktivumgebung
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> Einfach `dna` eingeben: die interaktive Shell

<img src="./assets/de/shell.svg" width="1200" alt="Die interaktive Shell von dna: Bei der Eingabe von acme erscheinen die freien Endungen mit Preisen, der Nutzer wählt acme.net, füllt ein kurzes Registrierungsformular aus, bestätigt den kostenpflichtigen Schritt im Bestätigungsdialog und die Domain wird registriert.">

Starten Sie `dna` ohne Argumente im Terminal. Sie müssen keinen einzigen Befehl kennen: Alles wird mit ↑↓ und Enter ausgewählt, Esc geht immer einen Schritt zurück.

|  | Was Sie bekommen |
|---|---|
| 🔎 **Live-Suche** | Tippen Sie einen Namen wie `acme`: freie Endungen und ihre Preise erscheinen schon beim Tippen. ↓ wählt eine aus und registriert sie über ein kurzes Formular. |
| ⚡ **Schnellaktionen** | ↓ in einer leeren Zeile: Ihre Domains, Domain registrieren, DNS-Einträge, Guthaben, Wechsel zwischen OTE und Produktion. |
| 🗂️ **Domain-Menüs** | Domains durchsuchen und Enter drücken: Details, DNS, Verlängern, Sperre, WHOIS-Datenschutz, Nameserver, Auth-Code. |
| 🛡️ **Nichts aus Versehen** | Jeder kostenpflichtige oder destruktive Schritt endet in einem Bestätigungsdialog mit der genauen Anfrage und dem Preis. Abbrechen bleibt ausgewählt, bis Sie sich anders entscheiden. |
| 🤖 **KI-Agent verbinden** | `/agent` listet über 20 Agenten: Claude Code, ChatGPT, Cursor, Claude Desktop, Codex und mehr. dna richtet ihn ein, prüft die Zugangsdaten und gibt Ihnen fertige Prompts. |

Skripte und KI-Agenten sehen die Shell nie: Über eine Pipe oder von einem Agenten gestartet gibt `dna` wie bisher JSON und die Hilfe aus.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> KI-Agenten verbinden

<img src="./assets/de/mcp.svg" width="1200" alt="Model Context Protocol: dna mcp serve stellt 65 Tools mit Sicherheitsannotationen bereit. Ein kostenpflichtiger Tool-Aufruf ohne confirm liefert eine Vorschau; mit confirm=true nach der Freigabe wird er ausgeführt.">

> [!TIP]
> **Am schnellsten:** Starten Sie `dna`, drücken Sie ↓ und wählen Sie **Connect an AI agent** (oder tippen Sie `/agent`). Wählen Sie einen von über 20 Agenten, OTE oder Produktion und wie viel er tun darf. dna führt die Einrichtung aus oder trägt sich in die Konfigurationsdatei des Agenten ein, prüft die Zugangsdaten und kopiert fertige Prompts für Sie.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # Snippets für alle anderen Clients
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
<summary><b>Agenten mit Shell (Skill und Guide)</b></summary>
<br>

```bash
dna agent skill --install                       # Skill für Claude Code: ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # Projekt-Skill, zusammen mit Ihrem Repository committen
dna agent guide                                 # kompakte Nutzungsanleitung für jeden Agenten
```
</details>

Sie legen fest, wie weit der Zugriff des Agenten reicht:

```bash
dna mcp serve --read-only            # nur ansehen, nichts anfassen
dna mcp serve --no-billable          # alles außer Käufen
dna mcp serve --groups domain,dns    # nur diese Befehlsgruppen
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> Fertige Prompts für Claude und ChatGPT

Kopieren Sie einen Prompt, fügen Sie ihn in Ihren KI-Assistenten ein und lassen Sie ihn die Arbeit erledigen. Jeder Prompt beginnt in der OTE-Testumgebung und hält vor jeder kostenpflichtigen Aktion an, bis Sie sie freigeben.

**dna in Claude Code einrichten**

```text
Richte die dna CLI für mich ein (https://github.com/domainreseller/cli).
1. Installiere sie mit dem offiziellen Installationsskript für mein Betriebssystem und
   führe dann `dna version` und `dna doctor` aus.
2. Installiere den Skill für Claude Code mit `dna agent skill --install` und füge den
   MCP-Server mit `claude mcp add dna -- dna mcp serve --env ote` hinzu.
3. Fordere mich nicht auf, meinen API-Key in diesen Chat einzufügen. Sag mir, dass ich
   `dna auth login --env ote` selbst ausführen soll, und prüfe dann die Verbindung mit
   `dna --ote account get`.
Arbeite nur in der OTE-Testumgebung und verwende `--yes` niemals ohne meine Freigabe.
```

**dna in ChatGPT Codex einrichten**

```text
Richte die dna CLI für mich ein (https://github.com/domainreseller/cli).
1. Installiere sie mit dem offiziellen Installationsskript für mein Betriebssystem und
   führe dann `dna version` und `dna doctor` aus.
2. Trage sie in meiner Codex-Konfiguration (`~/.codex/config.toml`) als MCP-Server mit
   dem Namen `dna`, dem Befehl `dna` und den Argumenten `mcp serve --env ote` ein.
3. Fordere mich nicht auf, meinen API-Key in diesen Chat einzufügen. Sag mir, dass ich
   `dna auth login --env ote` selbst ausführen soll, und prüfe dann die Verbindung mit
   `dna --ote account get`.
Arbeite nur in der OTE-Testumgebung und verwende `--yes` niemals ohne meine Freigabe.
```

**Kein Terminal? Fragen Sie ChatGPT oder Claude im Chat**

```text
Lies https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt und die dort
verlinkte dna-Dokumentation. Gib mir dann Schritt für Schritt die genauen dna-Befehle,
um eine Domain in der OTE-Testumgebung zu registrieren: installieren, anmelden,
suchen, mit `--dry-run` als Vorschau anzeigen und registrieren. Erkläre, was jeder
Befehl macht.
```

<details>
<summary><b>Weitere Prompts für die tägliche Arbeit</b></summary>
<br>

**Domain finden und registrieren**

```text
Prüfe mit dna, ob brandname.com, brandname.net und brandname.io verfügbar sind. Zeig
die Ergebnisse mit Preisen in einer Tabelle und zeig dann mit `--dry-run` eine
Vorschau der Registrierung für die günstigste verfügbare Domain. Registriere nichts,
bevor ich Ja sage.
```

**Anstehende Verlängerungen**

```text
Liste mit dna meine Domains auf, die in den nächsten 60 Tagen ablaufen, mit ihren
Verlängerungspreisen und sortiert nach Ablaufdatum. Verlängere nichts.
```

**Domain auf einen Server zeigen lassen**

```text
Lass example.com mit dna auf 203.0.113.10 zeigen: Lege A-Records für `@` und `www` an
und zeig mir anschließend die finale DNS-Zone. Frag mich, bevor du einen bestehenden
Eintrag änderst oder löschst.
```

**Domain in Ihr Konto transferieren**

```text
Prüfe mit dna, ob example.org in mein Konto transferiert werden kann, sag mir, was ich
noch brauche (Auth-Code, Entsperrung), und zeig mit `--dry-run` eine Vorschau des
Transfers.
```

**SSL-Zertifikat bestellen**

```text
Finde mit dna das günstigste DV-SSL-Zertifikat für example.com, erstelle einen CSR und
zeig mir eine Vorschau der Bestellung. Warte vor dem Kauf auf meine Freigabe.
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> So funktioniert es

<img src="./assets/de/flow.svg" width="1200" alt="Ein KI-Agent oder Sie selbst rufen dna über MCP oder die CLI auf; dna validiert die Eingabe, setzt die Sicherheitsschranke durch und kommuniziert mit der Domain Name API (Produktion oder OTE-Sandbox), die die Registries über EPP erreicht.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> Sicherheitsschranke

<img src="./assets/de/safety.svg" width="1200" alt="Lese- und Schreibbefehle werden sofort ausgeführt; destruktive und kostenpflichtige Befehle halten an, bis sie in der CLI mit --yes oder über MCP mit confirm=true bestätigt werden.">

| Klasse | Beispiele | Verhalten |
|---|---|---|
| `read` | Suche, Liste, Info, Preise | wird sofort ausgeführt |
| `write` | DNS-Einträge, Sperre, Datenschutz, Kontakte | wird sofort ausgeführt (kostenlos, umkehrbar) |
| `destructive` | Eintrag löschen, SSL widerrufen, DNSSEC deaktivieren | hält an, bis bestätigt |
| `billable` | Registrieren, Verlängern, Transferieren, Wiederherstellen, SSL/Hosting kaufen | hält an, bis bestätigt |

- **In einem Skript oder einer Agenten-Shell:** Ohne `--yes` beendet sich ein geschützter Befehl mit Code **9** und sendet nichts.
- **Über MCP:** Ohne `confirm=true` liefert das Tool eine Vorschau der exakten Anfrage.
- **Vorschau:** `--dry-run` gibt die HTTP-Anfrage mit maskiertem API-Key aus.
- **Absichern:** `--read-only`, `DNA_READ_ONLY=1` oder ein schreibgeschütztes Profil blockiert jede Änderung (Exit-Code **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> Ausgabe, Fehler und Exit-Codes

```bash
dna domain info example.com                   # Tabelle im Terminal
dna domain info example.com -o json           # JSON (Standard bei Weiterleitung per Pipe)
dna domain list --all -o jsonl                # eine Domain pro Zeile
dna domain info example.com --jq .authCode    # integriertes jq
dna dns add --body @record.json               # vollständige Eingabe aus einer Datei, oder --body - für stdin
dna schema domain.register                    # JSON-Schema einer beliebigen Eingabe
dna enums DomainStatus                        # was Status 7 bedeutet
```

Fehler werden als JSON auf stderr ausgegeben: `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| Code | Bedeutung | Code | Bedeutung |
|---:|---|---:|---|
| 0 | Erfolg | 6 | Rate-Limit erreicht |
| 1 | unerwarteter Fehler | 7 | Serverfehler |
| 2 | ungültige Eingabe, nichts gesendet | 8 | Netzwerk |
| 3 | Authentifizierung | 9 | Bestätigung erforderlich |
| 4 | nicht gefunden | 10 | Nur-Lese-Modus |
| 5 | von der API abgelehnt | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> Befehle

| Gruppe | Funktion |
|---|---|
| `dna` | die interaktive Shell: Live-Suche, Menüs, geführte Formulare, Einrichtung von KI-Agenten |
| `dna domain` | Suche, Massensuche, Liste, Info, Registrierung, Verlängerung, Transfer, Wiederherstellung, Sperre, Datenschutz, Auth-Code, Kontakte, Dokumente, Weiterleitung, Nameserver, Child-Nameserver |
| `dna dns` · `dna dnssec` | Zoneneinträge, DS-Einträge |
| `dna contact` | Kontakt-Handles und Verifizierung |
| `dna ssl` | CSR, Kauf, Konfiguration, Status, Download, Neuausstellung, Widerruf |
| `dna hosting` · `dna product` · `dna order` · `dna account` | Hosting, Katalog und Preise, Bestellungen, Guthaben und Transaktionen |
| `dna api` | direkte Anfrage an einen beliebigen Endpunkt |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | für Agenten und Tools |
| `dna auth` · `dna config` · `dna doctor` | Zugangsdaten, Profile, Diagnose |

Die `--help`-Ausgabe jedes Befehls zeigt dessen Sicherheitsklasse, API-Route und Beispiele. Siehe die vollständige [Befehlsreferenz](docs/reference/README.md).

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> API-Zugangsdaten

`dna` benötigt dieselben zwei Werte wie jedes Modul von Domain Name API:

| Wert | Bedeutung |
|---|---|
| **Reseller-ID** | eine UUID wie `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` |
| **API-Key** | Ihr API-Key; Produktion und OTE haben **unterschiedliche** Keys |

> [!TIP]
> Beide Werte finden Sie in Ihrem Panel von Domain Name API unter **API-Einstellungen**. Speichern Sie sie mit `dna auth login` oder verwenden Sie in CI die Umgebungsvariablen `DNA_RESELLER_ID` und `DNA_API_KEY`.

> [!CAUTION]
> Dies sind **API-Zugangsdaten**. Ihre E-Mail-Adresse und Ihr Passwort für die Anmeldung im Panel funktionieren hier **nicht**.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> FAQ

<details>
<summary><b>Ist es sicher, einen KI-Agenten dna verwenden zu lassen?</b></summary>
<br>
Kostenpflichtige und destruktive Operationen sind geschützt: Die CLI verweigert sie ohne <code>--yes</code>, und der MCP-Server liefert eine Vorschau, sofern der Aufruf nicht <code>confirm=true</code> enthält, was Agenten laut ihren Anweisungen erst nach Ihrer Freigabe setzen dürfen. Starten Sie Agenten mit <code>--env ote</code> und verwenden Sie <code>--read-only</code> oder <code>--no-billable</code>, um sie weiter einzuschränken.
</details>

<details>
<summary><b>Kann ich es ausprobieren, ohne Geld auszugeben?</b></summary>
<br>
Ja. <code>--env ote</code> nutzt die OTE-Testumgebung mit eigenem API-Key und eigenem Guthaben. Nichts dort berührt Live-Domains oder Ihr echtes Guthaben.
</details>

<details>
<summary><b>Welche KI-Tools funktionieren damit?</b></summary>
<br>
Jeder MCP-Client: Claude Code, Claude Desktop, Cursor, VS Code, Windsurf, OpenAI Codex CLI, Gemini CLI und weitere. Agenten, die nur eine Shell haben, können die CLI direkt aufrufen; <code>dna agent guide</code> und der Skill für Claude Code erklären sie ihnen.
</details>

<details>
<summary><b>Wo werden meine Zugangsdaten gespeichert?</b></summary>
<br>
<code>dna auth login</code> speichert sie in einer Profildatei in Ihrem Benutzer-Konfigurationsordner (<code>dna config path</code> zeigt den Speicherort), unter Linux und macOS mit Zugriffsrechten nur für den Eigentümer. Umgebungsvariablen haben stets Vorrang vor der Datei, und der API-Key wird in jeder Ausgabe maskiert.
</details>

<details>
<summary><b>Benötige ich Go?</b></summary>
<br>
Nein. Die Installationsskripte und die Release-Archive enthalten eine einzige Binärdatei. Go 1.26+ wird nur für <code>go install</code> oder zum Bauen aus dem Quellcode benötigt.
</details>

<details>
<summary><b>Gibt es eine Bibliothek für meinen eigenen Code?</b></summary>
<br>
<code>dna</code> ist eine CLI und ein MCP-Server. Um die API aus Ihrem Code aufzurufen, verwenden Sie die offiziellen Bibliotheken für <a href="https://github.com/domainreseller/php-dna">PHP</a>, <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>, <a href="https://github.com/domainreseller/python-dna">Python</a>, <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> oder <a href="https://github.com/domainreseller/go-dna">Go</a>.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> Dokumentation

| Anleitung | Inhalt |
|---|---|
| [Erste Schritte](docs/getting-started.md) | Installation, OTE-Anmeldung, erste Registrierung und erste DNS-Änderung |
| [Konfiguration](docs/configuration.md) | Zugangsdaten, Umgebungen, Profile, Umgebungsvariablen, Nur-Lese-Modus |
| [Verwendung](docs/usage.md) | Eingaben, `--body`, Ausgabeformate, `--jq`, Paginierung, Rate-Limits, Rezepte |
| [KI-Agenten und MCP](docs/ai-agents.md) | MCP-Einrichtung je Client, Bestätigungsprotokoll, Skill, Agent Guide |
| [Fehler und Exit-Codes](docs/errors.md) | Fehler-JSON, Fehlercodes, Fehlerbehebung |
| [Befehlsreferenz](docs/reference/README.md) · [Enums](docs/enums.md) | alle Befehle, Flags und codierten Werte (generiert) |
| [Architektur](docs/architecture.md) · [Mitwirken](CONTRIBUTING.md) | wie der Code organisiert ist |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> Kontakt

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API: Domain-Reseller-Programm"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="Domain-Reseller werden"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="Domain-Reseller-API & Module"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="Support"></a>
</p>

<p align="right">
  <sub>Folgen Sie uns</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · gemacht für Menschen und die Agenten, die mit ihnen arbeiten</sub></p>
