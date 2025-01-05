import { Outlet } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import { useAtomValue } from "jotai";
import BottomNav from "./BottomNav";

const AppLayout = () => {
  const user = useAtomValue(userAtom);

  return (
    <>
      <header>
      </header>
      <main className="relative w-full max-w-none sm:max-w-sm md:max-w-md lg:max-w-lg xl:max-w-xl mx-auto">
        <Outlet />
        {user && <BottomNav />}
      </main>
    </>
  );
};

export default AppLayout;
