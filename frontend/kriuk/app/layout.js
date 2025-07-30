import { Poppins } from "next/font/google"
import "./globals.css";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/ui/app-sidebar";
import PageContainer from "@/components/ui/page-container";
import { Toaster } from "sonner";

const poppins = Poppins({
  subsets: ["latin"],
  weight: ["100","200","300","500","600","700","800","900"],
})

export const metadata = {
  title: "Kriuk App",
  description: "Kriuk App Management",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en" className="dark">
      <body
        className={`${poppins} antialiased`}
      >
        <SidebarProvider>
          <AppSidebar/>
          <main>
            <SidebarTrigger />
            <PageContainer>
              <Toaster/>
              {children}
            </PageContainer>
          </main>
        </SidebarProvider>
      </body>
    </html>
  );
}
