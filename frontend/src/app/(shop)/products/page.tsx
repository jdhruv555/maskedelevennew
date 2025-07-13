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

export default function ProductsPage() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        {/* Header */}
        <div className="mb-8 text-center">
          <h1 className={`${bangers.className} text-5xl font-bold text-white mb-2 drop-shadow-lg`}>Shop Jerseys</h1>
          <p className="text-xl text-gray-200 font-semibold mb-4">Premium Football Jerseys – All Rs. 899</p>
        </div>

        {/* Jersey Grid */}
        <div className="flex flex-wrap justify-center gap-8">
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
    </div>
  );
} 