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

	result := struct{
		Name string
		Species string
		Origin string
	}{
		Name: characters.Result.

	}

	err = tmpl.Execute(w, characters)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to render template: %v", err), http.StatusInternalServerError)
	}
}
