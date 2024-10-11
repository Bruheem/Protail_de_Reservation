package main

import (
	"fmt"
	"net/http"
	"time"

	validator "github.com/Bruheem/Portail_de_Reservation/internal"
	"github.com/Bruheem/Portail_de_Reservation/internal/data"
)

func (app *application) showLibraryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	document := data.Document{
		ID:             id,
		Title:          "something",
		Author:         "ibrahim Ben.",
		YearPublished:  time.Now(),
		ISBN:           909090,
		LibraryID:      3,
		DocumentTypeID: 1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"document": document}, nil)
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

	fmt.Fprintf(w, "name: %s, createdby: %s", input.Name, input.CreatedBy)
}
