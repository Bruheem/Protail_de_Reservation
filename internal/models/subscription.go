package models

import (
	"database/sql"
	"fmt"
	"log"
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
		log.Printf("Error executing query: %v", err)
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

func (s *SubscriptionModel) GetSubscriptions(userID int64) ([]*Library, error) {
	query := `SELECT l.LibraryID, l.Name, l.CreatedBy
	FROM subscription s
	JOIN library l ON s.libraryID = l.LibraryID
	WHERE s.userID = ?`

	rows, err := s.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libs []*Library
	for rows.Next() {
		lib := &Library{}
		if err := rows.Scan(&lib.ID, &lib.Name, &lib.CreatedBy); err != nil {
			fmt.Println(err)
			return nil, err
		}
		libs = append(libs, lib)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return libs, nil
}
