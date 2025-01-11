INSERT INTO library (Name, CreatedBy) VALUES 
('Benbeghdad Library', 'Ibrahim Ben.'),
('Baba Library', 'Baba Farouk'),
('MatMat Library', 'Khooled Mat.');

INSERT INTO documentType (documentTypeName) VALUES 
('book'),
('audio book'),
('video book'),
('document');

INSERT INTO userRole (userRoleID, RoleName) VALUES 
(1, 'Admin'),
(2, 'LibAdmin'),
(3, 'User ');

INSERT INTO Genres (name) VALUES 
('adventure'),
('romance'),
('horror'),
('comedy'),
('fiction'),
('scientific'),
('history'),
('novel'),
('science-fiction');