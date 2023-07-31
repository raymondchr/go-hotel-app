package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raymondchr/go-hotel-app/cmd/pkg/config"
	"github.com/raymondchr/go-hotel-app/cmd/pkg/handlers"
	"github.com/raymondchr/go-hotel-app/cmd/pkg/render"
)

func main() {
	var app config.AppConfig

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Failed in creating template cache")
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("The server is running on port 8080")
	_ = http.ListenAndServe(":8080", nil)
}
