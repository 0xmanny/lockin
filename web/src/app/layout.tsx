import type { Metadata } from "next";
import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "lockin.sh — Block Distractions, Ship Code",
  description:
    "A lightweight CLI tool that blocks distracting websites and apps on macOS. Open source application firewall for developers.",
  openGraph: {
    title: "lockin.sh — Block Distractions, Ship Code",
    description:
      "A lightweight CLI tool that blocks distracting websites and apps on macOS.",
    url: "https://lockin.sh",
    siteName: "lockin.sh",
    type: "website",
  },
  twitter: {
    card: "summary_large_image",
    title: "lockin.sh — Block Distractions, Ship Code",
    description:
      "A lightweight CLI tool that blocks distracting websites and apps on macOS.",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
