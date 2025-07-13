import Link from "next/link";
import Image from "next/image";
import { Bangers } from "next/font/google";

const bangers = Bangers({
  weight: "400",
  subsets: ["latin"],
});

const jerseys = [
  "Atletico Madrid Home 25/26",
  "Barcelona Home 25/26 Fan Version",
  "Bayern Munich Home 25/26",
  "Borussia Dortmund Home 25/26",
  "Chelsea Home 25/26",
  "Inter Miami Away 25/26",
  "Inter Miami Home 25/26",
  "Juventus Home 25/26",
  "Liverpool Home 25/26",
  "Manchester City Home 25/26",
  "Manchester United Home 25/26",
  "Mexico Home 25/26",
  "Real Madrid 25/26 HOME FAN Version",
  "Real Madrid Away 25/26 Fan Version",
];

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 flex flex-col">
      {/* Graffiti Spray Paint Hero */}
      <section className="relative h-[80vh] flex items-center justify-center overflow-hidden">
        <div className="absolute inset-0 z-0">
          {/* Graffiti spray SVG background */}
          <svg width="100%" height="100%" className="absolute inset-0 w-full h-full" style={{ filter: "blur(2px)" }}>
            <defs>
              <radialGradient id="graffiti1" cx="50%" cy="50%" r="80%">
                <stop offset="0%" stopColor="#a21caf" stopOpacity="0.7" />
                <stop offset="100%" stopColor="#1e293b" stopOpacity="0" />
              </radialGradient>
            </defs>
            <circle cx="60%" cy="40%" r="400" fill="url(#graffiti1)" />
            <circle cx="30%" cy="70%" r="250" fill="#2563eb" fillOpacity="0.3" />
            <circle cx="80%" cy="80%" r="200" fill="#f59e42" fillOpacity="0.15" />
          </svg>
        </div>
        <div className="relative z-10 text-center text-white px-4 max-w-4xl mx-auto">
          <h1 className={`${bangers.className} text-6xl md:text-8xl font-bold mb-6 drop-shadow-lg tracking-tight`}>Masked 11</h1>
          <p className="text-2xl md:text-3xl mb-8 max-w-2xl mx-auto font-semibold tracking-wide" style={{ textShadow: "2px 2px 8px #000" }}>
            Premium Football Jerseys <span className="inline-block bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-transparent bg-clip-text font-extrabold">Rs. 899</span>
          </p>
          <div className="flex flex-col sm:flex-row gap-4 justify-center animate-fade-in delay-200">
            <Link 
              href="/products" 
              className="bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-white px-10 py-4 rounded-full font-bold text-xl shadow-xl hover:scale-105 transition-transform border-4 border-white/10"
            >
              Shop Jerseys
            </Link>
            <Link 
              href="/categories" 
              className="border-4 border-white text-white px-10 py-4 rounded-full font-bold text-xl shadow-xl hover:bg-white hover:text-purple-700 transition-colors"
            >
              Browse Teams
            </Link>
          </div>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 animate-fade-in delay-300">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className={`${bangers.className} text-4xl font-bold text-center mb-12 text-white drop-shadow-lg`}>Why Masked 11?</h2>
          <div className="grid md:grid-cols-4 gap-8">
            <div className="text-center">
              <div className="w-16 h-16 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-white">Authentic Quality</h3>
              <p className="text-gray-300">Official fan versions, stitched badges, and premium fabrics.</p>
            </div>
            <div className="text-center">
              <div className="w-16 h-16 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-white">Fast Shipping</h3>
              <p className="text-gray-300">Get your jersey delivered in 2-4 days, pan India.</p>
            </div>
            <div className="text-center">
              <div className="w-16 h-16 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-white">Secure Payments</h3>
              <p className="text-gray-300">Pay with UPI, cards, or COD. 100% safe and encrypted.</p>
            </div>
            <div className="text-center">
              <div className="w-16 h-16 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full flex items-center justify-center mx-auto mb-4 shadow-lg">
                <svg className="w-8 h-8 text-white" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <h3 className="text-xl font-semibold mb-2 text-white">Easy Returns</h3>
              <p className="text-gray-300">7-day hassle-free returns on all jerseys.</p>
            </div>
          </div>
        </div>
      </section>

      {/* Jersey List Preview */}
      <section className="py-16 bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 animate-fade-in delay-400">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className={`${bangers.className} text-4xl font-bold text-center mb-12 text-white drop-shadow-lg`}>Shop Jerseys</h2>
          <div className="flex flex-wrap justify-center gap-6">
            {jerseys.map((name, idx) => (
              <div key={name} className="bg-white/10 border-2 border-white/20 rounded-2xl p-6 min-w-[220px] max-w-xs flex flex-col items-center shadow-xl hover:scale-105 hover:border-yellow-400 transition-transform duration-200">
                <div className="w-32 h-32 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full flex items-center justify-center mb-4 shadow-lg overflow-hidden">
                  <Image src="/globe.svg" alt={name} width={96} height={96} className="object-contain" />
                </div>
                <h3 className="text-lg font-bold text-white text-center mb-2" style={{ textShadow: "2px 2px 8px #000" }}>{name}</h3>
                <span className="text-yellow-300 font-bold text-xl mb-2">Rs. 899</span>
                <Link href={`/products/${idx + 1}`} className="mt-2 bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-white px-6 py-2 rounded-full font-semibold shadow hover:scale-105 transition-transform">
                  View Jersey
                </Link>
              </div>
            ))}
          </div>
        </div>
      </section>
    </div>
  );
}
