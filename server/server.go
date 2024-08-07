package server

import (
	"fullstackdev42/sober/handlers"
	"fullstackdev42/sober/services"
	"net/http"
)

type Server struct {
	Handler http.Handler
}

func NewServer(pageService services.PageService) *Server { // Add pageService as a parameter
	handler := &handlers.DefaultHandler{
		PageService: pageService,
	}

	http.Handle("/", handler)

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	return &Server{
		Handler: handler,
	}
}
