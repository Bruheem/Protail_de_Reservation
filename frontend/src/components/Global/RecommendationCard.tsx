import React, { useState } from "react";
import Button from "react-bootstrap/Button";
import Card from "react-bootstrap/Card";
import Modal from "react-bootstrap/Modal";

interface RecommendationCardProps {
  document: {
    ID: number;
    Title: string;
    Author: string;
    ISBN: string;
    LibraryID: number;
    LibraryName: string;
  };
  index: number;
}

const RecommendationCard: React.FC<RecommendationCardProps> = ({
  document,
  index,
}) => {
  const [modalVisible, setModalVisible] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [isBorrowed, setIsBorrowed] = useState(false);

  const handleBorrow = async () => {
    const token = localStorage.getItem("authToken");

    if (!token) {
      alert("You must be logged in to borrow a document.");
      return;
    }

    try {
      const response = await fetch(`http://localhost:4000/v1/borrow`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({
          document_id: document.ID,
          library_id: document.LibraryID,
        }),
      });

      if (!response.ok) {
        throw new Error("You have to subscribe to the library first!");
      }

      setIsBorrowed(true);
      alert("Successfully borrowed the document!");
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError("An unknown error occurred");
      }
    }
  };

  return (
    <>
      <Card style={{ width: "18rem" }}>
        <Card.Img
          variant="top"
          src={`https://picsum.photos/300/300?random=${index}`}
        />
        <Card.Body>
          <Card.Title>{document.Title}</Card.Title>
          <Card.Text>{document.Author}</Card.Text>
          <Button variant="primary" onClick={() => setModalVisible(true)}>
            View Document
          </Button>
        </Card.Body>
      </Card>

      {/* Modal for Document Details */}
      <Modal show={modalVisible} onHide={() => setModalVisible(false)}>
        <Modal.Header closeButton>
          <Modal.Title>{document.Title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <p>Author: {document.Author}</p>
          <p>ISBN: {document.ISBN}</p>
          <p>Library: {document.LibraryName}</p>
          {isBorrowed ? (
            <p>You have borrowed this document.</p>
          ) : (
            <Button variant="success" onClick={handleBorrow}>
              Borrow
            </Button>
          )}
          {error && <p className="text-danger">{error}</p>}
        </Modal.Body>
      </Modal>
    </>
  );
};

export default RecommendationCard;
