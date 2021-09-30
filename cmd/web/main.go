package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/anhminh10a2hoa/go-course/pkg/config"
	"github.com/anhminh10a2hoa/go-course/pkg/handlers"
	"github.com/anhminh10a2hoa/go-course/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("can't create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHanlders(repo)

	render.NewTemplate(&app)
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
