package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Body template.HTML
}

// parses all the html templates once so that they aren't re-rendered repeatedly.
// This feature is known as Template Caching.
var templates = template.Must(template.ParseGlob("./pub/*.html"))

// loadPage returns a Page object containing the data for each resource.
func loadPage(title string) (*Page, error) {
	filename := "./pages/" + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{template.HTML(body)}, nil
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	p, err := loadPage("home")
	if err != nil {
		renderTemplate(w, "home", &Page{"The home page was not found."})
	}
	renderTemplate(w, "home", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	log.Println("Server listening at http://localhost:8080/")

	styles := http.FileServer(http.Dir("./pub/"))
	http.Handle("/pub/", http.StripPrefix("/pub/", styles))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusFound)
	})

	http.HandleFunc("/home", homeHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
