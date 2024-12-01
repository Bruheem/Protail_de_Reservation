package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.Handler(http.MethodGet, "/v1/library/:id", app.adminMiddleware(http.HandlerFunc(app.showLibraryHandler)))
	router.Handler(http.MethodPost, "/v1/library", app.adminMiddleware(http.HandlerFunc(app.createLibraryHandler)))
	router.Handler(http.MethodPut, "/v1/library/:id", app.adminMiddleware(http.HandlerFunc(app.updateLibraryHandler)))
	router.Handler(http.MethodDelete, "/v1/library/:id", app.adminMiddleware(http.HandlerFunc(app.deleteLibraryHandler)))

	router.Handler(http.MethodGet, "/v1/document/:id", app.libAdminMiddleware(http.HandlerFunc(app.showDocumentHandler)))
	router.Handler(http.MethodPost, "/v1/document", app.libAdminMiddleware(http.HandlerFunc(app.createDocumentHandler)))
	router.Handler(http.MethodPut, "/v1/document/:id", app.libAdminMiddleware(http.HandlerFunc(app.updateDocumentHandler)))
	router.Handler(http.MethodDelete, "/v1/document/:id", app.libAdminMiddleware(http.HandlerFunc(app.deleteDocumentHandler)))

	router.HandlerFunc(http.MethodGet, "/v1/document/suggestions/:id", app.getSuggestions)

	router.HandlerFunc(http.MethodPost, "/v1/document/borrow", app.borrowDocument)
	router.HandlerFunc(http.MethodPost, "/v1/document/return/:id", app.returnDocument)
	router.HandlerFunc(http.MethodGet, "/v1/document/status/:id", app.getBorrowedDocumentStatus)

	router.HandlerFunc(http.MethodPost, "/v1/user/signup", app.userSignup)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.authenticationTokenHandler)

	// middlware to accup CORS from the frontend server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(router)

	return corsHandler
}
