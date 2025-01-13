package models

import (
	"context"
	"database/sql"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail = errors.New("duplicate email")
)

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password password `json:"-"`
	Role     string   `json:"role"`
}

type password struct {
	plaintext *string
	hash      []byte
}

func (p *password) Set(plaintextPassword string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return err
	}

	p.plaintext = &plaintextPassword
	p.hash = hash

	return nil
}

func (p *password) Matches(plaintextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword(p.hash, []byte(plaintextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func ValidateEmail(v *validator.Validator, email string) {
	v.Check(email != "", "email", "must be provided")
	v.Check(validator.Matches(email, validator.EmailRX), "email", "must be a valid email address")
}

func ValidatePasswordPlaintext(v *validator.Validator, password string) {
	v.Check(password != "", "password", "must be provided")
	v.Check(len(password) >= 8, "password", "must be at least 8 characters long")
	v.Check(len(password) <= 72, "password", "must be at most 72 characters long")
}

func ValidateUser(v *validator.Validator, user *User) {
	v.Check(user.Username != "", "username", "must be provided")
	v.Check(len(user.Username) <= 500, "username", "must be at most 500 characters long")
	v.Check(user.Role != "", "role", "must be a valid role")
	ValidateEmail(v, user.Email)

	if user.Password.plaintext != nil {
		ValidatePasswordPlaintext(v, *user.Password.plaintext)
	}

	if user.Password.hash == nil {
		panic("missing password hash for user")
	}
}

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(user *User) error {

	roleMap := map[string]int{
		"admin":     1,
		"librarian": 2,
		"user":      3,
	}

	roleID, exists := roleMap[user.Role]
	if !exists {
		return errors.New("invalid role")
	}

	insertQuery := `INSERT INTO user (username, password, email, userRoleID) 
                    VALUES (?, ?, ?, ?);`

	args := []interface{}{
		user.Username,
		user.Password.hash,
		user.Email,
		roleID,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := m.DB.ExecContext(ctx, insertQuery, args...)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "email"):
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return err
	}

	user.ID = strconv.FormatInt(lastInsertID, 10)
	return nil
}

func (m *UserModel) Update(user *User) error {
	query := `UPDATE user SET username = ?, password = ?, email = ?, userRoleID = ? WHERE id = ?`
	err := m.DB.QueryRowContext(context.Background(), query, user.Username, user.Password, user.Email, user).Scan(&user.ID)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), `email`):
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}

func (m UserModel) GetByEmail(email string) (*User, error) {
	// query := `SELECT id, username, password, email, userRoleID FROM user WHERE email = ?`

	query := `
    SELECT u.id, u.username, u.password, u.email, r.RoleName as role
    FROM user u
    JOIN userRole r ON u.userRoleID = r.userRoleID
    WHERE u.email = ?
	`

	var user User

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password.hash,
		&user.Email,
		&user.Role,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("no record found")
		default:
			return nil, err
		}
	}

	return &user, nil
}

func (m *UserModel) GetByID(id int64) (*User, error) {

	var user User

	query := `
		SELECT * FROM user
		WHERE id = ?
	`

	err := m.DB.QueryRow(query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password.hash,
		&user.Email,
		&user.Role,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.New("no record found")
		default:
			return nil, err
		}
	}

	return &user, nil
}
