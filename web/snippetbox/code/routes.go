package main

import "net/http"

func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return mux
}
