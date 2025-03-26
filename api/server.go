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
        http.Error(w, "Error loading template: "+err.Error(), http.StatusInternalServerError)
        log.Println("Error loading template:", err)  // Log the actual error
        return
    }

    err = tpl.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, "Error rendering template: "+err.Error(), http.StatusInternalServerError)
        log.Println("Error rendering template:", err)  // Log the actual error
    }
}

func home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Home",
		Mode:  "",
	}
	renderTemplate(w, "index.html", data)
}
func about(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "About",
		Mode:  "",
	}
	renderTemplate(w, "about.html", data)
}
func skills(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Skills",
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
	log.Println("Server listening on port 8080 ...")
}
