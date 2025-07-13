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

const cartItems = [0, 2, 5]; // Example: jersey indices in cart
const sizes = ["M", "L", "XL"];

export default function CartPage() {
  const subtotal = cartItems.length * 899;
  const shipping = subtotal > 1500 ? 0 : 99;
  const tax = subtotal * 0.05;
  const total = subtotal + shipping + tax;

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-900 via-purple-900 to-blue-900 py-8">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="mb-8">
          <h1 className={`${bangers.className} text-4xl font-bold text-white mb-2 drop-shadow-lg`}>Shopping Cart</h1>
          <p className="text-gray-200 mt-2">
            {cartItems.length} jersey{cartItems.length !== 1 ? 's' : ''} in your cart
          </p>
        </div>

        {cartItems.length === 0 ? (
          <div className="text-center py-12">
            <svg className="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M3 3h2l.4 2M7 13h10l4-8H5.4m0 0L7 13m0 0l-2.5 5M7 13l2.5 5m6-5v6a2 2 0 01-2 2H9a2 2 0 01-2-2v-6m8 0V9a2 2 0 00-2-2H9a2 2 0 00-2 2v4.01" />
            </svg>
            <h3 className="mt-2 text-sm font-medium text-white">Your cart is empty</h3>
            <p className="mt-1 text-sm text-gray-300">Start shopping to add jerseys to your cart.</p>
            <div className="mt-6">
              <Link
                href="/products"
                className="inline-flex items-center px-4 py-2 border border-transparent shadow-sm text-sm font-medium rounded-md text-white bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 hover:scale-105"
              >
                Continue Shopping
              </Link>
            </div>
          </div>
        ) : (
          <div className="lg:grid lg:grid-cols-12 lg:gap-x-12 lg:items-start">
            {/* Cart Items */}
            <div className="lg:col-span-7">
              <div className="border-b border-yellow-400/30 pb-4 mb-4">
                <h2 className="text-lg font-medium text-white">Cart Items</h2>
              </div>
              <div className="space-y-4">
                {cartItems.map((idx, i) => (
                  <div key={jerseys[idx]} className="flex items-center space-x-4 py-4 border-b border-yellow-400/20 group">
                    {/* Product Image */}
                    <div className="flex-shrink-0">
                      <div className="w-20 h-20 bg-gradient-to-br from-yellow-400 via-pink-500 to-purple-600 rounded-full overflow-hidden relative">
                        <Image
                          src="/globe.svg"
                          alt={jerseys[idx]}
                          fill
                          className="object-contain group-hover:scale-105 transition-transform duration-300"
                          sizes="80px"
                          priority
                        />
                      </div>
                    </div>
                    {/* Product Details */}
                    <div className="flex-1 min-w-0">
                      <div className="flex justify-between">
                        <div>
                          <h3 className="text-sm font-bold text-white">
                            <Link href={`/products/${idx + 1}`} className="hover:text-yellow-300">
                              {jerseys[idx]}
                            </Link>
                          </h3>
                          <p className="text-sm text-gray-300">Size: {sizes[i % sizes.length]}</p>
                        </div>
                        <p className="text-sm font-bold text-yellow-300">Rs. 899</p>
                      </div>
                      {/* Quantity Controls */}
                      <div className="flex items-center justify-between mt-2">
                        <div className="flex items-center border border-yellow-400/30 rounded-lg">
                          <button className="px-3 py-1 text-gray-200 hover:text-white" aria-label="Decrease quantity">
                            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M20 12H4" />
                            </svg>
                          </button>
                          <span className="px-3 py-1 text-sm text-white">1</span>
                          <button className="px-3 py-1 text-gray-200 hover:text-white" aria-label="Increase quantity">
                            <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 4v16m8-8H4" />
                            </svg>
                          </button>
                        </div>
                        <button className="text-sm text-red-400 hover:text-red-600" aria-label="Remove from cart">
                          Remove
                        </button>
                      </div>
                    </div>
                  </div>
                ))}
              </div>
              {/* Continue Shopping */}
              <div className="mt-8">
                <Link
                  href="/products"
                  className="text-yellow-300 hover:text-white font-medium"
                >
                   Continue Shopping
                </Link>
              </div>
            </div>
            {/* Order Summary */}
            <div className="lg:col-span-5 mt-8 lg:mt-0">
              <div className="bg-white/10 rounded-2xl p-6 border-2 border-yellow-400/20">
                <h2 className="text-lg font-bold text-white mb-4">Order Summary</h2>
                <div className="space-y-4">
                  <div className="flex justify-between">
                    <span className="text-gray-200">Subtotal</span>
                    <span className="text-yellow-300 font-bold">Rs. {subtotal}</span>
                  </div>
                  <div className="flex justify-between">
                    <span className="text-gray-200">Shipping</span>
                    <span className="text-yellow-300 font-bold">
                      {shipping === 0 ? 'Free' : `Rs. ${shipping}`}
                    </span>
                  </div>
                  <div className="flex justify-between">
                    <span className="text-gray-200">Tax</span>
                    <span className="text-yellow-300 font-bold">Rs. {tax.toFixed(0)}</span>
                  </div>
                  <div className="border-t border-yellow-400/20 pt-4">
                    <div className="flex justify-between">
                      <span className="text-lg font-bold text-white">Total</span>
                      <span className="text-lg font-bold text-yellow-300">Rs. {total.toFixed(0)}</span>
                    </div>
                  </div>
                </div>
                {/* Checkout Button */}
                <div className="mt-6">
                  <Link
                    href="/checkout"
                    className="w-full bg-gradient-to-r from-yellow-400 via-pink-500 to-purple-600 text-white py-3 px-4 rounded-full font-bold hover:scale-105 transition-transform block text-center"
                  >
                    Proceed to Checkout
                  </Link>
                </div>
                {/* Security Notice */}
                <div className="mt-4 text-center">
                  <div className="flex items-center justify-center text-sm text-gray-300">
                    <svg className="w-4 h-4 mr-1" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z" />
                    </svg>
                    Secure checkout
                  </div>
                </div>
              </div>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
