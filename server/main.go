package main

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Title string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Home",
		}
		tmpl := template.Must(template.ParseFiles("src/base.html", "public/index.html", "src/header.html", "src/footer.html"))
		tmpl.ExecuteTemplate(w, "base.html", data)
	})

	http.HandleFunc("/community", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title: "Community",
		}
		tmpl := template.Must(template.ParseFiles("src/base.html", "public/community.html", "src/header.html", "src/footer.html"))
		tmpl.ExecuteTemplate(w, "base.html", data)
	})

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/style.css", fs)
	http.Handle("/sober.svg", fs)

	http.ListenAndServe(":8080", nil)
}
