package main

import (
	"fullstackdev42/sober/server"
	"fullstackdev42/sober/services"
	"log"
	"net/http"
)

func main() {
	pageService := services.PageService{}
	srv := server.NewServer(pageService)

	err := http.ListenAndServe(":8080", srv.Handler)
	if err != nil {
		log.Fatal(err)
	}
}
