package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/jonesrussell/sober/content"
	"github.com/jonesrussell/sober/handlers"
	"github.com/jonesrussell/sober/server"
	"github.com/jonesrussell/sober/services"

	"github.com/labstack/echo/v4"
)

// Config holds the application configuration
type Config struct {
	BasePath string
}

// loadConfig reads configuration from environment variables or config files
func loadConfig() (*Config, error) {
	return &Config{
		BasePath: os.Getenv("BASE_PATH"),
	}, nil
}

func main() {
	// Define and parse the "generate" flag
	generate := flag.Bool("generate", false, "Generate static site")
	flag.Parse()

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	pageService := services.NewPageService()

	if *generate {
		staticSiteService := services.NewStaticSiteService(pageService)
		staticSiteService.Generate(config.BasePath)
		return
	}

	// Setup server
	srv := setupServer(pageService, config)

	// Start server
	log.Printf("Starting server on port :8080 with generate flag: %v", *generate)
	log.Fatal(srv.Echo.Start(":8080"))
}

// setupServer initializes and configures the server
func setupServer(pageService services.PageService, config *Config) *server.Server {
	srv := server.NewServer(pageService)

	// Custom 404 handler
	srv.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if code == http.StatusNotFound {
			handler := &handlers.DefaultHandler{
				PageService: pageService,
				BasePath:    config.BasePath,
			}
			if err := handler.RenderPage(c, content.NotFound(), "404"); err != nil {
				c.Logger().Error(err)
			}
		} else {
			srv.Echo.DefaultHTTPErrorHandler(err, c)
		}
	}

	// Add middleware for security or other purposes
	// srv.Echo.Use(middleware.CORS())

	return srv
}
