package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ModernWebAppWithGo/sec3/booking/pkg/config"
	"github.com/ModernWebAppWithGo/sec3/booking/pkg/handlers"
	"github.com/ModernWebAppWithGo/sec3/booking/pkg/render"
)

const portNumber = ":8080"

// main is the main application
func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Startung application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}