package service

import "songlib/internal/sql/model"

// CreateRequest represents the payload to create a new song
// @Description Payload for creating a new song
type CreateRequest struct {
	Group string `json:"group" example:"Muse"`
	Song  string `json:"song" example:"Supermassive Black Hole"`
}

// UpdateRequest represents the payload to update a song
// @Description Payload for updating a song
type UpdateRequest struct {
	// example: {"song": "New Song Title"}
	Data map[string]interface{} `json:"data"`
}

// ListRequest represents the payload for filtering songs
// @Description Payload for filtering song list
type ListRequest struct {
	// example: {"like:group": "Mu%"}
	Filtermap map[string]interface{} `json:"filter"`
}

// ListResponse represents the response for the filtered song list
// @Description Response containing a list of songs
type ListResponse struct {
	Songs []*model.Song `json:"songs"`
}

// ErrorResponse represents an error response
// @Description Error response
type ErrorResponse struct {
	Message string `json:"message"`
}
