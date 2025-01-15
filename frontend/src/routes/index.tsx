import { Routes, Route } from "react-router-dom";
// import ProtectedRoute from "../utils/ProtectedRoute";
import Recommendations from "../Pages/User/Recommendations";
import Search from "../Pages/User/Search";
import ProtectedRoute from "../utils/ProtectedRoute";
import Login from "../Pages/Login";
import Register from "../Pages/Register";
import Profile from "../Pages/User/Profile";
import Collections from "../Pages/User/Collections";
import UserDashboard from "../Pages/User/UserDashboard";
import LibrarianDashboard from "../Pages/Librarian/LibrarianDashboard";
import AdminDashboard from "../Pages/Admin/AdminDashboard";

export default function AppRoutes() {
  return (
    <Routes>
      {/* public routes */}
      <Route path="/auth">
        <Route index element={<Login />} />
        <Route path="register" element={<Register />} />
      </Route>

      {/* user specific routes */}
      <Route path="/user" element={<ProtectedRoute allowedRoles={["user"]} />}>
        <Route index={true} element={<UserDashboard />} />
        <Route path="collections" element={<Collections />} />
        <Route path="search" element={<Search />} />
        <Route path="profile" element={<Profile />} />
        <Route path="recommendations" element={<Recommendations />} />
      </Route>

      {/* librarian specific routes */}
      <Route
        path="/librarian"
        element={<ProtectedRoute allowedRoles={["librarian"]} />}
      >
        <Route index={true} element={<LibrarianDashboard />} />
        <Route path="search" element={<Search />} />
        <Route path="profile" element={<Profile />} />
      </Route>
      {/* admin specific routes */}

      <Route
        path="/admin"
        element={<ProtectedRoute allowedRoles={["admin"]} />}
      >
        <Route index={true} element={<AdminDashboard />} />
        <Route path="search" element={<Search />} />
        <Route path="profile" element={<Profile />} />
      </Route>
    </Routes>
  );
}
