import { Drawer, DrawerTrigger, DrawerContent } from "@/components/ui/drawer";
import { Button } from "../ui/button";
import { ChevronsUpDown } from "lucide-react";
import AccountDropdown from "./AccountDropdown";
import { useNavigate } from "react-router-dom";

const NavDrawer = () => {
  const navigate = useNavigate()

  const handleNew = () => {
    navigate("/new")
  }

  const handleStorage = () => {
    navigate("/")
  }
  return (
    <Drawer>
      <DrawerTrigger asChild>
        <div className="fixed bottom-0 left-0 z-50">
          <Button variant="outline" className="m-4">
            <ChevronsUpDown />
          </Button>
        </div>
      </DrawerTrigger>
      <DrawerContent>
        <div className="flex w-full">
          <div className="flex w-1/3 h-20 items-center justify-center">
            <AccountDropdown />
          </div>
          <div className="flex w-1/3 items-center justify-center border-x border-gray-200">
            <Button onClick={handleNew} variant="ghost" className="text-center">
              New Note
            </Button>
          </div>
          <div className="flex w-1/3 items-center justify-center">
            <Button onClick={handleStorage} variant="ghost" className="text-center">
              Storage
            </Button>
          </div>
        </div>
      </DrawerContent>
    </Drawer>
  );
};

export default NavDrawer;
