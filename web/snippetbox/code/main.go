package main

import (
	"database/sql"
	"errors"
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

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, snips := range snippets {
		fmt.Fprintf(w, "%+v", snips)
	}

	//files := []string{
	//	"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\base.tmpl",
	//	"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\partials\\nav.tmpl",
	//	"C:\\Users\\Aspl-Kishore\\GolandProjects\\go-with-test\\web\\ui\\html\\pages\\home.tmpl",
	//}
	//ts, err := template.ParseFiles(files...)

	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}

	//err = ts.ExecuteTemplate(w, "base", nil)

	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/view.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.serverError(w, err)
	}
	err = ts.ExecuteTemplate(w, "base", snippet)
	if err != nil {
		app.serverError(w, err)
	}

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header().Set("Cache-Control", "public,max-age=31536000")
		w.Header()["Date"] = nil
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)

	if err != nil {
		app.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)

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
