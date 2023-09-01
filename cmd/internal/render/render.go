package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"
	"github.com/raymondchr/go-hotel-app/cmd/internal/config"
	"github.com/raymondchr/go-hotel-app/cmd/internal/models"
)

var app *config.AppConfig
var pathToTemplate = "./templates"
var functions = template.FuncMap{}

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.Flash = app.Session.PopString(r.Context(), "flash")
	td.Error = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) {
	//get template cache from app config
	var templateCache map[string]*template.Template

	if app.UseCache {
		templateCache = app.TemplateCache
		log.Println("Using tc")
	} else {
		templateCache, _ = CreateTemplateCache()
		log.Println("reload from disk")
	}

	//get reqeusted template from cache
	template, ok := templateCache[tmpl]
	if !ok {
		log.Fatal("Error in calling the cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	err := template.Execute(buf, td)
	if err != nil {
		log.Println(err)
	}

	//render the template

	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplate))
	if err != nil {
		return myCache, err
	}

	//range through all files with *.page.tmpl

	for _, page := range pages {
		name := filepath.Base(page)
		templateSet, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplate))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			templateSet, err = templateSet.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplate))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = templateSet
	}

	return myCache, nil
}
