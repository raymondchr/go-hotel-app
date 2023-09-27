package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raymondchr/go-hotel-app/internal/config"
	"github.com/raymondchr/go-hotel-app/internal/driver"
	"github.com/raymondchr/go-hotel-app/internal/handlers"
	"github.com/raymondchr/go-hotel-app/internal/helpers"
	"github.com/raymondchr/go-hotel-app/internal/models"
	"github.com/raymondchr/go-hotel-app/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println("The server is running on port 8080")

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.RoomRestriction{})

	//change to true when in production
	app.InProduction = false
	

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	//connect to database
	log.Println("connecting to database.....")

	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=qweasd123")
	if err != nil{
		log.Fatal("Cannot connect to database")
	}

	log.Println("connected to database!")

	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Failed in creating template cache")
		return nil,err
	}

	app.TemplateCache = templateCache
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
