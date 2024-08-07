package handlers

import (
	"fullstackdev42/sober/services"
	"net/http"
	"text/template"
)

type DefaultHandler struct {
	PageService services.PageService
}

func (h *DefaultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		h.Home(w, r)
	case "/community":
		h.Community(w, r)
	default:
		http.NotFound(w, r)
	}
}

func (h *DefaultHandler) Home(w http.ResponseWriter, r *http.Request) {
	data := services.PageData{
		Title: "Home",
	}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/index.html", "templates/header.html", "templates/footer.html"))
	tmpl.ExecuteTemplate(w, "base.html", data)
}

func (h *DefaultHandler) Community(w http.ResponseWriter, r *http.Request) {
	data := services.PageData{
		Title: "Community",
	}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/community.html", "templates/header.html", "templates/footer.html"))
	tmpl.ExecuteTemplate(w, "base.html", data)
}
