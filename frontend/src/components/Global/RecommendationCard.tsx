import React from "react";

interface RecommendationCardProps {
  document: {
    id: number;
    Title: string;
    Author: string;
    description: string;
  };
}

const RecommendationCard: React.FC<RecommendationCardProps> = ({
  document,
}) => {
  return (
    <div className="col-md-4 mb-4">
      <div className="card">
        <img
          src="/path/to/default-image.jpg"
          className="card-img-top"
          alt={document.Title}
        />
        <div className="card-body">
          <h5 className="card-title">{document.Title}</h5>
          <p className="card-text">{document.Author}</p>
          <a href={`/documents/${document.id}`} className="btn btn-primary">
            View Document
          </a>
        </div>
      </div>
    </div>
  );
};

export default RecommendationCard;
