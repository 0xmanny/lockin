import Link from "next/link";

export function Nav() {
  return (
    <nav className="fixed top-0 z-50 w-full border-b border-border/50 bg-background/80 backdrop-blur-md">
      <div className="mx-auto flex h-14 max-w-5xl items-center justify-between px-6">
        <Link href="/" className="font-mono text-sm font-bold tracking-tight">
          lockin<span className="text-accent">.sh</span>
        </Link>
        <div className="flex items-center gap-6">
          <Link
            href="/docs"
            className="text-sm text-muted hover:text-foreground transition-colors"
          >
            Docs
          </Link>
          <a
            href="https://github.com/0xmanny/lockin"
            target="_blank"
            rel="noopener noreferrer"
            className="text-sm text-muted hover:text-foreground transition-colors"
          >
            GitHub
          </a>
        </div>
      </div>
    </nav>
  );
}
