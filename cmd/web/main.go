package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raymondchr/go-hotel-app/cmd/internal/config"
	"github.com/raymondchr/go-hotel-app/cmd/internal/handlers"
	"github.com/raymondchr/go-hotel-app/cmd/internal/models"
	"github.com/raymondchr/go-hotel-app/cmd/internal/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The server is running on port 8080")

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(models.ReservationData{})

	//change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Failed in creating template cache")
		return err
	}

	app.TemplateCache = templateCache
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	
	render.NewTemplates(&app)

	return nil
}