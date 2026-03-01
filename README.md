# lockin.sh

**Block distractions. Ship code.**

A lightweight CLI tool that blocks distracting websites and apps on macOS. No subscriptions, no bloated apps — just a single binary that does its job.

[lockin.sh](https://lockin.sh) · [GitHub](https://github.com/0xmanny/lockin)

---

## Install

```sh
curl -fsSL https://lockin.sh/install | sh
```

The script detects your architecture (Apple Silicon or Intel) and installs the binary to `/usr/local/bin/lockin`.

### Build from source

```sh
git clone https://github.com/0xmanny/lockin.git
cd lockin
make install
```

Requires Go 1.21+.

---

## Usage

### Start blocking

```sh
lockin
```

This blocks distracting websites and apps immediately. lockin will prompt for your password (modifying `/etc/hosts` requires root), activate blocking, and return you to your terminal. App monitoring runs silently in the background.

### Check status

```sh
lockin status
```

Shows whether lockin is active and lists all blocked websites and apps.

### Change the blocklist

```sh
lockin config
```

Opens an interactive picker to toggle which websites and apps are blocked. If lockin is active, changes are applied immediately.

### Stop blocking

```sh
lockin stop
```

Restores `/etc/hosts` and stops all blocking. This works even if the start process crashed.

### Uninstall

```sh
lockin uninstall
```

Fully removes lockin: restores `/etc/hosts`, stops the daemon, removes the LaunchDaemon, config directory, and binary.

---

## Configuration

On first run, lockin creates a config file at `~/.lockin/config.yaml` with sensible defaults:

```yaml
blocked_websites:
  - twitter.com
  - x.com
  - facebook.com
  - instagram.com
  - tiktok.com
  - reddit.com
  - youtube.com
  - netflix.com
  - twitch.tv

blocked_apps:
  - Messages
  - Telegram
  - Discord
  - WhatsApp

settings:
  poll_interval: 3s
```

Edit this file to customize your blocklist. The `www.` variants of all websites are automatically included.

---

## How It Works

**Website blocking** — Adds entries to `/etc/hosts` that redirect blocked domains to `127.0.0.1`. This is kernel-level DNS interception with zero network performance impact. No proxy, no packet inspection. Persists until you run `lockin stop`.

**App blocking** — A lightweight background daemon polls running applications via macOS System Events every few seconds and gracefully quits blocked apps using AppleScript. Falls back to force-quit if needed.

### Safety

- Your `/etc/hosts` is backed up to `~/.lockin/hosts.bak` before any modification
- Managed entries are wrapped in `# LOCKIN-START` / `# LOCKIN-END` markers
- `lockin stop` restores your original hosts file and stops the background daemon
- The daemon PID is stored in `~/.lockin/lockin.pid` for clean shutdown
- A LaunchDaemon (`/Library/LaunchDaemons/sh.lockin.cleanup.plist`) automatically runs `lockin stop` on reboot, so your hosts file is always restored even after a crash or power loss

---

## Project Structure

```
lockin/
  cli/                  # Go CLI tool
    cmd/                # Cobra CLI commands
    internal/
      blocker/          # Website + app blocking logic
      config/           # YAML config management
    main.go
    go.mod
  web/                  # Next.js landing page + docs
  install.sh            # Curl installer script
  Makefile              # Build and cross-compile targets
```

---

## Development

### Prerequisites

- Go 1.21+
- Node.js 18+ (for the landing page)
- macOS (the CLI uses macOS-specific APIs)

### CLI

```sh
make build              # Build for current platform
make build-all          # Cross-compile for arm64 + amd64
make dev                # Cleans lockin state, build, run binary
```

### Run the landing page

```sh
cd web
npm install
npm run dev
```

---

## Contributing

Contributions are welcome. Here's how to get started:

1. Fork the repository
2. Create a feature branch: `git checkout -b my-feature`
3. Make your changes and test locally
4. Submit a pull request

Please open an issue first for major changes so we can discuss the approach.

---

## License

MIT
