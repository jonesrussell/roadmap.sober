package main

import (
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("public/index.html", "src/header.html", "src/footer.html"))
		tmpl.ExecuteTemplate(w, "index.html", nil)
	})

	// Serve static files
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/style.css", fs)
	http.Handle("/sober.svg", fs)

	http.ListenAndServe(":8080", nil)
}
