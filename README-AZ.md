<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa.svg" height="32" alt="العربية"></a>
  <a href="README-NL.md"><img src="./assets/lang/nl.svg" height="32" alt="Nederlands"></a>
  <a href="README-AZ.md"><img src="./assets/lang/az-active.svg" height="32" alt="Azərbaycanca"></a>
  <a href="README-CN.md"><img src="./assets/lang/cn.svg" height="32" alt="中文"></a>
  <a href="README-FR.md"><img src="./assets/lang/fr.svg" height="32" alt="Français"></a>
  <a href="README-IT.md"><img src="./assets/lang/it.svg" height="32" alt="Italiano"></a>
</p>

<!--
  dna: the official Domain Name API command line client and MCP server, built for AI agents and developers.
  Every animation on this page is a hand-made SVG in ./assets. No external widgets, no trackers.
-->

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/az/hero.svg" width="1200" alt="dna, AI agentlərinin idarə etdiyi Domain Name API CLI: agent domen axtarır, ödənişli qeydiyyatı --dry-run ilə önizləyir, təsdiq istəyir və domeni qeydiyyatdan keçirir."></a>

<h1 align="center">dna: AI agentləri üçün Domain Name API CLI</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="Son buraxılış"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b> <a href="https://www.domainnameapi.com">Domain Name API</a> reseller platformasını komanda sətrindən və AI agentləri vasitəsilə idarə edir.
  Bütün 64 API əməliyyatını əhatə edir: domen axtarışı, qeydiyyat, yenilənmə, transfer, DNS, DNSSEC, əlaqələr, SSL və hostinq.
  Onu <b>Claude Code, Claude Desktop, Cursor, Codex</b> və ya istənilən MCP klientinə qoşun və hər əməliyyat təhlükəsiz, yaxşı təsvir edilmiş bir alətə çevrilsin.
</p>

<img src="./assets/az/stats.svg" width="1200" alt="64 API əməliyyatı, AI agentləri üçün 65 MCP aləti, 4 məcburi təhlükəsizlik sinfi, tək binar fayldan 6 platforma.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> AI agentləri üçün hazırlanıb

|  | Agent nə əldə edir |
|---|---|
| 🧰 **Hər əməliyyat üçün bir alət** | `dna mcp serve` API spesifikasiyasından yaradılmış JSON Schema-larla 65 MCP aləti təqdim edir, beləliklə modellər girişləri ilk cəhddən düzgün doldurur. |
| 🛡️ **Təhlükəsizlik qapısı** | Ödənişli və dağıdıcı çağırışlar heç vaxt öz-özünə icra olunmur. CLI üçün `--yes`, MCP üçün `confirm=true` tələb olunur; bunlar olmadan agent sizə göstərmək üçün önizləmə alır. |
| 🧾 **Hər şey maşın tərəfindən oxuna bilir** | Pipe edildikdə JSON çıxışı, sabit çıxış kodları ilə JSON xətaları, dəqiq sorğunu görmək üçün `--dry-run` və hər giriş üçün `dna schema`. |
| 🧪 **Sınaq mühiti** | `--env ote` test mühitinə yönəldir: real API, real pul yoxdur. `--read-only` hər dəyişikliyi bloklayır. |
| 📘 **Özünü sənədləşdirən** | `dna agent guide` bir çağırışla agentə bütün CLI-ni öyrədir. `dna agent skill --install` Claude Code skill-i əlavə edir. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> Quraşdırma

> [!IMPORTANT]
> **Həmişə _son buraxılışdan_ quraşdırın.** Buraxılış binar faylları versiyalaşdırılır, test edilir və SHA-256 yoxlama cəmləri ilə dərc olunur.

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

**Go 1.26+ ilə**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>Əl ilə yükləmə</b></summary>
<br>

[Son buraxılışda](https://github.com/domainreseller/cli/releases/latest) arxivinizi seçin, onu `checksums.txt` ilə yoxlayın, arxivdən çıxarın və `dna` (və ya `dna.exe`) faylını `PATH`-ə əlavə edin.

| ƏS | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

Quraşdırmanı `dna version` və `dna doctor` ilə yoxlayın.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> Sürətli başlanğıc

```bash
# 1. Əvvəlcə test mühitinə daxil olun (onun öz API açarı var)
dna auth login --env ote

# 2. Ətrafa baxın: buradakı heç bir komanda heç nəyi dəyişmir
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. Alışı önizləyin, sonra həyata keçirin (təsdiq istəyəcək)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. Canlı mühit
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> Sadəcə `dna` yazın: interaktiv shell

<img src="./assets/az/shell.svg" width="1200" alt="dna interaktiv shell: acme yazanda boş uzantılar qiymətləri ilə görünür, istifadəçi acme.net-i seçir, qısa qeydiyyat formasını doldurur, ödənişli addımı təsdiq pəncərəsində təsdiqləyir və domen qeydiyyatdan keçir.">

Terminalda `dna`-nı tək başına işə salın. Heç bir əmri bilməyinizə ehtiyac yoxdur: hər şey ↑↓ və Enter ilə seçilir, Esc isə həmişə bir addım geri qaytarır.

|  | Nə əldə edirsiniz |
|---|---|
| 🔎 **Canlı axtarış** | `acme` kimi bir ad yazın: boş uzantılar və qiymətləri siz yazarkən görünür. ↓ ilə birini seçib qısa forma ilə qeydiyyatdan keçirirsiniz. |
| ⚡ **Sürətli əməliyyatlar** | Boş sətirdə ↓: domenləriniz, domen qeydiyyatı, DNS qeydləri, balans, OTE ilə production arasında keçid. |
| 🗂️ **Domen menyuları** | Domenlərinizdə axtarın və Enter basın: təfərrüatlar, DNS, yeniləmə, kilid, WHOIS məxfiliyi, nameserver-lər, auth kodu. |
| 🛡️ **Təsadüfən heç nə olmur** | Hər ödənişli və ya dağıdıcı addım dəqiq sorğunu və qiyməti göstərən təsdiq pəncərəsi ilə bitir. Siz başqasını seçməyincə Ləğv seçili qalır. |
| 🤖 **AI agentini qoşun** | `/agent` 20-dən çox agenti göstərir: Claude Code, ChatGPT, Cursor, Claude Desktop, Codex və digərləri. dna quraşdırmanı edir, giriş məlumatlarını yoxlayır və hazır promptlar verir. |

Skriptlər və AI agentləri shell-i heç vaxt görmür: pipe ilə və ya agent tərəfindən işə salındıqda `dna` əvvəlki kimi JSON və kömək mətni çıxarır.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> AI agentinizi qoşun

<img src="./assets/az/mcp.svg" width="1200" alt="Model Context Protocol: dna mcp serve təhlükəsizlik annotasiyaları ilə 65 alət təqdim edir. confirm olmadan ödənişli alət çağırışı önizləmə qaytarır; təsdiqdən sonra confirm=true ilə icra olunur.">

> [!TIP]
> **Ən sürətli yol:** `dna`-nı işə salın, ↓ basın və **Connect an AI agent** seçin (və ya `/agent` yazın). 20-dən çox agentdən birini, OTE və ya production mühitini və agentin nə qədər səlahiyyətli olacağını seçin. dna quraşdırmanı icra edir və ya özünü agentin konfiqurasiya faylına əlavə edir, giriş məlumatlarını yoxlayır və hazır promptları sizin üçün kopyalayır.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # digər bütün klientlər üçün konfiqurasiya nümunələri
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
<summary><b>Shell-ə çıxışı olan agentlər (skill və bələdçi)</b></summary>
<br>

```bash
dna agent skill --install                       # Claude Code skill-i: ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # layihə skill-i, onu repozitoriyanızla birlikdə commit edin
dna agent guide                                 # istənilən agent üçün yığcam istifadə bələdçisi
```
</details>

Agentin nəyə qədər çata biləcəyini siz müəyyən edirsiniz:

```bash
dna mcp serve --read-only            # bax, amma toxunma
dna mcp serve --no-billable          # alışlardan başqa hər şey
dna mcp serve --groups domain,dns    # yalnız bu komanda qrupları
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> Claude və ChatGPT üçün hazır promptlar

Promptu kopyalayın, AI assistentinizə yapışdırın və işi ona həvalə edin. Hər prompt OTE test mühitində başlayır və ödənişli əməliyyatdan əvvəl dayanıb təsdiqinizi gözləyir.

**dna-nı Claude Code-da quraşdırın**

```text
Mənim üçün dna CLI-ni quraşdır (https://github.com/domainreseller/cli).
1. Onu əməliyyat sistemimə uyğun rəsmi quraşdırma skripti ilə quraşdır, sonra
   `dna version` və `dna doctor` əmrlərini işə sal.
2. Claude Code skill-ini `dna agent skill --install` ilə quraşdır və MCP serverini
   `claude mcp add dna -- dna mcp serve --env ote` ilə əlavə et.
3. API açarımı bu söhbətə yapışdırmağımı istəmə. `dna auth login --env ote` əmrini
   özüm işə salmağımı söylə, sonra bağlantını `dna --ote account get` ilə yoxla.
Yalnız OTE test mühitində işlə və mənim təsdiqim olmadan heç vaxt `--yes`
parametrindən istifadə etmə.
```

**dna-nı ChatGPT Codex-də quraşdırın**

```text
Mənim üçün dna CLI-ni quraşdır (https://github.com/domainreseller/cli).
1. Onu əməliyyat sistemimə uyğun rəsmi quraşdırma skripti ilə quraşdır, sonra
   `dna version` və `dna doctor` əmrlərini işə sal.
2. Onu Codex konfiqurasiyama (`~/.codex/config.toml`) `dna` adlı MCP server kimi əlavə
   et: əmr `dna`, arqumentlər `mcp serve --env ote`.
3. API açarımı bu söhbətə yapışdırmağımı istəmə. `dna auth login --env ote` əmrini
   özüm işə salmağımı söylə, sonra bağlantını `dna --ote account get` ilə yoxla.
Yalnız OTE test mühitində işlə və mənim təsdiqim olmadan heç vaxt `--yes`
parametrindən istifadə etmə.
```

**Terminalınız yoxdur? ChatGPT və ya Claude söhbətində soruşun**

```text
https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt faylını və onun
keçid verdiyi dna sənədlərini oxu. Sonra OTE test mühitində bir domeni qeydiyyatdan
keçirmək üçün lazım olan dəqiq dna əmrlərini addım-addım ver: quraşdırma, giriş,
axtarış, `--dry-run` ilə önizləmə və qeydiyyat. Hər əmrin nə etdiyini izah et.
```

<details>
<summary><b>Gündəlik işlər üçün daha çox prompt</b></summary>
<br>

**Domen tapın və qeydiyyatdan keçirin**

```text
dna ilə brandname.com, brandname.net və brandname.io domenlərinin boş olub-olmadığını
yoxla. Nəticələri qiymətləri ilə birlikdə cədvəldə göstər, sonra boş olanlardan ən
ucuzunun qeydiyyatını `--dry-run` ilə önizlə. Mən hə deyənə qədər heç nəyi
qeydiyyatdan keçirmə.
```

**Yaxınlaşan yeniləmələr**

```text
dna ilə növbəti 60 gün ərzində müddəti bitəcək domenlərimi yeniləmə qiymətləri ilə
birlikdə, bitmə tarixinə görə sıralayaraq sadala. Heç nəyi yeniləmə.
```

**Domeni serverə yönləndirin**

```text
dna ilə example.com domenini 203.0.113.10 ünvanına yönləndir: `@` və `www` üçün A
qeydləri əlavə et, sonra mənə son DNS zonasını göstər. Hər hansı mövcud qeydi
dəyişməzdən və ya silməzdən əvvəl məndən soruş.
```

**Domeni hesabınıza transfer edin**

```text
dna ilə example.org domeninin hesabıma transfer edilə bilib-bilmədiyini yoxla, mənə
hələ nələrin lazım olduğunu de (auth kodu, kilidin açılması) və transferi `--dry-run`
ilə önizlə.
```

**SSL sertifikatı sifariş edin**

```text
dna ilə example.com üçün ən ucuz DV SSL sertifikatını tap, CSR yarat və sifarişi
önizlə. Almazdan əvvəl təsdiqimi gözlə.
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> Necə işləyir

<img src="./assets/az/flow.svg" width="1200" alt="AI agenti və ya siz MCP ya da CLI vasitəsilə dna-ya müraciət edirsiniz; dna girişi yoxlayır, təhlükəsizlik qapısını tətbiq edir və Domain Name API ilə (canlı mühit və ya OTE sınaq mühiti) əlaqə qurur, o da EPP üzərindən registrlərə qoşulur.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> Təhlükəsizlik qapısı

<img src="./assets/az/safety.svg" width="1200" alt="Oxuma və yazma komandaları dərhal icra olunur; dağıdıcı və ödənişli komandalar CLI-də --yes və ya MCP üzərindən confirm=true ilə təsdiqlənənə qədər dayanır.">

| sinif | nümunələr | davranış |
|---|---|---|
| `read` | axtarış, siyahı, məlumat, qiymətlər | dərhal icra olunur |
| `write` | DNS qeydləri, kilid, məxfilik, əlaqələr | dərhal icra olunur (pulsuz, geri qaytarıla bilən) |
| `destructive` | qeydin silinməsi, SSL-in ləğvi, DNSSEC-in deaktiv edilməsi | təsdiqlənənə qədər dayanır |
| `billable` | qeydiyyat, yenilənmə, transfer, bərpa, SSL/hostinq alışı | təsdiqlənənə qədər dayanır |

- **Skriptdə və ya agent shell-ində:** `--yes` olmadan qorunan komanda **9** kodu ilə başa çatır və heç nə göndərmir.
- **MCP üzərindən:** `confirm=true` olmadan alət dəqiq sorğunun önizləməsini qaytarır.
- **Önizləmə:** `--dry-run` HTTP sorğusunu API açarı maskalanmış şəkildə çap edir.
- **Tam bağlamaq:** `--read-only`, `DNA_READ_ONLY=1` və ya yalnız oxuma profili hər dəyişikliyi bloklayır (çıxış kodu **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> Çıxış formatı, xətalar və çıxış kodları

```bash
dna domain info example.com                   # terminalda cədvəl
dna domain info example.com -o json           # JSON (pipe edildikdə standart)
dna domain list --all -o jsonl                # hər sətirdə bir domen
dna domain info example.com --jq .authCode    # daxili jq
dna dns add --body @record.json               # bütün giriş fayldan, stdin üçün isə --body -
dna schema domain.register                    # istənilən girişin JSON Schema-sı
dna enums DomainStatus                        # 7 statusunun nə demək olduğu
```

Xətalar stderr-də JSON formatında verilir: `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| kod | mənası | kod | mənası |
|---:|---|---:|---|
| 0 | uğurlu | 6 | sorğu limiti aşılıb |
| 1 | gözlənilməz xəta | 7 | server xətası |
| 2 | yanlış giriş, heç nə göndərilməyib | 8 | şəbəkə |
| 3 | autentifikasiya | 9 | təsdiq tələb olunur |
| 4 | tapılmadı | 10 | yalnız oxuma rejimi |
| 5 | API tərəfindən rədd edildi | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> Komandalar

| qrup | nə edir |
|---|---|
| `dna` | interaktiv shell: canlı axtarış, menyular, addım-addım formalar, AI agent quraşdırması |
| `dna domain` | axtarış, toplu axtarış, siyahı, məlumat, qeydiyyat, yenilənmə, transfer, bərpa, kilid, məxfilik, auth kodu, əlaqələr, sənədlər, yönləndirmə, nameserver-lər, child host-lar |
| `dna dns` · `dna dnssec` | zona qeydləri, DS qeydləri |
| `dna contact` | əlaqə handle-ları və doğrulama |
| `dna ssl` | CSR, alış, konfiqurasiya, status, yükləmə, yenidən buraxılış, ləğv |
| `dna hosting` · `dna product` · `dna order` · `dna account` | hostinq, kataloq və qiymətlər, sifarişlər, balans və əməliyyatlar |
| `dna api` | istənilən endpoint-ə xam sorğu |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | agentlər və alətlər üçün |
| `dna auth` · `dna config` · `dna doctor` | giriş məlumatları, profillər, diaqnostika |

Hər komandanın `--help` çıxışı onun təhlükəsizlik sinfini, API marşrutunu və nümunələrini göstərir. Tam [komanda arayışına](docs/reference/README.md) baxın.

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> API giriş məlumatları

`dna` bütün Domain Name API modulları kimi eyni iki dəyərə ehtiyac duyur:

| dəyər | nədir |
|---|---|
| **Reseller ID** | `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` formatında UUID |
| **API Key** | API açarınız; canlı mühit və OTE üçün açarlar **fərqlidir** |

> [!TIP]
> Hər ikisini Domain Name API panelinizdə **API Settings** bölməsində tapa bilərsiniz. Onları `dna auth login` ilə yadda saxlayın və ya CI-da `DNA_RESELLER_ID` və `DNA_API_KEY` mühit dəyişənlərindən istifadə edin.

> [!CAUTION]
> Bunlar **API giriş məlumatlarıdır**. Panelə daxil olmaq üçün istifadə etdiyiniz e-poçt və şifrə burada **işləməyəcək**.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> Tez-tez verilən suallar

<details>
<summary><b>AI agentinin dna-dan istifadə etməsinə icazə vermək təhlükəsizdirmi?</b></summary>
<br>
Ödənişli və dağıdıcı əməliyyatlar qorunur: CLI <code>--yes</code> olmadan onları rədd edir, MCP serveri isə çağırışda <code>confirm=true</code> olmadıqca yalnız önizləmə qaytarır; agentlərə bu parametri yalnız siz təsdiq etdikdən sonra təyin etmək tapşırılır. Agentləri <code>--env ote</code> ilə başladın və onları daha da məhdudlaşdırmaq üçün <code>--read-only</code> və ya <code>--no-billable</code> istifadə edin.
</details>

<details>
<summary><b>Pul xərcləmədən sınaya bilərəmmi?</b></summary>
<br>
Bəli. <code>--env ote</code> öz API açarı və balansı olan OTE test mühitindən istifadə edir. Orada heç nə canlı domenlərə və ya real balansınıza təsir etmir.
</details>

<details>
<summary><b>Hansı AI alətləri onunla işləyir?</b></summary>
<br>
İstənilən MCP klienti: Claude Code, Claude Desktop, Cursor, VS Code, Windsurf, OpenAI Codex CLI, Gemini CLI və digərləri. Yalnız shell-ə çıxışı olan agentlər CLI-ni birbaşa çağıra bilər; <code>dna agent guide</code> və Claude Code skill-i bunu onlara izah edir.
</details>

<details>
<summary><b>Giriş məlumatlarım harada saxlanılır?</b></summary>
<br>
<code>dna auth login</code> onları istifadəçi konfiqurasiya qovluğunuzdakı profil faylında saxlayır (yerini <code>dna config path</code> göstərir); Linux və macOS-da bu fayl yalnız sahibinə aid icazələrlə yazılır. Mühit dəyişənləri həmişə fayldan üstün tutulur və API açarı bütün çıxışlarda maskalanır.
</details>

<details>
<summary><b>Go lazımdırmı?</b></summary>
<br>
Xeyr. Quraşdırma skriptləri və buraxılış arxivləri tək binar fayldan ibarətdir. Go 1.26+ yalnız <code>go install</code> üçün və ya mənbə kodundan qurmaq üçün lazımdır.
</details>

<details>
<summary><b>Öz kodum üçün kitabxana varmı?</b></summary>
<br>
<code>dna</code> CLI və MCP serveridir. API-ni koddan çağırmaq üçün rəsmi <a href="https://github.com/domainreseller/php-dna">PHP</a>, <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>, <a href="https://github.com/domainreseller/python-dna">Python</a>, <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> və ya <a href="https://github.com/domainreseller/go-dna">Go</a> kitabxanalarından istifadə edin.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> Sənədlər

| bələdçi | mövzu |
|---|---|
| [Başlanğıc](docs/getting-started.md) | quraşdırma, OTE-yə giriş, ilk qeydiyyat və DNS dəyişikliyi |
| [Konfiqurasiya](docs/configuration.md) | giriş məlumatları, mühitlər, profillər, mühit dəyişənləri, yalnız oxuma rejimi |
| [İstifadə](docs/usage.md) | girişlər, `--body`, çıxış formatları, `--jq`, səhifələmə, sorğu limitləri, hazır reseptlər |
| [AI agentləri və MCP](docs/ai-agents.md) | hər klient üçün MCP quraşdırması, təsdiq protokolu, skill, agent bələdçisi |
| [Xətalar və çıxış kodları](docs/errors.md) | xəta JSON-u, xəta kodları, problemlərin həlli |
| [Komanda arayışı](docs/reference/README.md) · [Enum-lar](docs/enums.md) | bütün komandalar, flag-lar və kodlaşdırılmış dəyərlər (avtomatik yaradılıb) |
| [Arxitektura](docs/architecture.md) · [Töhfə vermək](CONTRIBUTING.md) | kodun necə təşkil olunduğu |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> Əlaqə

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API: domen reseller proqramı"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="Domen reseller-i olun"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="Domen reseller API-si və modulları"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="Dəstək"></a>
</p>

<p align="right">
  <sub>Bizi izləyin</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · insanlar və onlarla işləyən agentlər üçün hazırlanıb</sub></p>
