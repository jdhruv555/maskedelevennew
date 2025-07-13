import Link from "next/link";
import Image from "next/image";
import { Bangers } from "next/font/google";

const bangers = Bangers({
  weight: "400",
  subsets: ["latin"],
});

const jerseys = [
  { name: "Atletico Madrid Home 25/26", image: "/atletico-madrid-home-25-26.jpg" },
  { name: "Barcelona Home 25/26 Fan Version", image: "/barcelona-home-25-26.jpg" },
  { name: "Bayern Munich Home 25/26", image: "/bayern-munich-home-25-26.jpg" },
  { name: "Borussia Dortmund Home 25/26", image: "/borussia-dortmund-home-25-26.jpg" },
  { name: "Chelsea Home 25/26", image: "/chelsea-home-25-26.jpg" },
  { name: "Inter Miami Away 25/26", image: "/inter-miami-away-25-26.jpg" },
  { name: "Inter Miami Home 25/26", image: "/inter-miami-home-25-26.jpg" },
  { name: "Juventus Home 25/26", image: "/juventus-home-25-26.jpg" },
  { name: "Liverpool Home 25/26", image: "/liverpool-home-25-26.jpg" },
  { name: "Manchester City Home 25/26", image: "/manchester-city-home-25-26.jpg" },
  { name: "Manchester United Home 25/26", image: "/manchester-united-home-25-26.jpg" },
  { name: "Mexico Home 25/26", image: "/mexico-home-25-26.jpg" },
  { name: "Real Madrid 25/26 HOME FAN Version", image: "/rm-home-25-26.jpg" },
  { name: "Real Madrid Away 25/26 KIT", image: "/rm-away-25-26.jpg" },
];

export default function Home() {
  return (
    <div className="min-h-screen bg-white flex flex-col">
      {/* Banner below header */}
      <div className="w-full relative">
        <img src="/banner-masked-eleven.jpg" alt="Masked Eleven Banner" className="w-full h-auto object-cover max-h-[320px] md:max-h-[400px] lg:max-h-[480px]" />
      </div>
      {/* Shop Jerseys Section - moved to top */}
      <section className="py-16 bg-white">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className={`${bangers.className} text-3xl font-bold text-center mb-2 text-gray-900`}>Shop Jerseys</h2>
          <p className="text-center text-lg text-gray-500 mb-10">Premium Football Jerseys – All Rs. 899</p>
          <div className="flex flex-wrap justify-center gap-8">
            {jerseys.map((jersey, idx) => (
              <div key={jersey.name} className="bg-white border border-gray-200 rounded-2xl p-6 min-w-[220px] max-w-xs flex flex-col items-center shadow hover:shadow-lg transition">
                <div className="w-40 h-40 mb-4 flex items-center justify-center">
                  <Image src={jersey.image} alt={jersey.name} width={180} height={180} className="object-contain rounded-xl bg-gray-50" />
                </div>
                <h3 className="text-lg font-bold text-gray-900 text-center mb-2">{jersey.name}</h3>
                <span className="text-blue-600 font-semibold text-xl mb-2">Rs. 899</span>
                <Link href={`/products/${idx + 1}`} className="mt-2 bg-blue-600 text-white px-6 py-2 rounded-lg font-semibold shadow hover:bg-blue-700 transition">
                  View Jersey
                </Link>
              </div>
            ))}
          </div>
        </div>
      </section>

      {/* Hero Section */}
      <section className="flex flex-col items-center justify-center py-20 bg-white">
        <h1 className={`${bangers.className} text-5xl md:text-7xl font-extrabold text-gray-900 mb-4`}>MASKED 11</h1>
        <p className="text-xl text-gray-700 mb-8">Premium Football Jerseys <span className="text-blue-600 font-bold">Rs. 899</span></p>
        <div className="flex gap-4">
          <Link 
            href="/products" 
            className="bg-blue-600 text-white px-8 py-3 rounded-lg font-semibold shadow hover:bg-blue-700 transition"
          >
            Shop Jerseys
          </Link>
          <Link 
            href="/categories" 
            className="border border-blue-600 text-blue-600 px-8 py-3 rounded-lg font-semibold hover:bg-blue-50 transition"
          >
            Browse Teams
          </Link>
        </div>
      </section>

      {/* Features Section */}
      <section className="py-16 bg-gray-50">
        <div className="max-w-6xl mx-auto px-4">
          <h2 className={`${bangers.className} text-3xl font-bold text-center mb-12 text-gray-900`}>Why Masked 11?</h2>
          <div className="grid md:grid-cols-4 gap-8">
            <div className="text-center">
              <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4 shadow">
                <svg className="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4" />
                </svg>
              </div>
              <h3 className="text-lg font-semibold mb-2 text-gray-900">Authentic Quality</h3>
              <p className="text-gray-600">Official fan versions, stitched badges, and premium fabrics.</p>
            </div>
            <div className="text-center">
              <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4 shadow">
                <svg className="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" />
                </svg>
              </div>
              <h3 className="text-lg font-semibold mb-2 text-gray-900">Fast Shipping</h3>
              <p className="text-gray-600">Get your jersey delivered in 2-4 days, pan India.</p>
            </div>
            <div className="text-center">
              <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4 shadow">
                <svg className="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                </svg>
              </div>
              <h3 className="text-lg font-semibold mb-2 text-gray-900">Secure Payments</h3>
              <p className="text-gray-600">Pay with UPI, cards, or COD. 100% safe and encrypted.</p>
            </div>
            <div className="text-center">
              <div className="w-16 h-16 bg-blue-100 rounded-full flex items-center justify-center mx-auto mb-4 shadow">
                <svg className="w-8 h-8 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24" aria-hidden="true">
                  <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                </svg>
              </div>
              <h3 className="text-lg font-semibold mb-2 text-gray-900">Easy Returns</h3>
              <p className="text-gray-600">7-day hassle-free returns on all jerseys.</p>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
