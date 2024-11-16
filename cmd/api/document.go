package main

import (
	"net/http"

	"github.com/Bruheem/Portail_de_Reservation/internal/data"
	"github.com/Bruheem/Portail_de_Reservation/internal/models"
	"github.com/Bruheem/Portail_de_Reservation/internal/validator"
)

func (app *application) createDocumentHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title          string `json:"title"`
		Author         string `json:"author"`
		YearPublished  int    `json:"yearPublished`
		ISBN           string `json:"isbn"`
		LibraryID      int    `json:"libraryid"`
		DocumentTypeID int    `json:"documenttypeid"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	document := &models.Document{
		Title:          input.Title,
		Author:         input.Author,
		YearPublished:  input.YearPublished,
		ISBN:           input.ISBN,
		LibraryID:      input.LibraryID,
		DocumentTypeID: input.DocumentTypeID,
	}

	v := validator.New()
	if data.ValidateDocument(v, document); !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	id, err := app.document.InsertDocument(document)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.logger.Printf("new document added with success! (id = %d)", id)
}

func (app *application) showDocumentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	doc, err := app.document.GetDocument(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"document": doc}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateDocumentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	doc, err := app.document.GetDocument(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	var input struct {
		Title          string `json:"title"`
		Author         string `json:"author"`
		YearPublished  int    `json:"yearPublished"`
		ISBN           string `json:"isbn"`
		LibraryID      int    `json:"libraryID"`
		DocumentTypeID int    `json:"documentTypeID"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	doc.Title = input.Title
	doc.Author = input.Author
	doc.YearPublished = input.YearPublished
	doc.ISBN = input.ISBN
	doc.LibraryID = input.LibraryID
	doc.DocumentTypeID = input.DocumentTypeID

	v := validator.New()
	if data.ValidateDocument(v, doc); !v.IsValid() {
		app.failedValidatorResponse(w, r, v.Errors)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"document": doc}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteDocumentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.document.DeleteDocument(id)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Document deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
