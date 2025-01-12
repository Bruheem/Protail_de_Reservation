import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./Pages/Login";
import Register from "./Pages/Register";
import ProctectedRoute from "./utils/ProtectedRoute";
import { AuthProvider } from "./utils/AuthContext";
import UserDashboard from "./Pages/User/UserDashboard";
import LibrarianDashboard from "./Pages/Librarian/LibrarianDashboard";
import AdminDashboard from "./Pages/Admin/AdminDashboard";

const App: React.FC = () => {
  return (
    <AuthProvider>
      <Router>
        <Routes>
          <Route path="/login" element={<Login></Login>} />
          <Route path="/register" element={<Register></Register>} />
          <Route
            path="/dashboard/user"
            element={
              <ProctectedRoute allowedRoles={["user"]}>
                <UserDashboard />
              </ProctectedRoute>
            }
          />

          <Route
            path="/dashboard/librarian"
            element={
              <ProctectedRoute allowedRoles={["librarian"]}>
                <LibrarianDashboard />
              </ProctectedRoute>
            }
          />

          <Route
            path="/dashboard/admin"
            element={
              <ProctectedRoute allowedRoles={["admin"]}>
                <AdminDashboard />
              </ProctectedRoute>
            }
          />
        </Routes>
      </Router>
    </AuthProvider>
  );
};

export default App;
