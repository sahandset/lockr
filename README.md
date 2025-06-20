# Lockr 🔒

**Lockr** is a fast, pluggable CLI tool written in Go to scan `.env`, `.yaml`, `.json`, and other config files for hardcoded secrets and insecure values.

It supports built-in and custom rules, redacted output, machine-readable results, and CI integration.

---

## ✨ Features

- Scans config files for secrets using regex-based detection
- Load custom rules from JSON or YAML
- `--test` mode for CI pipelines (non-zero exit if issues found)
- `--redact` mode to mask sensitive values in output
- Color-coded, tab-aligned output for easy readability

---

## 🚀 Installation

### Clone and Build:

```bash
git clone https://github.com/sahandset/lockr.git
cd lockr
go build -o lockr
```

### Run:

```bash
./lockr .env --redact --test --ruleset rules.json
```

---

## 🧪 Example Usage

To test the scanner locally:

1. Copy the provided example env file:
   ```bash
   cp .env.example .env
   ```

2. Run the scanner:
   ```bash
   go run main.go .env --redact
   ```

The `.env.example` file contains synthetic secrets that trigger both built-in and custom rules without risking exposure of real credentials. It is safe to commit and use for testing.

---

## 📘 Custom Rules

You can extend detection using your own rules:

### JSON Format

```json
{
  "Slack Webhook": "https://hooks.slack.com/services/\w+/\w+/\w+",
  "Stripe Secret": "sk_live_[0-9a-zA-Z]+"
}
```

### YAML Format

```yaml
Slack Webhook: "https://hooks.slack.com/services/\w+/\w+/\w+"
Stripe Secret: "sk_live_[0-9a-zA-Z]+"
```

Pass the file using:

```bash
./lockr .env --ruleset rules.yaml
```

---

## ⚙️ CLI Flags

| Flag         | Description                                           |
|--------------|-------------------------------------------------------|
| `--ruleset`  | Path to JSON or YAML custom ruleset                   |
| `--redact`   | Redacts secret values in the output                   |
| `--test`     | Exits with code 1 if any issues are found (CI usage)  |

---

## ✅ Example Output

```
[Password]    config.env    line 3   | DB_PASSWORD=***REDACTED***
[JWT]         config.env    line 7   | JWT=***REDACTED***
```

---

## 🤝 Contributing

Contributions are welcome!

If you’d like to help improve lockr:

- Open an issue or feature request
- Write unit or integration tests
- Add support for new secret patterns or file formats
- Improve CLI UX or documentation

### Development Setup

```bash
git clone https://github.com/your-username/lockr.git
cd lockr
go run main.go .env
```

### Guidelines

- Format your code using `gofmt` before committing
- Use semantic commit messages (`feat:`, `fix:`, `refactor:`, etc.)
- Ensure all tests pass: `go test ./...`
- Be clear and concise in your PR descriptions

---

## 📄 License

MIT License. See `LICENSE` for details.
