package main

import (
	"go-with-test/web/snippetbox/internal/models"
	"html/template"
	"net/http"
	"path/filepath"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     *models.Snippet
	Snippets    []*models.Snippet
	Form        any
	Flash       string
}

func humanDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var temp_functions = template.FuncMap{"humanDate": humanDate}

func newTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./web/ui/html/pages/*.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		file_name := filepath.Base(page)

		ts, err := template.New(file_name).Funcs(temp_functions).ParseFiles("web/ui/html/base.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseGlob("web/ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		cache[file_name] = ts
	}
	return cache, nil
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{CurrentYear: time.Now().Year(), Flash: app.sessionManager.PopString(r.Context(), "flash")}
}
