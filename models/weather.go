package models

type Weather struct {
	City          string  `json:"city"`
	Temperature   float64 `json:"temperature"`
	Humidity      int     `json:"humidity"`
	Illumination  float64 `json:"illumination"`
	WaterPerMonth int     `json:"water_per_month"`
}
