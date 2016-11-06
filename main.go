package main

import (
	"net/http"
	"os"

	"github.com/russross/blackfriday"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/markdown", generateMarkdown)

	// Catch all route
	http.Handle("/", http.FileServer(http.Dir("public")))

	// Kicks off HTTP server into action
	http.ListenAndServe(":"+port, nil)
}

func generateMarkdown(rw http.ResponseWriter, r *http.Request) {
	markdown := blackfriday.MarkdownCommon([]byte(r.FormValue("body")))
	rw.Write(markdown)
}
