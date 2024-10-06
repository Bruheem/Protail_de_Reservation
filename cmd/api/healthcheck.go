package main

import (
	"fmt"
	"net/http"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status: OK\n")
	fmt.Fprintf(w, "env: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}
