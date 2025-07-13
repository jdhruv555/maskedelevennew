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
        <header className="sticky top-0 z-50 bg-gradient-to-r from-gray-900 via-purple-900 to-blue-900 border-b-4 border-yellow-400/40 shadow-xl">
          <div className="max-w-7xl mx-auto flex items-center justify-between px-4 py-3">
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
        {/* Footer */}
        <footer className="bg-gradient-to-r from-gray-900 via-purple-900 to-blue-900 border-t-4 border-yellow-400/40 py-8 mt-12">
          <div className="max-w-7xl mx-auto px-4 flex flex-col md:flex-row items-center justify-between gap-4">
            <div className="flex items-center gap-2 text-2xl font-bold text-yellow-300 drop-shadow-lg" style={{ fontFamily: bangers.style.fontFamily }}>
              <svg className="w-7 h-7 text-yellow-400" fill="none" stroke="currentColor" viewBox="0 0 24 24"><circle cx="12" cy="12" r="10" strokeWidth="2" /><path d="M8 12l2 2 4-4" strokeWidth="2" /></svg>
              Masked 11
            </div>
            <nav className="flex gap-6 text-white text-sm font-bold uppercase tracking-wide">
              <Link href="/products" className="hover:text-yellow-300 transition-colors">Shop</Link>
              <Link href="/categories" className="hover:text-yellow-300 transition-colors">Teams</Link>
              <Link href="/about" className="hover:text-yellow-300 transition-colors">About</Link>
              <Link href="/contact" className="hover:text-yellow-300 transition-colors">Contact</Link>
              <Link href="/profile" className="hover:text-yellow-300 transition-colors">Profile</Link>
            </nav>
            <div className="flex gap-4">
              <a href="#" aria-label="Instagram" className="hover:text-yellow-300 transition-colors"><svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><path d="M7.75 2h8.5A5.75 5.75 0 0122 7.75v8.5A5.75 5.75 0 0116.25 22h-8.5A5.75 5.75 0 012 16.25v-8.5A5.75 5.75 0 017.75 2zm0 1.5A4.25 4.25 0 003.5 7.75v8.5A4.25 4.25 0 007.75 20.5h8.5a4.25 4.25 0 004.25-4.25v-8.5A4.25 4.25 0 0016.25 3.5zm4.25 2.25a5.25 5.25 0 11-5.25 5.25 5.25 5.25 0 015.25-5.25zm0 1.5a3.75 3.75 0 103.75 3.75A3.75 3.75 0 0012 5.25zm5.5 1.25a1 1 0 11-1 1 1 1 0 011-1z" /></svg></a>
              <a href="#" aria-label="Facebook" className="hover:text-yellow-300 transition-colors"><svg className="w-5 h-5" fill="currentColor" viewBox="0 0 24 24"><path d="M22 12c0-5.523-4.477-10-10-10S2 6.477 2 12c0 5.019 3.676 9.163 8.438 9.877v-6.987h-2.54v-2.89h2.54V9.797c0-2.506 1.492-3.89 3.777-3.89 1.094 0 2.238.195 2.238.195v2.46h-1.26c-1.242 0-1.63.771-1.63 1.562v1.875h2.773l-.443 2.89h-2.33v6.987C18.324 21.163 22 17.019 22 12z" /></svg></a>
            </div>
          </div>
        </footer>
      </body>
    </html>
  );
}
