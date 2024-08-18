package server

import (
	"errors"
	"net/http"

	"github.com/jonesrussell/sober/handlers"
	"github.com/jonesrussell/sober/services"
	"github.com/jonesrussell/sober/ui/components"

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

	e.GET("/roadmap/:id", func(c echo.Context) error {
		id := c.Param("id")
		// Assuming you have a function to convert the id to the corresponding step
		step, err := getStepById(id)
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

func getStepById(id string) (*components.Step, error) {
	// Implement logic to retrieve the step based on the id
	// This could involve querying a database or another form of storage
	// For demonstration, returning hardcoded values
	switch id {
	case "0":
		return &components.Step{Text: "Admit the Problem", Content: "Initial step...", PanelTitle: "Step 1"}, nil
	case "1":
		return &components.Step{Text: "Seek Help", Content: "Find support...", PanelTitle: "Step 2"}, nil
	// Add more cases as needed
	default:
		return nil, errors.New("step not found")
	}
}
