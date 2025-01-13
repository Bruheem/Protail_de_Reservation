CREATE TABLE DocGenres (
    doc_id INT,
    genre_id INT,
    PRIMARY KEY (doc_id, genre_id),
    FOREIGN KEY (doc_id) REFERENCES document(DocumentID) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (genre_id) REFERENCES Genres(id) ON DELETE CASCADE ON UPDATE CASCADE
)

