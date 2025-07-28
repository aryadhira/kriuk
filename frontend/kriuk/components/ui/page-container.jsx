// components/PageContainer.tsx
"use client"

export default function PageContainer({ children }) {
  return (
    <div className="w-full min-h-screen p-4 md:p-6 lg:p-8">
      {/* Responsive, dark, rounded container */}
      <div className="bg-gray-900 text-white rounded-2xl shadow-xl border border-gray-800 w-full min-w-[320px] md:min-w-[calc(100vw-21rem)] max-w-screen-2xl min-h-[calc(100vh-2rem)] md:min-h-[calc(100vh-6rem)] lg:min-h-[calc(100vh-4rem)] transition-all duration-300">
        {/* Slot for page content */}
        <div className="p-3 md:p-6 w-full h-full">{children}</div>
      </div>
    </div>
  )
}