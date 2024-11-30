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
                to="/admin/libraries"
              >
                <i className="fas fa-cogs mr-2"></i> Manage Libraries
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
                to="/libadmin/documents"
              >
                <i className="fas fa-book mr-2"></i> Manage Documents
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/libadmin/profile"
              >
                <i className="fas fa-cogs mr-2"></i> Profile
              </NavLink>
            </li>
          </>
        )}
        {role === "user" && (
          <>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/user/home"
              >
                <i className="fas fa-user mr-2"></i> Home
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/user/collections"
              >
                <i className="fas fa-book mr-2"></i> Collections
              </NavLink>
            </li>
            <li className="nav-item">
              <NavLink
                className="nav-link text-white mb-2 rounded hover-shadow"
                to="/user/profile"
              >
                <i className="fas fa-cogs mr-2"></i> Profile
              </NavLink>
            </li>
          </>
        )}
      </ul>
    </div>
  );
};

export default Sidebar;
