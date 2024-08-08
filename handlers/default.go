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

func (h *DefaultHandler) Home(c echo.Context) error {
	data, err := h.PageService.GetPageData("home")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return components.ContentPage(data.Title, components.Unsafe(data.Content)).Render(c.Request().Context(), c.Response().Writer)
}

func (h *DefaultHandler) Community(c echo.Context) error {
	data, err := h.PageService.GetPageData("community")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return components.ContentPage(data.Title, components.Unsafe(data.Content)).Render(c.Request().Context(), c.Response().Writer)
}
