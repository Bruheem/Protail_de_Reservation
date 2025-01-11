package models

import (
	"database/sql"
	"errors"
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

func (m *DocumentModel) SearchDocuments(query, genre, docType string, libraryID int) ([]*Document, error) {

	querySQL := `
		SELECT d.*, g.name AS genre, dt.documentTypeName AS document_type
		FROM document d
		LEFT JOIN DocGenres dg ON d.DocumentID = dg.doc_id
		LEFT JOIN Genres g ON dg.genre_id = g.id
		LEFT JOIN documentType dt ON d.documentTypeID = dt.documentTypeID
		WHERE
			(d.title LIKE CONCAT('%', ?, '%') OR d.author LIKE CONCAT('%', ?, '%'))
			AND (? IS NULL OR g.name = ?)
			AND (? IS NULL OR dt.documentTypeName = ?)
			AND (? IS NULL OR d.libraryID = ?)
		ORDER BY d.title
	`

	rows, err := m.DB.Query(querySQL, query, query, genre, genre, docType, docType, libraryID, libraryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []*Document
	for rows.Next() {
		doc := &Document{}
		err := rows.Scan(&doc.ID, &doc.Title, &doc.Author, &doc.YearPublished, &doc.ISBN, &doc.LibraryID, &doc.DocumentTypeID)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}

func (m *DocumentModel) GetPopular() ([]*Document, error) {
	sqlQuery := `
		SELECT d.DocumentID, d.title, COUNT(l.id) AS num_borrows
		FROM document d
		LEFT JOIN lending l ON d.DocumentID = l.document_id
		GROUP BY d.id
		ORDER BY num_borrows DESC
		LIMIT 10
	`

	rows, err := m.DB.Query(sqlQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documents []*Document
	for rows.Next() {
		doc := &Document{}
		err := rows.Scan(&doc.ID, &doc.Title, &doc.NumBorrows)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	return documents, nil
}
