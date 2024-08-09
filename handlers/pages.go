package handlers

import (
	"fullstackdev42/sober/content"

	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) Home(c echo.Context) error {
	return h.RenderPage(c, content.Home(), "Home")
}

func (h *DefaultHandler) Community(c echo.Context) error {
	return h.RenderPage(c, content.Community(), "Community")
}
