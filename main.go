package main

import (
	"fullstackdev42/sober/server"
	"fullstackdev42/sober/services"
	"log"
)

func main() {
	pageService := services.PageService{} // Create PageService here
	srv := server.NewServer(pageService)  // Pass it to NewServer

	// Start server
	log.Fatal(srv.Echo.Start(":8080"))
}
