import React, { useState, useEffect } from "react";
import RecommendationCard from "../../components/Global/RecommendationCard";

const Recommendations: React.FC = () => {
  const [recommendedDocs, setRecommendedDocs] = useState<any[]>([]);
  const [recommendedLibs, setRecommendedLibs] = useState<any[]>([]);

  useEffect(() => {
    const fetchRecommendations = async () => {
      const docRes = await fetch("http://localhost:4000/v1/recommendations/documents");
      const libRes = await fetch("http://localhost:4000/v1/recommendations/libraries");

      const docs = await docRes.json();
      const libs = await libRes.json();

      setRecommendedDocs(docs.recommended_documents); 
      setRecommendedLibs(libs.recommended_libraries);
    };

    fetchRecommendations();
  }, []);

  return (
    <div>
      <h2>Recommended Libraries</h2>
      <div className="row">
        {recommendedLibs.map(lib => (
          <div key={lib.ID} className="col-md-4 mb-4">
            <div className="card">
              <div className="card-body">
                <h5 className="card-title">{lib.Name}</h5>
                <p className="card-text">
                  Subscribers: {lib.NumSubscribers}
                </p>
                <a href={`/libraries/${lib.ID}`} className="btn btn-primary">
                  View Library
                </a>
              </div>
            </div>
          </div>
        ))}
      </div>

      <h2>Recommended Documents</h2>
      <div className="row">
        {recommendedDocs.map(doc => (
          <RecommendationCard key={doc.ID} document={doc} />
        ))}
      </div>
    </div>
  );
};

export default Recommendations;
