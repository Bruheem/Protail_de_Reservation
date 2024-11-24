package main

import "net/http"

func (app *application) adminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userRole := r.Header.Get("userRole")
		if userRole != "admin" {
			app.unauthorizedAccessResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) libAdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userRole := r.Header.Get("userRole")
		if userRole != "libadmin" {
			app.unauthorizedAccessResponse(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
