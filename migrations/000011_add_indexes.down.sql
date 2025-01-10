DROP INDEX idx_user_email ON user;

DROP INDEX idx_user_role_id ON user;

DROP INDEX idx_library_name ON library;

DROP INDEX idx_document_title ON document;

DROP INDEX idx_document_author ON document;

DROP INDEX idx_document_type_id ON document;

DROP INDEX idx_document_library_id ON document;

DROP INDEX idx_genres_name ON genres;

DROP INDEX idx_docGenres_document_id ON docGenres;

DROP INDEX idx_docGenres_genre_id ON docGenres;

DROP INDEX idx_lending_user_id ON lending;

DROP INDEX idx_lending_document_id ON lending;

DROP INDEX idx_lending_due_date ON lending;

DROP INDEX idx_lending_status ON lending;

DROP INDEX idx_request_user_id ON request;

DROP INDEX idx_request_library_id ON request;

DROP INDEX idx_request_status ON request;

DROP INDEX idx_subscription_user_id ON subscription;

DROP INDEX idx_subscription_library_id ON subscription;

DROP INDEX idx_userRole_role_name ON userRole;
