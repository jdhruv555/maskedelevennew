import Image from "next/image";
import Link from "next/link";
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

const sizes = ["S", "M", "L", "XL", "XXL"];

export default function ProductDetailPage({ params }: { params: { slug: string } }) {
  // Use slug as index for demo; in real app, use slug/id lookup
  const idx = parseInt(params.slug) - 1;
  const name = jerseys[idx] || "Jersey";
  const related = jerseys.filter((_, i) => i !== idx).slice(0, 3);

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 py-12">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="grid md:grid-cols-2 gap-12">
          {/* Image Gallery */}
          <div className="flex flex-col gap-4">
            <div className="relative w-full aspect-square rounded-2xl overflow-hidden shadow-2xl border-4 border-yellow-400/30">
              <Image
                src="/globe.svg"
                alt={name}
                fill
                className="object-contain"
                sizes="(max-width: 768px) 100vw, 50vw"
                priority
              />
            </div>
          </div>
          {/* Product Info */}
          <div>
            <h1 className={`${bangers.className} text-4xl font-bold text-white mb-4 drop-shadow-lg`}>{name}</h1>
            <div className="flex items-center gap-4 mb-6">
              <span className="text-2xl font-bold text-yellow-300">Rs. 899</span>
              <span className="text-sm text-gray-200">In stock</span>
            </div>
            <p className="text-gray-200 mb-6">Premium quality football jersey. Official fan version. Fast shipping, easy returns.</p>
            <div className="mb-6">
              <label htmlFor="size" className="block text-sm font-medium text-gray-200 mb-2">Size</label>
              <select id="size" className="w-full px-4 py-2 border-2 border-yellow-400/30 rounded-lg focus:ring-2 focus:ring-yellow-400 focus:border-transparent bg-white/10 text-white">
                {sizes.map((size) => (
                  <option key={size} value={size}>{size}</option>
                ))}
              </select>
            </div>
            <button className="w-full bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-white py-3 px-4 rounded-full font-bold text-lg shadow-xl hover:scale-105 transition-transform mb-4" aria-label="Add to cart">
              Add to Cart
            </button>
            <div className="flex gap-2 text-sm">
              <Link href="/products" className="text-gray-200 hover:text-yellow-300">Back to Shop</Link>
              <span className="text-gray-400">|</span>
              <Link href="/cart" className="text-gray-200 hover:text-yellow-300">Go to Cart</Link>
            </div>
          </div>
        </div>
        {/* Related Jerseys */}
        <div className="mt-16">
          <h2 className={`${bangers.className} text-2xl font-bold text-white mb-6 drop-shadow-lg`}>Related Jerseys</h2>
          <div className="flex flex-wrap gap-6">
            {related.map((rel, i) => (
              <div key={rel} className="bg-white/10 border-2 border-white/20 rounded-2xl p-6 min-w-[220px] max-w-xs flex flex-col items-center shadow-xl hover:scale-105 hover:border-yellow-400 transition-transform duration-200">
                <div className="w-24 h-24 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full flex items-center justify-center mb-4 shadow-lg overflow-hidden">
                  <Image src="/globe.svg" alt={rel} width={72} height={72} className="object-contain" />
                </div>
                <h3 className="text-base font-bold text-white text-center mb-2" style={{ textShadow: "2px 2px 8px #000" }}>{rel}</h3>
                <span className="text-yellow-300 font-bold text-lg mb-2">Rs. 899</span>
                <Link href={`/products/${jerseys.indexOf(rel) + 1}`} className="mt-2 bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-white px-4 py-2 rounded-full font-semibold shadow hover:scale-105 transition-transform">
                  View Jersey
                </Link>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}
