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
  title: "Masked 11 - Premium Fashion & Accessories",
  description: "Discover the latest trends in fashion, accessories, and lifestyle products at Masked 11. Fast shipping, secure payments, and exceptional customer service.",
  keywords: "fashion, accessories, clothing, shoes, bags, jewelry, lifestyle, online shopping",
  authors: [{ name: "Masked 11" }],
  openGraph: {
    title: "Masked 11 - Premium Fashion & Accessories",
    description: "Discover the latest trends in fashion, accessories, and lifestyle products",
    type: "website",
    locale: "en_US",
  },
  twitter: {
    card: "summary_large_image",
    title: "Masked 11 - Premium Fashion & Accessories",
    description: "Discover the latest trends in fashion, accessories, and lifestyle products",
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
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        {children}
      </body>
    </html>
  );
}
