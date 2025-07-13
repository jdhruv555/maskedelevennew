import Link from "next/link";
import { Bangers } from "next/font/google";

const bangers = Bangers({ weight: "400", subsets: ["latin"] });

export default function NotFound() {
  return (
    <div className="min-h-screen flex flex-col items-center justify-center bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 text-white">
      <h1 className={`${bangers.className} text-7xl font-bold mb-4 drop-shadow-lg`}>404</h1>
      <h2 className="text-2xl font-bold mb-2">Page Not Found</h2>
      <p className="mb-6 text-gray-300">Sorry, the page you are looking for does not exist.</p>
      <Link href="/products" className="bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-white px-8 py-3 rounded-full font-bold shadow hover:scale-105 transition-transform">
        Go to Shop
      </Link>
    </div>
  );
}
