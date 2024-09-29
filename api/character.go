package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/denilany/Rick-n-Morty/models"
)


func GetCharacterApi(CharacterURL string) (models.CharacterResponse, error) {
	client := &http.Client{
		Timeout: 10 * time.Second, // Ten second timeout
	}

	req, err := http.NewRequest("GET", CharacterURL, nil)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to create new request: %w", err)
	}

	// Set headers
	req.Header.Set("Accept", "application/json")

	// Make request using the client
	resp, err := client.Do(req)
	if err != nil {
		return models.CharacterResponse{}, fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.CharacterResponse{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var baseResponse models.CharacterResponse
	if err := json.NewDecoder(resp.Body).Decode(&baseResponse); err != nil {
		return models.CharacterResponse{}, fmt.Errorf("failed to decode response: %w", err)
	}

	return baseResponse, nil
}
