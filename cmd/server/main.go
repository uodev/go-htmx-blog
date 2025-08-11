package main

import (
	"html/template"
	"log"
	"net/http"

	routerpkg "github.com/uodev/gotemplatehtmxblog/internal/http"
)

func render(w http.ResponseWriter, name string, data any) {

	tpl := template.Must(template.ParseFiles(
		"views/layouts/base.html",
		"views/partials/nav.html",
		"views/pages/"+name+".html",
	))

	if err := tpl.ExecuteTemplate(w, "base", data); err != nil {
		http.Error(w, err.Error(), 500)
	}
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/app.css", http.FileServer(http.Dir("public")))

	app := routerpkg.New(render)

	mux.Handle("/", app)
	log.Println("http://localhost:8081")
	_ = http.ListenAndServe("localhost:8081", mux)
}
