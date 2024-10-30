package main

import (
	"log"
	"net/http"

	"github.com/denilany/Rick-n-Morty/customerrors"
	"github.com/denilany/Rick-n-Morty/route"
)

func main() {
	http.HandleFunc("/characters", route.CharacterHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// A default handler to pass a 404 error for any unmatched or undefined route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			customerrors.NotFoundHandler(w, r)
			log.Println("Page not found.")
			return
		}
		// Optionally redirect "/" to /characters
		http.Redirect(w, r, "/characters", http.StatusSeeOther)
	})

	log.Println("server starting at http://localhost:8080/characters")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
