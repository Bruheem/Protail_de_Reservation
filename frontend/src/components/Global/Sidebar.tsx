import React from "react";
import { NavLink } from "react-router-dom";

interface SidebarProps {
  links: { name: string; path: string }[];
}

const Sidebar: React.FC<SidebarProps> = ({ links }) => {
  return (
    <div className="bg-light border-end vh-100" style={{ width: "250px" }}>
      <div className="list-group list-group-flush">
        {links.map((link, idx) => (
          <NavLink
            key={idx}
            to={link.path}
            className={({ isActive }) =>
              isActive
                ? "list-group-item list-group-item-action active"
                : "list-group-item list-group-item-action"
            }
          >
            {link.name}
          </NavLink>
        ))}
      </div>
    </div>
  );
};

export default Sidebar;
