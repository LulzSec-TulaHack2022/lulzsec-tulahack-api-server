package models

// ! Тестовая модель пользователя

type User struct {
	ID       int    `json:"id,omitempty"`
	UserID   string `json:"userid"`
}
