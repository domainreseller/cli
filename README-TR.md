<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr-active.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
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

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/tr/hero.svg" width="1200" alt="dna, yapay zeka ajanlarının kullandığı Domain Name API CLI: bir ajan alan adını arar, ücretli kaydı --dry-run ile önizler, onay ister ve alan adını kaydeder."></a>

<h1 align="center">dna: yapay zeka ajanları için Domain Name API CLI</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="Son sürüm"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b>, <a href="https://www.domainnameapi.com">Domain Name API</a> bayi platformunu komut satırından ve yapay zeka ajanları üzerinden yönetir.
  64 API işleminin tamamını kapsar: alan adı arama, kayıt, yenileme, transfer, DNS, DNSSEC, iletişim bilgileri, SSL ve hosting.
  <b>Claude Code, Claude Desktop, Cursor, Codex</b> veya herhangi bir MCP istemcisine bağlayın; her işlem güvenli ve iyi tanımlanmış bir araca dönüşsün.
</p>

<img src="./assets/tr/stats.svg" width="1200" alt="64 API işlemi, yapay zeka ajanları için 65 MCP aracı, zorunlu kılınan 4 güvenlik sınıfı, tek bir ikili dosyayla 6 platform.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> Yapay zeka ajanları için tasarlandı

|  | Bir ajanın elde ettikleri |
|---|---|
| 🧰 **Her işlem için bir araç** | `dna mcp serve`, API spesifikasyonundan üretilen JSON Schema'larla 65 MCP aracı sunar; böylece modeller girdileri ilk seferde doğru doldurur. |
| 🛡️ **Bir güvenlik kapısı** | Ücretli ve yıkıcı çağrılar asla kendiliğinden çalışmaz. CLI `--yes`, MCP ise `confirm=true` ister; bu olmadan ajan, size göstermesi için bir önizleme alır. |
| 🧾 **Her şey makine tarafından okunabilir** | Pipe ile aktarıldığında JSON çıktı, sabit çıkış kodlarıyla JSON hatalar, isteğin birebir hali için `--dry-run` ve her girdi için `dna schema`. |
| 🧪 **Bir test alanı** | `--env ote` test ortamını (OTE) hedefler: gerçek API, gerçek para yok. `--read-only` her değişikliği engeller. |
| 📘 **Kendi kendini belgeler** | `dna agent guide`, tek bir çağrıyla bir ajana tüm CLI kullanımını öğretir. `dna agent skill --install` bir Claude Code skill'i ekler. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> Kurulum

> [!IMPORTANT]
> **Her zaman _en son sürümden_ kurun.** Yayınlanan ikili dosyalar sürüm numaralıdır, test edilmiştir ve SHA-256 sağlama toplamlarıyla birlikte yayımlanır.

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

**Go 1.26+ ile**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>Manuel indirme</b></summary>
<br>

[En son sürümden](https://github.com/domainreseller/cli/releases/latest) size uygun arşivi seçin, `checksums.txt` ile doğrulayın, arşivi açın ve `dna` (veya `dna.exe`) dosyasını `PATH` üzerindeki bir dizine koyun.

| İşletim sistemi | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

Kurulumu `dna version` ve `dna doctor` ile kontrol edin.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> Hızlı başlangıç

```bash
# 1. Önce test ortamına giriş yapın (kendine ait bir API anahtarı vardır)
dna auth login --env ote

# 2. Etrafa göz atın: buradaki hiçbir komut bir şey değiştirmez
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. Bir satın almayı önizleyin, ardından gerçekleştirin (onay ister)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. Canlı ortam
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> Sadece `dna` yazın: etkileşimli kabuk

<img src="./assets/tr/shell.svg" width="1200" alt="dna etkileşimli kabuğu: acme yazınca boş uzantılar fiyatlarıyla görünür, kullanıcı acme.net’i seçer, kısa kayıt formunu doldurur, ücretli adımı onay kutusunda onaylar ve alan adı kaydedilir.">

Terminalde `dna` komutunu tek başına çalıştırın. Tek bir komut bilmeniz gerekmez: her şey ↑↓ ve Enter ile seçilir, Esc her zaman bir adım geri gider.

|  | Neler var |
|---|---|
| 🔎 **Canlı arama** | `acme` gibi bir ad yazın: boş uzantılar ve fiyatları siz yazarken görünür. ↓ ile birini seçip kısa bir formla kaydedersiniz. |
| ⚡ **Hızlı eylemler** | Boş satırda ↓: alan adlarınız, alan adı kaydı, DNS kayıtları, bakiye, OTE ile production arasında geçiş. |
| 🗂️ **Alan adı menüleri** | Alan adlarınızda arayın ve Enter’a basın: detay, DNS, yenileme, kilit, WHOIS gizliliği, nameserver’lar, auth kodu. |
| 🛡️ **Yanlışlıkla hiçbir şey olmaz** | Her ücretli veya yıkıcı adım, gönderilecek isteği ve fiyatı gösteren bir onay kutusuyla biter. Siz değiştirmedikçe İptal seçili kalır. |
| 🤖 **AI ajanını bağlayın** | `/agent` 20’den fazla ajanı listeler: Claude Code, ChatGPT, Cursor, Claude Desktop, Codex ve diğerleri. dna kurulumu yapar, kimlik bilgilerini kontrol eder ve hazır promptları verir. |

Betikler ve AI ajanları kabuğu hiç görmez: çıktı bir pipe’a gidiyorsa ya da komutu bir ajan çalıştırıyorsa `dna` eskisi gibi JSON ve yardım metni basar.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> Yapay zeka ajanınızı bağlayın

<img src="./assets/tr/mcp.svg" width="1200" alt="Model Context Protocol: dna mcp serve, güvenlik etiketleriyle 65 araç sunar. confirm olmadan yapılan ücretli bir araç çağrısı önizleme döndürür; onaydan sonra confirm=true ile çalışır.">

> [!TIP]
> **En hızlı yol:** `dna` çalıştırın, ↓ tuşuna basın ve **Connect an AI agent** seçin (ya da `/agent` yazın). 20’den fazla ajandan birini, OTE ya da production ortamını ve ajanın ne kadar yetkili olacağını seçin. dna kurulumu çalıştırır ya da kendini ajanın yapılandırma dosyasına ekler, kimlik bilgilerini kontrol eder ve hazır promptları sizin için kopyalar.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # diğer tüm istemciler için yapılandırma parçacıkları
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
<summary><b>Kabuk erişimi olan ajanlar (skill ve rehber)</b></summary>
<br>

```bash
dna agent skill --install                       # Claude Code skill'i: ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # proje skill'i; deponuzla birlikte commit edin
dna agent guide                                 # her ajan için kısa kullanım rehberi
```
</details>

Ajanın ne kadar ileri gidebileceğine siz karar verirsiniz:

```bash
dna mcp serve --read-only            # bak ama dokunma
dna mcp serve --no-billable          # satın almalar hariç her şey
dna mcp serve --groups domain,dns    # yalnızca bu komut grupları
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> Claude ve ChatGPT için hazır promptlar

Bir promptu kopyalayın, AI asistanınıza yapıştırın ve gerisini ona bırakın. Her prompt OTE test ortamında başlar ve ücretli bir işlemden önce onayınızı bekler.

**dna'yı Claude Code'da kurun**

```text
Benim için dna CLI'yi kur (https://github.com/domainreseller/cli).
1. İşletim sistemime uygun resmi kurulum betiğiyle kur, ardından `dna version` ve
   `dna doctor` komutlarını çalıştır.
2. Claude Code skill'ini `dna agent skill --install` ile kur ve MCP sunucusunu
   `claude mcp add dna -- dna mcp serve --env ote` ile ekle.
3. API anahtarımı bu sohbete yapıştırmamı isteme. `dna auth login --env ote` komutunu
   benim çalıştırmamı söyle, sonra bağlantıyı `dna --ote account get` ile kontrol et.
Yalnızca OTE test ortamında çalış ve onayım olmadan asla `--yes` kullanma.
```

**dna'yı ChatGPT Codex'te kurun**

```text
Benim için dna CLI'yi kur (https://github.com/domainreseller/cli).
1. İşletim sistemime uygun resmi kurulum betiğiyle kur, ardından `dna version` ve
   `dna doctor` komutlarını çalıştır.
2. Codex yapılandırmama (`~/.codex/config.toml`) `dna` adında bir MCP sunucusu olarak
   ekle: komut `dna`, argümanlar `mcp serve --env ote`.
3. API anahtarımı bu sohbete yapıştırmamı isteme. `dna auth login --env ote` komutunu
   benim çalıştırmamı söyle, sonra bağlantıyı `dna --ote account get` ile kontrol et.
Yalnızca OTE test ortamında çalış ve onayım olmadan asla `--yes` kullanma.
```

**Terminaliniz yok mu? ChatGPT veya Claude sohbetine sorun**

```text
https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt dosyasını ve
bağlantı verdiği dna dokümanlarını oku. Sonra OTE test ortamında bir alan adı
kaydetmek için gereken dna komutlarını adım adım ve eksiksiz ver: kurulum, giriş,
sorgulama, `--dry-run` ile önizleme ve kayıt. Her komutun ne yaptığını açıkla.
```

<details>
<summary><b>Günlük işler için daha fazla prompt</b></summary>
<br>

**Alan adı bulun ve kaydedin**

```text
dna ile brandname.com, brandname.net ve brandname.io alan adlarının müsait olup
olmadığını kontrol et. Sonuçları fiyatlarıyla bir tabloda göster, ardından müsait
olanlar arasında en ucuzunun kaydını `--dry-run` ile önizle. Ben evet demeden hiçbir
şey kaydetme.
```

**Yaklaşan yenilemeler**

```text
dna ile önümüzdeki 60 gün içinde süresi dolacak alan adlarımı yenileme fiyatlarıyla
birlikte, bitiş tarihine göre sıralı listele. Hiçbir şeyi yenileme.
```

**Alan adını bir sunucuya yönlendirin**

```text
dna ile example.com alan adını 203.0.113.10 adresine yönlendir: `@` ve `www` için A
kayıtları ekle, ardından son DNS bölgesini göster. Mevcut bir kaydı değiştirmeden veya
silmeden önce bana sor.
```

**Alan adı transferi**

```text
dna ile example.org alan adının hesabıma transfer edilip edilemeyeceğini kontrol et,
eksik olanları söyle (auth kodu, kilidin kaldırılması) ve transferi `--dry-run` ile
önizle.
```

**SSL sertifikası sipariş edin**

```text
dna ile example.com için en ucuz DV SSL sertifikasını bul, bir CSR oluştur ve siparişi
önizle. Satın almadan önce onayımı bekle.
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> Nasıl çalışır

<img src="./assets/tr/flow.svg" width="1200" alt="Bir yapay zeka ajanı ya da siz dna'yı MCP veya CLI üzerinden çağırırsınız; dna girdiyi doğrular, güvenlik kapısını uygular ve Domain Name API ile (canlı ortam veya OTE test ortamı) iletişim kurar; Domain Name API de EPP üzerinden kayıt otoritelerine ulaşır.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> Güvenlik kapısı

<img src="./assets/tr/safety.svg" width="1200" alt="Okuma ve yazma komutları hemen çalışır; yıkıcı ve ücretli komutlar, CLI tarafında --yes veya MCP üzerinden confirm=true ile onaylanana kadar durur.">

| sınıf | örnekler | davranış |
|---|---|---|
| `read` | arama, listeleme, bilgi, fiyatlar | hemen çalışır |
| `write` | DNS kayıtları, kilit, gizlilik, iletişim bilgileri | hemen çalışır (ücretsiz, geri alınabilir) |
| `destructive` | kayıt silme, SSL iptali, DNSSEC'i devre dışı bırakma | onaylanana kadar durur |
| `billable` | kayıt, yenileme, transfer, geri yükleme, SSL/hosting satın alma | onaylanana kadar durur |

- **Bir betikte veya ajan kabuğunda:** `--yes` olmadan, onay gerektiren bir komut **9** koduyla çıkar ve hiçbir şey göndermez.
- **MCP üzerinden:** `confirm=true` olmadan araç, isteğin birebir önizlemesini döndürür.
- **Önizleme:** `--dry-run`, HTTP isteğini API anahtarı maskelenmiş olarak yazdırır.
- **Tamamen kilitleme:** `--read-only`, `DNA_READ_ONLY=1` veya salt-okunur bir profil her değişikliği engeller (çıkış kodu **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> Çıktı, hatalar ve çıkış kodları

```bash
dna domain info example.com                   # terminalde tablo
dna domain info example.com -o json           # JSON (pipe ile aktarıldığında varsayılan)
dna domain list --all -o jsonl                # her satırda bir alan adı
dna domain info example.com --jq .authCode    # yerleşik jq
dna dns add --body @record.json               # tüm girdi bir dosyadan; stdin için --body -
dna schema domain.register                    # herhangi bir girdinin JSON Schema'sı
dna enums DomainStatus                        # 7 numaralı durumun anlamı
```

Hatalar stderr'e JSON olarak yazılır: `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| kod | anlamı | kod | anlamı |
|---:|---|---:|---|
| 0 | başarılı | 6 | hız sınırına takıldı |
| 1 | beklenmeyen hata | 7 | sunucu hatası |
| 2 | geçersiz girdi, hiçbir şey gönderilmedi | 8 | ağ hatası |
| 3 | kimlik doğrulama | 9 | onay gerekli |
| 4 | bulunamadı | 10 | salt-okunur mod |
| 5 | API tarafından reddedildi | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> Komutlar

| grup | ne yapar |
|---|---|
| `dna` | etkileşimli kabuk: canlı arama, menüler, adım adım formlar, AI ajanı kurulumu |
| `dna domain` | arama, toplu arama, listeleme, bilgi, kayıt, yenileme, transfer, geri yükleme, kilit, gizlilik, transfer kodu (auth code), iletişim bilgileri, belgeler, yönlendirme, ad sunucuları (nameserver), child host'lar |
| `dna dns` · `dna dnssec` | bölge (zone) kayıtları, DS kayıtları |
| `dna contact` | iletişim kişisi tanımlayıcıları (handle) ve doğrulama |
| `dna ssl` | CSR, satın alma, yapılandırma, durum, indirme, yeniden düzenleme (reissue), iptal |
| `dna hosting` · `dna product` · `dna order` · `dna account` | hosting, katalog ve fiyatlar, siparişler, bakiye ve hesap hareketleri |
| `dna api` | herhangi bir uç noktaya ham istek |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | ajanlar ve araçlar için |
| `dna auth` · `dna config` · `dna doctor` | kimlik bilgileri, profiller, tanılama |

Her komutun `--help` çıktısı, o komutun güvenlik sınıfını, API rotasını ve örneklerini gösterir. Tüm komutlar için [komut referansına](docs/reference/README.md) bakın.

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> API kimlik bilgileri

`dna`, tüm Domain Name API modüllerinin kullandığı aynı iki değere ihtiyaç duyar:

| değer | nedir |
|---|---|
| **Reseller ID** | `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` biçiminde bir UUID |
| **API Key** | API anahtarınız; canlı ortam ve OTE için anahtarlar **farklıdır** |

> [!TIP]
> Her ikisini de Domain Name API panelinizde **API Ayarları** bölümünde bulabilirsiniz. Bunları `dna auth login` ile kaydedin veya CI ortamında `DNA_RESELLER_ID` ve `DNA_API_KEY` ortam değişkenlerini kullanın.

> [!CAUTION]
> Bunlar **API kimlik bilgileridir**. Panele giriş yaptığınız e-posta adresi ve şifre burada **çalışmaz**.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> Sıkça sorulan sorular

<details>
<summary><b>Bir yapay zeka ajanının dna'yı kullanmasına izin vermek güvenli mi?</b></summary>
<br>
Ücretli ve yıkıcı işlemler bir onay kapısının arkasındadır: CLI bunları <code>--yes</code> olmadan reddeder; MCP sunucusu ise çağrı <code>confirm=true</code> içermedikçe yalnızca bir önizleme döndürür. Ajanlara bu değeri yalnızca siz onay verdikten sonra ayarlamaları talimatı verilir. Ajanları <code>--env ote</code> ile başlatın; onları daha da sınırlamak için <code>--read-only</code> veya <code>--no-billable</code> kullanın.
</details>

<details>
<summary><b>Para harcamadan deneyebilir miyim?</b></summary>
<br>
Evet. <code>--env ote</code>, kendine ait API anahtarı ve bakiyesi olan OTE test ortamını kullanır. Orada yapılan hiçbir işlem canlı alan adlarınıza veya gerçek bakiyenize dokunmaz.
</details>

<details>
<summary><b>Hangi yapay zeka araçlarıyla çalışır?</b></summary>
<br>
Herhangi bir MCP istemcisiyle: Claude Code, Claude Desktop, Cursor, VS Code, Windsurf, OpenAI Codex CLI, Gemini CLI ve diğerleri. Yalnızca kabuk erişimi olan ajanlar CLI komutlarını doğrudan çalıştırabilir; <code>dna agent guide</code> ve Claude Code skill'i kullanımı onlara anlatır.
</details>

<details>
<summary><b>Kimlik bilgilerim nerede saklanır?</b></summary>
<br>
<code>dna auth login</code> bunları kullanıcı yapılandırma klasörünüzdeki bir profil dosyasında saklar (konumunu <code>dna config path</code> gösterir); bu dosya Linux ve macOS'ta yalnızca sahibinin erişebileceği izinlerle yazılır. Ortam değişkenleri her zaman dosyadan önceliklidir ve API anahtarı tüm çıktılarda maskelenir.
</details>

<details>
<summary><b>Go gerekli mi?</b></summary>
<br>
Hayır. Kurulum betikleri ve sürüm arşivleri tek bir ikili dosya içerir. Go 1.26+ yalnızca <code>go install</code> ya da kaynak koddan derleme için gereklidir.
</details>

<details>
<summary><b>Kendi kodum için bir kütüphane var mı?</b></summary>
<br>
<code>dna</code> bir CLI ve MCP sunucusudur. API'yi kendi kodunuzdan çağırmak için resmi <a href="https://github.com/domainreseller/php-dna">PHP</a>, <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>, <a href="https://github.com/domainreseller/python-dna">Python</a>, <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> veya <a href="https://github.com/domainreseller/go-dna">Go</a> kütüphanelerini kullanın.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> Dokümantasyon

| rehber | kapsam |
|---|---|
| [Başlarken](docs/getting-started.md) | kurulum, OTE girişi, ilk kayıt ve DNS değişikliği |
| [Yapılandırma](docs/configuration.md) | kimlik bilgileri, ortamlar, profiller, ortam değişkenleri, salt-okunur mod |
| [Kullanım](docs/usage.md) | girdiler, `--body`, çıktı biçimleri, `--jq`, sayfalama, hız sınırları, örnek tarifler |
| [Yapay zeka ajanları ve MCP](docs/ai-agents.md) | istemciye göre MCP kurulumu, onay protokolü, skill, ajan rehberi |
| [Hatalar ve çıkış kodları](docs/errors.md) | hata JSON'u, hata kodları, sorun giderme |
| [Komut referansı](docs/reference/README.md) · [Enum'lar](docs/enums.md) | tüm komutlar, bayraklar ve kodlu değerler (otomatik üretilir) |
| [Mimari](docs/architecture.md) · [Katkıda bulunma](CONTRIBUTING.md) | kodun nasıl düzenlendiği |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> Bağlantılar

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API: alan adı bayilik programı"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="Alan adı bayisi olun"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="Alan adı bayi API'si ve modülleri"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="Destek"></a>
</p>

<p align="right">
  <sub>Bizi takip edin</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · insanlar ve onlarla birlikte çalışan ajanlar için yapıldı</sub></p>
