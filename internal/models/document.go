package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type Document struct {
	ID             uint64
	Title          string
	Author         string
	ISBN           string
	YearPublished  int
	LibraryID      int
	DocumentTypeID int
	NumBorrows     int
	LibraryName    string
}

type DocumentModel struct {
	DB *sql.DB
}

func (m *DocumentModel) GetDocument(id uint64) (*Document, error) {
	query := "SELECT * FROM document  WHERE id = ?"

	row := m.DB.QueryRow(query, id)
	doc := &Document{}

	err := row.Scan(&doc.ID, &doc.Title, &doc.Author, &doc.YearPublished, &doc.ISBN, &doc.LibraryID, &doc.DocumentTypeID)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func (m *DocumentModel) InsertDocument(doc *Document) (uint64, error) {
	query := "INSERT INTO document (title, author, yearPublished, ISBN, libraryID, documentTypeID) VALUES (?, ?, ?, ?, ?, ?)"
	result, err := m.DB.Exec(query, doc.Title, doc.Author, doc.YearPublished, doc.ISBN, doc.LibraryID, doc.DocumentTypeID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), err
}

func (m *DocumentModel) UpdateDocument(doc *Document) {
	query := `UPDATE document SET title = ?, author = ?, yearPublished = ?, ISBN = ?, libraryID = ?, documentTypeID = ? WHERE id = ?`
	args := []interface{}{
		doc.Title,
		doc.Author,
		doc.YearPublished,
		doc.ISBN,
		doc.LibraryID,
		doc.DocumentTypeID,
		doc.ID,
	}
	m.DB.QueryRow(query, args...)
}

func (m *DocumentModel) DeleteDocument(id uint64) error {
	query := "DELETE FROM document WHERE id = ?"

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("error: document not found")
	}

	return nil
}

func (m *DocumentModel) IsAvailable(documentID uint64) (bool, error) {
	query := "SELECT DocumentID FROM document WHERE DocumentID = ?"

	_, err := m.DB.Exec(query, documentID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("document not found")
		}
		return false, err
	}

	return true, nil
}

func (m *DocumentModel) SearchDocuments(query string) ([]*Document, error) {

	querySQL := `
    SELECT d.*, l.Name
    FROM document d
	JOIN library l ON l.LibraryID = d.libraryID
    WHERE 
        (d.title LIKE CONCAT('%', ?, '%') OR d.author LIKE CONCAT('%', ?, '%'))
    ORDER BY d.title
	`

	rows, err := m.DB.Query(querySQL, query, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []*Document
	for rows.Next() {
		doc := &Document{}
		err := rows.Scan(&doc.ID, &doc.Title, &doc.Author, &doc.YearPublished, &doc.ISBN, &doc.LibraryID, &doc.DocumentTypeID, &doc.LibraryName)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (m *DocumentModel) GetPopular() ([]*Document, error) {
	sqlQuery := `
		SELECT d.DocumentID, d.title, d.author, d.libraryID, d.ISBN, li.Name, COUNT(l.id) AS num_borrows
		FROM document d
		LEFT JOIN lending l ON d.DocumentID = l.document_id
		JOIN library li ON li.LibraryID = d.libraryID
		GROUP BY d.DocumentID
		ORDER BY num_borrows DESC
		LIMIT 8
	`

	rows, err := m.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []*Document
	for rows.Next() {
		doc := &Document{}
		err := rows.Scan(&doc.ID, &doc.Title, &doc.Author, &doc.LibraryID, &doc.ISBN, &doc.LibraryName, &doc.NumBorrows)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

type BorrowedDocuments struct {
	DocumentID  int64
	Title       string
	Author      string
	Borrow_date time.Time
	Return_date time.Time
	Due_date    time.Time
	Status      string
}

func (m *DocumentModel) GetBorrowedDocuments(id int64) ([]*BorrowedDocuments, error) {

	query := `
        SELECT d.DocumentID, d.title, d.author, bd.borrow_date, bd.due_date, bd.return_date, bd.status
        FROM lending bd
        JOIN document d ON bd.document_id = d.DocumentID
        WHERE bd.user_id = ?
    `

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var docs []*BorrowedDocuments
	for rows.Next() {
		doc := &BorrowedDocuments{}
		if err := rows.Scan(&doc.DocumentID, &doc.Title, &doc.Author, &doc.Borrow_date, &doc.Due_date, &doc.Return_date, &doc.Status); err != nil {
			fmt.Println(err)
			return nil, err
		}
		docs = append(docs, doc)
	}

	return docs, nil
}
