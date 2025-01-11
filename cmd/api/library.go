package main

import (
	"net/http"

	"github.com/Bruheem/Portail_de_Reservation/internal/data"
	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
)

func (app *application) showLibraryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	lib, err := app.library.Get(id)
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

	library := &models.Library{
		Name:      input.Name,
		CreatedBy: input.CreatedBy,
	}

	v := validator.New()
	if data.ValidateLibrary(v, library); !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	id, err := app.library.Insert(library.Name, library.CreatedBy)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.logger.Printf("new library added with success! (id = %d)", id)
}

func (app *application) updateLibraryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	lib, err := app.library.Get(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var input struct {
		Name      string `json:"name"`
		CreatedBy string `json:"createdby"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	lib.Name = input.Name
	lib.CreatedBy = input.CreatedBy

	v := validator.New()
	if data.ValidateLibrary(v, lib); !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	app.library.Update(lib)

	err = app.writeJSON(w, http.StatusOK, envelope{"library": lib}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteLibraryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.library.Delete(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "library deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) searchLibraries(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")

	libs, err := app.library.Search(query)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"libraries": libs}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) recommendLibraries(w http.ResponseWriter, r *http.Request) {
	libraries, err := app.library.GetPopular()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"recommended_libraries": libraries}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
