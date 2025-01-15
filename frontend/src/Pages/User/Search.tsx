import React, { useState, useEffect } from "react";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.bundle.min.js";

const Search: React.FC = () => {
  const [query, setQuery] = useState("");
  const [libraries, setLibraries] = useState<any[]>([]);
  const [documents, setDocuments] = useState<any[]>([]);
  const [selectedItem, setSelectedItem] = useState<any>(null);
  const [userSubscriptions, setUserSubscriptions] = useState<any[]>([]);

  const handleSearch = async () => {
    if (query.trim() === "") return;

    try {
      const [libRes, docRes] = await Promise.all([
        fetch(`http://localhost:4000/v1/search/libraries?query=${query}`),
        fetch(`http://localhost:4000/v1/search/documents?query=${query}`),
      ]);

      const libs = await libRes.json();
      const docs = await docRes.json();

      setLibraries(libs.libraries || []);
      setDocuments(docs.documents || []);
    } catch (error) {
      console.error("Error fetching search results:", error);
    }
  };

  useEffect(() => {
    if (query.length > 2) {
      handleSearch();
    }
  }, [query]);

  useEffect(() => {
    const fetchSubscriptions = async () => {
      try {
        const token = localStorage.getItem("authToken");
        const response = await fetch(`http://localhost:4000/v1/subscriptions`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
        const data = await response.json();
        setUserSubscriptions(data.subscriptions || []);
      } catch (error) {
        console.error("Error fetching subscriptions:", error);
      }
    };

    fetchSubscriptions();
  }, []);

  const handleSubscribe = async (library_id: string) => {
    try {
      const token = localStorage.getItem("authToken");
      const response = await fetch(`http://localhost:4000/v1/subscribe`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ library_id }),
      });

      if (response.ok) {
        alert("Subscribed successfully!");
        const updatedSubscriptions = await fetch(
          `http://localhost:4000/v1/subscriptions`,
          {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }
        );
        const data = await updatedSubscriptions.json();
        setUserSubscriptions(data.subscriptions || []);
      } else {
        alert("You are already subscribed to this library!");
      }
    } catch (error) {
      console.error("Error subscribing to library:", error);
    }
  };

  const handleBorrow = async (document_id: string, library_id: string) => {
    try {
      const token = localStorage.getItem("authToken");
      const response = await fetch(`http://localhost:4000/v1/borrow`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ document_id, library_id }),
      });

      if (response.ok) {
        alert("Document borrowed successfully!");
      } else {
        alert("Failed to borrow the document. Please try again.");
      }
    } catch (error) {
      console.error("Error borrowing document:", error);
    }
  };

  const handleModalClose = () => {
    setSelectedItem(null);
  };

  const isSubscribedToLibrary = (libraryId: string) => {
    return userSubscriptions.some((sub) => sub.libraryId === libraryId);
  };

  return (
    <div className="container mt-5">
      <h2 className="text-center mb-4">Search Libraries and Documents</h2>
      <div className="d-flex justify-content-center mb-4">
        <input
          type="text"
          className="form-control w-50"
          placeholder="Search for libraries or documents..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
      </div>

      {query && (
        <div>
          <h4 className="text-center mb-3">Results</h4>
          <div className="row">
            {libraries.length > 0 && (
              <div className="col-12 mb-4">
                <h5 className="text-primary">Libraries</h5>
                <div className="row">
                  {libraries.map((lib) => (
                    <div key={lib.ID} className="col-md-4 mb-4">
                      <div className="card shadow-sm border-light h-100">
                        <div className="card-body">
                          <h5 className="card-title">{lib.Name}</h5>
                          <p className="card-text">
                            <strong>Subscribers:</strong> {lib.NumSubscribers}
                          </p>
                          <button
                            className="btn btn-primary"
                            onClick={() => setSelectedItem(lib)}
                          >
                            View Library
                          </button>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}

            {documents.length > 0 && (
              <div className="col-12 mb-4">
                <h5 className="text-primary">Documents</h5>
                <div className="row">
                  {documents.map((doc) => (
                    <div key={doc.ID} className="col-md-4 mb-4">
                      <div className="card shadow-sm border-light h-100">
                        <div className="card-body">
                          <h5 className="card-title">{doc.Title}</h5>
                          <p className="card-text">{doc.Author}</p>
                          <button
                            className="btn btn-primary"
                            onClick={() => setSelectedItem(doc)}
                          >
                            View Document
                          </button>
                        </div>
                      </div>
                    </div>
                  ))}
                </div>
              </div>
            )}
          </div>
        </div>
      )}

      {selectedItem && (
        <>
          {/* Modal Backdrop */}
          <div className="modal-backdrop show"></div>

          {/* Modal */}
          <div
            className="modal show d-block"
            tabIndex={-1}
            role="dialog"
            aria-labelledby="modalTitle"
          >
            <div className="modal-dialog modal-dialog-centered" role="document">
              <div className="modal-content">
                <div className="modal-header">
                  <h5 className="modal-title">
                    {selectedItem.Name || selectedItem.Title}
                  </h5>
                  <button
                    type="button"
                    className="btn-close"
                    onClick={handleModalClose}
                  ></button>
                </div>
                <div className="modal-body">
                  {selectedItem.Name ? (
                    <>
                      <p>
                        <strong>Subscribers:</strong>{" "}
                        {selectedItem.NumSubscribers}
                      </p>
                      <button
                        className="btn btn-success"
                        onClick={() => handleSubscribe(selectedItem.ID)}
                      >
                        Subscribe
                      </button>
                    </>
                  ) : (
                    <>
                      <p>
                        <strong>Author:</strong> {selectedItem.Author}
                      </p>
                      <p>
                        <strong>Library:</strong> {selectedItem.LibraryName}
                      </p>
                      <button
                        className="btn btn-warning"
                        onClick={() =>
                          handleBorrow(selectedItem.ID, selectedItem.LibraryID)
                        }
                      >
                        Borrow
                      </button>
                    </>
                  )}
                </div>
              </div>
            </div>
          </div>
        </>
      )}
    </div>
  );
};

export default Search;
