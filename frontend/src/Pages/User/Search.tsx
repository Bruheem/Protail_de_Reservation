import React, { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Search: React.FC = () => {
  const [query, setQuery] = useState("");
  const [libraries, setLibraries] = useState<any[]>([]);
  const [documents, setDocuments] = useState<any[]>([]);
  const navigate = useNavigate();

  const handleSearch = async () => {
    if (query.trim() === "") return;

    const [libRes, docRes] = await Promise.all([
      fetch(`http://localhost:4000/v1/search/libraries?query=${query}`),
      fetch(`http://localhost:4000/v1/search/documents?query=${query}`),
    ]);

    const libs = await libRes.json();
    const docs = await docRes.json();

    setLibraries(libs.libraries || []);
    setDocuments(docs.documents || []);
  };

  useEffect(() => {
    if (query.length > 2) {
      handleSearch();
    }
  }, [query]);

  return (
    <div className="container">
      <h2>Search Libraries and Documents</h2>
      <div className="mb-4">
        <input
          type="text"
          className="form-control"
          placeholder="Search for libraries or documents..."
          value={query}
          onChange={(e) => setQuery(e.target.value)}
        />
      </div>
      <div>
        {query && (
          <div>
            <h4>Libraries</h4>
            <div className="row">
              {libraries.map((lib) => (
                <div key={lib.ID} className="col-md-4 mb-4">
                  <div className="card">
                    <div className="card-body">
                      <h5 className="card-title">{lib.Name}</h5>
                      <p className="card-text">
                        Subscribers: {lib.NumSubscribers}
                      </p>
                      <button
                        className="btn btn-primary"
                        onClick={() => navigate(`/libraries/${lib.ID}`)}
                      >
                        View Library
                      </button>
                    </div>
                  </div>
                </div>
              ))}
            </div>

            <h4>Documents</h4>
            <div className="row">
              {documents.map((doc) => (
                <div key={doc.ID} className="col-md-4 mb-4">
                  <div className="card">
                    <div className="card-body">
                      <h5 className="card-title">{doc.Title}</h5>
                      <p className="card-text">{doc.Author}</p>
                      <button
                        className="btn btn-primary"
                        onClick={() => navigate(`/documents/${doc.ID}`)}
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
  );
};

export default Search;
