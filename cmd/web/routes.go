package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/command", app.getCommand)
	mux.HandleFunc("/command/create", app.createCommand)
	mux.HandleFunc("/command/exec", app.execCommand)
	mux.HandleFunc("/command/delete", app.deleteCommand)

	return mux
}
