package models

import (
	"database/sql"
	"time"
)

type Subscription struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	LibraryID int64     `json:"library_id"`
	CreatedAt time.Time `json:"created_at"`
}

type SubscriptionModel struct {
	DB *sql.DB
}

func (s *SubscriptionModel) Exists(userID, libraryID int64) (bool, error) {

	query := `SELECT COUNT(*) FROM subscription WHERE userID = ? AND libraryID = ?`

	var count int
	err := s.DB.QueryRow(query, userID, libraryID).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (s *SubscriptionModel) Insert(userID, libraryID int64) error {

	query := `INSERT INTO subscription (userID, libraryID, subscriptionDate) VALUES (?, ?, UTC_TIMESTAMP())`

	_, err := s.DB.Exec(query, userID, libraryID)
	return err
}

func (s *SubscriptionModel) Delete(userID, libraryID int64) error {

	query := `DELETE FROM subscription WHERE userID = ? AND libraryID = ?`

	_, err := s.DB.Exec(query, userID, libraryID)
	return err
}
