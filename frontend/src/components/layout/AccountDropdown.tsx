import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { Button } from "../ui/button";
import { useAtom } from "jotai";
import { authAtom } from "@/services/auth";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

const AccountDropdown = () => {
  const [authState, setAuthState] = useAtom(authAtom);
  const navigate = useNavigate();

  useEffect(() => {
    if (!authState.token) {
      navigate("/");
    }
  }, [authState.token, navigate]);

  const logout = () => {
    setAuthState({ token: null, username: null });
    navigate("/");
  };

  if (!authState.token) return null;

  return (
    <div>
      <DropdownMenu>
        <DropdownMenuTrigger>
          <Button variant="ghost">Account</Button>
        </DropdownMenuTrigger>
        <DropdownMenuContent>
          <div className="p-4 flex flex-col gap-2">
            <Button onClick={logout}>Logout</Button>
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
