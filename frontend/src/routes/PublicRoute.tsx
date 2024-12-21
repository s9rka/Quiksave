import { ReactNode } from "react";
import { Navigate, Outlet } from "react-router-dom";

interface PublicRouteProps {
  children: ReactNode;
}

const PublicRoute = ({ children }: PublicRouteProps) => {
  const authToken = localStorage.getItem("authToken");
  const username = localStorage.getItem("username")
    ? JSON.parse(localStorage.getItem("username") as string)
    : null;
  if (authToken) {
    return <Navigate to={`/${username}`} replace />;
  }
  return <>{children}</>;
};

export const PublicRoutes = () => {
  return (
    <PublicRoute>
      <Outlet />
    </PublicRoute>
  );
};
