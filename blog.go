package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title, Heading string
	Body           template.HTML
}

// parses all the template html files once so that they aren't re-loaded repeatedly.
// This feature is known as Template Caching.
var templates = template.Must(template.ParseGlob("./pub/*.html"))

func pageElement(src []byte, elem string) string {
	elemString := regexp.MustCompile(
		fmt.Sprintf("\\[\\[%s]]([a-zA-Z0-9'!-=+$£\"\\[\\]:;,.<>\\s]*)\\[\\[%s]]", elem, elem),
	)
	m := elemString.FindStringSubmatch(string(src))
	if m == nil {
		return ""
	}
	// Here we return m[1] because it is the first sub-match. Each sub-match is represented by parentheses AKA
	// "([a-zA-Z0-9'!-...])". m[0] is the exact string that was matched, however, this would be useless to us because
	// we aren't making use of the custom tags I have designed.
	return m[1]
}

// loadPage returns the Page data for each resource.
func loadHome(title string) (*Page, error) {
	filepath := "./pages/" + title + ".txt"
	contents, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	tabTitle := pageElement(contents, "Title")
	heading := pageElement(contents, "Heading")
	body := pageElement(contents, "Body")
	body = strings.Replace(body, "\n", "<br>", -1)
	return &Page{tabTitle, heading, template.HTML(body)}, nil
}

func homeHandler(w http.ResponseWriter, _ *http.Request) {
	p, err := loadHome("home")
	if err != nil {
		renderTemplate(w, "home",
			&Page{"Home", "Sorry!", "The home page is empty. Consider contacting the website author."},
		)
	}
	renderTemplate(w, "home", p)
}

func cvHandler(_ http.ResponseWriter, _ *http.Request) {
	// render cv to user
	log.Println("Successful request for CV!")
}

func blogHandler(_ http.ResponseWriter, _ *http.Request) {
	// render blog page
	log.Println("Successful request for blog!")
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
	http.HandleFunc("/cv", cvHandler)
	http.HandleFunc("/blog", blogHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
