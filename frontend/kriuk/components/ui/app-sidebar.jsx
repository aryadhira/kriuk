"use client"

import { LayoutGrid, Receipt, Package, User } from "lucide-react"
import { usePathname } from "next/navigation"

import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
} from "@/components/ui/sidebar"

// Menu items.
const items = [
  {
    title: "Dashboard",
    url: "/",
    icon: LayoutGrid,
  },
  {
    title: "Transactions",
    url: "/transaction",
    icon: Receipt,
  },
  {
    title: "Stocks",
    url: "/stock",
    icon: Package,
  },
  {
    title: "Employees",
    url: "/employee",
    icon: User,
  },
]

export function AppSidebar() {
    const pathname = usePathname() 
    return (
        <Sidebar>
            <SidebarHeader>
                <div className="font-bold text-2xl p-2">Kriuk</div>
            </SidebarHeader>
            <SidebarContent>
                <SidebarGroup>
                <SidebarGroupLabel>Menu</SidebarGroupLabel>
                <SidebarGroupContent>
                    <SidebarMenu>
                    {items.map((item) => (
                        <SidebarMenuItem key={item.title}>
                        <SidebarMenuButton asChild  isActive={pathname === item.url}>
                            <a href={item.url}>
                            <item.icon />
                            <span>{item.title}</span>
                            </a>
                        </SidebarMenuButton>
                        </SidebarMenuItem>
                    ))}
                    </SidebarMenu>
                </SidebarGroupContent>
                </SidebarGroup>
            </SidebarContent>
        </Sidebar>
    )
}