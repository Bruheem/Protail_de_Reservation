import React from "react";
import Header from "../../components/Global/Header";

const AdminDashboard: React.FC = () => {
  return (
    <div className="d-flex">
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
