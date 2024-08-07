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
	data := services.PageData{
		Title: "Home",
	}
	tmpl := template.Must(template.ParseFiles("templates/base.html", "templates/index.html", "templates/header.html", "templates/footer.html"))
	tmpl.ExecuteTemplate(w, "base.html", data)
}

// Add more methods for other routes...
