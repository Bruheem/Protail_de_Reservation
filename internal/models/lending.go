package models

import (
	"database/sql"
	"time"
)

type Lending struct {
	ID          int64     `json:"id"`
	User_id     int64     `json:"user_id"`
	Document_id int64     `json:"document_id"`
	Borrow_date time.Time `json:"borrow_date"`
	Due_date    time.Time `json:"due_date"`
	Return_date time.Time `json:"return_date"`
	Status      string    `json:"status"`
}

type BorrowedDocument struct {
	DocumentID    int64      `json:"document_id"`
	Title         string     `json:"title"`
	Author        string     `json:"author"`
	YearPublished int        `json:"year_published"`
	ISBN          string     `json:"isbn"`
	LibraryID     int64      `json:"library_id"`
	BorrowedAt    time.Time  `json:"borrowed_at"`
	DueAt         time.Time  `json:"due_at"`
	ReturnedAt    *time.Time `json:"returned_at,omitempty"`
	Status        string     `json:"status"`
}

type LendingModel struct {
	DB *sql.DB
}

func (m *LendingModel) GetBorrowedDocuments(userID int64) ([]*BorrowedDocument, error) {
	query := `
		SELECT 
			d.DocumentID AS document_id,
			d.title,
			d.author,
			d.yearPublished,
			d.ISBN,
			d.libraryID,
			l.borrowed_date,
			l.due_date,
			l.returned_date,
			l.status
		FROM 
			lending l
		JOIN 
			document d ON l.document_id = d.DocumentID
		WHERE 
			l.user_id = ?
		ORDER BY 
			l.borrowed_at DESC
	`

	rows, err := m.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var borrowedDocs []*BorrowedDocument
	for rows.Next() {
		doc := &BorrowedDocument{}
		err := rows.Scan(
			&doc.DocumentID,
			&doc.Title,
			&doc.Author,
			&doc.YearPublished,
			&doc.ISBN,
			&doc.LibraryID,
			&doc.BorrowedAt,
			&doc.DueAt,
			&doc.ReturnedAt,
			&doc.Status,
		)
		if err != nil {
			return nil, err
		}
		borrowedDocs = append(borrowedDocs, doc)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return borrowedDocs, nil
}

func (m *LendingModel) GetBorrowingID(id int64) (*Lending, error) {

	query := `SELECT id, user_id FROM lending WHERE id = ?`
	row := m.DB.QueryRow(query, id)
	borrow := &Lending{}

	err := row.Scan(&borrow.ID, &borrow.User_id)
	if err != nil {
		return nil, err
	}

	return borrow, nil
}

func (m *LendingModel) BorrowDocument(userID, documentID int64) (int64, error) {
	query := `
        INSERT INTO lending (user_id, document_id, borrowed_date, due_date, status)
        VALUES (?, ?, NOW(), DATE_ADD(NOW(), INTERVAL 14 DAY), ?)`
	result, err := m.DB.Exec(query, userID, documentID, "borrowed")
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (m *LendingModel) MarkAsReturned(borrowID int64) error {
	query := `
        UPDATE lending
        SET return_date = NOW(), status = 'returned'
        WHERE id = ?`
	_, err := m.DB.Exec(query, borrowID)
	return err
}
