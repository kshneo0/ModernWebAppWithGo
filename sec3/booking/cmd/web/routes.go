package main

import (
	"net/http"

	"github.com/ModernWebAppWithGo/sec3/booking/pkg/config"
	"github.com/ModernWebAppWithGo/sec3/booking/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}