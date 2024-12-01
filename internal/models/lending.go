package models

import (
	"database/sql"
	"time"
)

type Lending struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	DocumentID int       `json:"document_id"`
	BorrowDate time.Time `json:"borrow_date"`
	DueDate    time.Time `json:"due_date"`
	ReturnDate time.Time `json:"return_date"`
	Status     string    `json:"status"`
}

type LendingModel struct {
	DB *sql.DB
}

func (m *LendingModel) BorrowDocument(userID, documentID, dueDays int) error {

	dueDate := time.Now().AddDate(0, 0, dueDays)
	query := `
		INSERT INTO lending (user_id, document_id, borrow_date, due_date, status)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := m.DB.Exec(query, userID, documentID, time.Now(), dueDate, "borrowed")
	if err != nil {
		return err
	}

	return nil
}

func (m *LendingModel) ReturnDocument(id int) error {

	query := `
		UPDATE lending
		SET return_date = ?, status = ?
		WHERE id = ?
	`

	_, err := m.DB.Exec(query, time.Now(), "returned", id)
	if err != nil {
		return err
	}
	return nil
}

func (m *LendingModel) GetLendingStatus(id int) (string, error) {
	query := `
		SELECT status
		FROM lending
		WHERE id = ?
	`

	row := m.DB.QueryRow(query, id)

	var status string
	err := row.Scan(&status)
	if err != nil {
		return "", err
	}

	return status, nil
}
