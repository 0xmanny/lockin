import { readFileSync } from "fs";
import { join } from "path";

export function GET() {
  const script = readFileSync(
    join(process.cwd(), "public", "install.sh"),
    "utf-8"
  );

  return new Response(script, {
    headers: {
      "Content-Type": "text/plain; charset=utf-8",
      "Cache-Control": "no-cache",
    },
  });
}
