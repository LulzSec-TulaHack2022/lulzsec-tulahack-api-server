package models

type UserFlower struct {
	ID               int    `json:"-"`
	NameNomenclature string `json:"name_nomenclature"`
	FlowerID         string `json:"id"`
	OwnerID          string `json:"owner_id"`
	Name             string `json:"name"`
	Alive            bool   `json:"alive"`
	NeedWater        bool   `json:"need_water"`
}
