<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa.svg" height="32" alt="العربية"></a>
  <a href="README-NL.md"><img src="./assets/lang/nl.svg" height="32" alt="Nederlands"></a>
  <a href="README-AZ.md"><img src="./assets/lang/az.svg" height="32" alt="Azərbaycanca"></a>
  <a href="README-CN.md"><img src="./assets/lang/cn.svg" height="32" alt="中文"></a>
  <a href="README-FR.md"><img src="./assets/lang/fr-active.svg" height="32" alt="Français"></a>
  <a href="README-IT.md"><img src="./assets/lang/it.svg" height="32" alt="Italiano"></a>
</p>

<!--
  dna: the official Domain Name API command line client and MCP server, built for AI agents and developers.
  Every animation on this page is a hand-made SVG in ./assets. No external widgets, no trackers.
-->

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/fr/hero.svg" width="1200" alt="dna, la CLI Domain Name API pilotée par des agents IA : un agent recherche un domaine, prévisualise l'enregistrement facturable avec --dry-run, demande votre approbation et l'enregistre."></a>

<h1 align="center">dna : la CLI Domain Name API pour les agents IA</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="Dernière version"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b> pilote la plateforme revendeur <a href="https://www.domainnameapi.com">Domain Name API</a> depuis la ligne de commande et depuis des agents IA.
  L'outil couvre les 64 opérations de l'API : recherche, enregistrement, renouvellement et transfert de domaines, DNS, DNSSEC, contacts, SSL et hébergement.
  Connectez-le à <b>Claude Code, Claude Desktop, Cursor, Codex</b> ou à tout client MCP, et chaque opération devient un outil sûr et clairement décrit.
</p>

<img src="./assets/fr/stats.svg" width="1200" alt="64 opérations d'API, 65 outils MCP pour les agents IA, 4 classes de sécurité appliquées, 6 plateformes avec un seul binaire.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> Conçu pour les agents IA

|  | Ce qu'obtient un agent |
|---|---|
| 🧰 **Un outil par opération** | `dna mcp serve` expose 65 outils MCP dotés de JSON Schemas générés à partir de la spécification de l'API, afin que les modèles remplissent correctement les entrées du premier coup. |
| 🛡️ **Une barrière de sécurité** | Les appels facturables et destructifs ne s'exécutent jamais d'eux-mêmes. La CLI exige `--yes`, MCP exige `confirm=true`, et sans cela l'agent reçoit un aperçu à vous présenter. |
| 🧾 **Tout est lisible par machine** | Sortie JSON lorsque la sortie est redirigée vers un pipe, erreurs JSON avec des codes de sortie stables, `--dry-run` pour la requête exacte et `dna schema` pour chaque entrée. |
| 🧪 **Un bac à sable** | `--env ote` pointe vers l'environnement de test : une vraie API, sans argent réel. `--read-only` bloque toute modification. |
| 📘 **Autodocumenté** | `dna agent guide` enseigne l'ensemble de la CLI à un agent en un seul appel. `dna agent skill --install` ajoute une skill Claude Code. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> Installation

> [!IMPORTANT]
> **Installez toujours depuis la _dernière version publiée_.** Les binaires publiés sont versionnés, testés et accompagnés de sommes de contrôle SHA-256.

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

**Avec Go 1.26+**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>Téléchargement manuel</b></summary>
<br>

Choisissez votre archive sur la [dernière version publiée](https://github.com/domainreseller/cli/releases/latest), vérifiez-la à l'aide de `checksums.txt`, décompressez-la et placez `dna` (ou `dna.exe`) dans votre `PATH`.

| OS | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

Vérifiez l'installation avec `dna version` et `dna doctor`.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> Démarrage rapide

```bash
# 1. Connectez-vous d'abord à l'environnement de test (il dispose de sa propre clé API)
dna auth login --env ote

# 2. Explorez : rien ici ne modifie quoi que ce soit
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. Prévisualisez un achat, puis effectuez-le (une confirmation vous est demandée)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. Production
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> Tapez simplement `dna` : le shell interactif

<img src="./assets/fr/shell.svg" width="1200" alt="Le shell interactif de dna : en tapant acme, les extensions libres s’affichent avec leurs prix, l’utilisateur choisit acme.net, remplit un court formulaire d’enregistrement, confirme l’étape payante dans une boîte de confirmation et le domaine est enregistré.">

Lancez `dna` seul dans un terminal. Inutile de connaître la moindre commande : tout se choisit avec ↑↓ et Entrée, et Échap revient toujours d’une étape.

|  | Ce que vous obtenez |
|---|---|
| 🔎 **Recherche en direct** | Tapez un nom comme `acme` : les extensions libres et leurs prix s’affichent pendant la saisie. ↓ en choisit une et l’enregistre via un court formulaire. |
| ⚡ **Actions rapides** | ↓ sur une ligne vide : vos domaines, enregistrer un domaine, enregistrements DNS, solde, passage entre OTE et production. |
| 🗂️ **Menus de domaine** | Cherchez dans vos domaines et appuyez sur Entrée : détails, DNS, renouvellement, verrouillage, confidentialité WHOIS, serveurs de noms, code d’autorisation. |
| 🛡️ **Rien par accident** | Chaque étape payante ou destructrice se termine par une boîte de confirmation avec la requête exacte et le prix. Annuler reste sélectionné tant que vous ne choisissez pas autre chose. |
| 🤖 **Connecter un agent IA** | `/agent` propose plus de 20 agents : Claude Code, ChatGPT, Cursor, Claude Desktop, Codex et d’autres. dna le configure, vérifie les identifiants et vous donne des prompts prêts à l’emploi. |

Les scripts et les agents IA ne voient jamais le shell : redirigé ou lancé par un agent, `dna` affiche du JSON et l’aide exactement comme avant.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> Connectez votre agent IA

<img src="./assets/fr/mcp.svg" width="1200" alt="Model Context Protocol : dna mcp serve expose 65 outils avec des annotations de sécurité. Un appel d'outil facturable sans confirmation renvoie un aperçu ; avec confirm=true après approbation, il s'exécute.">

> [!TIP]
> **Le plus rapide :** lancez `dna`, appuyez sur ↓ et choisissez **Connect an AI agent** (ou tapez `/agent`). Choisissez l’un des plus de 20 agents, OTE ou production et ce qu’il a le droit de faire. dna lance la configuration ou s’ajoute au fichier de configuration de l’agent, vérifie les identifiants et copie pour vous des prompts prêts à l’emploi.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # extraits de configuration pour tous les autres clients
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

Claude Desktop : `claude_desktop_config.json` · Cursor : `~/.cursor/mcp.json` · Windsurf : `~/.codeium/windsurf/mcp_config.json` · Gemini CLI : `~/.gemini/settings.json`
</details>

<details>
<summary><b>VS Code</b></summary>
<br>

`.vscode/mcp.json` :

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

`~/.codex/config.toml` :

```toml
[mcp_servers.dna]
command = "dna"
args = ["mcp", "serve", "--env", "ote"]
```
</details>

<details>
<summary><b>Agents disposant d'un shell (skill et guide)</b></summary>
<br>

```bash
dna agent skill --install                       # skill Claude Code : ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # skill de projet, à committer avec votre dépôt
dna agent guide                                 # guide d'utilisation compact pour tout agent
```
</details>

Vous choisissez jusqu'où l'agent peut aller :

```bash
dna mcp serve --read-only            # consulter sans rien modifier
dna mcp serve --no-billable          # tout sauf les achats
dna mcp serve --groups domain,dns    # uniquement ces groupes de commandes
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> Prompts prêts à l'emploi pour Claude et ChatGPT

Copiez un prompt, collez-le dans votre assistant IA et laissez-le faire le travail. Chaque prompt commence dans l'environnement de test OTE et s'arrête pour attendre votre accord avant toute opération payante.

**Installer dna dans Claude Code**

```text
Mets en place la CLI dna pour moi (https://github.com/domainreseller/cli).
1. Installe-la avec le script d'installation officiel pour mon système d'exploitation,
   puis lance `dna version` et `dna doctor`.
2. Installe le skill Claude Code avec `dna agent skill --install` et ajoute le serveur
   MCP avec `claude mcp add dna -- dna mcp serve --env ote`.
3. Ne me demande pas de coller ma clé API dans cette conversation. Dis-moi de lancer
   moi-même `dna auth login --env ote`, puis vérifie la connexion avec
   `dna --ote account get`.
Travaille uniquement dans l'environnement de test OTE et n'utilise jamais `--yes` sans
mon accord.
```

**Installer dna dans ChatGPT Codex**

```text
Mets en place la CLI dna pour moi (https://github.com/domainreseller/cli).
1. Installe-la avec le script d'installation officiel pour mon système d'exploitation,
   puis lance `dna version` et `dna doctor`.
2. Ajoute-la à ma config Codex (`~/.codex/config.toml`) comme serveur MCP nommé `dna`,
   avec la commande `dna` et les arguments `mcp serve --env ote`.
3. Ne me demande pas de coller ma clé API dans cette conversation. Dis-moi de lancer
   moi-même `dna auth login --env ote`, puis vérifie la connexion avec
   `dna --ote account get`.
Travaille uniquement dans l'environnement de test OTE et n'utilise jamais `--yes` sans
mon accord.
```

**Pas de terminal ? Passez par le chat de ChatGPT ou de Claude**

```text
Lis https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt et la
documentation dna vers laquelle il renvoie. Ensuite, donne-moi les commandes dna
exactes, étape par étape, pour enregistrer un domaine dans l'environnement de test OTE
: installation, connexion, recherche, prévisualisation avec `--dry-run` et
enregistrement. Explique ce que fait chaque commande.
```

<details>
<summary><b>D'autres prompts pour le travail au quotidien</b></summary>
<br>

**Trouver et enregistrer un domaine**

```text
Avec dna, vérifie si brandname.com, brandname.net et brandname.io sont disponibles.
Affiche les résultats avec les prix dans un tableau, puis prévisualise avec
`--dry-run` l'enregistrement du moins cher des domaines disponibles. N'enregistre rien
tant que je n'ai pas dit oui.
```

**Renouvellements à venir**

```text
Avec dna, liste mes domaines qui expirent dans les 60 prochains jours avec leur prix
de renouvellement, triés par date d'expiration. Ne renouvelle rien.
```

**Faire pointer un domaine vers un serveur**

```text
Avec dna, fais pointer example.com vers 203.0.113.10 : ajoute des enregistrements A
pour `@` et `www`, puis montre-moi la zone DNS finale. Demande-moi avant de modifier
ou de supprimer le moindre enregistrement existant.
```

**Transférer un domaine vers votre compte**

```text
Avec dna, vérifie si example.org peut être transféré vers mon compte, dis-moi ce qu'il
me manque encore (code d'autorisation, déverrouillage) et prévisualise le transfert
avec `--dry-run`.
```

**Commander un certificat SSL**

```text
Avec dna, trouve le certificat DV SSL le moins cher pour example.com, crée un CSR et
prévisualise la commande. Attends mon accord avant d'acheter.
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> Fonctionnement

<img src="./assets/fr/flow.svg" width="1200" alt="Un agent IA, ou vous-même, appelle dna via MCP ou la CLI ; dna valide les entrées, applique la barrière de sécurité et communique avec Domain Name API (production ou bac à sable OTE), qui atteint les registres via EPP.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> Barrière de sécurité

<img src="./assets/fr/safety.svg" width="1200" alt="Les commandes de lecture et d'écriture s'exécutent immédiatement ; les commandes destructives et facturables s'arrêtent jusqu'à leur confirmation avec --yes dans la CLI ou confirm=true via MCP.">

| classe | exemples | comportement |
|---|---|---|
| `read` | recherche, liste, informations, prix | s'exécute immédiatement |
| `write` | enregistrements DNS, verrouillage, confidentialité, contacts | s'exécute immédiatement (gratuit, réversible) |
| `destructive` | suppression d'un enregistrement, révocation SSL, désactivation de DNSSEC | s'arrête jusqu'à confirmation |
| `billable` | enregistrement, renouvellement, transfert, restauration, achat de SSL/hébergement | s'arrête jusqu'à confirmation |

- **Dans un script ou un shell d'agent :** sans `--yes`, une commande protégée se termine avec le code **9** et n'envoie rien.
- **Via MCP :** sans `confirm=true`, l'outil renvoie un aperçu de la requête exacte.
- **Prévisualisation :** `--dry-run` affiche la requête HTTP avec la clé API masquée.
- **Verrouillage :** `--read-only`, `DNA_READ_ONLY=1` ou un profil en lecture seule bloque toute modification (code de sortie **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> Sortie, erreurs et codes de sortie

```bash
dna domain info example.com                   # tableau dans un terminal
dna domain info example.com -o json           # JSON (par défaut lorsque la sortie est redirigée vers un pipe)
dna domain list --all -o jsonl                # un domaine par ligne
dna domain info example.com --jq .authCode    # jq intégré
dna dns add --body @record.json               # entrée complète depuis un fichier, ou --body - pour stdin
dna schema domain.register                    # JSON Schema de n'importe quelle entrée
dna enums DomainStatus                        # signification du statut 7
```

Les erreurs sont renvoyées en JSON sur stderr : `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| code | signification | code | signification |
|---:|---|---:|---|
| 0 | succès | 6 | limite de débit atteinte |
| 1 | erreur inattendue | 7 | erreur serveur |
| 2 | entrée invalide, rien n'a été envoyé | 8 | réseau |
| 3 | authentification | 9 | confirmation requise |
| 4 | introuvable | 10 | mode lecture seule |
| 5 | rejeté par l'API | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> Commandes

| groupe | rôle |
|---|---|
| `dna` | le shell interactif : recherche en direct, menus, formulaires guidés, configuration des agents IA |
| `dna domain` | recherche, recherche groupée, liste, informations, enregistrement, renouvellement, transfert, restauration, verrouillage, confidentialité, code d'autorisation, contacts, documents, redirection, serveurs de noms, hôtes enfants |
| `dna dns` · `dna dnssec` | enregistrements de zone, enregistrements DS |
| `dna contact` | handles de contact et vérification |
| `dna ssl` | CSR, achat, configuration, statut, téléchargement, réémission, révocation |
| `dna hosting` · `dna product` · `dna order` · `dna account` | hébergement, catalogue et prix, commandes, solde et transactions |
| `dna api` | requête brute vers n'importe quel endpoint |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | pour les agents et les outils |
| `dna auth` · `dna config` · `dna doctor` | identifiants, profils, diagnostics |

L'option `--help` de chaque commande indique sa classe de sécurité, sa route d'API et des exemples. Consultez la [référence complète des commandes](docs/reference/README.md).

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> Identifiants API

`dna` a besoin des deux mêmes valeurs que tous les modules Domain Name API :

| valeur | description |
|---|---|
| **Reseller ID** | un UUID de la forme `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` |
| **API Key** | votre clé API ; la production et l'OTE utilisent des clés **différentes** |

> [!TIP]
> Vous trouverez ces deux valeurs dans votre panneau Domain Name API, sous **API Settings**. Enregistrez-les avec `dna auth login`, ou utilisez les variables d'environnement `DNA_RESELLER_ID` et `DNA_API_KEY` en CI.

> [!CAUTION]
> Il s'agit d'**identifiants API**. L'adresse e-mail et le mot de passe de connexion à votre panneau ne fonctionneront **pas** ici.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> FAQ

<details>
<summary><b>Est-il sûr de laisser un agent IA utiliser dna ?</b></summary>
<br>
Les opérations facturables et destructives sont protégées : la CLI les refuse sans <code>--yes</code>, et le serveur MCP renvoie un aperçu tant que l'appel ne contient pas <code>confirm=true</code>, que les agents ont pour consigne de ne définir qu'après votre approbation. Démarrez les agents sur <code>--env ote</code>, et utilisez <code>--read-only</code> ou <code>--no-billable</code> pour les limiter davantage.
</details>

<details>
<summary><b>Puis-je l'essayer sans dépenser d'argent ?</b></summary>
<br>
Oui. <code>--env ote</code> utilise l'environnement de test OTE, avec sa propre clé API et son propre solde. Rien de ce qui s'y passe n'affecte les domaines en production ni votre solde réel.
</details>

<details>
<summary><b>Quels outils d'IA sont compatibles ?</b></summary>
<br>
Tout client MCP : Claude Code, Claude Desktop, Cursor, VS Code, Windsurf, OpenAI Codex CLI, Gemini CLI et d'autres. Les agents qui ne disposent que d'un shell peuvent appeler directement la CLI ; <code>dna agent guide</code> et la skill Claude Code leur en expliquent l'utilisation.
</details>

<details>
<summary><b>Où mes identifiants sont-ils stockés ?</b></summary>
<br>
<code>dna auth login</code> les conserve dans un fichier de profil situé dans votre dossier de configuration utilisateur (<code>dna config path</code> indique son emplacement), écrit avec des permissions réservées au propriétaire sous Linux et macOS. Les variables d'environnement ont toujours la priorité sur le fichier, et la clé API est masquée dans toutes les sorties.
</details>

<details>
<summary><b>Ai-je besoin de Go ?</b></summary>
<br>
Non. Les scripts d'installation et les archives publiées contiennent un binaire unique. Go 1.26+ n'est nécessaire que pour <code>go install</code> ou pour compiler depuis les sources.
</details>

<details>
<summary><b>Existe-t-il une bibliothèque pour mon propre code ?</b></summary>
<br>
<code>dna</code> est une CLI et un serveur MCP. Pour appeler l'API depuis votre code, utilisez les bibliothèques officielles <a href="https://github.com/domainreseller/php-dna">PHP</a>, <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>, <a href="https://github.com/domainreseller/python-dna">Python</a>, <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> ou <a href="https://github.com/domainreseller/go-dna">Go</a>.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> Documentation

| guide | contenu |
|---|---|
| [Prise en main](docs/getting-started.md) | installation, connexion à l'OTE, premier enregistrement et première modification DNS |
| [Configuration](docs/configuration.md) | identifiants, environnements, profils, variables d'environnement, mode lecture seule |
| [Utilisation](docs/usage.md) | entrées, `--body`, formats de sortie, `--jq`, pagination, limites de débit, recettes |
| [Agents IA et MCP](docs/ai-agents.md) | configuration MCP par client, protocole de confirmation, skill, guide pour agents |
| [Erreurs et codes de sortie](docs/errors.md) | JSON d'erreur, codes d'erreur, dépannage |
| [Référence des commandes](docs/reference/README.md) · [Énumérations](docs/enums.md) | toutes les commandes, options et valeurs codées (contenu généré) |
| [Architecture](docs/architecture.md) · [Contribuer](CONTRIBUTING.md) | organisation du code |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> Restons en contact

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API : programme de revendeur de domaines"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="Devenez revendeur de domaines"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="API et modules pour revendeurs de domaines"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="Assistance"></a>
</p>

<p align="right">
  <sub>Suivez-nous</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · conçu pour les personnes et les agents qui travaillent avec elles</sub></p>
