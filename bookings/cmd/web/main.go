package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"time"

	"github.com/ModernWebAppWithGo/bookings/internal/config"
	"github.com/ModernWebAppWithGo/bookings/internal/driver"
	"github.com/ModernWebAppWithGo/bookings/internal/handlers"
	"github.com/ModernWebAppWithGo/bookings/internal/helpers"
	"github.com/ModernWebAppWithGo/bookings/internal/models"
	"github.com/ModernWebAppWithGo/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main application
func main() {

	db, err := run()
	if err != nil {
		log.Fatal(err)
	}

	defer db.SQL.Close()

	//brew services start mailhog
	from := "kshneo@gmail.com"
	auth := smtp.PlainAuth("", from, "", "localhost")
	err = smtp.SendMail("localhost:1025", auth, from, []string{"you@there.com"}, []byte("Hello, world"))
	if err != nil {
		log.Println(err)
	}

	fmt.Println(fmt.Sprintf("Startung application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB, error) {

	// what am I going to put in the session
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	//change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	//session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=sanghoonkim password=")
	if err != nil {
		log.Fatal("Cannot copnnect to database! Dying ...!")
	}

	log.Println("Connected to database!")

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil, err
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}
