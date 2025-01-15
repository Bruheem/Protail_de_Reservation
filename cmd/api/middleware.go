package main

import (
	"net/http"
)

func (app *application) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isAuthorized := app.authorizationMiddleware(r, "admin")
		if !isAuthorized {
			app.unauthorizedAccessResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) librarianMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		isAuthorized := app.authorizationMiddleware(r, "librarian")
		if !isAuthorized {
			app.unauthorizedAccessResponse(w, r)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (app *application) authorizationMiddleware(r *http.Request, role string) bool {

	app.logger.Println("User is trying to authenticate!")

	userID, err := app.extractUserIDFromToken(r)
	if err != nil {
		return false
	}

	user, err := app.user.GetByID(userID)
	if err != nil {
		return false
	}

	if user.Role != role {
		return false
	}

	return true
}
