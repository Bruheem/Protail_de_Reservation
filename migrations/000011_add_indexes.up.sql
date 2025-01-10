CREATE UNIQUE INDEX idx_user_email ON user (email);

CREATE INDEX idx_user_role_id ON user (userRoleID);

CREATE FULLTEXT INDEX idx_library_name ON library (Name);

CREATE FULLTEXT INDEX idx_document_title ON document (title);

CREATE INDEX idx_document_author ON document (author);

CREATE INDEX idx_document_type_id ON document (documentTypeID);

CREATE INDEX idx_document_library_id ON document (LibraryID);

CREATE UNIQUE INDEX idx_genres_name ON Genres (name);

CREATE INDEX idx_docGenres_document_id ON DocGenres (doc_id);

CREATE INDEX idx_docGenres_genre_id ON DocGenres (genre_id);

CREATE INDEX idx_lending_user_id ON lending (user_id);

CREATE INDEX idx_lending_document_id ON lending (document_id);

CREATE INDEX idx_lending_due_date ON lending (due_date);

CREATE INDEX idx_lending_status ON lending (status);

CREATE INDEX idx_request_user_id ON request (userID);

CREATE INDEX idx_request_library_id ON request (libraryID);

CREATE INDEX idx_request_status ON request (requestStatus);

CREATE INDEX idx_subscription_user_id ON subscription (userID);

CREATE INDEX idx_subscription_library_id ON subscription (libraryID);

CREATE UNIQUE INDEX idx_userRole_role_name ON userRole (RoleName);
