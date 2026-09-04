package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\base.tmpl",
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\partials\\nav.tmpl",
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\home.tmpl",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Set("Cache-Control", "public,max-age=31536000")
		w.Header()["Date"] = nil
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("create a new snippet"))
}

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)

	log.Fatal(err)

}
