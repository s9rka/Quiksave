import { Navigate, Outlet } from "react-router-dom";
import { useAtomValue } from "jotai";
import { userAtom } from "@/context/UserContext";

export const PublicRoute = () => {
  const user = useAtomValue(userAtom);
  console.log("PublicRoute, user", user?.username);

  if (user) {
    return <Navigate to={`/${user.username}`} replace />;
  } else {
    return <Outlet />;
  }
};
