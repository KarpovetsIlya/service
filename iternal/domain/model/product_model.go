package model

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Volume      string `json:"volume"`
	Alcohol     string `json:"alcohol"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
	Category    string `json:"category"`
}
