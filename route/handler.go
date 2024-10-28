package route

import (
	"fmt"
	"net/http"
	"net/url"
	"text/template"

	"github.com/denilany/Rick-n-Morty/api"
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
		return ""
	}

	// Get the "page" parameter from the query string
	return parsedURL.Query().Get("page")
}

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/characters" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query()
	page := query.Get("page")

	characterResponse, err := api.GetCharacters(page)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to fetch character data: %v", err), http.StatusInternalServerError)
		return
	}

	if characterResponse == nil {
		http.Error(w, "no character data available", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to load template: %v", err), http.StatusInternalServerError)
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
		http.Error(w, fmt.Sprintf("failed to render template: %v", err), http.StatusInternalServerError)
	}
}
