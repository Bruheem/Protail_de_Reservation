import React from "react";
import Header from "../../components/Global/Header";
import Sidebar from "../../components/Global/Sidebar";

const LibrarianDashboard: React.FC = () => {
  const links = [
    { name: "Add Document", path: "/add-document" },
    { name: "Manage Library", path: "/manage-library" },
  ];

  return (
    <div className="d-flex">
      <Sidebar routes={links} />
      <div className="flex-grow-1">
        <Header />
        <div className="container mt-4">
          <h2>Librarian Dashboard</h2>
          <p className="lead">Manage your library's content:</p>
          <ul>
            <li>Create new documents</li>
            <li>Update existing documents</li>
          </ul>
        </div>
      </div>
    </div>
  );
};

export default LibrarianDashboard;
