package domain

type Message struct {
	Header      string  `json:"header"`
	Description string  `json:"description"`
	Image       *string `json:"image"` 
	Status      string  `json:"status"`
}