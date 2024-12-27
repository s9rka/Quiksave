import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "../ui/button";
import { useNavigate } from "react-router-dom";
import { useLogout } from "@/services/mutations";

const AccountDropdown = () => {
  const navigate = useNavigate();
  const logoutMutation = useLogout()

  const handleLogout = async () => {
    logoutMutation.mutate()
  };

  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger>
          <Button variant="ghost">Account</Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <div className="p-4 flex flex-col gap-2">
            <Button onClick={handleLogout}>Logout</Button>
            <Button variant="outline" onClick={() => navigate("/settings")}>
              Settings
            </Button>
          </div>
        </DropdownMenuContent>
      </DropdownMenu>
    </div>
  );
};

export default AccountDropdown;
