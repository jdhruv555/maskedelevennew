import type { Metadata } from "next";
import { Geist, Geist_Mono, Bangers } from "next/font/google";
import Link from "next/link";
import "./globals.css";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

const bangers = Bangers({
  weight: "400",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "Masked 11 | Premium Sports Jerseys & Football Kits",
  description: "Shop authentic football jerseys, soccer shirts, and fan kits at Masked 11. Fast shipping, secure payments, and the best selection of club and international jerseys in India.",
  keywords: "sports jerseys, football kits, soccer shirts, fan version, club jerseys, international jerseys, Masked 11, buy football jersey India, premium jerseys",
  authors: [{ name: "Masked 11" }],
  openGraph: {
    title: "Masked 11 | Premium Sports Jerseys & Football Kits",
    description: "Shop authentic football jerseys, soccer shirts, and fan kits at Masked 11. Fast shipping, secure payments, and the best selection of club and international jerseys in India.",
    type: "website",
    locale: "en_IN",
    url: "https://masked11.com/",
    images: [
      {
        url: "/globe.svg",
        width: 1200,
        height: 630,
        alt: "Masked 11 - Premium Sports Jerseys",
      },
    ],
  },
  twitter: {
    card: "summary_large_image",
    title: "Masked 11 | Premium Sports Jerseys & Football Kits",
    description: "Shop authentic football jerseys, soccer shirts, and fan kits at Masked 11.",
    images: ["/globe.svg"],
  },
  robots: {
    index: true,
    follow: true,
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased bg-gray-50 min-h-screen flex flex-col`}
      >
        {/* Header */}
        <header className="sticky top-0 z-50 bg-gradient-to-r from-black via-red-900 to-red-600 bg-opacity-70 backdrop-blur-md border-b border-red-400 shadow-xl relative overflow-hidden">
          {/* Full-width background image */}
          <div className="absolute inset-0 w-full h-full z-0 pointer-events-none">
            <img src="/globe.svg" alt="Header Background" className="w-full h-full object-cover opacity-20" />
          </div>
          <div className="relative z-10 max-w-7xl mx-auto flex items-center justify-between px-4 py-3">
            <Link href="/" className="flex items-center gap-2 text-3xl font-bold text-yellow-300 drop-shadow-lg" style={{ fontFamily: bangers.style.fontFamily }}>
              <span className="sr-only">Masked 11 Home</span>
              <svg className="w-9 h-9 text-yellow-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="12" r="10" strokeWidth="2" /><path d="M8 12l2 2 4-4" strokeWidth="2" /></svg>
              Masked 11
            </Link>
            <nav className="hidden md:flex gap-8 text-lg font-bold uppercase tracking-wide text-white">
              <Link href="/products" className="hover:text-yellow-300 transition-colors">Shop</Link>
              <Link href="/categories" className="hover:text-yellow-300 transition-colors">Teams</Link>
              <Link href="/about" className="hover:text-yellow-300 transition-colors">About</Link>
              <Link href="/contact" className="hover:text-yellow-300 transition-colors">Contact</Link>
            </nav>
            <div className="flex items-center gap-4">
              <Link href="/cart" className="relative group">
                <svg className="w-7 h-7 text-yellow-300 group-hover:text-white transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="9" cy="21" r="1" /><circle cx="20" cy="21" r="1" /><path d="M1 1h4l2.68 13.39a2 2 0 002 1.61h9.72a2 2 0 002-1.61L23 6H6" strokeWidth="2" /></svg>
                <span className="sr-only">Cart</span>
              </Link>
              <Link href="/profile" className="group">
                <svg className="w-7 h-7 text-yellow-300 group-hover:text-white transition-colors" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="8" r="4" /><path d="M6 20v-2a4 4 0 014-4h4a4 4 0 014 4v2" strokeWidth="2" /></svg>
                <span className="sr-only">Profile</span>
              </Link>
            </div>
            {/* Mobile menu button */}
            <button className="md:hidden ml-2 p-2 rounded hover:bg-yellow-400/10" aria-label="Open menu">
              <svg className="w-7 h-7 text-yellow-300" fill="none" stroke="currentColor" viewBox="0 0 24 24"><path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M4 6h16M4 12h16M4 18h16" /></svg>
            </button>
          </div>
        </header>
        {/* Main content */}
        <main className="flex-1 w-full">
          {children}
        </main>
      </body>
    </html>
  );
}
