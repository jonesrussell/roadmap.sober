package handlers

import (
	"github.com/labstack/echo/v4"
)

func (h *DefaultHandler) Home(c echo.Context) error {
	page, err := h.PageService.GetWebpage("home")
	if err != nil {
		return err
	}
	return h.RenderPage(c, page.Content, page.Title)
}

func (h *DefaultHandler) Community(c echo.Context) error {
	page, err := h.PageService.GetWebpage("community")
	if err != nil {
		return err
	}
	return h.RenderPage(c, page.Content, page.Title)
}
