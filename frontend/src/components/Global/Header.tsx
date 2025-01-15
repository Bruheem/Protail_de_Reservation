import React from "react";
import { useAuth } from "../../utils/AuthContext";
import { useNavigate } from "react-router-dom";

const Header: React.FC = () => {
  const { user, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = () => {
    logout();
    navigate("/auth");
  };

  return (
    <nav className="navbar navbar-expand-lg navbar-dark bg-primary shadow">
      <div className="container-fluid">
        <a className="navbar-brand" href="/user">
          <span className="ms-2">PDRE</span>
        </a>
        <div className="d-flex justify-content-between align-items-center w-100">
          <div className="flex-grow-1 text-center">
            <span className="text-white">Welcome, {user?.username}</span>
          </div>
          <button
            className="btn btn-outline-light btn-sm"
            onClick={handleLogout}
          >
            Logout
          </button>
        </div>
      </div>
    </nav>
  );
};

export default Header;
