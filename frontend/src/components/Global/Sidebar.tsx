// src/Sidebar.tsx
import React from "react";
import { NavLink } from "react-router-dom";

interface SidebarProps {
  role: "admin" | "libadmin" | "user";
}

const Sidebar: React.FC<SidebarProps> = ({ role }) => {
  return (
    <div
      className="sidebar bg-dark text-white p-3 border-right shadow-sm"
      style={{ minWidth: "250px", height: "100vh" }}
    >
      <h2 className="text-center mt-4 mb-5 font-weight-bold">
        {role === "admin"
          ? "Admin Sidebar"
          : role === "libadmin"
          ? "Library Admin Sidebar"
          : "User Sidebar"}
      </h2>
      <ul className="nav flex-column mt-3">
        {role === "admin" && (
          <>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/admin/dashboard"
              >
                <i className="fas fa-tachometer-alt mr-2"></i> Dashboard
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/admin/users"
              >
                <i className="fas fa-users mr-2"></i> Manage Users
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/admin/settings"
              >
                <i className="fas fa-cogs mr-2"></i> Settings
              </NavLink>
            </li>
          </>
        )}
        {role === "libadmin" && (
          <>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/libadmin/dashboard"
              >
                <i className="fas fa-tachometer-alt mr-2"></i> Dashboard
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/libadmin/books"
              >
                <i className="fas fa-book mr-2"></i> Manage Books
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/libadmin/settings"
              >
                <i className="fas fa-cogs mr-2"></i> Settings
              </NavLink>
            </li>
          </>
        )}
        {role === "user" && (
          <>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/user/profile"
              >
                <i className="fas fa-user mr-2"></i> Profile
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/user/books"
              >
                <i className="fas fa-book mr-2"></i> My Books
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/user/settings"
              >
                <i className="fas fa-cogs mr-2"></i> Settings
              </NavLink>
            </li>
          </>
        )}
      </ul>
    </div>
  );
};

export default Sidebar;
