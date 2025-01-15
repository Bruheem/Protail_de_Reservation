package main

import "net/http"

func (app *application) subscribe(w http.ResponseWriter, r *http.Request) {
	var input struct {
		LibraryID int64 `json:"library_id"`
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

	exists, err := app.subscription.Exists(userID, input.LibraryID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if exists {
		app.errorResponse(w, r, http.StatusConflict, "You are already subscribed to this library")
		return
	}

	err = app.subscription.Insert(userID, input.LibraryID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"message": "Successfully subscribed"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) unsubscribe(w http.ResponseWriter, r *http.Request) {

	var input struct {
		LibraryID int64 `json:"library_id"`
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

	exists, err := app.subscription.Exists(userID, input.LibraryID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	if !exists {
		app.errorResponse(w, r, http.StatusNotFound, "No subscription found for this library")
		return
	}

	err = app.subscription.Delete(userID, input.LibraryID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "Successfully unsubscribed"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) getUserSubscriptions(w http.ResponseWriter, r *http.Request) {
	userID, err := app.extractUserIDFromToken(r)
	if err != nil {
		app.unauthorizedAccessResponse(w, r)
		return
	}

	libs, err := app.library.GetSubscribedLibraries(userID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"subscriptions": libs}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
