package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {

	// router configuration
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/library/:id", app.showLibraryHandler)
	router.HandlerFunc(http.MethodPost, "/v1/library", app.createLibraryHandler)
	router.HandlerFunc(http.MethodPut, "/v1/library/:id", app.updateLibraryHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/library/:id", app.deleteLibraryHandler)

	router.HandlerFunc(http.MethodGet, "/v1/document/:id", app.showDocumentHandler)
	router.HandlerFunc(http.MethodPost, "/v1/document", app.createDocumentHandler)
	router.HandlerFunc(http.MethodPut, "/v1/document/:id", app.updateDocumentHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/document/:id", app.deleteDocumentHandler)

	router.HandlerFunc(http.MethodPost, "/v1/user/signup", app.userSignup)
	router.HandlerFunc(http.MethodPost, "/v1/user/logout", app.userLogout)
	router.HandlerFunc(http.MethodPost, "/v1/user/login", app.userLogin)

	return router
}
