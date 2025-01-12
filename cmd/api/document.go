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
		YearPublished  int    `json:"yearpublished"`
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

func (app *application) borrowDocument(w http.ResponseWriter, r *http.Request) {

	var input struct {
		DocumentID int64 `json:"document_id"`
		LibraryID  int64 `json:"library_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	userID, err := app.extractUserIDFromToken(r)
	if err != nil {
		app.unauthorizedAccessResponse(w, r)
		return
	}

	subscribed, err := app.subscription.Exists(userID, input.LibraryID)
	if err != nil || !subscribed {
		app.errorResponse(w, r, http.StatusForbidden, "You are not subscribed to this library")
		return
	}

	available, err := app.document.IsAvailable(uint64(input.DocumentID))
	if err != nil || !available {
		app.errorResponse(w, r, http.StatusBadRequest, "Document is not available for borrowing")
		return
	}

	borrowID, err := app.lending.BorrowDocument(userID, input.DocumentID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusCreated, envelope{"borrow_id": borrowID}, nil)
}

func (app *application) returnDocument(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	userID, err := app.extractUserIDFromToken(r)
	if err != nil {
		app.unauthorizedAccessResponse(w, r)
		return
	}

	borrow, err := app.lending.GetBorrowingID(int64(id))
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	if borrow.User_id != userID {
		app.errorResponse(w, r, http.StatusForbidden, "You cannot borrow this document")
		return
	}

	err = app.lending.MarkAsReturned(borrow.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"message": "Document returned successfully"}, nil)
}

func (app *application) showBorrowedDocument(w http.ResponseWriter, r *http.Request) {

	userID, err := app.extractUserIDFromToken(r)
	if err != nil {
		app.unauthorizedAccessResponse(w, r)
		return
	}

	borrowedDocs, err := app.lending.GetBorrowedDocuments(userID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"borrowed_documents": borrowedDocs}, nil)
}

func (app *application) searchDocuments(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")

	documents, err := app.document.SearchDocuments(query)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	app.writeJSON(w, http.StatusOK, envelope{"documents": documents}, nil)
}

func (app *application) recommendDocuments(w http.ResponseWriter, r *http.Request) {
	documents, err := app.document.GetPopular()
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"recommended_documents": documents}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
