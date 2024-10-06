package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/home", app.homePageHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/library", app.libraryPageHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/collections", app.collectionsPageHandler)
	// router.HandlerFunc(http.MethodGet, "/v1/profile", app.profilePageHandler)

	return router
}
