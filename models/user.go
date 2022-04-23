package models

type User struct {
	ID      int `json:"id,omitempty"`
	UserID  string `json:"userid"`
	Flowers []int `json:"flowers,omitempty"`
}
