package models

type Flower struct {
	ID           int    `json:"id,omitempty"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	Temperature  string `json:"temperature"`
	Humidity     string `json:"humidity"`
	Illumination string `json:"illumination"`
}
