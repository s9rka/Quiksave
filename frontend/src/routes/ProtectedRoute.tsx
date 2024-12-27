import { Navigate, Outlet } from "react-router-dom";
import { useAtomValue } from "jotai";
import { userAtom } from "@/context/UserContext";

const ProtectedRoute = () => {
  const user = useAtomValue(userAtom);

  console.log("ProtectedRoute, user", user?.username);

  if (!user) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
};

export default ProtectedRoute;
