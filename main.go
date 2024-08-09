package main

import (
	"fullstackdev42/sober/content"
	"fullstackdev42/sober/handlers"
	"fullstackdev42/sober/server"
	"fullstackdev42/sober/services"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	pageService := &services.PageServiceImpl{}
	srv := server.NewServer(pageService)

	// Set the NotFoundHandler to render your custom 404 page
	srv.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}
		if code == http.StatusNotFound {
			// Use your DefaultHandler to render the 404 page
			handler := &handlers.DefaultHandler{
				PageService: pageService,
			}
			if err := handler.RenderPage(c, content.NotFound(), "404"); err != nil {
				c.Logger().Error(err)
			}
		} else {
			srv.Echo.DefaultHTTPErrorHandler(err, c)
		}
	}

	// Start server
	log.Fatal(srv.Echo.Start(":8080"))
}
