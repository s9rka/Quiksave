import { Outlet } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import { useAtomValue } from "jotai";
import BottomNav from "./ReadWriteNav";
import { useVault } from "@/context/VaultContext";

const AppLayout = () => {
  const user = useAtomValue(userAtom);
  const { vaultId } = useVault();

  console.log("vaultId: ",vaultId)

  return (
    <div className="flex flex-col min-h-screen">
      <main className="flex-1 w-full max-w-none sm:max-w-md md:max-w-lg lg:max-w-xl xl:max-w-2xl mx-auto">
        <Outlet />
      </main>
      {vaultId !== null && <BottomNav />}
    </div>
  );
};

export default AppLayout;
