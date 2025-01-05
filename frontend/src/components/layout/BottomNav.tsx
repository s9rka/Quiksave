import { NavLink } from "react-router-dom";
import { Plus, Library, User } from 'lucide-react';
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

const navItems = [
  { icon: Library, label: "Storage", href: "/" },
  { icon: User, label: "Account", href: "/account" },
  { icon: Plus, label: "Create Note", href: "/new" },
];

export default function BottomNav() {
  return (
    <nav className="sticky w-full bottom-2 left-2 right-2 bg-[#EAEFF3]/50 p-4 backdrop-blur-sm">
      <ul className="flex justify-around items-center w-full">
        {navItems.map((item) => (
          <li key={item.href}>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <NavLink
                    to={item.href}
                    className="flex h-12 w-12 items-center justify-center rounded-2xl bg-[#D9EFBD] text-[#335F68] transition-colors hover:bg-[#D8F0C8]"
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

