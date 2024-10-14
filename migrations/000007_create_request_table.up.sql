CREATE TABLE IF NOT EXISTS request (
    requestID SERIAL PRIMARY KEY,
    userID INT NOT NULL,
    libraryID INT NOT NULL,
    requestStatus SET('accepted', 'rejected', 'ongoing'),
    requestDate DATE NOT NULL,

    FOREIGN KEY (libraryID) REFERENCES library(libraryID)  ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (userID) REFERENCES user(userID)  ON DELETE CASCADE ON UPDATE CASCADE,
);