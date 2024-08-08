package handlers

import (
	"fullstackdev42/sober/components"
	"fullstackdev42/sober/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	PageService services.PageService
}

func (h *DefaultHandler) RenderPage(c echo.Context, pageName string) error {
	data, err := h.PageService.GetWebpage(pageName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return components.ContentPage(data.Title, components.Unsafe(data.Content)).Render(c.Request().Context(), c.Response().Writer)
}

func (h *DefaultHandler) Home(c echo.Context) error {
	return h.RenderPage(c, "home")
}

func (h *DefaultHandler) Community(c echo.Context) error {
	return h.RenderPage(c, "community")
}
