package main

import (
	"fullstackdev42/sober/server"
	"fullstackdev42/sober/services"
	"log"
)

func main() {
	pageService := &services.PageServiceImpl{}
	srv := server.NewServer(pageService)

	// Start server
	log.Fatal(srv.Echo.Start(":8080"))
}
