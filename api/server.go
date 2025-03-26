package api

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title string
	Mode  string // e.g., "dark" or "" (for light mode)
}

func renderTemplate(w http.ResponseWriter, templateName string, data PageData) {
	tpl, err := template.ParseFiles("templates/base.html", "templates/"+templateName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	tpl.ExecuteTemplate(w, "base", data)
}

func home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Magic!",
		Mode:  "",
	}
	renderTemplate(w, "index.html", data)
}
func about(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "About Me",
		Mode:  "",
	}
	renderTemplate(w, "about.html", data)
}
func skills(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "My Skills",
		Mode:  "",
	}
	renderTemplate(w, "skills.html", data)
}

func Server() {
	// Serve static files (css, js, images) from /static and /assets directories
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Serve main pages
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/skills", skills)

	// Run server
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
