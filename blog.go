package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed pub
var content embed.FS

func main() {
	log.Println("Server listening at http://localhost:8080/")

	pubFS, err := fs.Sub(content, "pub")
	if err != nil {
		log.Fatal(err)
	}

	// We need to strip /pub/ because the pubFS already responds to URIs containing /pub/<path>.
	// In effect, we are trimming up the URL /pub/pub/<path> which comes from the request sent by the client.
	http.Handle("/pub/", http.StripPrefix("/pub/", http.FileServer(http.FS(pubFS))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "./pub/home/", http.StatusFound)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
