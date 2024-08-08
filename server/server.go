package server

import (
	"fullstackdev42/sober/handlers"
	"fullstackdev42/sober/services"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer(pageService services.PageService) *Server {
	e := echo.New()

	handler := &handlers.DefaultHandler{
		PageService: pageService,
	}

	// Routes
	e.GET("/", handler.Home)
	e.GET("/community", handler.Community)

	// Serve static files
	e.Static("/static", "public")

	return &Server{
		Echo: e,
	}
}
