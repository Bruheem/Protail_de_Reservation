CREATE TABLE IF NOT EXISTS lending (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    document_id INT,
    borrow_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    due_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP DEFAULT NULL,
    status ENUM ('borrowed', 'returned', 'overdue') DEFAULT 'borrowed',
    FOREIGN KEY (user_id) REFERENCES user (id) ON DELETE CASCADE,
    FOREIGN KEY (document_id) REFERENCES document (DocumentID) ON DELETE CASCADE
);
