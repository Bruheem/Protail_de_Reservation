package models

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
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

func (m *DocumentModel) GetSuggestedContent(userID uint64) ([]Document, error) {

	// Get all user books

	query := `
		SELECT d.DocumentID
		FROM lending l
		JOIN user u ON l.user_id = u.id
		JOIN document d ON l.document_id = d.DocumentID
		WHERE l.user_id = ?
	`

	docs := []uint64{}
	result, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	for result.Next() {
		var docID uint64
		if err := result.Scan(&docID); err != nil {
			return nil, err
		}
		docs = append(docs, docID)
	}

	// Get all suggested content without the ones that the user has already lent

	suggestedDocs := []Document{}

	if len(docs) > 0 {
		IDs := make([]string, len(docs))
		for i, id := range docs {
			IDs[i] = fmt.Sprintf("%d", id)
		}
		docsString := strings.Join(IDs, ",")

		suggestQuery := fmt.Sprintf(`
			SELECT d.DocumentID, d.title, d.author
			FROM document d
			WHERE d.DocumentID NOT IN (%s)
			ORDER BY RAND()
			LIMIT 10
		`, docsString)

		rows, err := m.DB.Query(suggestQuery)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			var doc Document
			if err := rows.Scan(&doc.ID, &doc.Title, &doc.Author); err != nil {
				return nil, err
			}
			suggestedDocs = append(suggestedDocs, doc)
		}
	} else {
		// no books lent, no suggestions
		return suggestedDocs, nil
	}

	return suggestedDocs, nil
}
