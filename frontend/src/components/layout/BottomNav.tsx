import { NavLink } from "react-router-dom";
import { Plus, Library } from "lucide-react";
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip";

const navItems = [
  { icon: Library, label: "Storage", href: `/` },
  { icon: Plus, label: "Create Note", href: "/new" },
];

export default function EnhancedBottomNav() {
  return (
    <nav className="fixed bottom-2 left-1/2 transform -translate-x-1/2 z-50 bg-white/80">
      <ul className="flex justify-around gap-1 w-auto sm:max-w-sm items-center py-2 px-6">
        {navItems.map((item) => (
          <li key={item.href} className="relative group">
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger>
                  <NavLink
                    to={item.href}
                    className={({ isActive }) =>
                      `flex flex-col items-center p-3 transition-all duration-300 border rounded ${
                        isActive
                          ? "bg-gray-300 text-gray-900"
                          : "text-gray-600 hover:text-gray-900"
                      }`
                    }
                  >
                    <item.icon className="w-6 h-6 relative z-10" />
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
