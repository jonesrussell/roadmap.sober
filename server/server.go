package server

import (
	"github.com/jonesrussell/sober/handlers"
	"github.com/jonesrussell/sober/services"

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
	e.GET("/", func(c echo.Context) error {
		c.SetParamNames("page")
		c.SetParamValues("home")
		return handler.Page(c)
	})

	e.GET("/:page", handler.Page)

	e.GET("/sober/:id", handler.LoadContent)

	// Serve static files
	e.Static("/static", "public")

	return &Server{
		Echo: e,
	}
}
