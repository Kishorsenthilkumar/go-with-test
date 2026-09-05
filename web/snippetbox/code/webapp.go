package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	files := []string{
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\base.tmpl",
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\partials\\nav.tmpl",
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\home.bak",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errorlog.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.errorlog.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

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

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{errorlog: errorlog, infolog: infolog}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  mux,
	}

	infolog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()

	errorlog.Fatal(err)

}
