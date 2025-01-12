import React from "react";
import { Link } from "react-router-dom";

type SidebarRoute = {
  path: string;
  name: string;
};

type SidebarProps = {
  routes: SidebarRoute[];
};

const Sidebar: React.FC<SidebarProps> = ({ routes }) => {
  return (
    <div className="sidebar bg-light p-3">
      <ul className="nav flex-column">
        {routes.map((route) => (
          <li className="nav-item" key={route.path}>
            <Link to={route.path} className="nav-link">
              {route.name}
            </Link>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Sidebar;
