package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
	"github.com/dgrijalva/jwt-go"
)

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Role:     input.Role,
	}

	err = user.Password.Set(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	v := validator.New()

	if models.ValidateUser(v, user); !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	err = app.user.Insert(user)
	if err != nil {
		switch {
		case errors.Is(err, models.ErrDuplicateEmail):
			v.Add("email", "email already in use")
			app.failedValidatorResponse(w, r, v.Errors)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

type Token struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {

	app.logger.Println("a user is attempting to login")

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	models.ValidateEmail(v, input.Email)
	models.ValidatePasswordPlaintext(v, input.Password)

	if !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	user, err := app.user.GetByEmail(input.Email)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "found"):
			app.invalidCredentialsResponse(w, r)
			app.logger.Println("user failed to authenticate: Bad credentials")
		default:
			app.serverErrorResponse(w, r, err)
			app.logger.Println("the server couldn't resolve user's authentication")
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		app.logger.Println("unable to convert password to hash")
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		app.logger.Println("mismatched password")
		return
	}

	id, _ := strconv.ParseInt(user.ID, 10, 64)
	token, err := app.generateToken(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"token": token, "user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

	app.logger.Printf("a user has successfully authenticated (id = %s)", user.ID)
}

func (app *application) generateToken(userID int64) (Token, error) {

	var token Token

	token.Expiration = time.Now().Add(2 * time.Hour)

	claims := &jwt.StandardClaims{
		Issuer:    fmt.Sprintf("%s:%d", app.config.jwt.issuer, app.config.jwt.port),
		Subject:   strconv.FormatInt(userID, 10),
		ExpiresAt: token.Expiration.Unix(),
	}

	secretKey := []byte(app.config.jwt.secret)

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(secretKey)
	if err != nil {
		app.logger.Printf("Error signing token: %v", err)
		return Token{}, err
	}

	token.Token = tokenString
	return token, nil
}
