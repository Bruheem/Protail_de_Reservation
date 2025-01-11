import React from "react";
import Header from "../../components/Global/Header";
import Sidebar from "../../components/Global/Sidebar";

const UserDashboard: React.FC = () => {
  const links = [
    { name: "Search Libraries", path: "/search" },
    { name: "Borrowed Documents", path: "/borrowed" },
    { name: "Recommendations", path: "/recommendations" },
  ];

  return (
    <div className="d-flex">
      <Sidebar links={links} />
      <div className="flex-grow-1">
        <Header />
        <div className="container mt-4">
          <h2>User Dashboard</h2>
          <p className="lead">Manage your library interactions:</p>
          <ul>
            <li>Search for libraries</li>
            <li>View borrowed documents</li>
            <li>Receive personalized recommendations</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default UserDashboard;
