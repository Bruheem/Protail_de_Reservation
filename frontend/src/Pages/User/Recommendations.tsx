import React, { useState, useEffect } from "react";
import RecommendationCard from "../../components/Global/RecommendationCard";
import RecommendationLibrary from "../../components/Global/RecommendationLibrary";

const Recommendations: React.FC = () => {
  const [recommendedDocs, setRecommendedDocs] = useState<any[]>([]);
  const [recommendedLibs, setRecommendedLibs] = useState<any[]>([]);

  useEffect(() => {
    const fetchRecommendations = async () => {
      const docRes = await fetch(
        "http://localhost:4000/v1/recommendations/documents"
      );
      const libRes = await fetch(
        "http://localhost:4000/v1/recommendations/libraries"
      );

      const docs = await docRes.json();
      const libs = await libRes.json();

      setRecommendedDocs(docs.recommended_documents);
      setRecommendedLibs(libs.recommended_libraries);
    };

    fetchRecommendations();
  }, []);

  return (
    <div className="container mt-5">
      {/* Recommended Libraries Section */}
      <div className="mb-5">
        <h2 className="mb-4">Recommended Libraries</h2>
        <div className="row g-4">
          {recommendedLibs.map((lib, index) => (
            <div className="col-md-3 mb-4" key={lib.ID}>
              <div className="p-2">
                <RecommendationLibrary
                  library_id={lib.ID}
                  library={lib}
                  index={index}
                />
              </div>
            </div>
          ))}
        </div>
      </div>

      {/* Separator */}
      <div className="my-5">
        <hr style={{ borderTop: "2px solid #007bff" }} />
      </div>

      {/* Recommended Documents Section */}
      <div>
        <h2 className="mt-5 mb-4">Recommended Documents</h2>
        <div className="row g-4">
          {recommendedDocs.map((doc, index) => (
            <div className="col-md-3 mb-4" key={doc.ID}>
              <div className="p-2">
                <RecommendationCard document={doc} index={index + 8} />
              </div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Recommendations;
