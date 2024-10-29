package route

import (
	"log"
	"net/http"
	"net/url"
	"text/template"

	"github.com/denilany/Rick-n-Morty/api"
	"github.com/denilany/Rick-n-Morty/customerrors"
)

type TemplateData struct {
	Characters  []CharacterInfo
	NextPage    string
	PrevPage    string
	CurrentPage string
}

type CharacterInfo struct {
	Name    string
	Image   string
	Species string
	Origin  string
}

func extractPageNumber(fullurl string) string {
	if fullurl == "" {
		return ""
	}

	parsedURL, err := url.Parse(fullurl)
	if err != nil {
		log.Printf("Failed parsing URL %s: %s\n", fullurl, err)
		return ""
	}

	// Get the "page" parameter from the query string
	return parsedURL.Query().Get("page")
}

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/characters" {
		customerrors.NotFoundHandler(w, r)
		log.Println("Page not found.")
		return
	}

	if r.Method != http.MethodGet {
		customerrors.MethodNotAllowedHandler(w, r)
		log.Println("Invalid method while accessing resource.")
		return
	}

	query := r.URL.Query()
	// Gets the page query parameter if it exists
	page := query.Get("page")

	characterResponse, err := api.GetCharacters(page)
	if err != nil {
		customerrors.NotFoundHandler(w, r)
		log.Printf("Failed fetching character information: %s\n", err)
		return
	}

	if characterResponse == nil {
		customerrors.NotFoundHandler(w, r)
		log.Println("No character data availabe.")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		customerrors.InternalServerErrorHandler(w, r)
		log.Printf("Failed to load template: %v\n", err)
		return
	}

	var characters []CharacterInfo

	for _, character := range characterResponse.Result {
		charInfo := CharacterInfo{
			Name:    character.Name,
			Image:   character.Image,
			Species: character.Species,
			Origin:  character.Origin.Name,
		}
		characters = append(characters, charInfo)
	}

	nextPage := extractPageNumber(characterResponse.Info.Next)
	prevPage := extractPageNumber(characterResponse.Info.Prev)

	templateData := TemplateData{
		Characters:  characters,
		NextPage:    nextPage,
		PrevPage:    prevPage,
		CurrentPage: page,
	}

	err = tmpl.Execute(w, templateData)
	if err != nil {
		customerrors.InternalServerErrorHandler(w, r)
		log.Printf("Failed to render template: %v\n", err)
		return
	}
}
