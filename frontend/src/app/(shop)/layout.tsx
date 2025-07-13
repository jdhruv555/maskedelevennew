import { Suspense } from "react";

export default function ShopLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <>
      <main>
        <Suspense fallback={<div className="flex justify-center items-center h-64">Loading...</div>}>
          {children}
        </Suspense>
      </main>
    </>
  );
}
