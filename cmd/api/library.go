package main

import (
	"net/http"

	validator "github.com/Bruheem/Portail_de_Reservation/internal"
	"github.com/Bruheem/Portail_de_Reservation/internal/data"
)

func (app *application) showLibraryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	lib, err := app.library.GetLibrary(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"library": lib}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createLibraryHandler(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name      string `json:"name"`
		CreatedBy string `json:"createdby"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	library := &data.Library{
		Name:      input.Name,
		CreatedBy: input.CreatedBy,
	}

	v := validator.New()
	if data.ValidateLibrary(v, library); !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	id, err := app.library.InsertLibrary(library.Name, library.CreatedBy)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.logger.Printf("new library added with success! (id = %d)", id)
}
