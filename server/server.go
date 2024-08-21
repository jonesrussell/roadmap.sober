package server

import (
	"net/http"

	"github.com/jonesrussell/sober/handlers"
	"github.com/jonesrussell/sober/services"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo *echo.Echo
}

func NewServer(pageService services.PageService, stepService services.StepService) *Server {
	e := echo.New()

	handler := &handlers.DefaultHandler{
		PageService: pageService,
		StepService: stepService,
	}

	// Routes
	e.GET("/", func(c echo.Context) error {
		c.SetParamNames("page")
		c.SetParamValues("home")
		return handler.Page(c)
	})

	e.GET("/:page", handler.Page)

	e.GET("/sober/:id", handler.LoadContent)

	e.GET("/roadmap/:id", func(c echo.Context) error {
		id := c.Param("id")
		step, err := stepService.GetStepById(id) // Use the stepService to get the step
		if err != nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "step not found"})
		}
		return c.JSON(http.StatusOK, step)
	})

	// Serve static files
	e.Static("/static", "public")

	return &Server{
		Echo: e,
	}
}
