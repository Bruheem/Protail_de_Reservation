package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserSeeder struct {
	DB *sql.DB
}

type LibrarySeeder struct {
	DB *sql.DB
}

type DocumentSeeder struct {
	DB *sql.DB
}

type GenreSeeder struct {
	DB *sql.DB
}

type SubscriptionSeeder struct {
	DB *sql.DB
}

type DocumentTypeSeeder struct {
	DB *sql.DB
}

type UserRoleSeeder struct {
	DB *sql.DB
}

type Seeder struct {
	UserSeeder         *UserSeeder
	LibrarySeeder      *LibrarySeeder
	DocumentSeeder     *DocumentSeeder
	GenreSeeder        *GenreSeeder
	SubscriptionSeeder *SubscriptionSeeder
	DocumentTypeSeeder *DocumentTypeSeeder
	UserRoleSeeder     *UserRoleSeeder
}

func NewSeeder(db *sql.DB) *Seeder {
	return &Seeder{
		UserSeeder:         &UserSeeder{DB: db},
		LibrarySeeder:      &LibrarySeeder{DB: db},
		DocumentSeeder:     &DocumentSeeder{DB: db},
		GenreSeeder:        &GenreSeeder{DB: db},
		SubscriptionSeeder: &SubscriptionSeeder{DB: db},
		DocumentTypeSeeder: &DocumentTypeSeeder{DB: db},
		UserRoleSeeder:     &UserRoleSeeder{DB: db},
	}
}

func (s *DocumentTypeSeeder) Seed() error {
	documentTypes := []struct {
		Type string
	}{
		{"Book"},
		{"Journal"},
		{"Magazine"},
		{"Newspaper"},
	}

	for _, docType := range documentTypes {
		query := `INSERT INTO documentType (documentTypeName) VALUES (?)`
		_, err := s.DB.Exec(query, docType.Type)
		if err != nil {
			return fmt.Errorf("error seeding document types: %v", err)
		}
	}
	return nil
}

func (s *LibrarySeeder) Seed() error {
	libraries := []struct {
		Name      string
		CreatedBy string
	}{
		{"Central Library", "admin"},
		{"City Library", "johndoe"},
	}

	for _, library := range libraries {
		query := `INSERT INTO library (Name, CreatedBy) VALUES (?, ?)`
		_, err := s.DB.Exec(query, library.Name, library.CreatedBy)
		if err != nil {
			return fmt.Errorf("error seeding libraries: %v", err)
		}
	}
	return nil
}

func (s *DocumentSeeder) Seed() error {
	documents := []struct {
		Title          string
		Author         string
		YearPublished  int
		ISBN           string
		LibraryID      int
		DocumentTypeID int
	}{
		{"Golang Basics", "John Smith", 2020, "1234567890", 1, 1},
		{"Advanced Golang", "Jane Doe", 2021, "9876543210", 1, 2},
	}

	for _, doc := range documents {
		query := `INSERT INTO document (title, author, yearPublished, ISBN, libraryID, documentTypeID) 
                  VALUES (?, ?, ?, ?, ?, ?)`
		_, err := s.DB.Exec(query, doc.Title, doc.Author, doc.YearPublished, doc.ISBN, doc.LibraryID, doc.DocumentTypeID)
		if err != nil {
			return fmt.Errorf("error seeding documents: %v", err)
		}
	}
	return nil
}

func (s *GenreSeeder) Seed() error {
	genres := []struct {
		Name string
	}{
		{"Fiction"},
		{"Non-Fiction"},
		{"Science"},
		{"Technology"},
	}

	for _, genre := range genres {
		query := `INSERT INTO Genres (name) VALUES (?)`
		_, err := s.DB.Exec(query, genre.Name)
		if err != nil {
			return fmt.Errorf("error seeding genres: %v", err)
		}
	}
	return nil
}

func (s *SubscriptionSeeder) Seed() error {
	subscriptionDate := time.Now()

	subscriptions := []struct {
		UserID           int
		LibraryID        int
		SubscriptionDate time.Time
	}{
		{1, 1, subscriptionDate},
		{2, 2, subscriptionDate},
		{3, 1, subscriptionDate},
	}

	for _, subscription := range subscriptions {
		query := `INSERT INTO subscription (userID, libraryID, subscriptionDate) VALUES (?, ?, ?)`
		_, err := s.DB.Exec(query, subscription.UserID, subscription.LibraryID, subscription.SubscriptionDate)
		if err != nil {
			return fmt.Errorf("error seeding subscriptions: %v", err)
		}
	}

	return nil
}

func (s *UserRoleSeeder) Seed() error {
	roles := []struct {
		RoleName string
	}{
		{"admin"},
		{"librarian"},
		{"user"},
	}

	for _, role := range roles {
		query := `INSERT INTO userRole (RoleName) VALUES (?)`
		_, err := s.DB.Exec(query, role.RoleName)
		if err != nil {
			return fmt.Errorf("error seeding user roles: %v", err)
		}
	}

	return nil
}

func (s *UserSeeder) Seed() error {

	var adminRoleID, librarianRoleID, userRoleID int

	err := s.DB.QueryRow(`SELECT userRoleID FROM userRole WHERE RoleName = ?`, "admin").Scan(&adminRoleID)
	if err != nil {
		return fmt.Errorf("could not find admin role: %v", err)
	}
	err = s.DB.QueryRow(`SELECT userRoleID FROM userRole WHERE RoleName = ?`, "librarian").Scan(&librarianRoleID)
	if err != nil {
		return fmt.Errorf("could not find librarian role: %v", err)
	}
	err = s.DB.QueryRow(`SELECT userRoleID FROM userRole WHERE RoleName = ?`, "user").Scan(&userRoleID)
	if err != nil {
		return fmt.Errorf("could not find user role: %v", err)
	}

	users := []struct {
		Username string
		Password string
		Email    string
		RoleID   int
	}{
		{"johndoe", "password123", "johndoe@example.com", userRoleID},
		{"janedoe", "password123", "janedoe@example.com", userRoleID},
		{"admin", "adminpass", "admin@example.com", adminRoleID},
		{"librarian", "librarianpass", "librarian@example.com", librarianRoleID},
	}

	for _, user := range users {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
		if err != nil {
			return fmt.Errorf("could not hash password: %v", err)
		}

		query := `INSERT INTO user (username, password, email, userRoleID) VALUES (?, ?, ?, ?)`
		_, err = s.DB.Exec(query, user.Username, hashedPassword, user.Email, user.RoleID)
		if err != nil {
			return fmt.Errorf("error seeding users: %v", err)
		}
	}

	return nil
}
