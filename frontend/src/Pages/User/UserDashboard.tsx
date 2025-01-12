import React from "react";
import { Routes, Route, Link } from "react-router-dom";
import Header from "../../components/Global/Header";
import Sidebar from "../../components/Global/Sidebar";
import Recommendations from "./Recommendations";
import Search from "./Search";
import Profile from "./Profile";

const UserDashboard: React.FC = () => {
  const userRoutes = [
    { path: "/dashboard/user", name: "Recommendations" },
    { path: "/dashboard/user/search", name: "Search" },
    { path: "/dashboard/user/collections", name: "Collections" },
    { path: "/dashboard/user/profile", name: "Profile" },
  ];

  return (
    <div className="dashboard-container">
      <Header />
      <div className="dashboard-content d-flex">
        <div className="sidebar bg-light p-3">
          <ul className="nav flex-column">
            {userRoutes.map((route) => (
              <li className="nav-item" key={route.path}>
                <Link to={route.path} className="nav-link">
                  {route.name}
                </Link>
              </li>
            ))}
          </ul>
        </div>
        <div className="flex-grow-1 p-3">
          <Routes>
            <Route path="/" element={<Recommendations />} />
            <Route path="/search" element={<Search />} />
            <Route path="/profile" element={<Profile />} />
          </Routes>
        </div>
      </div>
    </div>
  );
};

export default UserDashboard;
