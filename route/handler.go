package route

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/denilany/Rick-n-Morty/api"
	apiconfig "github.com/denilany/Rick-n-Morty/constant"
)

func CharacterHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/characters" {
		err := fmt.Sprintf("wrong URL endpoint: %s", r.URL.Path)
		http.Error(w, err, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		err := fmt.Sprintf("invalid HTTP method: %s. Only GET is allowed", r.Method)
		http.Error(w, err, http.StatusMethodNotAllowed)
		return
	}

	url := apiconfig.BaseURL + apiconfig.Character
	characters, err := api.GetCharacterApi(url)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to retrieve character data: %v", err), http.StatusInternalServerError)
		return
	}

	if characters == nil {
		http.Error(w, "no character data available", http.StatusNotFound)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to load template: %v", err), http.StatusInternalServerError)
		return
	}

	var results []struct {
		Name    string
		Image   string
		Species string
		Origin  string
	}

	for _, character := range characters.Result {
		result := struct {
			Name    string
			Image   string
			Species string
			Origin  string
		}{
			Name:    character.Name,
			Image:   character.Image,
			Species: character.Species,
			Origin:  character.Origin.Name,
		}
		results = append(results, result)
	}

	err = tmpl.Execute(w, results)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to render template: %v", err), http.StatusInternalServerError)
	}
}
