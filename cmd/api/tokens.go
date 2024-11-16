package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	Token      string    `json:"token"`
	Expiration time.Time `json:"expiration"`
}

func (app *application) authenticationTokenHandler(w http.ResponseWriter, r *http.Request) {
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
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if !match {
		app.invalidCredentialsResponse(w, r)
		return
	}

	id, _ := strconv.ParseInt(user.ID, 10, 64)
	token, err := app.generateToken(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"authentication_token:": token}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
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
