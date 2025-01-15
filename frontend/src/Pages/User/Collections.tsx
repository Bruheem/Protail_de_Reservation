import React, { useEffect, useState } from "react";
import { useAuth } from "../../utils/AuthContext";
import { useNavigate } from "react-router-dom";

const Collections: React.FC = () => {
  const { user } = useAuth();
  const [libraries, setLibraries] = useState<any[]>([]);
  const [lentBooks, setLentBooks] = useState<any[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchCollections = async () => {
      try {
        const token = localStorage.getItem("authToken");

        const [libRes, bookRes] = await Promise.all([
          fetch(`http://localhost:4000/v1/libraries`, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }),
          fetch(`http://localhost:4000/v1/documents`, {
            headers: {
              Authorization: `Bearer ${token}`,
            },
          }),
        ]);

        if (!libRes.ok || !bookRes.ok) {
          throw new Error("Failed to fetch data");
        }

        const librariesData = await libRes.json();
        const lentBooksData = await bookRes.json();

        setLibraries(librariesData.libraries || []);
        setLentBooks(lentBooksData.borrowed_document || []);
      } catch (err) {
        if (err instanceof Error) {
          setError(err.message);
        } else {
          setError("An unknown error occurred");
        }
      } finally {
        setLoading(false);
      }
    };

    if (user) {
      fetchCollections();
    }
  }, [user]);

  const handleUnsubscribe = async (libraryId: string) => {
    try {
      const token = localStorage.getItem("authToken");

      const res = await fetch(`http://localhost:4000/v1/unsubscribe`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify({ library_id: libraryId }),
      });

      if (!res.ok) {
        throw new Error("Failed to unsubscribe from the library");
      }

      setLibraries((prevLibraries) =>
        prevLibraries.filter((lib) => lib.ID !== libraryId)
      );
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError("An unknown error occurred");
      }
    }
  };

  const handleReturn = async (bookId: string) => {
    try {
      const token = localStorage.getItem("authToken");

      const res = await fetch(`http://localhost:4000/v1/return/${bookId}`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) {
        throw new Error("Failed to return the book");
      }

      setLentBooks((prevBooks) =>
        prevBooks.filter((book) => book.ID !== bookId)
      );
    } catch (err) {
      if (err instanceof Error) {
        setError(err.message);
      } else {
        setError("An unknown error occurred");
      }
    }
  };

  if (loading) {
    return <div className="text-center">Loading...</div>;
  }

  if (error) {
    return <div className="text-danger text-center">{error}</div>;
  }

  return (
    <div className="container mt-5">
      <h2 className="text-center mb-4">Your Collections</h2>

      <div className="mb-5">
        <h4>Subscribed Libraries</h4>
        <div className="row">
          {libraries.length > 0 ? (
            libraries.map((lib, index) => (
              <div key={`library-${lib.ID}-${index}`} className="col-md-4 mb-4">
                <div className="card shadow-sm border-light">
                  <div className="card-body">
                    <h5 className="card-title">{lib.Name}</h5>
                    <p className="card-text">{lib.CreatedBy}</p>
                    <button
                      className="btn btn-primary"
                      onClick={() => navigate(`/libraries/${lib.ID}`)}
                    >
                      View Library
                    </button>
                    <button
                      className="btn btn-danger ml-2"
                      onClick={() => handleUnsubscribe(lib.ID)}
                    >
                      Unsubscribe
                    </button>
                  </div>
                </div>
              </div>
            ))
          ) : (
            <p>No subscribed libraries found.</p>
          )}
        </div>
      </div>

      <div>
        <h4>Lent Books</h4>
        <div className="row">
          {lentBooks.length > 0 ? (
            lentBooks
              .filter((book) => book.Status !== "returned")
              .map((book, index) => (
                <div key={`book-${book.ID}-${index}`} className="col-md-4 mb-4">
                  <div className="card shadow-sm border-light">
                    <div className="card-body">
                      <h5 className="card-title">{book.Title}</h5>
                      <p className="card-text">Author: {book.Author}</p>
                      <p className="card-text">
                        Borrow Date: {book.Borrow_date}
                      </p>
                      <p className="card-text">Due Date: {book.Due_date}</p>
                      <button
                        className="btn btn-secondary"
                        onClick={() => navigate(`/books/${book.DocumentID}`)}
                      >
                        View Book
                      </button>
                      <button
                        className="btn btn-danger ml-2"
                        onClick={() => handleReturn(book.DocumentID)}
                      >
                        Return
                      </button>
                    </div>
                  </div>
                </div>
              ))
          ) : (
            <p>No lent books found.</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default Collections;
