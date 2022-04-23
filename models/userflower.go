package models

type UserFlower struct {
	ID        int `json:"id,omitempty"`
	CatalogID int `json:"catalog_id"`
	FlowerID  string `json:"flower_id"`
	OwnerID   string `json:"owner_id"`
	Name      string `json:"name"`
	Alive     bool `json:"alive,omitempty"`
}
