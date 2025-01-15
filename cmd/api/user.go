package main

import (
	"errors"
	"net/http"

	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
)

func (app *application) updateUser(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	user := &models.User{
		ID:       input.Id,
		Username: input.Username,
		Email:    input.Email,
		Role:     "user",
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

	updatedUser, err := app.user.Update(user)
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

	err = app.writeJSON(w, http.StatusCreated, envelope{"user": updatedUser}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.logger.Println("User updated successfully")
}
