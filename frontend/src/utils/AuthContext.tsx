import { createContext, useState, useContext, ReactNode } from "react";
import { useNavigate } from "react-router-dom";

type User = {
  id: number;
  username: string;
  email: string;
  role: "user" | "librarian" | "admin";
  token: string;
};

type AuthContextType = {
  user: User | null;
  setUser: (user: User | null) => void; // Added setUser
  login: (token: string, user: User) => void;
  logout: () => void;
  loading: boolean;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const UserProvider = ({ children }: { children: ReactNode }) => {
  const [user, setUser] = useState<User | null>(() => {
    const currentUser = localStorage.getItem("user");

    if (!currentUser) return null;

    const user1 = JSON.parse(currentUser);
    return user1;
  });

  const [loading, setLoading] = useState<boolean>(false);
  const navigate = useNavigate();

  const login = (token: string, user: User) => {
    setUser(user); // Update the user in the context

    localStorage.setItem("authToken", token);
    localStorage.setItem("user", JSON.stringify(user));

    navigateToDashboard(user.role);
  };

  const logout = () => {
    setUser(null); // Reset the user in the context
    localStorage.removeItem("authToken");
    localStorage.removeItem("user");
    navigate("/auth");
  };

  const navigateToDashboard = (role: string) => {
    if (role === "admin") {
      navigate("/admin");
    } else if (role === "librarian") {
      navigate("/librarian");
    } else if (role === "user") {
      navigate("/user");
    }
  };

  return (
    <AuthContext.Provider value={{ user, setUser, login, logout, loading }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
};
