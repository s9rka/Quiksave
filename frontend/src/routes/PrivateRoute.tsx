import { ReactNode } from "react";
import { Navigate, Outlet } from "react-router-dom";

interface PrivateRouteProps {
  children: ReactNode;
}

const PrivateRoute = ({ children }: PrivateRouteProps) => {
  const authToken = localStorage.getItem("authToken");
  if (!authToken) {
    return <Navigate to="/" replace />;
  }
  return <>{children}</>;
}

export const PrivateRoutes = () => {
    return (
      <PrivateRoute>
        <Outlet />
      </PrivateRoute>
    );
  }
  
