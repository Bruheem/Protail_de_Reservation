import React, { useState } from "react";
import { useAuth } from "../../utils/AuthContext"; // Adjust the import path as necessary
import { useNavigate } from "react-router-dom";

const Profile: React.FC = () => {
  const { user, logout, setUser } = useAuth(); // Assuming `setUser` is available in context
  const navigate = useNavigate();

  const [formData, setFormData] = useState({
    id: user?.id,
    username: user?.username || "",
    email: user?.email || "",
    password: "",
  });

  const [statusMessage, setStatusMessage] = useState<string | null>(null);
  const [isSuccess, setIsSuccess] = useState<boolean | null>(null);

  const handleLogout = () => {
    logout();
    navigate("/auth");
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    try {
      const response = await fetch("http://localhost:4000/v1/users", {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(formData),
      });

      if (response.ok) {
        // Successfully updated, set the user in state and localStorage
        const updatedUser = await response.json();
        setStatusMessage("Profile updated successfully!");
        setIsSuccess(true);

        // Update the user context and localStorage with the new data
        setUser(updatedUser); // Update user in context
        localStorage.setItem("user", JSON.stringify(updatedUser)); // Update user in localStorage
      } else {
        const errorData = await response.json();
        setStatusMessage(errorData.message || "Failed to update profile.");
        setIsSuccess(false);
      }
    } catch (error) {
      setStatusMessage("An error occurred. Please try again later.");
      setIsSuccess(false);
      console.error("Error updating profile:", error);
    }
  };

  return (
    <div className="container mt-5">
      <div className="row">
        {/* Profile Section */}
        <div className="col-md-6 d-flex justify-content-center">
          <div className="card rounded shadow-lg" style={{ width: "400px" }}>
            <div className="card-body text-center">
              {/* Profile Picture */}
              <img
                src={`https://robohash.org/${user?.username || "user"}`}
                alt="Profile"
                className="img-fluid rounded-circle mb-3"
                style={{ width: "150px", height: "150px" }}
              />
              <h4 className="card-title">{user?.username}</h4>
              <p className="card-text text-muted">{user?.email}</p>
              <button className="btn btn-danger mt-3" onClick={handleLogout}>
                Logout
              </button>
            </div>
          </div>
        </div>

        {/* Update Form Section */}
        <div className="col-md-6">
          <div className="card rounded shadow-lg p-4">
            <h5 className="card-title mb-4">Update Profile</h5>
            <form onSubmit={handleSubmit}>
              <div className="mb-3">
                <label htmlFor="username" className="form-label">
                  Username
                </label>
                <input
                  type="text"
                  className="form-control"
                  id="username"
                  name="username"
                  value={formData.username}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="mb-3">
                <label htmlFor="email" className="form-label">
                  Email
                </label>
                <input
                  type="email"
                  className="form-control"
                  id="email"
                  name="email"
                  value={formData.email}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className="mb-3">
                <label htmlFor="password" className="form-label">
                  Password
                </label>
                <input
                  type="password"
                  className="form-control"
                  id="password"
                  name="password"
                  value={formData.password}
                  onChange={handleChange}
                  required
                />
              </div>
              <button type="submit" className="btn btn-primary w-100">
                Apply Changes
              </button>
            </form>
            {statusMessage && (
              <div
                className={`alert mt-3 ${
                  isSuccess ? "alert-success" : "alert-danger"
                }`}
              >
                {statusMessage}
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  );
};

export default Profile;
