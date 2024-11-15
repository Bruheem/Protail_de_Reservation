package models

import (
	"database/sql"
	"errors"
)

type Library struct {
	ID        uint64
	Name      string
	CreatedBy string
}

type LibraryModel struct {
	DB *sql.DB
}

func (m *LibraryModel) GetLibrary(id uint64) (*Library, error) {

	query := "SELECT LibraryID, Name, CreatedBy FROM library WHERE LibraryID = ?"

	row := m.DB.QueryRow(query, id)
	lib := &Library{}

	err := row.Scan(&lib.ID, &lib.Name, &lib.CreatedBy)
	if err != nil {
		return nil, err
	}

	return lib, nil
}

func (m *LibraryModel) InsertLibrary(title, createdby string) (uint64, error) {

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

func (m *LibraryModel) UpdateLibrary(lib *Library) {
	query := "UPDATE library SET Name = ?, CreatedBy = ? WHERE LibraryID = ?"

	args := []interface{}{
		lib.Name,
		lib.CreatedBy,
		lib.ID,
	}

	m.DB.QueryRow(query, args...)
}

func (m *LibraryModel) DeleteLibrary(id uint64) error {
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
