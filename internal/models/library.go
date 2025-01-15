package models

import (
	"database/sql"
	"errors"
	"fmt"
)

type Library struct {
	ID             uint64
	Name           string
	CreatedBy      string
	NumSubscribers int
}

type LibraryModel struct {
	DB *sql.DB
}

func (m *LibraryModel) Get(id uint64) (*Library, error) {

	query := "SELECT LibraryID, Name, CreatedBy FROM library WHERE LibraryID = ?"

	row := m.DB.QueryRow(query, id)
	lib := &Library{}

	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedBy)
	if err != nil {
		return nil, err
	}

	return lib, nil
}

func (m *LibraryModel) Insert(title, createdby string) (uint64, error) {

	query := "INSERT INTO library (Name, CreatedBy) VALUES (?, ?)"
	result, err := m.DB.Exec(query, title, createdby)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), err
}

func (m *LibraryModel) Update(lib *Library) {
	query := "UPDATE library SET Name = ?, CreatedBy = ? WHERE LibraryID = ?"

	args := []interface{}{
		lib.Name,
		lib.CreatedBy,
		lib.ID,
	}

	m.DB.QueryRow(query, args...)
}

func (m *LibraryModel) Delete(id uint64) error {
	query := "DELETE FROM library WHERE LibraryID = ?"

	result, err := m.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("error: no records found")
	}

	return nil
}

func (m *LibraryModel) Search(query string) ([]*Library, error) {

	sqlQuery := `
		SELECT library.*, COUNT(subscription.subscriptionID) AS subscription_count
		FROM library
		LEFT JOIN subscription 
		ON library.LibraryID = subscription.libraryID
		WHERE library.Name LIKE CONCAT('%', ?, '%')
		GROUP BY library.LibraryID
		ORDER BY library.Name
	`

	rows, err := m.DB.Query(sqlQuery, "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libraries []*Library
	for rows.Next() {
		library := &Library{}

		if err := rows.Scan(&library.ID, &library.Name, &library.CreatedBy, &library.NumSubscribers); err != nil {
			return nil, err
		}

		libraries = append(libraries, library)
	}

	return libraries, nil
}

func (m *LibraryModel) GetPopular() ([]*Library, error) {
	query := `
		SELECT l.LibraryID, l.name, l.CreatedBy,COUNT(s.subscriptionID) AS num_subscribers
		FROM library l
		LEFT JOIN subscription s ON l.LibraryID = s.libraryID
		GROUP BY l.LibraryID
		ORDER BY num_subscribers DESC
		LIMIT 8
	`

	rows, err := m.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libraries []*Library
	for rows.Next() {
		lib := &Library{}
		err := rows.Scan(&lib.ID, &lib.Name, &lib.CreatedBy, &lib.NumSubscribers)
		if err != nil {
			return nil, err
		}
		libraries = append(libraries, lib)
	}

	return libraries, nil
}

func (m *LibraryModel) GetSubscribedLibraries(id int64) ([]*Library, error) {
	query := `
        SELECT l.LibraryID, l.Name, l.CreatedBy
        FROM library l
        JOIN subscription s ON l.LibraryID = s.libraryID
        WHERE s.userID = ?
    `

	rows, err := m.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var libraries []*Library
	for rows.Next() {
		lib := &Library{}
		if err := rows.Scan(&lib.ID, &lib.Name, &lib.CreatedBy); err != nil {
			fmt.Println(err)
			return nil, err
		}
		libraries = append(libraries, lib)
	}

	return libraries, nil
}
