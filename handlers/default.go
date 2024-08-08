package handlers

import (
	"context"
	"fullstackdev42/sober/components"
	"fullstackdev42/sober/services"
	"io"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	PageService services.PageService
}

func (h *DefaultHandler) RenderPage(c echo.Context, pageName string) error {
	data, err := h.PageService.GetWebpage(pageName)
	if err != nil {
		return err
	}

	// Create a ContentPage with the title and content
	page := components.ContentPage(data.Title, h.Unsafe(data.Content))

	// Get the request context and response writer
	ctx := c.Request().Context()
	writer := c.Response().Writer

	// Render the page
	return page.Render(ctx, writer)
}

func (h *DefaultHandler) Home(c echo.Context) error {
	return h.RenderPage(c, "home")
}

func (h *DefaultHandler) Community(c echo.Context) error {
	return h.RenderPage(c, "community")
}

func (h *DefaultHandler) Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}
