import React from "react";
import Header from "../../components/Global/Header";
import Sidebar from "../../components/Global/Sidebar";

const AdminDashboard: React.FC = () => {
  const links = [
    { name: "Add Library", path: "/add-library" },
    { name: "Manage Users", path: "/manage-users" },
  ];

  return (
    <div className="d-flex">
      <Sidebar links={links} />
      <div className="flex-grow-1">
        <Header />
        <div className="container mt-4">
          <h2>Admin Dashboard</h2>
          <p className="lead">Manage libraries and system users:</p>
          <ul>
            <li>Add new libraries to the system</li>
            <li>Manage user roles and permissions</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default AdminDashboard;
