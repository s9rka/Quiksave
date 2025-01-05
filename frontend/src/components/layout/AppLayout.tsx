import { Outlet } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import { useAtomValue } from "jotai";
import BottomNav from "./BottomNav";

const AppLayout = () => {

  const user = useAtomValue(userAtom);

  return (
    <div>
      <header>
      </header>
      <main className="w-full max-w-none px-2 sm:max-w-sm md:max-w-md lg:max-w-lg xl:max-w-xl mx-auto">
        <Outlet />
        {user && <BottomNav />}
      </main>
    </div>
  );
};

export default AppLayout;
