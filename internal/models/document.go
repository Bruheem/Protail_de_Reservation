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
