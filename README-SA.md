<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa-active.svg" height="32" alt="العربية"></a>
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

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/sa/hero.svg" width="1200" alt="dna، أداة CLI الخاصة بـ Domain Name API التي يقودها وكلاء الذكاء الاصطناعي: يبحث الوكيل عن نطاق، ويعاين التسجيل المدفوع باستخدام --dry-run، ويطلب الموافقة ثم يسجّله."></a>

<h1 align="center">dna: أداة CLI لـ Domain Name API مخصّصة لوكلاء الذكاء الاصطناعي</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="أحدث إصدار"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  تتحكّم أداة <b>dna</b> في منصة الموزعين التابعة لـ <a href="https://www.domainnameapi.com">Domain Name API</a> من سطر الأوامر وعبر وكلاء الذكاء الاصطناعي.
  وهي تغطي جميع عمليات API البالغ عددها 64 عملية: البحث عن النطاقات، والتسجيل، والتجديد، والنقل، وDNS، وDNSSEC، وجهات الاتصال، وSSL، والاستضافة.
  اربطها بـ <b>Claude Code, Claude Desktop, Cursor, Codex</b> أو بأي عميل MCP، لتصبح كل عملية أداةً آمنة وموصوفة بدقة.
</p>

<img src="./assets/sa/stats.svg" width="1200" alt="64 عملية API، و65 أداة MCP لوكلاء الذكاء الاصطناعي، و4 فئات أمان مُطبَّقة إلزاميًا، و6 منصات من ملف تنفيذي واحد.">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> مصمَّمة لوكلاء الذكاء الاصطناعي

|  | ما يحصل عليه الوكيل |
|---|---|
| 🧰 **أداة لكل عملية** | يوفّر الأمر `dna mcp serve` ما مجموعه 65 أداة MCP مزوّدة بمخططات JSON Schema مبنية من مواصفات API، فتملأ النماذج المدخلات بشكل صحيح من المحاولة الأولى. |
| 🛡️ **بوابة أمان** | لا تُنفَّذ الاستدعاءات المدفوعة والتدميرية من تلقاء نفسها أبدًا. تتطلب واجهة CLI الخيار `--yes`، ويتطلب MCP المعامل `confirm=true`، وبدونهما يحصل الوكيل على معاينة يعرضها عليك. |
| 🧾 **كل شيء قابل للقراءة آليًا** | مخرجات JSON عند تمرير المخرجات عبر أنبوب (pipe)، وأخطاء بصيغة JSON مع رموز خروج ثابتة، و`--dry-run` لعرض الطلب كما هو بالضبط، و`dna schema` لكل مدخل. |
| 🧪 **بيئة اختبار معزولة** | يوجّه الخيار `--env ote` الطلبات إلى بيئة الاختبار: API حقيقية دون أموال حقيقية. ويمنع الخيار `--read-only` أي تغيير. |
| 📘 **توثيق ذاتي** | يُعلّم الأمر `dna agent guide` الوكيلَ واجهة CLI بالكامل في استدعاء واحد. ويضيف الأمر `dna agent skill --install` مهارة (skill) إلى Claude Code. |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> التثبيت

> [!IMPORTANT]
> **ثبّت دائمًا من _أحدث إصدار_.** الملفات التنفيذية للإصدارات مرقّمة بإصدارات، ومُختبرة، ومنشورة مع مجاميع تحقق SHA-256.

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

**باستخدام Go 1.26 أو أحدث**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>التنزيل اليدوي</b></summary>
<br>

اختر الأرشيف المناسب لك من [أحدث إصدار](https://github.com/domainreseller/cli/releases/latest)، وتحقّق منه بمطابقته مع `checksums.txt`، ثم فُكّ ضغطه وضع `dna` (أو `dna.exe`) ضمن `PATH` لديك.

| نظام التشغيل | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

تحقّق من التثبيت باستخدام `dna version` و`dna doctor`.

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> البدء السريع

```bash
# 1. سجّل الدخول إلى بيئة الاختبار أولًا (لها مفتاح API خاص بها)
dna auth login --env ote

# 2. استكشف بحرية: لا شيء هنا يُغيّر أي شيء
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. عاين عملية الشراء، ثم نفّذها (سيُطلب منك التأكيد)
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. بيئة الإنتاج
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> اكتب `dna` فقط: الواجهة التفاعلية

<img src="./assets/sa/shell.svg" width="1200" alt="الواجهة التفاعلية لـ dna: عند كتابة acme تظهر الامتدادات المتاحة مع أسعارها، يختار المستخدم acme.net ويملأ نموذج تسجيل قصيرًا ويؤكد الخطوة المدفوعة في نافذة التأكيد فيُسجَّل النطاق.">

شغّل `dna` وحده في الطرفية. لا تحتاج إلى معرفة أي أمر: كل شيء اختيار بالأسهم ↑↓ و Enter، و Esc يعيدك دائمًا خطوة إلى الوراء.

|  | ما تحصل عليه |
|---|---|
| 🔎 **بحث فوري** | اكتب اسمًا مثل `acme`: تظهر الامتدادات المتاحة وأسعارها أثناء الكتابة. يختار السهم ↓ أحدها ويسجّله عبر نموذج قصير. |
| ⚡ **إجراءات سريعة** | السهم ↓ في سطر فارغ: نطاقاتك، تسجيل نطاق، سجلات DNS، الرصيد، التبديل بين OTE والإنتاج. |
| 🗂️ **قوائم النطاقات** | ابحث في نطاقاتك واضغط Enter: التفاصيل، DNS، التجديد، القفل، خصوصية WHOIS، خوادم الأسماء، رمز النقل. |
| 🛡️ **لا شيء بالخطأ** | كل خطوة مدفوعة أو مدمّرة تنتهي بنافذة تأكيد تعرض الطلب الدقيق والسعر. يبقى خيار الإلغاء محددًا حتى تختار غيره. |
| 🤖 **اربط وكيل ذكاء اصطناعي** | يعرض `/agent` أكثر من 20 وكيلًا: Claude Code و ChatGPT و Cursor و Claude Desktop و Codex وغيرها. يجهّز dna الإعداد ويتحقق من بيانات الاعتماد ويعطيك أوامر جاهزة. |

لا ترى السكربتات ولا وكلاء الذكاء الاصطناعي الواجهة التفاعلية أبدًا: عند التمرير عبر pipe أو التشغيل بواسطة وكيل، يطبع `dna` مخرجات JSON والمساعدة كما في السابق.

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> اربط وكيل الذكاء الاصطناعي الخاص بك

<img src="./assets/sa/mcp.svg" width="1200" alt="Model Context Protocol: يتيح dna mcp serve استخدام 65 أداة مزوّدة بتعليقات توضيحية للأمان. استدعاء أداة مدفوعة دون confirm يُعيد معاينة؛ ومع confirm=true بعد الموافقة يُنفَّذ.">

> [!TIP]
> **الطريقة الأسرع:** شغّل `dna` واضغط ↓ واختر **Connect an AI agent** (أو اكتب `/agent`). اختر واحدًا من أكثر من 20 وكيلًا، وحدد OTE أو الإنتاج ومدى الصلاحيات. يشغّل dna الإعداد أو يضيف نفسه إلى ملف إعداد الوكيل، ويتحقق من بيانات الاعتماد وينسخ لك الأوامر الجاهزة.

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # مقتطفات الإعداد لجميع العملاء الآخرين
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
<summary><b>الوكلاء الذين يعملون عبر shell (المهارة والدليل)</b></summary>
<br>

```bash
dna agent skill --install                       # مهارة Claude Code: ~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # مهارة على مستوى المشروع، أضفها (commit) إلى مستودعك
dna agent guide                                 # دليل استخدام موجز لأي وكيل
```
</details>

أنت من يحدد مدى ما يمكن للوكيل الوصول إليه:

```bash
dna mcp serve --read-only            # اطّلاع فقط دون أي تعديل
dna mcp serve --no-billable          # كل شيء ما عدا عمليات الشراء
dna mcp serve --groups domain,dns    # مجموعات الأوامر هذه فقط
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> موجّهات جاهزة لـ Claude وChatGPT

انسخ موجّهًا، والصقه في مساعد الذكاء الاصطناعي لديك، واتركه ينجز العمل. يبدأ كل موجّه في بيئة الاختبار OTE، ويتوقف لانتظار موافقتك قبل أي عملية مدفوعة.

**إعداد dna في Claude Code**

<div dir="rtl">

```text
قم بإعداد أداة dna CLI لي (https://github.com/domainreseller/cli).
1. ثبّتها باستخدام سكربت التثبيت الرسمي الخاص بنظام التشغيل لديّ، ثم شغّل
   ‎`dna version`‎ و‎`dna doctor`‎.
2. ثبّت مهارة Claude Code باستخدام ‎`dna agent skill --install`‎، وأضف خادم MCP باستخدام
   ‎`claude mcp add dna -- dna mcp serve --env ote`‎.
3. لا تطلب مني لصق مفتاح API الخاص بي في هذه المحادثة. اطلب مني تشغيل
   ‎`dna auth login --env ote`‎ بنفسي، ثم تحقّق من الاتصال باستخدام
   ‎`dna --ote account get`‎.
اعمل في بيئة الاختبار OTE فقط، ولا تستخدم ‎`--yes`‎ أبدًا دون موافقتي.
```

</div>

**إعداد dna في ChatGPT Codex**

<div dir="rtl">

```text
قم بإعداد أداة dna CLI لي (https://github.com/domainreseller/cli).
1. ثبّتها باستخدام سكربت التثبيت الرسمي الخاص بنظام التشغيل لديّ، ثم شغّل
   ‎`dna version`‎ و‎`dna doctor`‎.
2. أضفها إلى ملف إعدادات Codex لديّ (‎`~/.codex/config.toml`‎) كخادم MCP باسم ‎`dna`‎، على
   أن يكون الأمر ‎`dna`‎ والوسائط ‎`mcp serve --env ote`‎.
3. لا تطلب مني لصق مفتاح API الخاص بي في هذه المحادثة. اطلب مني تشغيل
   ‎`dna auth login --env ote`‎ بنفسي، ثم تحقّق من الاتصال باستخدام
   ‎`dna --ote account get`‎.
اعمل في بيئة الاختبار OTE فقط، ولا تستخدم ‎`--yes`‎ أبدًا دون موافقتي.
```

</div>

**لا تتوفر لديك طرفية؟ اسأل ChatGPT أو Claude عبر الدردشة**

<div dir="rtl">

```text
اقرأ الملف https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt وصفحات
توثيق dna التي يحيل إليها. ثم أعطني أوامر dna الدقيقة، خطوةً بخطوة، لتسجيل نطاق في
بيئة الاختبار OTE: التثبيت، وتسجيل الدخول، والبحث، والمعاينة باستخدام ‎`--dry-run`‎،
والتسجيل. واشرح ما يفعله كل أمر.
```

</div>

<details>
<summary dir="rtl"><b>المزيد من الموجّهات للمهام اليومية</b></summary>
<br>

**البحث عن نطاق وتسجيله**

<div dir="rtl">

```text
باستخدام dna، تحقّق مما إذا كانت النطاقات brandname.com وbrandname.net وbrandname.io
متاحة. اعرض النتائج مع أسعارها في جدول، ثم عاين تسجيل أرخص نطاق متاح منها باستخدام
‎`--dry-run`‎. لا تسجّل أي شيء حتى أقول نعم.
```

</div>

**التجديدات القادمة**

<div dir="rtl">

```text
باستخدام dna، اعرض قائمة بنطاقاتي التي تنتهي صلاحيتها خلال الأيام الـ60 القادمة مع
أسعار تجديدها، مرتبةً حسب تاريخ الانتهاء. لا تجدّد أي شيء.
```

</div>

**توجيه نطاق إلى خادم**

<div dir="rtl">

```text
باستخدام dna، وجّه example.com إلى 203.0.113.10: أضف سجلات A لـ ‎`@`‎ و‎`www`‎، ثم اعرض لي
منطقة DNS النهائية. اسألني قبل تعديل أي سجل موجود أو حذفه.
```

</div>

**نقل نطاق إلى حسابك**

<div dir="rtl">

```text
باستخدام dna، تحقّق مما إذا كان يمكن نقل example.org إلى حسابي، وأخبرني بما لا يزال
ينقصني (رمز التفويض auth code، وإلغاء القفل)، وعاين عملية النقل باستخدام ‎`--dry-run`‎.
```

</div>

**طلب شهادة SSL**

<div dir="rtl">

```text
باستخدام dna، ابحث عن أرخص شهادة DV SSL لـ example.com، وأنشئ CSR، وعاين الطلب. انتظر
موافقتي قبل الشراء.
```

</div>

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> آلية العمل

<img src="./assets/sa/flow.svg" width="1200" alt="يستدعي وكيل الذكاء الاصطناعي، أو أنت، أداة dna عبر MCP أو CLI؛ فتتحقق dna من صحة المدخلات، وتطبّق بوابة الأمان، وتتواصل مع Domain Name API (بيئة الإنتاج أو بيئة OTE التجريبية)، التي تصل بدورها إلى السجلات (registries) عبر EPP.">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> بوابة الأمان

<img src="./assets/sa/safety.svg" width="1200" alt="تُنفَّذ أوامر القراءة والكتابة فورًا؛ أما الأوامر التدميرية والمدفوعة فتتوقف إلى أن يتم تأكيدها باستخدام --yes في CLI أو confirm=true عبر MCP.">

| الفئة | أمثلة | السلوك |
|---|---|---|
| `read` | البحث، القوائم، المعلومات، الأسعار | يُنفَّذ فورًا |
| `write` | سجلات DNS، القفل، الخصوصية، جهات الاتصال | يُنفَّذ فورًا (مجاني وقابل للتراجع) |
| `destructive` | حذف سجل، إلغاء شهادة SSL، تعطيل DNSSEC | يتوقف حتى يتم التأكيد |
| `billable` | التسجيل، التجديد، النقل، الاستعادة، شراء SSL أو استضافة | يتوقف حتى يتم التأكيد |

- **في سكربت أو في shell الخاص بالوكيل:** بدون `--yes`، يخرج الأمر الخاضع للبوابة برمز **9** دون أن يرسل أي شيء.
- **عبر MCP:** بدون `confirm=true`، تُعيد الأداة معاينة للطلب كما هو بالضبط.
- **المعاينة:** يطبع الخيار `--dry-run` طلب HTTP مع إخفاء مفتاح API.
- **الإقفال التام:** يمنع `--read-only` أو `DNA_READ_ONLY=1` أو ملف تعريف للقراءة فقط أي تغيير (رمز الخروج **10**).

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> المخرجات والأخطاء ورموز الخروج

```bash
dna domain info example.com                   # جدول في الطرفية
dna domain info example.com -o json           # JSON (الافتراضي عند التمرير عبر أنبوب)
dna domain list --all -o jsonl                # نطاق واحد في كل سطر
dna domain info example.com --jq .authCode    # أداة jq مدمجة
dna dns add --body @record.json               # المدخلات كاملة من ملف، أو --body - للقراءة من stdin
dna schema domain.register                    # مخطط JSON Schema لأي مدخل
dna enums DomainStatus                        # معنى الحالة 7
```

تُكتب الأخطاء بصيغة JSON على stderr: `{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| الرمز | المعنى | الرمز | المعنى |
|---:|---|---:|---|
| 0 | نجاح | 6 | تجاوز حدّ معدل الطلبات |
| 1 | خطأ غير متوقع | 7 | خطأ في الخادم |
| 2 | مدخلات غير صالحة، ولم يُرسَل شيء | 8 | الشبكة |
| 3 | المصادقة | 9 | التأكيد مطلوب |
| 4 | غير موجود | 10 | وضع القراءة فقط |
| 5 | رفضته API | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> الأوامر

| المجموعة | الوظيفة |
|---|---|
| `dna` | الواجهة التفاعلية: بحث فوري، قوائم، نماذج موجّهة، إعداد وكلاء الذكاء الاصطناعي |
| `dna domain` | البحث، البحث المجمّع، القوائم، المعلومات، التسجيل، التجديد، النقل، الاستعادة، القفل، الخصوصية، رمز التفويض (auth code)، جهات الاتصال، المستندات، إعادة التوجيه، خوادم الأسماء، المضيفات الفرعية (child hosts) |
| `dna dns` · `dna dnssec` | سجلات المنطقة (zone)، سجلات DS |
| `dna contact` | معرّفات جهات الاتصال (handles) والتحقق منها |
| `dna ssl` | CSR، الشراء، الإعداد، الحالة، التنزيل، إعادة الإصدار، الإلغاء |
| `dna hosting` · `dna product` · `dna order` · `dna account` | الاستضافة، الكتالوج والأسعار، الطلبات، الرصيد والمعاملات |
| `dna api` | طلب خام إلى أي نقطة نهاية (endpoint) |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | للوكلاء والأدوات |
| `dna auth` · `dna config` · `dna doctor` | بيانات الاعتماد، ملفات التعريف، التشخيص |

يعرض الخيار `--help` لكل أمر فئة الأمان الخاصة به ومسار API وأمثلة على الاستخدام. راجع [مرجع الأوامر](docs/reference/README.md) الكامل.

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> بيانات اعتماد API

تحتاج `dna` إلى القيمتين نفسيهما اللتين تحتاجهما كل وحدة من وحدات Domain Name API:

| القيمة | ماهيتها |
|---|---|
| **Reseller ID** | معرّف الموزع، وهو UUID بالشكل `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` |
| **API Key** | مفتاح API الخاص بك؛ ولكلٍّ من بيئة الإنتاج وبيئة OTE مفاتيح **مختلفة** |

> [!TIP]
> ستجد القيمتين كلتيهما في لوحة Domain Name API الخاصة بك ضمن **API Settings** (إعدادات API). احفظهما باستخدام `dna auth login`، أو استخدم متغيرَي البيئة `DNA_RESELLER_ID` و`DNA_API_KEY` في CI.

> [!CAUTION]
> هذه **بيانات اعتماد API**. **لن** يعمل هنا البريد الإلكتروني وكلمة المرور اللذان تستخدمهما لتسجيل الدخول إلى اللوحة.

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> الأسئلة الشائعة

<details>
<summary><b>هل من الآمن السماح لوكيل ذكاء اصطناعي باستخدام dna؟</b></summary>
<br>
العمليات المدفوعة والتدميرية محمية ببوابة أمان: ترفضها واجهة CLI بدون <code>--yes</code>، ويُعيد خادم MCP معاينة ما لم يتضمن الاستدعاء <code>confirm=true</code>، وهو ما يُوجَّه الوكلاء إلى تعيينه فقط بعد موافقتك. ابدأ تشغيل الوكلاء على <code>--env ote</code>، واستخدم <code>--read-only</code> أو <code>--no-billable</code> لتقييدهم أكثر.
</details>

<details>
<summary><b>هل يمكنني تجربتها دون إنفاق أي أموال؟</b></summary>
<br>
نعم. يستخدم الخيار <code>--env ote</code> بيئة الاختبار OTE بمفتاح API ورصيد خاصين بها. ولا شيء هناك يمسّ النطاقات الفعلية أو رصيدك الحقيقي.
</details>

<details>
<summary><b>ما أدوات الذكاء الاصطناعي التي تعمل معها؟</b></summary>
<br>
أي عميل MCP، مثل: Claude Code، وClaude Desktop، وCursor، وVS Code، وWindsurf، وOpenAI Codex CLI، وGemini CLI وغيرها. أما الوكلاء الذين لا يملكون سوى shell فيمكنهم استدعاء CLI مباشرةً؛ إذ يشرحها لهم <code>dna agent guide</code> ومهارة Claude Code.
</details>

<details>
<summary><b>أين تُخزَّن بيانات الاعتماد الخاصة بي؟</b></summary>
<br>
يحفظها الأمر <code>dna auth login</code> في ملف تعريف داخل مجلد إعدادات المستخدم لديك (يعرض <code>dna config path</code> موقعه)، ويُكتب هذا الملف بصلاحيات تقتصر على المالك وحده في Linux وmacOS. وتتقدّم متغيرات البيئة دائمًا على الملف، ويُخفى مفتاح API في جميع المخرجات.
</details>

<details>
<summary><b>هل أحتاج إلى Go؟</b></summary>
<br>
لا. تحتوي سكربتات التثبيت وأرشيفات الإصدارات على ملف تنفيذي واحد. ولا يلزم Go 1.26 أو أحدث إلا عند استخدام <code>go install</code> أو البناء من الشيفرة المصدرية.
</details>

<details>
<summary><b>هل توجد مكتبة لاستخدامها في الشيفرة الخاصة بي؟</b></summary>
<br>
تُعدّ <code>dna</code> أداة CLI وخادم MCP. لاستدعاء API من الشيفرة البرمجية، استخدم المكتبات الرسمية الخاصة بـ <a href="https://github.com/domainreseller/php-dna">PHP</a> أو <a href="https://github.com/domainreseller/nodejs-dna">Node.js</a> أو <a href="https://github.com/domainreseller/python-dna">Python</a> أو <a href="https://github.com/domainreseller/ruby-dna">Ruby</a> أو <a href="https://github.com/domainreseller/go-dna">Go</a>.
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> التوثيق

| الدليل | ما يتناوله |
|---|---|
| [البدء](docs/getting-started.md) | التثبيت، وتسجيل الدخول إلى OTE، وأول تسجيل لنطاق وأول تغيير في DNS |
| [الإعداد](docs/configuration.md) | بيانات الاعتماد، البيئات، ملفات التعريف، متغيرات البيئة، وضع القراءة فقط |
| [الاستخدام](docs/usage.md) | المدخلات، `--body`، صيغ المخرجات، `--jq`، ترقيم الصفحات، حدود معدل الطلبات، وصفات جاهزة |
| [وكلاء الذكاء الاصطناعي وMCP](docs/ai-agents.md) | إعداد MCP لكل عميل، بروتوكول التأكيد، المهارة، دليل الوكيل |
| [الأخطاء ورموز الخروج](docs/errors.md) | صيغة JSON للأخطاء، رموز الأخطاء، استكشاف المشكلات وإصلاحها |
| [مرجع الأوامر](docs/reference/README.md) · [القيم المُعدَّدة (Enums)](docs/enums.md) | كل أمر وخيار وقيمة مرمّزة (مُولَّد تلقائيًا) |
| [البنية المعمارية](docs/architecture.md) · [المساهمة](CONTRIBUTING.md) | كيفية تنظيم الشيفرة |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> تواصل معنا

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API: برنامج موزعي النطاقات"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="كن موزعًا للنطاقات"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="واجهة API والوحدات لموزعي النطاقات"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="الدعم"></a>
</p>

<p align="right">
  <sub>تابعنا</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · صُنعت للناس وللوكلاء الذين يعملون معهم</sub></p>
