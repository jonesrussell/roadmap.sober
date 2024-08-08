package handlers

import (
	"fullstackdev42/sober/services"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type DefaultHandler struct {
	PageService services.PageService
}

func (h *DefaultHandler) Home(c echo.Context) error {
	data := services.PageData{
		Title: "Home",
	}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/index.html", "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}

func (h *DefaultHandler) Community(c echo.Context) error {
	data := services.PageData{
		Title: "Community",
	}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/community.html", "templates/header.html", "templates/footer.html"))
	if err := tmpl.ExecuteTemplate(c.Response().Writer, "base.html", data); err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return nil
}
