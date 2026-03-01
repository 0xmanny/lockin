export function Footer() {
  return (
    <footer className="border-t border-border px-6 py-8">
      <div className="mx-auto flex max-w-5xl flex-col items-center justify-between gap-4 sm:flex-row">
        <span className="font-mono text-sm text-muted">
          lockin<span className="text-accent">.sh</span>
        </span>
        <div className="flex items-center gap-6 text-sm text-muted">
          <a
            href="https://github.com/0xmanny/lockin"
            target="_blank"
            rel="noopener noreferrer"
            className="hover:text-foreground transition-colors"
          >
            GitHub
          </a>
          <span>MIT License</span>
        </div>
      </div>
    </footer>
  );
}
