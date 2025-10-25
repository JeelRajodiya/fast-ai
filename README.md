# Fast AI

A command-line AI for quick questions with acceptable [Time to First Token (TTFT)](https://docs.nvidia.com/nim/benchmarking/llm/latest/metrics.html#time-to-first-token-ttft), powered by the [Groq API](https://groq.com/)

![Screencast_20251021_224055](https://github.com/user-attachments/assets/222f102a-50cf-4bd2-891a-4dc6057d4c6f)

## Installing (Linux)

Download and install with a single command:

```bash
curl -fsSL https://raw.githubusercontent.com/JeelRajodiya/fast-ai/main/install.sh | bash
```

![Screencast_20251021_225003](https://github.com/user-attachments/assets/e71f059f-4a4e-4b6c-a205-f455055cfabb)

> [!TIP]
> You can also use the same command to update the tool to latest version.

## Installing (MacOS/Windows)

Please download the binary from the [Releases](https://github.com/JeelRajodiya/fast-ai/releases) page and add it to your system PATH.

## Usage

Get A Groq API key from: [https://console.groq.com/keys](https://console.groq.com/keys)

### Single Query mode

```bash
ai "How to exit vim?"
```

### Interactive mode

![Screencast_20251021_225624](https://github.com/user-attachments/assets/62fdb0e9-ca02-4bd5-ace2-a33080ece395)

```bash
ai
# Then type your question and press enter
```

## Uninstalling

```bash
curl -fsSL https://raw.githubusercontent.com/JeelRajodiya/fast-ai/main/uninstall.sh | bash
```
