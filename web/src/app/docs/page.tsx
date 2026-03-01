import type { Metadata } from "next";
import { Nav } from "@/components/nav";
import { Footer } from "@/components/footer";

export const metadata: Metadata = {
  title: "Docs — lockin.sh",
  description: "Installation, usage, configuration, and how lockin.sh works.",
};

export default function DocsPage() {
  return (
    <div className="min-h-screen font-sans">
      <Nav />

      <div className="mx-auto max-w-3xl px-6 pt-28 pb-24">
        <h1 className="text-3xl font-bold tracking-tight sm:text-4xl">
          Documentation
        </h1>
        <p className="mt-3 text-muted">
          Everything you need to install, configure, and use lockin.
        </p>

        {/* Installation */}
        <div className="mt-16">
          <h2 className="text-xl font-semibold">Installation</h2>
          <p className="mt-3 text-sm leading-relaxed text-muted">
            Install lockin with a single command. The script detects your
            architecture (Apple Silicon or Intel) and downloads the correct
            binary.
          </p>
          <div className="mt-4 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
            <span className="text-accent select-none">$ </span>
              curl -fsSL https://lockin.sh/install | sh
          </div>
          <p className="mt-3 text-sm leading-relaxed text-muted">
            The binary is installed to{" "}
            <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
              /usr/local/bin/lockin
            </code>
            . You can also build from source:
          </p>
          <div className="mt-4 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
            <div>
              <span className="text-accent select-none">$ </span>
              git clone https://github.com/0xmanny/lockin.git
            </div>
            <div>
              <span className="text-accent select-none">$ </span>cd lockin
              && make install
            </div>
          </div>
        </div>

        {/* Usage */}
        <div className="mt-16">
          <h2 className="text-xl font-semibold">Usage</h2>
          <div className="mt-4 space-y-4">
            <div>
              <p className="text-sm text-muted">Start blocking distractions:</p>
              <div className="mt-2 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
                <span className="text-accent select-none">$ </span>lockin
              </div>
            </div>
            <div>
              <p className="text-sm text-muted">
                Check what&apos;s being blocked:
              </p>
              <div className="mt-2 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
                <span className="text-accent select-none">$ </span>lockin
                status
              </div>
            </div>
            <div>
              <p className="text-sm text-muted">
                Change what&apos;s blocked:
              </p>
              <div className="mt-2 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
                <span className="text-accent select-none">$ </span>lockin
                config
              </div>
            </div>
            <div>
              <p className="text-sm text-muted">Stop blocking:</p>
              <div className="mt-2 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
                <span className="text-accent select-none">$ </span>lockin stop
              </div>
            </div>
            <div>
              <p className="text-sm text-muted">
                Fully remove lockin and all its data:
              </p>
              <div className="mt-2 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
                <span className="text-accent select-none">$ </span>lockin
                uninstall
              </div>
            </div>
          </div>
          <p className="mt-4 text-sm leading-relaxed text-muted">
            lockin prompts for your password, activates blocking, and returns you
            to your terminal. App monitoring runs silently in the background.
            Your original{" "}
            <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
              /etc/hosts
            </code>{" "}
            is backed up to{" "}
            <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
              ~/.lockin/hosts.bak
            </code>{" "}
            and fully restored on stop.
          </p>
        </div>

        {/* Config */}
        <div className="mt-16">
          <h2 className="text-xl font-semibold">Configuration</h2>
          <p className="mt-3 text-sm leading-relaxed text-muted">
            On first run, lockin creates a config at{" "}
            <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
              ~/.lockin/config.yaml
            </code>
            . Edit it to customize what gets blocked:
          </p>
          <div className="mt-4 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm leading-relaxed">
            <pre className="overflow-x-auto">{`blocked_websites:
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
  poll_interval: 3s`}</pre>
          </div>
          <p className="mt-3 text-sm leading-relaxed text-muted">
            The{" "}
            <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
              www.
            </code>{" "}
            variants of all websites are automatically included. The{" "}
            <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
              poll_interval
            </code>{" "}
            controls how often lockin checks for blocked apps.
          </p>
        </div>

        {/* How It Works */}
        <div className="mt-16">
          <h2 className="text-xl font-semibold">How It Works</h2>
          <div className="mt-4 space-y-4 text-sm leading-relaxed text-muted">
            <p>
              <strong className="text-foreground">Website blocking</strong>{" "}
              works by adding entries to{" "}
              <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                /etc/hosts
              </code>{" "}
              that redirect blocked domains to{" "}
              <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                127.0.0.1
              </code>
              . This is kernel-level DNS interception with zero network
              performance impact. No proxy, no packet inspection.
            </p>
            <p>
              <strong className="text-foreground">App blocking</strong> runs as
              a lightweight background daemon that polls running applications via
              macOS System Events and gracefully quits any blocked apps. Your
              terminal is returned to you immediately.
            </p>
            <p>
              All changes are fully reversible. Your{" "}
              <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                /etc/hosts
              </code>{" "}
              is backed up before modification and restored when you stop lockin.
              Managed entries are wrapped in marker comments so they never
              interfere with your existing configuration.
            </p>
          </div>
        </div>

        {/* Safety */}
        <div className="mt-16">
          <h2 className="text-xl font-semibold">Safety</h2>
          <ul className="mt-4 space-y-3 text-sm leading-relaxed text-muted">
            <li className="flex gap-2">
              <span className="text-accent mt-0.5">-</span>
              <span>
                Your{" "}
                <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                  /etc/hosts
                </code>{" "}
                is backed up to{" "}
                <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                  ~/.lockin/hosts.bak
                </code>{" "}
                before any modification
              </span>
            </li>
            <li className="flex gap-2">
              <span className="text-accent mt-0.5">-</span>
              <span>
                Managed entries are wrapped in{" "}
                <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                  LOCKIN-START
                </code>{" "}
                /{" "}
                <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                  LOCKIN-END
                </code>{" "}
                markers
              </span>
            </li>
            <li className="flex gap-2">
              <span className="text-accent mt-0.5">-</span>
              <span>
                A LaunchDaemon automatically runs{" "}
                <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                  lockin stop
                </code>{" "}
                on reboot, so your hosts file is always restored — even after a
                crash or power loss
              </span>
            </li>
            <li className="flex gap-2">
              <span className="text-accent mt-0.5">-</span>
              <span>
                <code className="rounded bg-card px-1.5 py-0.5 text-xs font-mono border border-border">
                  lockin uninstall
                </code>{" "}
                fully removes everything: hosts entries, daemon, config, and the
                binary
              </span>
            </li>
          </ul>
        </div>
      </div>

      <Footer />
    </div>
  );
}
