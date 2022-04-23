package models

type Weather struct {
	City         string  `json:"city"`
	Temperature  float64 `json:"temperature"`
	Humidity     int     `json:"humidity"`
	Illumination int     `json:"illumination"`
}
