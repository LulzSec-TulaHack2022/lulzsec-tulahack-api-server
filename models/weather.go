package models

type Weather struct {
	Temperature  string `json:"temperature"`
	Humidity     string `json:"humidity"`
	Illumination int    `json:"illumination"`
}
