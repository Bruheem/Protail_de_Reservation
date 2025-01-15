import React, { useState } from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import Modal from "react-bootstrap/Modal";
import { useAuth } from "../../utils/AuthContext"; // Ensure the auth context is imported for user info

interface RecommendationCardProps {
  library_id: number;
  library: {
    LibraryID: number;
    Name: string;
    CreatedBy: string;
    NumSubscribers: number;
  };
  index: number;
}

const RecommendationCard: React.FC<RecommendationCardProps> = ({
  library_id,
  library,
  index,
}) => {
  const [showModal, setShowModal] = useState(false);
  const [libraryDetails, setLibraryDetails] = useState<any>(null);
  const [isSubscribed, setIsSubscribed] = useState(false);

  const { user } = useAuth();

  const handleShow = () => setShowModal(true);
  const handleClose = () => setShowModal(false);

  const fetchLibraryDetails = async () => {
    try {
      const response = await fetch(
        `http://localhost:4000/v1/libraries/${library_id}`
      );
      const data = await response.json();
      setLibraryDetails(data);
    } catch (error) {
      console.error("Error fetching library details:", error);
    }
  };

  // Handle subscription action
  const handleSubscribe = async () => {
    const token = localStorage.getItem("authToken"); // Extract token from localStorage
    if (!token) {
      alert("You must be logged in to subscribe.");
      return;
    }

    try {
      const response = await fetch(`http://localhost:4000/v1/subscribe`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ library_id: library_id }),
      });
      const result = await response.json();

      if (response.ok) {
        setIsSubscribed(true);
        alert("Successfully subscribed to the library!");
      } else {
        alert(result.message || "Failed to subscribe.");
      }
    } catch (error) {
      console.error("Subscription error:", error);
      alert("An error occurred while subscribing.");
    }
  };

  return (
    <div>
      <Card style={{ width: "18rem" }}>
        <Card.Img
          variant="top"
          src={`https://picsum.photos/300/300?random=${index}`}
        />
        <Card.Body>
          <Card.Title>{library.Name}</Card.Title>
          <Card.Text>{library.CreatedBy}</Card.Text>
          <Card.Text>Subscribers: {library.NumSubscribers}</Card.Text>
          <Button
            variant="primary"
            onClick={() => {
              fetchLibraryDetails();
              handleShow();
            }}
          >
            View
          </Button>
        </Card.Body>
      </Card>

      {/* Modal for displaying detailed library info */}
      <Modal show={showModal} onHide={handleClose} size="lg">
        <Modal.Header closeButton>
          <Modal.Title>{library.Name} Details</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          {libraryDetails ? (
            <>
              <h5>Created By: {libraryDetails.CreatedBy}</h5>
              <p>{libraryDetails.Description}</p>
              <p>Location: {libraryDetails.Location}</p>
              {/* Add any additional library details here */}
            </>
          ) : (
            <p>Loading...</p>
          )}
        </Modal.Body>
        <Modal.Footer>
          <Button variant="secondary" onClick={handleClose}>
            Close
          </Button>
          <Button
            variant={isSubscribed ? "success" : "primary"}
            onClick={handleSubscribe}
            disabled={isSubscribed} // Disable the button if already subscribed
          >
            {isSubscribed ? "Subscribed" : "Subscribe"}
          </Button>
        </Modal.Footer>
      </Modal>
    </div>
  );
};

export default RecommendationCard;
