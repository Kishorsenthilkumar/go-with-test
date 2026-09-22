package main

import "net/http"

func (app *application) routes() http.Handler {

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("web/ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return secureHeaders(mux)
}
