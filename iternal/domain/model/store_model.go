package model

type Store struct {
	ID          int    `json:"id"`
	Address     string `json:"address"`
	Coordinates string `json:"coordinates"`
}
