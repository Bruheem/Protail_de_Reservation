import React from "react";

interface SidebarProps {
  routes: { path: string; name: string; icon?: React.ReactNode }[];
  onSelect: (route: { path: string; name: string }) => void;
}

const Sidebar: React.FC<SidebarProps> = ({ routes, onSelect }) => {
  const profileRoute = routes.find((route) => route.name === "Profile");
  const otherRoutes = routes.filter((route) => route.name !== "Profile");

  return (
    <div
      className="bg-light border-right"
      style={{
        minHeight: "90vh",
        position: "sticky",
        top: "0",
        zIndex: 100,
        display: "flex",
        flexDirection: "column", // Arrange elements vertically
        justifyContent: "space-between", // Space elements between start and end
        height: "92vh",
        overflowY: "auto", // Makes the sidebar scrollable if necessary
      }}
    >
      <div className="p-3">
        <h4 className="text-center">User Dashboard</h4>
      </div>
      <ul className="nav flex-column" style={{ flexGrow: 1 }}>
        {otherRoutes.map((route) => (
          <li
            className="nav-item"
            key={route.path}
            style={{ marginBottom: "12px" }} // Added margin between buttons
          >
            <button
              className="nav-link btn text-left"
              onClick={() => onSelect(route)}
              style={{
                width: "100%",
                borderRadius: "0.375rem", // Rounded corners for a modern look
                padding: "12px 15px",
                backgroundColor: "#28a745", // Vibrant Green
                color: "#fff", // White text for contrast
                transition:
                  "background-color 0.3s, color 0.3s, transform 0.2s ease-in-out", // Smooth transition
                border: "none", // Remove border
              }}
              onMouseEnter={(e) => {
                e.currentTarget.style.backgroundColor = "#218838"; // Darker green on hover
                e.currentTarget.style.transform = "scale(1.05)"; // Slightly enlarge the button
              }}
              onMouseLeave={(e) => {
                e.currentTarget.style.backgroundColor = "#28a745"; // Restore original color
                e.currentTarget.style.transform = "scale(1)"; // Reset size
              }}
            >
              {route.icon && <span className="mr-2">{route.icon}</span>}
              {route.name}
            </button>
          </li>
        ))}
      </ul>

      {profileRoute && (
        <div>
          <ul className="nav flex-column">
            <li
              className="nav-item"
              style={{ marginBottom: "12px" }} // Added margin here as well
            >
              <button
                className="nav-link btn text-left"
                onClick={() => onSelect(profileRoute)}
                style={{
                  width: "100%",
                  borderRadius: "0.375rem", // Same rounded corners for consistency
                  padding: "12px 15px",
                  backgroundColor: "#007bff", // Vibrant Blue
                  color: "#fff", // White text for contrast
                  transition:
                    "background-color 0.3s, color 0.3s, transform 0.2s ease-in-out", // Smooth transition
                  border: "none", // Remove border
                }}
                onMouseEnter={(e) => {
                  e.currentTarget.style.backgroundColor = "#0056b3"; // Darker blue on hover
                  e.currentTarget.style.transform = "scale(1.05)"; // Slightly enlarge the button
                }}
                onMouseLeave={(e) => {
                  e.currentTarget.style.backgroundColor = "#007bff"; // Restore original color
                  e.currentTarget.style.transform = "scale(1)"; // Reset size
                }}
              >
                {profileRoute.icon && (
                  <span className="mr-2">{profileRoute.icon}</span>
                )}
                {profileRoute.name}
              </button>
            </li>
          </ul>
        </div>
      )}
    </div>
  );
};

export default Sidebar;
