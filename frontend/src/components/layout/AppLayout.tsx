import { Outlet } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import { useAtomValue } from "jotai";
import BottomNav from "./BottomNav";

const AppLayout = () => {
  const user = useAtomValue(userAtom);

  return (
    <div className="flex flex-col min-h-screen">
      <main className="flex-1 w-full max-w-none sm:max-w-md md:max-w-lg lg:max-w-xl xl:max-w-2xl mx-auto px-4 pb-24">
        <Outlet />
      </main>
      {user && <BottomNav />}
    </div>
  );
};

export default AppLayout;
