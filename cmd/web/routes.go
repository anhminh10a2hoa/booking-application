package main

import (
	"net/http"
	"github.com/bmizerany/pat"

	"github.com/anhminh10a2hoa/go-course/pkg/handlers"
	"github.com/anhminh10a2hoa/go-course/pkg/config"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}