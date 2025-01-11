import { useAuth } from "./AuthContext";
import { Navigate } from "react-router-dom";

type ProtectedRouteProps = {
    children: JSX.Element;
    allowedRoles: string[];
}

const ProctectedRoute: React.FC<ProtectedRouteProps> = ({children, allowedRoles}) => {
    const { user } = useAuth();

    if (!user || !allowedRoles.includes(user.role)) {
        return <Navigate to="/login" />
    }

    return <>{children}</>;
};

export default ProctectedRoute;
