import { Outlet } from "react-router-dom";
import { userAtom } from "@/context/UserContext";
import NavDrawer from "./NavDrawer";
import { useAtomValue } from "jotai";

const AppLayout = () => {

  const user = useAtomValue(userAtom);

  return (
    <div>
      <header>
        <h1>My App</h1>
      </header>
      <main>
        <Outlet />
        {user && <NavDrawer />}
      </main>
    </div>
  );
};

export default AppLayout;
