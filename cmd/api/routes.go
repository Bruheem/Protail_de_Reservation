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

	// API Endpoints
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	// Library Endpoints
	router.HandlerFunc(http.MethodGet, "/v1/libraries", app.searchLibraries)
	router.HandlerFunc(http.MethodGet, "/v1/libraries/:id", app.showLibraryHandler)
	router.Handler(http.MethodPost, "/v1/libraries", app.adminMiddleware(http.HandlerFunc(app.createLibraryHandler)))
	router.Handler(http.MethodPut, "/v1/libraries/:id", app.adminMiddleware(http.HandlerFunc(app.updateLibraryHandler)))
	router.Handler(http.MethodDelete, "/v1/libraries/:id", app.adminMiddleware(http.HandlerFunc(app.deleteLibraryHandler)))

	// Document Endpoints
	router.HandlerFunc(http.MethodPost, "/v1/document/:id/request", app.borrowDocument)
	router.HandlerFunc(http.MethodGet, "/v1/documents", app.searchDocuments)
	router.HandlerFunc(http.MethodGet, "/v1/documents/:id", app.showDocumentHandler)
	router.Handler(http.MethodPost, "/v1/documents", app.libAdminMiddleware(http.HandlerFunc(app.createDocumentHandler)))
	router.Handler(http.MethodPut, "/v1/documents/:id", app.libAdminMiddleware(http.HandlerFunc(app.updateDocumentHandler)))
	router.Handler(http.MethodDelete, "/v1/documents/:id", app.libAdminMiddleware(http.HandlerFunc(app.deleteDocumentHandler)))

	// Authentication Endpoints
	router.HandlerFunc(http.MethodPost, "/v1/auth/register", app.register)
	router.HandlerFunc(http.MethodPost, "/v1/auth/login", app.login)
	// router.HandlerFunc(http.MethodGet, "/v1/auth/profile", app.showProfileHandler)

	// User Management
	// router.HandlerFunc(http.MethodGet, "/v1/users/:id", app.getUser)
	// router.HandlerFunc(http.MethodPut, "/v1/users/:id", app.updateUser)
	// router.HandlerFunc(http.MethodDelete, "/v1/users/:id", app.deleteUser)

	// Subscription Management
	router.HandlerFunc(http.MethodPost, "/v1/subscriptions", app.subscribe)
	router.HandlerFunc(http.MethodDelete, "/v1/subscriptions", app.unsubscribe)

	// Borrowing Management
	router.HandlerFunc(http.MethodGet, "/v1/borrow", app.showBorrowedDocument)
	router.HandlerFunc(http.MethodPost, "/v1/borrow", app.borrowDocument)
	router.HandlerFunc(http.MethodPut, "/v1/borrow/:id/return", app.returnDocument)

	// Recommendations Management
	router.HandlerFunc(http.MethodGet, "/v1/recommendations/libraries", app.recommendLibraries)
	router.HandlerFunc(http.MethodGet, "/v1/recommendations/libraries", app.recommendDocuments)

	// middlware to accept CORS from the frontend server
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowCredentials: true,
	})

	corsHandler := c.Handler(router)

	return corsHandler
}
