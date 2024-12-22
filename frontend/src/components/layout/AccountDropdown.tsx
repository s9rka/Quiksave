import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "../ui/button";

const AccountDropdown = () => {
  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger>
          <Button variant="ghost">Account</Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <div className="p-4 flex flex-col gap-2">
            <Button>Logout</Button>
            <Button variant="outline">Settings</Button>
          </div>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

export default AccountDropdown;
