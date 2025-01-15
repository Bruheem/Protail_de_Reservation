// import React, { JSX } from "react";
import { Navigate, Outlet } from "react-router-dom";
import { useAuth } from "./AuthContext";

interface ProtectedRouteProps {
  allowedRoles: string[];
}

const ProtectedRoute = ({ allowedRoles }: ProtectedRouteProps) => {
  const { user, loading } = useAuth();

  if (loading) return <div>Loading...</div>;

  if (!allowedRoles.includes(user?.role as string)) {
    return <Navigate to="/auth" />;
  }

  return <Outlet />;
};

export default ProtectedRoute;
