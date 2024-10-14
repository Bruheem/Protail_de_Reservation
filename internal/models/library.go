package models

import "database/sql"

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
