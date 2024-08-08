package handlers

import (
	"context"
	"fullstackdev42/sober/components"
	"fullstackdev42/sober/services"
	"io"
	"net/http"

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

	// Render the page into a string
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)
	if err := page.Render(c.Request().Context(), buf); err != nil {
		return err
	}

	// Send the HTTP response
	return c.HTML(http.StatusOK, buf.String())
}

func (h *DefaultHandler) Unsafe(html string) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		_, err = io.WriteString(w, html)
		return
	})
}
