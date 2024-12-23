import { Navigate, Outlet } from "react-router-dom";
import { useAtom } from "jotai";
import { isAuthenticatedAtom } from "@/services/auth";

export const ProtectedRoute = () => {
  const [isAuthenticated] = useAtom(isAuthenticatedAtom);

  if (!isAuthenticated) {
    return <Navigate to="/" replace />;
  }

  return <Outlet />;
};
