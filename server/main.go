package main

import (
	"fullstackdev42/sober/handlers"
	"fullstackdev42/sober/services"
	"log"
	"net/http"
)

func main() {
	pageService := services.PageService{}
	handler := handlers.DefaultHandler{
		PageService: pageService,
	}

	http.Handle("/", &handler)

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/style.css", fs)
	http.Handle("/sober.svg", fs)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
