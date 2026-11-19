<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa.svg" height="32" alt="العربية"></a>
  <a href="README-NL.md"><img src="./assets/lang/nl.svg" height="32" alt="Nederlands"></a>
  <a href="README-AZ.md"><img src="./assets/lang/az.svg" height="32" alt="Azərbaycanca"></a>
  <a href="README-CN.md"><img src="./assets/lang/cn.svg" height="32" alt="中文"></a>
  <a href="README-FR.md"><img src="./assets/lang/fr.svg" height="32" alt="Français"></a>
  <a href="README-IT.md"><img src="./assets/lang/it-active.svg" height="32" alt="Italiano"></a>
</p>

<!--
  dna: the official Domain Name API command line client and MCP server, built for AI agents and developers.
  Every animation on this page is a hand-made SVG in ./assets. No external widgets, no trackers.
-->

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/it/hero.svg" width="1200" alt="dna, la CLI di Domain Name API guidata dagli agenti AI: un agente cerca un dominio, visualizza in anteprima la registrazione a pagamento con --dry-run, chiede l'approvazione e lo registra."></a>

<h1 align="center">dna: la CLI di Domain Name API per agenti AI</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="Ultima versione"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b> gestisce la piattaforma reseller di <a href="https://www.domainnameapi.com">Domain Name API</a> dalla riga di comando e tramite agenti AI.
  Copre tutte le 64 operazioni API: ricerca, registrazione, rinnovo e trasferimento di domini, DNS, DNSSEC, contatti, SSL e hosting.
  Collegalo a <b>Claude Code, Claude Desktop, Cursor, Codex</b> o a qualsiasi client MCP e ogni operazione diventa uno strumento sicuro e ben descritto.
</p>

<img src="./assets/it/stats.svg" width="1200" alt="64 operazioni API, 65 strumenti MCP per agenti AI, 4 classi di sicurezza applicate, 6 piattaforme da un unico binario.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> Pensato per gli agenti AI

|  | Cosa ottiene un agente |
|---|---|
| 🧰 **Uno strumento per ogni operazione** | `dna mcp serve` espone 65 strumenti MCP con JSON Schema generati dalla specifica dell'API, così i modelli compilano correttamente gli input al primo tentativo. |
| 🛡️ **Un controllo di sicurezza** | Le chiamate a pagamento e distruttive non vengono mai eseguite da sole. La CLI richiede `--yes`, MCP richiede `confirm=true` e, in loro assenza, l'agente riceve un'anteprima da mostrarti. |
| 🧾 **Tutto leggibile dalle macchine** | Output JSON quando l'output passa in una pipe, errori JSON con codici di uscita stabili, `--dry-run` per vedere la richiesta esatta e `dna schema` per ogni input. |
| 🧪 **Una sandbox** | `--env ote` punta all'ambiente di test: API reale, nessun denaro reale. `--read-only` blocca qualsiasi modifica. |
| 📘 **Autodocumentato** | `dna agent guide` insegna a un agente l'intera CLI con una sola chiamata. `dna agent skill --install` aggiunge una skill per Claude Code. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> Installazione

> [!IMPORTANT]
> **Installa sempre dalla _release più recente_.** I binari delle release sono versionati, testati e pubblicati con checksum SHA-256.

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

**Con Go 1.26+**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>Download manuale</b></summary>
<br>

Scegli il tuo archivio nella [release più recente](https://github.com/domainreseller/cli/releases/latest), verificalo con `checksums.txt`, estrailo e metti `dna` (o `dna.exe`) nel tuo `PATH`.

| OS | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

Verifica l'installazione con `dna version` e `dna doctor`.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> Avvio rapido

```bash
# 1. Accedi prima all'ambiente di test (ha una propria chiave API)
dna auth login --env ote

# 2. Dai un'occhiata: nessuno di questi comandi modifica nulla
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. Visualizza l'anteprima di un acquisto, poi effettualo (chiede conferma)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. Produzione
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> Basta digitare `dna`: la shell interattiva

<img src="./assets/it/shell.svg" width="1200" alt="La shell interattiva di dna: digitando acme compaiono le estensioni libere con i prezzi, l'utente sceglie acme.net, compila un breve modulo di registrazione, conferma il passaggio a pagamento in una finestra di conferma e il dominio viene registrato.">

Avvia `dna` da solo in un terminale. Non serve conoscere nessun comando: tutto si sceglie con ↑↓ e Invio, ed Esc torna sempre indietro di un passo.

|  | Cosa ottieni |
|---|---|
| 🔎 **Ricerca live** | Digita un nome come `acme`: le estensioni libere e i loro prezzi compaiono mentre scrivi. Con ↓ ne scegli una e la registri con un breve modulo. |
| ⚡ **Azioni rapide** | ↓ su una riga vuota: i tuoi domini, registra un dominio, record DNS, saldo, passaggio tra OTE e produzione. |
| 🗂️ **Menu dei domini** | Cerca tra i tuoi domini e premi Invio: dettagli, DNS, rinnovo, blocco, privacy WHOIS, nameserver, codice di autorizzazione. |
| 🛡️ **Niente per sbaglio** | Ogni passaggio a pagamento o distruttivo termina in una finestra di conferma con la richiesta esatta e il prezzo. Annulla resta selezionato finché non scegli altrimenti. |
| 🤖 **Collega un agente IA** | `/agent` elenca oltre 20 agenti: Claude Code, ChatGPT, Cursor, Claude Desktop, Codex e altri. dna lo configura, verifica le credenziali e ti fornisce prompt pronti all'uso. |

Script e agenti IA non vedono mai la shell: in una pipe o avviato da un agente, `dna` stampa JSON e l'aiuto esattamente come prima.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> Collega il tuo agente AI

<img src="./assets/it/mcp.svg" width="1200" alt="Model Context Protocol: dna mcp serve espone 65 strumenti con annotazioni di sicurezza. Una chiamata a uno strumento a pagamento senza conferma restituisce un'anteprima; con confirm=true, dopo l'approvazione, viene eseguita.">

> [!TIP]
> **Il modo più rapido:** avvia `dna`, premi ↓ e scegli **Connect an AI agent** (oppure digita `/agent`). Scegli uno degli oltre 20 agenti, OTE o produzione e quanto può fare. dna esegue la configurazione o si aggiunge al file di configurazione dell'agente, verifica le credenziali e copia per te i prompt pronti all'uso.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # snippet per tutti gli altri client
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
<summary><b>Agenti con una shell (skill e guida)</b></summary>
<br>

```bash
dna agent skill --install                       # skill di Claude Code: ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # skill di progetto, da versionare insieme al tuo repository
dna agent guide                                 # guida d'uso compatta per qualsiasi agente
```
</details>

Sei tu a decidere fin dove può arrivare l'agente:

```bash
dna mcp serve --read-only            # guardare, non toccare
dna mcp serve --no-billable          # tutto tranne gli acquisti
dna mcp serve --groups domain,dns    # solo questi gruppi di comandi
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> Prompt pronti all'uso per Claude e ChatGPT

Copia un prompt, incollalo nel tuo assistente IA e lascia che faccia il lavoro. Ogni prompt inizia nell'ambiente di test OTE e si ferma in attesa della tua approvazione prima di qualsiasi operazione a pagamento.

**Configura dna in Claude Code**

```text
Configurami la CLI dna (https://github.com/domainreseller/cli).
1. Installala con lo script di installazione ufficiale per il mio sistema operativo,
   poi esegui `dna version` e `dna doctor`.
2. Installa la skill di Claude Code con `dna agent skill --install` e aggiungi il
   server MCP con `claude mcp add dna -- dna mcp serve --env ote`.
3. Non chiedermi di incollare la mia chiave API in questa chat. Dimmi di eseguire
   personalmente `dna auth login --env ote`, poi verifica la connessione con
   `dna --ote account get`.
Lavora solo nell'ambiente di test OTE e non usare mai `--yes` senza la mia
approvazione.
```

**Configura dna in ChatGPT Codex**

```text
Configurami la CLI dna (https://github.com/domainreseller/cli).
1. Installala con lo script di installazione ufficiale per il mio sistema operativo,
   poi esegui `dna version` e `dna doctor`.
2. Aggiungila alla mia configurazione di Codex (`~/.codex/config.toml`) come server
   MCP chiamato `dna`, con il comando `dna` e gli argomenti `mcp serve --env ote`.
3. Non chiedermi di incollare la mia chiave API in questa chat. Dimmi di eseguire
   personalmente `dna auth login --env ote`, poi verifica la connessione con
   `dna --ote account get`.
Lavora solo nell'ambiente di test OTE e non usare mai `--yes` senza la mia
approvazione.
```

**Niente terminale? Chiedi a ChatGPT o a Claude in chat**

```text
Leggi https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt e la
documentazione di dna a cui rimanda. Poi dammi i comandi dna esatti, passo dopo passo,
per registrare un dominio nell'ambiente di test OTE: installazione, accesso, ricerca,
anteprima con `--dry-run` e registrazione. Spiega cosa fa ogni comando.
```

<details>
<summary><b>Altri prompt per il lavoro di tutti i giorni</b></summary>
<br>

**Trova e registra un dominio**

```text
Con dna, controlla se brandname.com, brandname.net e brandname.io sono disponibili.
Mostra i risultati con i prezzi in una tabella, poi fai l'anteprima della
registrazione del più economico tra quelli disponibili con `--dry-run`. Non registrare
nulla finché non ti dico di sì.
```

**Prossimi rinnovi**

```text
Con dna, elenca i miei domini che scadono nei prossimi 60 giorni con i prezzi di
rinnovo, ordinati per data di scadenza. Non rinnovare nulla.
```

**Fai puntare un dominio a un server**

```text
Con dna, fai puntare example.com a 203.0.113.10: aggiungi record A per `@` e `www`,
poi mostrami la zona DNS finale. Chiedimi prima di modificare o eliminare qualsiasi
record esistente.
```

**Trasferisci un dominio sul tuo account**

```text
Con dna, controlla se example.org può essere trasferito sul mio account, dimmi cosa mi
serve ancora (codice di autorizzazione, sblocco) e fai l'anteprima del trasferimento
con `--dry-run`.
```

**Ordina un certificato SSL**

```text
Con dna, trova il certificato DV SSL più economico per example.com, crea una CSR e fai
l'anteprima dell'ordine. Aspetta la mia approvazione prima di acquistare.
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> Come funziona

<img src="./assets/it/flow.svg" width="1200" alt="Un agente AI, o tu stesso, chiama dna tramite MCP o la CLI; dna convalida l'input, applica il controllo di sicurezza e comunica con Domain Name API (produzione o la sandbox OTE), che raggiunge i registri tramite EPP.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> Controllo di sicurezza

<img src="./assets/it/safety.svg" width="1200" alt="I comandi di lettura e scrittura vengono eseguiti immediatamente; i comandi distruttivi e a pagamento si fermano finché non vengono confermati con --yes nella CLI o con confirm=true tramite MCP.">

| classe | esempi | comportamento |
|---|---|---|
| `read` | ricerca, elenco, informazioni, prezzi | eseguito immediatamente |
| `write` | record DNS, blocco, privacy, contatti | eseguito immediatamente (gratuito, reversibile) |
| `destructive` | eliminare un record, revocare un SSL, disattivare DNSSEC | si ferma finché non viene confermato |
| `billable` | registrazione, rinnovo, trasferimento, ripristino, acquisto di SSL/hosting | si ferma finché non viene confermato |

- **In uno script o nella shell di un agente:** senza `--yes`, un comando protetto termina con il codice **9** e non invia nulla.
- **Tramite MCP:** senza `confirm=true`, lo strumento restituisce un'anteprima della richiesta esatta.
- **Anteprima:** `--dry-run` stampa la richiesta HTTP con la chiave API mascherata.
- **Blocco totale:** `--read-only`, `DNA_READ_ONLY=1` o un profilo di sola lettura bloccano qualsiasi modifica (codice di uscita **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> Output, errori e codici di uscita

```bash
dna domain info example.com                   # tabella nel terminale
dna domain info example.com -o json           # JSON (predefinito quando l'output passa in una pipe)
dna domain list --all -o jsonl                # un dominio per riga
dna domain info example.com --jq .authCode    # jq integrato
dna dns add --body @record.json               # input completo da un file, oppure --body - per stdin
dna schema domain.register                    # JSON Schema di qualsiasi input
dna enums DomainStatus                        # cosa significa lo stato 7
```

Gli errori vengono restituiti in JSON su stderr: `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| codice | significato | codice | significato |
|---:|---|---:|---|
| 0 | successo | 6 | limite di richieste superato |
| 1 | errore imprevisto | 7 | errore del server |
| 2 | input non valido, nulla inviato | 8 | rete |
| 3 | autenticazione | 9 | conferma richiesta |
| 4 | non trovato | 10 | modalità di sola lettura |
| 5 | rifiutato dall'API | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> Comandi

| gruppo | cosa fa |
|---|---|
| `dna` | la shell interattiva: ricerca live, menu, moduli guidati, configurazione degli agenti IA |
| `dna domain` | ricerca, ricerca multipla, elenco, informazioni, registrazione, rinnovo, trasferimento, ripristino, blocco, privacy, codice di autorizzazione, contatti, documenti, inoltro, nameserver, host figli |
| `dna dns` · `dna dnssec` | record della zona, record DS |
| `dna contact` | handle dei contatti e verifica |
| `dna ssl` | CSR, acquisto, configurazione, stato, download, riemissione, revoca |
| `dna hosting` · `dna product` · `dna order` · `dna account` | hosting, catalogo e prezzi, ordini, saldo e transazioni |
| `dna api` | richiesta diretta a qualsiasi endpoint |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | per agenti e strumenti |
| `dna auth` · `dna config` · `dna doctor` | credenziali, profili, diagnostica |

Il `--help` di ogni comando mostra la sua classe di sicurezza, la route dell'API ed esempi d'uso. Consulta il [riferimento completo dei comandi](docs/reference/README.md).

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> Credenziali API

`dna` richiede gli stessi due valori di ogni modulo di Domain Name API:

| valore | cos'è |
|---|---|
| **Reseller ID** | un UUID nel formato `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` |
| **API Key** | la tua chiave API; produzione e OTE hanno chiavi **diverse** |

> [!TIP]
> Li trovi entrambi nel tuo pannello Domain Name API, nella sezione **API Settings**. Salvali con `dna auth login` oppure, in CI, usa le variabili d'ambiente `DNA_RESELLER_ID` e `DNA_API_KEY`.

> [!CAUTION]
> Queste sono **credenziali API**. L'e-mail e la password con cui accedi al pannello **non** funzionano qui.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> Domande frequenti

<details>
<summary><b>È sicuro lasciare che un agente AI usi dna?</b></summary>
<br>
Le operazioni a pagamento e distruttive sono protette: la CLI le rifiuta senza <code>--yes</code> e il server MCP restituisce un'anteprima a meno che la chiamata non includa <code>confirm=true</code>, che gli agenti sono istruiti a impostare solo dopo la tua approvazione. Avvia gli agenti con <code>--env ote</code> e usa <code>--read-only</code> o <code>--no-billable</code> per limitarli ulteriormente.
</details>

<details>
<summary><b>Posso provarlo senza spendere soldi?</b></summary>
<br>
Sì. <code>--env ote</code> usa l'ambiente di test OTE, con una propria chiave API e un proprio saldo. Nulla di ciò che avviene lì tocca i domini reali o il tuo saldo effettivo.
</details>

<details>
<summary><b>Con quali strumenti AI funziona?</b></summary>
<br>
Con qualsiasi client MCP: Claude Code, Claude Desktop, Cursor, VS Code, Windsurf, OpenAI Codex CLI, Gemini CLI e altri. Gli agenti che dispongono solo di una shell possono chiamare direttamente la CLI; <code>dna agent guide</code> e la skill di Claude Code spiegano loro come usarla.
</details>

<details>
<summary><b>Dove vengono salvate le mie credenziali?</b></summary>
<br>
<code>dna auth login</code> le conserva in un file di profilo nella cartella di configurazione del tuo utente (<code>dna config path</code> mostra dove), scritto con permessi riservati al solo proprietario su Linux e macOS. Le variabili d'ambiente hanno sempre la precedenza sul file e la chiave API viene mascherata in ogni output.
</details>

<details>
<summary><b>Serve Go?</b></summary>
<br>
No. Gli script di installazione e gli archivi delle release contengono un unico binario. Go 1.26+ serve solo per <code>go install</code> o per compilare dai sorgenti.
</details>

<details>
<summary><b>Esiste una libreria da usare nel mio codice?</b></summary>
<br>
<code>dna</code> è una CLI e un server MCP. Per chiamare l'API dal codice, usa le librerie ufficiali per <a href="https://github.com/domainreseller/php-dna">PHP</a>, <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>, <a href="https://github.com/domainreseller/python-dna">Python</a>, <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> o <a href="https://github.com/domainreseller/go-dna">Go</a>.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> Documentazione

| guida | contenuto |
|---|---|
| [Primi passi](docs/getting-started.md) | installazione, accesso a OTE, prima registrazione e prima modifica DNS |
| [Configurazione](docs/configuration.md) | credenziali, ambienti, profili, variabili d'ambiente, modalità di sola lettura |
| [Utilizzo](docs/usage.md) | input, `--body`, formati di output, `--jq`, paginazione, limiti di richieste, esempi pratici |
| [Agenti AI e MCP](docs/ai-agents.md) | configurazione MCP per ciascun client, protocollo di conferma, skill, guida per agenti |
| [Errori e codici di uscita](docs/errors.md) | JSON degli errori, codici di errore, risoluzione dei problemi |
| [Riferimento dei comandi](docs/reference/README.md) · [Enum](docs/enums.md) | ogni comando, flag e valore codificato (generato automaticamente) |
| [Architettura](docs/architecture.md) · [Come contribuire](CONTRIBUTING.md) | come è organizzato il codice |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> Contatti

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API: programma reseller di domini"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="Diventa un rivenditore di domini"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="API e moduli per rivenditori di domini"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="Supporto"></a>
</p>

<p align="right">
  <sub>Seguici</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · creato per le persone e per gli agenti che lavorano con loro</sub></p>
