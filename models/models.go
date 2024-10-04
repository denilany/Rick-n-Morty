package models

type CharacterResponse struct {
	Info   ResponseInfo         `json:"info"`
	Result []ResponseCharacters `json:"results"`
}

type ResponseInfo struct {
	Count    int    `json:"count"`
	Pages    int    `json:"pages"`
	NextPage string `json:"next"`
	PrevPage string `json:"prev"`
}

type ResponseCharacters struct {
	ID           int               `json:"id"`
	Name         string            `json:"name"`
	Status       string            `json:"status"`
	Species      string            `json:"species"`
	Type         string            `json:"type"`
	Gender       string            `json:"gender"`
	Origin       CharacterOrigin   `json:"origin"`
	Location     CharacterLocation `json:"location"`
	Image        string            `json:"image"`
	Episode      []string          `json:"episode"`
	CharacterURL string            `json:"url"`
	Created      string            `json:"created"`
}

type CharacterOrigin struct {
	Name      string `json:"name"`
	OriginURL string `json:"url"`
}

type CharacterLocation struct {
	Name        string `json:"name"`
	LocationURL string `json:"url"`
}
