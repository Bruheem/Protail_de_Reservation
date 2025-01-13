INSERT INTO library (Name, CreatedBy)
VALUES
    ('Central Library', 'admin_user'),
    ('Downtown Library', 'manager_john'),
    ('Campus Library', 'librarian_sarah'),
    ('Community Library', 'admin_jane'),
    ('Westside Library', 'manager_lee'),
    ('Eastside Library', 'librarian_emma'),
    ('Northwood Library', 'manager_kate'),
    ('Riverside Library', 'admin_michael'),
    ('Greenfield Library', 'librarian_david'),
    ('Harmony Library', 'manager_claire'),
    ('Sunrise Library', 'librarian_nina'),
    ('Lakeside Library', 'manager_oliver'),
    ('Hillside Library', 'admin_charles'),
    ('Oakwood Library', 'librarian_jessica'),
    ('Elmwood Library', 'manager_luke'),
    ('Seaside Library', 'admin_diana'),
    ('Brighton Library', 'librarian_amy'),
    ('Mountainview Library', 'manager_steve'),
    ('Skyline Library', 'librarian_rachel'),
    ('Grandview Library', 'admin_george');

INSERT INTO documentType (documentTypeName)
VALUES
    ('Book'),
    ('Magazine'),
    ('Newspaper'),
    ('Journal'),
    ('E-Book'),
    ('Research Paper');

INSERT INTO userRole (RoleName)
VALUES
    ('admin'),
    ('librarian'),
    ('user');

INSERT INTO document (title, author, yearPublished, ISBN, libraryID, documentTypeID)
VALUES
    ('Introduction to Algorithms', 'Thomas H. Cormen', 2009, '9780262033848', 1, 1),
    ('Clean Code', 'Robert C. Martin', 2008, '9780132350884', 1, 1),
    ('The Pragmatic Programmer', 'Andrew Hunt', 1999, '9780201616224', 2, 1),
    ('Design Patterns', 'Erich Gamma', 1994, '9780201633610', 2, 1),
    ('Time Magazine - June 2022', 'Editorial Team', 2022, '9780012920192', 3, 2),
    ('Nature Journal - Vol 23', 'Editorial Team', 2021, '9780034567892', 3, 4),
    ('The New York Times - Jan 2023', 'Editorial Team', 2023, '9780025897861', 4, 3),
    ('Eloquent JavaScript', 'Marijn Haverbeke', 2018, '9781593279509', 4, 5),
    ('Artificial Intelligence: A Modern Approach', 'Stuart Russell', 2020, '9780134610993', 5, 1),
    ('Deep Learning', 'Ian Goodfellow', 2016, '9780262035613', 5, 1),
    ('Effective Python', 'Brett Slatkin', 2021, '9780134853987', 6, 1),
    ('Scientific American - Feb 2023', 'Editorial Team', 2023, '9780060987654', 6, 4),
    ('Harry Potter and the Philosophers Stone', 'J.K. Rowling', 1997, '9780747532699', 7, 1),
    ('The Lord of the Rings', 'J.R.R. Tolkien', 1954, '9780261102385', 7, 1),
    ('The Economist - May 2022', 'Editorial Team', 2022, '9780041234567', 8, 2),
    ('Scientific Computing', 'Michael Heath', 2018, '9780072399103', 8, 1),
    ('National Geographic - April 2022', 'Editorial Team', 2022, '9780063456789', 9, 2),
    ('Digital Design', 'Morris Mano', 2017, '9780132129383', 9, 1),
    ('Python Crash Course', 'Eric Matthes', 2019, '9781593279288', 10, 5),
    ('Introduction to Databases', 'Hector Garcia-Molina', 2008, '9780130319963', 10, 1);

INSERT INTO Genres (name)
VALUES
    ('Fiction'),
    ('Non-Fiction'),
    ('Science'),
    ('Technology'),
    ('History'),
    ('Biography'),
    ('Fantasy'),
    ('Mystery'),
    ('Romance'),
    ('Education');


INSERT INTO DocGenres (doc_id, genre_id)
VALUES
    (1, 3),
    (1, 10),

    (2, 4),
    (2, 10),

    (3, 4),
    (3, 10),

    (4, 4),

    (5, 2),

    (6, 3),

    (7, 2),

    (8, 4),
    (8, 10),

    (9, 3),
    (9, 4),

    (10, 3),

    (11, 4),
    (11, 10),

    (12, 3),

    (13, 7),

    (14, 7),

    (15, 2),

    (16, 3),
    (16, 4),

    (17, 2),
    (17, 3),

    (18, 4),

    (19, 4),
    (19, 10),

    (20, 4),
    (20, 10);
