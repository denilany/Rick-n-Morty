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
		log.Println("Link to fetch data from cannot be an empty string.")
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
	// Display characters from both "/" and "/characters" route
	if r.URL.Path != "/" && r.URL.Path != "/characters" {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusNotFound,
			Message:    "Not Found",
		})
		log.Println("Page not found.")
		return
	}

	if r.Method != http.MethodGet {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusMethodNotAllowed,
			Message:    "Method Not Allowed",
		})
		log.Println("Invalid method while accessing resource.")
		return
	}

	query := r.URL.Query()
	// Gets the page query parameter if it exists
	page := query.Get("page")

	characterResponse, err := api.GetCharacters(w, page)
	if err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Failed fetching character information: %s\n", err)
		return
	}

	if characterResponse == nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusNotFound,
			Message:    "Page Not Found",
		})
		log.Println("No character data availabe.")
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
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
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		log.Printf("Failed to render template: %v\n", err)
		return
	}
}
