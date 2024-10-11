package main

import (
	"net/http"
	"time"

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
