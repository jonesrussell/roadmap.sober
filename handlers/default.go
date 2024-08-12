package handlers

import (
	"context"
	"io"
	"net/http"

	"github.com/jonesrussell/sober/services"
	"github.com/jonesrussell/sober/ui"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	PageService services.PageService
	BasePath    string
}

func (h *DefaultHandler) RenderPage(c echo.Context, page templ.Component, title string) error {
	// Wrap the page content with the ContentPage component
	page = ui.ContentPage(h.BasePath, title, page)

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
