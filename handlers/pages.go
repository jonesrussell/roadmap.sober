package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) Home(c echo.Context) error {
	return h.RenderPage(c, "home")
}

func (h *DefaultHandler) Community(c echo.Context) error {
	return h.RenderPage(c, "community")
}
