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
		return ""
	}

	// Get the "page" parameter from the query string
	return parsedURL.Query().Get("page")
}

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/characters" {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		log.Printf("Page not found.")
		return
	}

	if r.Method != http.MethodGet {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		log.Printf("Invalid method while accessing resource.")
		return
	}

	query := r.URL.Query()
	page := query.Get("page")

	characterResponse, err := api.GetCharacters(w, page)
	if err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Failed fetching character information: %s", err)
		return
	}

	if characterResponse == nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		log.Printf("No character data  availabe.")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Failed to load template: %v", err)
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
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Failed to render template: %v", err)
		return
	}
}
