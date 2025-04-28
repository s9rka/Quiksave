import { NavLink, useLocation } from "react-router-dom";
import { Plus, Library } from "lucide-react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";
import { useVault } from "@/context/VaultContext";

export default function BottomNav() {
  const { vaultId } = useVault();
  const location = useLocation();
  
  // Determine if we're inside a vault by checking the URL pattern
  const isInVault = location.pathname.startsWith('/vault/');

  // Dynamic nav items - "Create Note" href changes based on location
  const navItems = [
    { 
      icon: Library, 
      label: "Storage", 
      href: isInVault && vaultId ? `/vault/${vaultId}` : "/vaults" 
    },
    { 
      icon: Plus, 
      label: "Create Note", 
      href: isInVault && vaultId ? `/vault/${vaultId}/new` : "/vaults"
    },
  ];

  return (
    <nav
      className="fixed bottom-4 left-1/2 -translate-x-1/2 z-50
                    max-w-md
                    rounded-xl bg-[#EAEFF3]/50 backdrop-blur-sm
                    p-4 shadow-md"
    >
      <ul className="flex justify-around items-center w-full">
        {navItems.map((item) => (
          <li key={item.label}>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <NavLink
                    to={item.href}
                    className="flex h-12 w-12 items-center 
                               justify-center rounded-2xl 
                               bg-[#D9EFBD] text-[#335F68] 
                               transition-colors hover:bg-[#D8F0C8]"
                  >
                    <item.icon className="w-6 h-6" />
                  </NavLink>
                </TooltipTrigger>
                <TooltipContent>
                  <p>{item.label}</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </li>
        ))}
      </ul>
    </nav>
  );
}