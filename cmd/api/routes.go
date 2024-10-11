package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/library/:id", app.showLibraryHandler)
	router.HandlerFunc(http.MethodPost, "/v1/library", app.createLibraryHandler)

	// router.HandlerFunc(http.MethodGet, "/v1/home", app.homePageHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/library", app.libraryPageHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/collections", app.collectionsPageHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/profile", app.profilePageHandler)

	return router
}
