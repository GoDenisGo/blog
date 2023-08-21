package main

import (
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pub/home.html")
}

func cvHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pub/cv.html")
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./pub/blog.html")
}

func main() {
	log.Println("Server listening at http://localhost:8080/")

	webpages := http.FileServer(http.Dir("./pub/"))
	http.Handle("/pub/", http.StripPrefix("/pub/", webpages))

	// /home is intended as the starting point of the web app.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/home", http.StatusFound)
	})

	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/cv", cvHandler)
	http.HandleFunc("/blog", blogHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
