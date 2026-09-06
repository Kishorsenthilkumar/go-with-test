package main

import (
	"database/sql"
	"flag"
	"fmt"
	"go-with-test/web/snippetbox/internal/models"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorlog *log.Logger
	infolog  *log.Logger
	snippets *models.SnippetModel
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	files := []string{
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\base.tmpl",
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\partials\\nav.tmpl",
		"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\home.tmpl",
	}
	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)

	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	fmt.Fprintf(w, "Display a specific snippet with ID %d", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Set("Cache-Control", "public,max-age=31536000")
		w.Header()["Date"] = nil
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("create a new snippet"))
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")

	flag.Parse()

	infolog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorlog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)

	if err != nil {
		errorlog.Fatal(err)
	}

	defer db.Close()

	app := &application{errorlog: errorlog, infolog: infolog, snippets: &models.SnippetModel{DB: db}}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorlog,
		Handler:  app.routes(),
	}

	infolog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()

	errorlog.Fatal(err)

}
