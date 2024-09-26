package service

import "songlib/internal/sql/model"

type CreateRequest struct {
	Group string `json:"group"`
	Song  string `json:"song"`
}

type UpdateRequest struct {
	Data map[string]interface{} `json:"data"`
}

type ListRequest struct {
	Filtermap map[string]interface{} `json:"filter"`
}
type ListResponse struct {
	Songs []*model.Song `json:"songs"` // Мне лень дальше инкапсулировать логику по отношению к выходам в json
}
