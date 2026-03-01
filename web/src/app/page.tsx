import { CopyButton } from "@/components/copy-button";
import { Nav } from "@/components/nav";
import { Footer } from "@/components/footer";

const INSTALL_CMD = "curl -fsSL https://lockin.sh/install | sh";

const features = [
  {
    title: "Block Websites",
    description:
      "Modifies /etc/hosts at the kernel level to block distracting sites. Zero network overhead — no proxies, no packet inspection.",
    icon: (
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <circle cx="12" cy="12" r="10" />
        <line x1="2" y1="12" x2="22" y2="12" />
        <path d="M12 2a15.3 15.3 0 014 10 15.3 15.3 0 01-4 10 15.3 15.3 0 01-4-10 15.3 15.3 0 014-10z" />
      </svg>
    ),
  },
  {
    title: "Block Apps",
    description:
      "Monitors running applications and gracefully quits distracting ones. Messages, Discord, Telegram — gone.",
    icon: (
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <rect x="2" y="3" width="20" height="14" rx="2" ry="2" />
        <line x1="8" y1="21" x2="16" y2="21" />
        <line x1="12" y1="17" x2="12" y2="21" />
      </svg>
    ),
  },
  {
    title: "Fully Customizable",
    description:
      "Edit ~/.lockin/config.yaml to add or remove sites and apps. Ship with sane defaults, tweak to your workflow.",
    icon: (
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <circle cx="12" cy="12" r="3" />
        <path d="M19.4 15a1.65 1.65 0 00.33 1.82l.06.06a2 2 0 010 2.83 2 2 0 01-2.83 0l-.06-.06a1.65 1.65 0 00-1.82-.33 1.65 1.65 0 00-1 1.51V21a2 2 0 01-4 0v-.09A1.65 1.65 0 009 19.4a1.65 1.65 0 00-1.82.33l-.06.06a2 2 0 01-2.83-2.83l.06-.06A1.65 1.65 0 004.68 15a1.65 1.65 0 00-1.51-1H3a2 2 0 010-4h.09A1.65 1.65 0 004.6 9a1.65 1.65 0 00-.33-1.82l-.06-.06a2 2 0 012.83-2.83l.06.06A1.65 1.65 0 009 4.68a1.65 1.65 0 001-1.51V3a2 2 0 014 0v.09a1.65 1.65 0 001 1.51 1.65 1.65 0 001.82-.33l.06-.06a2 2 0 012.83 2.83l-.06.06A1.65 1.65 0 0019.4 9a1.65 1.65 0 001.51 1H21a2 2 0 010 4h-.09a1.65 1.65 0 00-1.51 1z" />
      </svg>
    ),
  },
  {
    title: "Single Binary",
    description:
      "Written in Go. No runtime dependencies, no bloated apps hogging resources. Install in seconds, runs instantly.",
    icon: (
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        strokeWidth="2"
        strokeLinecap="round"
        strokeLinejoin="round"
      >
        <polyline points="4 17 10 11 4 5" />
        <line x1="12" y1="19" x2="20" y2="19" />
      </svg>
    ),
  },
];

const steps = [
  {
    step: "1",
    title: "Install",
    code: "curl -fsSL https://lockin.sh/install | sh",
  },
  {
    step: "2",
    title: "Lock in",
    code: "lockin",
  },
  {
    step: "3",
    title: "Ship code",
    code: "# distractions blocked. focus mode on.",
  },
];

export default function Home() {
  return (
    <div className="min-h-screen font-sans">
      <Nav />

      {/* Hero */}
      <section className="flex flex-col items-center justify-center px-6 pt-36 pb-24">
        <div className="mb-6 inline-flex items-center gap-2 rounded-full border border-border px-3 py-1 text-xs text-muted">
          <span className="h-1.5 w-1.5 rounded-full bg-accent" />
          Open source &middot; macOS
        </div>
        <h1 className="max-w-2xl text-center text-4xl font-bold tracking-tight sm:text-6xl">
          Block distractions.
          <br />
          <span className="text-accent">Ship code.</span>
        </h1>
        <p className="mt-6 max-w-lg text-center text-lg text-muted">
          A lightweight CLI that blocks distracting websites and apps on macOS.
          One command to focus. Zero performance impact.
        </p>

        <div className="mt-10 flex w-full max-w-lg items-center gap-3 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm">
          <span className="text-accent select-none">$</span>
          <code className="flex-1 overflow-x-auto text-foreground">
            {INSTALL_CMD}
          </code>
          <CopyButton text={INSTALL_CMD} />
        </div>

        <div className="mt-4 flex gap-4">
          <a
            href="/docs"
            className="rounded-lg bg-accent px-5 py-2.5 text-sm font-medium text-background transition-colors hover:bg-accent-muted"
          >
            Get Started
          </a>
          <a
            href="https://github.com/0xmanny/lockin"
            target="_blank"
            rel="noopener noreferrer"
            className="rounded-lg border border-border px-5 py-2.5 text-sm font-medium text-foreground transition-colors hover:bg-card"
          >
            View Source
          </a>
        </div>
      </section>

      {/* Features */}
      <section className="border-t border-border px-6 py-24">
        <div className="mx-auto max-w-5xl">
          <h2 className="text-center text-2xl font-bold tracking-tight sm:text-3xl">
            Everything you need to focus
          </h2>
          <p className="mt-3 text-center text-muted">
            No bloated app. No subscription. Just a CLI that does its job.
          </p>

          <div className="mt-16 grid gap-6 sm:grid-cols-2">
            {features.map((f) => (
              <div
                key={f.title}
                className="rounded-xl border border-border bg-card p-6"
              >
                <div className="mb-4 flex h-10 w-10 items-center justify-center rounded-lg border border-border text-accent">
                  {f.icon}
                </div>
                <h3 className="text-lg font-semibold">{f.title}</h3>
                <p className="mt-2 text-sm leading-relaxed text-muted">
                  {f.description}
                </p>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* How It Works */}
      <section className="border-t border-border px-6 py-24">
        <div className="mx-auto max-w-3xl">
          <h2 className="text-center text-2xl font-bold tracking-tight sm:text-3xl">
            Three commands. That&apos;s it.
          </h2>
          <div className="mt-16 flex flex-col gap-8">
            {steps.map((s) => (
              <div key={s.step} className="flex gap-6">
                <div className="flex h-10 w-10 shrink-0 items-center justify-center rounded-full border border-accent/30 text-sm font-bold text-accent">
                  {s.step}
                </div>
                <div className="flex-1">
                  <h3 className="text-lg font-semibold">{s.title}</h3>
                  <div className="mt-2 rounded-lg border border-border bg-card px-4 py-3 font-mono text-sm text-muted">
                    <span className="text-accent select-none">$ </span>
                    {s.code}
                  </div>
                </div>
              </div>
            ))}
          </div>

          <div className="mt-12 text-center">
            <a
              href="/docs"
              className="text-sm text-accent hover:text-accent-muted transition-colors"
            >
              Read the full documentation &rarr;
            </a>
          </div>
        </div>
      </section>

      <Footer />
    </div>
  );
}
