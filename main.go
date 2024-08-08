package main

import (
	"fullstackdev42/sober/server"
	"fullstackdev42/sober/services"
	"log"
)

func main() {
	fileReader := &services.OSFileReader{}
	pageService := &services.PageServiceImpl{
		FileReader: fileReader,
	}
	srv := server.NewServer(pageService)

	// Start server
	log.Fatal(srv.Echo.Start(":8080"))
}
