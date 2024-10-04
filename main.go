package main

import (
	"log"
	"net/http"

	"github.com/denilany/Rick-n-Morty/route"
)

func main() {
	http.HandleFunc("/characters", route.CharacterHandler)

	log.Println("server starting at http://localhost:8080/characters")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
