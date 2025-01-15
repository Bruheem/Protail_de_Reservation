import React, { useState } from "react";
import Header from "../../components/Global/Header";
import Sidebar from "../../components/Global/Sidebar";
import Recommendations from "./Recommendations";
import Search from "./Search";
import Collections from "./Collections";
import Profile from "./Profile";

const UserDashboard: React.FC = () => {
  const [activeComponent, setActiveComponent] =
    useState<string>("recommendations");

  const renderComponent = () => {
    switch (activeComponent) {
      case "recommendations":
        return <Recommendations />;
      case "search":
        return <Search />;
      case "collections":
        return <Collections />;
      case "profile":
        return <Profile />;
      default:
        return <Recommendations />;
    }
  };

  const userRoutes = [
    { path: "recommendations", name: "Recommendations" },
    { path: "search", name: "Search" },
    { path: "collections", name: "Collections" },
    { path: "profile", name: "Profile" },
  ];

  return (
    <div className="dashboard-container">
      <Header />
      <div className="dashboard-content d-flex">
        <div className="sidebar bg-light p-3">
          <Sidebar
            routes={userRoutes}
            onSelect={(route) => setActiveComponent(route.path)}
          />
        </div>
        <div className="flex-grow-1 p-3">{renderComponent()}</div>
      </div>
    </div>
  );
};

export default UserDashboard;
