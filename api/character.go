package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	apiconfig "github.com/denilany/Rick-n-Morty/constant"
	"github.com/denilany/Rick-n-Morty/customerrors"
	"github.com/denilany/Rick-n-Morty/models"
)

func GetCharacters(w http.ResponseWriter, page string) (*models.CharacterResponse, error) {
	url := apiconfig.BaseURL + apiconfig.Character

	if page != "" {
		url += "?page=" + page
	}

	client := &http.Client{
		Timeout: 10 * time.Second, // Ten second timeout
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		return nil, fmt.Errorf("failed to create new request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")

	// Make request using the client
	resp, err := client.Do(req)
	if err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		return nil, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusBadRequest,
			Message:    "Bad Request",
		})
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var characterResponse models.CharacterResponse
	if err := json.NewDecoder(resp.Body).Decode(&characterResponse); err != nil {
		customerrors.ServeErrorPage(w, customerrors.ErrorPage{
			StatusCode: http.StatusInternalServerError,
			Message:    "An Unexpected Error Occurred. Try Again Later",
		})
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &characterResponse, nil
}
