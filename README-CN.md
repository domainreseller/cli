<p align="center">
  <a href="README-TR.md"><img src="./assets/lang/tr.svg" height="32" alt="Türkçe"></a>
  <a href="README.md"><img src="./assets/lang/en.svg" height="32" alt="English"></a>
  <a href="README-DE.md"><img src="./assets/lang/de.svg" height="32" alt="Deutsch"></a>
  <a href="README-SA.md"><img src="./assets/lang/sa.svg" height="32" alt="العربية"></a>
  <a href="README-NL.md"><img src="./assets/lang/nl.svg" height="32" alt="Nederlands"></a>
  <a href="README-AZ.md"><img src="./assets/lang/az.svg" height="32" alt="Azərbaycanca"></a>
  <a href="README-CN.md"><img src="./assets/lang/cn-active.svg" height="32" alt="中文"></a>
  <a href="README-FR.md"><img src="./assets/lang/fr.svg" height="32" alt="Français"></a>
  <a href="README-IT.md"><img src="./assets/lang/it.svg" height="32" alt="Italiano"></a>
</p>

<!--
  dna: the official Domain Name API command line client and MCP server, built for AI agents and developers.
  Every animation on this page is a hand-made SVG in ./assets. No external widgets, no trackers.
-->

<a href="https://github.com/domainreseller/cli/releases/latest"><img src="./assets/cn/hero.svg" width="1200" alt="dna，由 AI 智能体驱动的 Domain Name API CLI：智能体搜索一个域名，用 --dry-run 预览需付费的注册，请求批准后完成注册。"></a>

<h1 align="center">dna：面向 AI 智能体的 Domain Name API CLI</h1>

<p align="center">
  <a href="https://github.com/domainreseller/cli/releases/latest"><img src="https://img.shields.io/github/v/release/domainreseller/cli?label=version&style=flat-square&color=12B385" alt="最新版本"></a>
  <a href="https://github.com/domainreseller/cli/actions/workflows/ci.yml"><img src="https://img.shields.io/github/actions/workflow/status/domainreseller/cli/ci.yml?branch=main&label=CI&style=flat-square" alt="CI"></a>
  <a href="LICENSE"><img src="https://img.shields.io/badge/license-Apache--2.0-0B8DB5?style=flat-square" alt="Apache-2.0"></a>
</p>

<p align="center">
  <b>dna</b> 让你可以通过命令行和 AI 智能体操作 <a href="https://www.domainnameapi.com">Domain Name API</a> 经销商平台。
  它覆盖全部 64 个 API 操作：域名搜索、注册、续费、转移、DNS、DNSSEC、联系人、SSL 和主机。
  将它接入 <b>Claude Code、Claude Desktop、Cursor、Codex</b> 或任意 MCP 客户端，每个操作都会成为一个安全且描述清晰的工具。
</p>

<img src="./assets/cn/stats.svg" width="1200" alt="64 个 API 操作，65 个面向 AI 智能体的 MCP 工具，4 个强制执行的安全类别，一个二进制文件支持 6 个平台。">

## <img src="./assets/h-agent.svg" width="30" height="30" alt=""> 专为 AI 智能体打造

|  | 智能体能获得什么 |
|---|---|
| 🧰 **每个操作对应一个工具** | `dna mcp serve` 提供 65 个 MCP 工具，其 JSON Schema 均根据 API 规范生成，模型第一次就能正确填写输入。 |
| 🛡️ **安全闸门** | 需付费和破坏性的调用绝不会自行执行。CLI 需要 `--yes`，MCP 需要 `confirm=true`；缺少时，智能体只会得到一份可以展示给你的预览。 |
| 🧾 **一切皆可机器读取** | 通过管道输出时为 JSON，错误同样是 JSON 并带有稳定的退出码，`--dry-run` 显示确切的请求，`dna schema` 提供每个输入的 Schema。 |
| 🧪 **沙盒环境** | `--env ote` 指向测试环境：真实的 API，不花真钱。`--read-only` 阻止一切更改。 |
| 📘 **自带文档** | `dna agent guide` 只需一次调用即可让智能体掌握整个 CLI。`dna agent skill --install` 会添加一个 Claude Code 技能（skill）。 |

## <img src="./assets/h-install.svg" width="30" height="30" alt=""> 安装

> [!IMPORTANT]
> **请始终从 _最新版本_ 安装。** 发布版二进制文件均带有版本号、经过测试，并随附 SHA-256 校验和发布。

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

**使用 Go 1.26+**

```bash
go install github.com/domainreseller/cli/cmd/dna@latest
```

<details>
<summary><b>手动下载</b></summary>
<br>

在 [最新版本](https://github.com/domainreseller/cli/releases/latest) 页面选择对应的压缩包，用 `checksums.txt` 校验，解压后将 `dna`（或 `dna.exe`）放到你的 `PATH` 中。

| 操作系统 | x64 | ARM |
|---|---|---|
| Linux | `dna_<version>_linux_amd64.tar.gz` | `dna_<version>_linux_arm64.tar.gz` |
| Linux (.rpm) | `dna_linux_amd64.rpm` | `dna_linux_arm64.rpm` |
| Linux (.deb) | `dna_linux_amd64.deb` | `dna_linux_arm64.deb` |
| macOS | `dna_<version>_darwin_amd64.tar.gz` | `dna_<version>_darwin_arm64.tar.gz` |
| Windows | `dna_<version>_windows_amd64.zip` | `dna_<version>_windows_arm64.zip` |

</details>

使用 `dna version` 和 `dna doctor` 检查安装情况。

## <img src="./assets/h-terminal.svg" width="30" height="30" alt=""> 快速开始

```bash
# 1. 先登录测试环境（它有单独的 API Key）
dna auth login --env ote

# 2. 先随便看看：这里的命令不会更改任何内容
dna --ote account get
dna --ote domain search example.com
dna --ote product tlds --filter-text com

# 3. 先预览购买，再正式执行（会要求确认）
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE> --dry-run
dna --ote domain register example.com --ns tr.apiname.com,eu.apiname.com --handle <HANDLE>

# 4. 正式环境
dna auth login
dna domain list
```

<a id="interactive-shell"></a>

## <img src="./assets/h-shell.svg" width="30" height="30" alt=""> 只需输入 `dna`：交互式终端

<img src="./assets/cn/shell.svg" width="1200" alt="dna 交互式终端：输入 acme 即显示可注册的后缀及价格，用户选择 acme.net，填写简短的注册表单，在确认对话框中确认付费步骤，域名注册完成。">

在终端中单独运行 `dna`。你无需记住任何命令：一切都用 ↑↓ 和 Enter 选择，Esc 总是返回上一步。

|  | 你能得到什么 |
|---|---|
| 🔎 **实时搜索** | 输入一个名称，例如 `acme`：可注册的后缀及其价格会在输入时出现。用 ↓ 选择一个，通过简短的表单完成注册。 |
| ⚡ **快捷操作** | 在空行按 ↓：我的域名、注册域名、DNS 记录、余额、在 OTE 与生产环境之间切换。 |
| 🗂️ **域名菜单** | 搜索你的域名并按 Enter：详情、DNS、续费、锁定、WHOIS 隐私、域名服务器、转移授权码。 |
| 🛡️ **不会误操作** | 每个付费或破坏性步骤都以确认对话框结束，显示确切的请求和价格。在你作出其他选择之前，“取消”始终处于选中状态。 |
| 🤖 **连接 AI 智能体** | `/agent` 列出 20 多个智能体：Claude Code、ChatGPT、Cursor、Claude Desktop、Codex 等。dna 会完成配置、检查凭据并提供现成的提示词。 |

脚本和 AI 智能体永远不会看到交互式终端：通过管道或由智能体运行时，`dna` 仍像以前一样输出 JSON 和帮助信息。

<a id="connect-your-ai-agent"></a>

## <img src="./assets/h-mcp.svg" width="30" height="30" alt=""> 连接你的 AI 智能体

<img src="./assets/cn/mcp.svg" width="1200" alt="Model Context Protocol：dna mcp serve 提供 65 个带安全注解的工具。不带 confirm 的付费工具调用会返回预览；获得批准后带上 confirm=true 才会执行。">

> [!TIP]
> **最快的方法：** 运行 `dna`，按 ↓ 并选择 **Connect an AI agent**（或输入 `/agent`）。从 20 多个智能体中选择一个，选择 OTE 或生产环境以及它的权限范围。dna 会执行配置或把自己加入智能体的配置文件，检查凭据，并为你复制现成的提示词。

```bash
claude mcp add dna -- dna mcp serve --env ote      # Claude Code
dna mcp config                                     # 其他所有客户端的配置片段
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
<summary><b>可使用 shell 的智能体（技能与指南）</b></summary>
<br>

```bash
dna agent skill --install                       # Claude Code 技能：~/.claude/skills/dna/SKILL.md
dna agent skill --install --dir .claude/skills  # 项目技能，随仓库一起提交
dna agent guide                                 # 适用于任何智能体的精简使用指南
```
</details>

由你决定智能体的权限范围：

```bash
dna mcp serve --read-only            # 只看不动
dna mcp serve --no-billable          # 除购买外的所有操作
dna mcp serve --groups domain,dns    # 仅限这些命令组
```

<a id="ready-made-prompts"></a>

## <img src="./assets/h-prompt.svg" width="30" height="30" alt=""> 适用于 Claude 和 ChatGPT 的现成提示词

复制一条提示词，粘贴到您的 AI 助手中，让它替您完成工作。每条提示词都从 OTE 测试环境开始，并会在任何付费操作之前停下来等待您的确认。

**在 Claude Code 中设置 dna**

```text
请帮我安装配置 dna CLI (https://github.com/domainreseller/cli)。
1. 使用适合我操作系统的官方安装脚本安装，然后运行 `dna version` 和 `dna doctor`。
2. 用 `dna agent skill --install` 安装 Claude Code 技能，并用
   `claude mcp add dna -- dna mcp serve --env ote` 添加 MCP 服务器。
3. 不要要求我把 API 密钥粘贴到这个对话中。提示我自己运行 `dna auth login --env ote`，
   然后用 `dna --ote account get` 检查连接。
只在 OTE 测试环境中操作，未经我同意，绝不要使用 `--yes`。
```

**在 ChatGPT Codex 中设置 dna**

```text
请帮我安装配置 dna CLI (https://github.com/domainreseller/cli)。
1. 使用适合我操作系统的官方安装脚本安装，然后运行 `dna version` 和 `dna doctor`。
2. 把它作为 MCP 服务器添加到我的 Codex 配置 (`~/.codex/config.toml`) 中：名称为 `dna`，
   命令为 `dna`，参数为 `mcp serve --env ote`。
3. 不要要求我把 API 密钥粘贴到这个对话中。提示我自己运行 `dna auth login --env ote`，
   然后用 `dna --ote account get` 检查连接。
只在 OTE 测试环境中操作，未经我同意，绝不要使用 `--yes`。
```

**没有终端？在 ChatGPT 或 Claude 聊天中提问**

```text
请阅读 https://raw.githubusercontent.com/domainreseller/cli/main/llms.txt 及其中链接的
dna 文档。然后一步一步地给出在 OTE 测试环境中注册域名的确切 dna 命令：安装、登录、查询、
用 `--dry-run` 预览，以及注册。解释每条命令的作用。
```

<details>
<summary><b>更多日常工作提示词</b></summary>
<br>

**查询并注册域名**

```text
用 dna 检查 brandname.com、brandname.net 和 brandname.io 是否可以注册。用表格列出结果
和价格，然后用 `--dry-run` 预览注册其中最便宜的可注册域名。在我同意之前，什么都不要注
册。
```

**近期续费**

```text
用 dna 列出我未来 60 天内到期的域名及其续费价格，按到期日期排序。什么都不要续费。
```

**将域名指向服务器**

```text
用 dna 将 example.com 指向 203.0.113.10：为 `@` 和 `www` 添加 A 记录，然后给我看最终的
DNS 区域。修改或删除任何现有记录之前，先问我。
```

**转入域名**

```text
用 dna 检查 example.org 能否转入我的账户，告诉我还缺什么（授权码、解锁），并用
`--dry-run` 预览转入。
```

**订购 SSL 证书**

```text
用 dna 为 example.com 找到最便宜的 DV SSL 证书，生成 CSR 并预览订单。购买前等待我的确
认。
```

</details>

## <img src="./assets/h-flow.svg" width="30" height="30" alt=""> 工作原理

<img src="./assets/cn/flow.svg" width="1200" alt="AI 智能体或你本人通过 MCP 或 CLI 调用 dna；dna 校验输入、执行安全闸门，并与 Domain Name API（正式环境或 OTE 沙盒）通信，后者再通过 EPP 与各注册局交互。">

## <img src="./assets/h-shield.svg" width="30" height="30" alt=""> 安全闸门

<img src="./assets/cn/safety.svg" width="1200" alt="读取和写入命令会立即执行；破坏性和付费命令会暂停，直到在 CLI 中用 --yes 或通过 MCP 用 confirm=true 确认为止。">

| 类别 | 示例 | 行为 |
|---|---|---|
| `read` | 搜索、列表、信息、价格 | 立即执行 |
| `write` | DNS 记录、锁定、隐私保护、联系人 | 立即执行（免费、可撤销） |
| `destructive` | 删除记录、吊销 SSL、禁用 DNSSEC | 确认前暂停 |
| `billable` | 注册、续费、转移、恢复、购买 SSL/主机 | 确认前暂停 |

- **在脚本或智能体 shell 中：** 没有 `--yes` 时，受管控的命令会以退出码 **9** 退出，且不发送任何请求。
- **通过 MCP：** 没有 `confirm=true` 时，工具会返回确切请求的预览。
- **预览：** `--dry-run` 会打印 HTTP 请求，其中 API Key 已做掩码处理。
- **完全锁定：** `--read-only`、`DNA_READ_ONLY=1` 或只读配置文件会阻止一切更改（退出码 **10**）。

## <img src="./assets/h-output.svg" width="30" height="30" alt=""> 输出、错误与退出码

```bash
dna domain info example.com                   # 在终端中显示为表格
dna domain info example.com -o json           # JSON（通过管道输出时的默认格式）
dna domain list --all -o jsonl                # 每行一个域名
dna domain info example.com --jq .authCode    # 内置 jq
dna dns add --body @record.json               # 从文件读取完整输入，或用 --body - 从 stdin 读取
dna schema domain.register                    # 任意输入的 JSON Schema
dna enums DomainStatus                        # 查看状态 7 的含义
```

错误以 JSON 形式输出到 stderr：`{"error":{"kind","code","message","status","fields","hint","correlationId","retryable","exitCode"}}`

| 代码 | 含义 | 代码 | 含义 |
|---:|---|---:|---|
| 0 | 成功 | 6 | 触发速率限制 |
| 1 | 意外错误 | 7 | 服务器错误 |
| 2 | 输入无效，未发送任何请求 | 8 | 网络错误 |
| 3 | 身份验证失败 | 9 | 需要确认 |
| 4 | 未找到 | 10 | 只读模式 |
| 5 | 被 API 拒绝 | | |

## <img src="./assets/h-commands.svg" width="30" height="30" alt=""> 命令

| 命令组 | 功能 |
|---|---|
| `dna` | 交互式终端：实时搜索、菜单、引导式表单、AI 智能体配置 |
| `dna domain` | 搜索、批量搜索、列表、信息、注册、续费、转移、恢复、锁定、隐私保护、授权码、联系人、文档、转发、名称服务器、子名称服务器 |
| `dna dns` · `dna dnssec` | 区域记录、DS 记录 |
| `dna contact` | 联系人句柄（handle）与验证 |
| `dna ssl` | CSR、购买、配置、状态、下载、重新签发、吊销 |
| `dna hosting` · `dna product` · `dna order` · `dna account` | 主机、产品目录与价格、订单、余额与交易记录 |
| `dna api` | 向任意端点发送原始请求 |
| `dna schema` · `dna enums` · `dna agent` · `dna mcp` | 供智能体和工具使用 |
| `dna auth` · `dna config` · `dna doctor` | 凭据、配置文件、诊断 |

每个命令的 `--help` 都会显示其安全类别、API 路由和示例。完整内容请参阅 [命令参考](docs/reference/README.md)。

## <img src="./assets/h-key.svg" width="30" height="30" alt=""> API 凭据

`dna` 需要的两个值与所有 Domain Name API 模块相同：

| 值 | 说明 |
|---|---|
| **Reseller ID** | 一个 UUID，形如 `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx` |
| **API Key** | 你的 API 密钥；正式环境和 OTE 使用**不同**的密钥 |

> [!TIP]
> 这两个值都可以在 Domain Name API 面板的 **API Settings**（API 设置）中找到。使用 `dna auth login` 保存它们，或在 CI 中使用 `DNA_RESELLER_ID` 和 `DNA_API_KEY` 环境变量。

> [!CAUTION]
> 这些是 **API 凭据**。你的面板登录邮箱和密码在这里**无法**使用。

## <img src="./assets/h-faq.svg" width="30" height="30" alt=""> 常见问题

<details>
<summary><b>让 AI 智能体使用 dna 安全吗？</b></summary>
<br>
需付费和破坏性的操作都受到管控：没有 <code>--yes</code> 时 CLI 会拒绝执行；除非调用中带有 <code>confirm=true</code>，否则 MCP 服务器只会返回预览，而智能体被要求仅在你批准后才设置该参数。建议先让智能体在 <code>--env ote</code> 上运行，并使用 <code>--read-only</code> 或 <code>--no-billable</code> 进一步限制其权限。
</details>

<details>
<summary><b>可以不花钱试用吗？</b></summary>
<br>
可以。<code>--env ote</code> 使用 OTE 测试环境，它有独立的 API Key 和余额。在那里的任何操作都不会影响真实域名或你的真实余额。
</details>

<details>
<summary><b>哪些 AI 工具可以使用它？</b></summary>
<br>
任何 MCP 客户端都可以：Claude Code、Claude Desktop、Cursor、VS Code、Windsurf、OpenAI Codex CLI、Gemini CLI 等。只有 shell 的智能体可以直接调用 CLI；<code>dna agent guide</code> 和 Claude Code 技能会向它们说明用法。
</details>

<details>
<summary><b>我的凭据存储在哪里？</b></summary>
<br>
<code>dna auth login</code> 会将它们保存在用户配置目录下的配置文件中（<code>dna config path</code> 会显示具体位置），在 Linux 和 macOS 上以仅所有者可访问的权限写入。环境变量始终优先于该文件，且 API Key 在所有输出中都会被掩码处理。
</details>

<details>
<summary><b>我需要安装 Go 吗？</b></summary>
<br>
不需要。安装脚本和发布压缩包中只包含一个二进制文件。只有使用 <code>go install</code> 或从源码构建时才需要 Go 1.26+。
</details>

<details>
<summary><b>有没有可以在自己代码中使用的库？</b></summary>
<br>
<code>dna</code> 是一个 CLI 和 MCP 服务器。如需在代码中调用 API，请使用官方的 <a href="https://github.com/domainreseller/php-dna">PHP</a>、<a href="https://github.com/domainreseller/nodejs-dna">Node.js</a>、<a href="https://github.com/domainreseller/python-dna">Python</a>、<a href="https://github.com/domainreseller/ruby-dna">Ruby</a> 或 <a href="https://github.com/domainreseller/go-dna">Go</a> 库。
</details>

## <img src="./assets/h-docs.svg" width="30" height="30" alt=""> 文档

| 指南 | 内容 |
|---|---|
| [入门指南](docs/getting-started.md) | 安装、OTE 登录、首次注册和 DNS 修改 |
| [配置](docs/configuration.md) | 凭据、环境、配置文件、环境变量、只读模式 |
| [用法](docs/usage.md) | 输入、`--body`、输出格式、`--jq`、分页、速率限制、实用示例 |
| [AI 智能体与 MCP](docs/ai-agents.md) | 各客户端的 MCP 配置、确认协议、技能、智能体指南 |
| [错误与退出码](docs/errors.md) | 错误 JSON、错误代码、故障排查 |
| [命令参考](docs/reference/README.md) · [枚举](docs/enums.md) | 所有命令、选项和编码值（自动生成） |
| [架构](docs/architecture.md) · [贡献指南](CONTRIBUTING.md) | 代码的组织方式 |

## <img src="./assets/h-connect.svg" width="30" height="30" alt=""> 联系我们

<p align="center">
  <a href="https://www.domainnameapi.com"><img src="./assets/btn-web.svg" height="46" alt="Domain Name API：域名经销商计划"></a>
  <a href="https://www.domainnameapi.com/domain-reseller"><img src="./assets/btn-signup.svg" height="46" alt="成为域名经销商"></a>
  <a href="https://www.domainnameapi.com/api-integration"><img src="./assets/btn-api.svg" height="46" alt="域名经销商 API 与模块"></a>
  <a href="https://www.domainnameapi.com/contacts"><img src="./assets/btn-support.svg" height="46" alt="技术支持"></a>
</p>

<p align="right">
  <sub>关注我们</sub><br>
  <a href="https://www.atakdomain.com"><img src="./assets/pill-atak.svg" height="34" alt="atakdomain.com"></a>
  <a href="https://www.linkedin.com/company/domain-name-reseller/"><img src="./assets/pill-linkedin.svg" height="34" alt="LinkedIn"></a>
  <a href="https://x.com/dna_resellerapi"><img src="./assets/pill-x.svg" height="34" alt="X"></a>
  <a href="https://www.youtube.com/DomainReseller"><img src="./assets/pill-youtube.svg" height="34" alt="YouTube"></a>
</p>

<p align="center"><sub>Apache-2.0 · 为人们以及与他们协作的智能体而打造</sub></p>
