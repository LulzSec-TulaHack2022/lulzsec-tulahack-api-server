package application

// ! Хелперы для работы с парсингом и отправкой JSON, нужно переписать, недостаточно гибкая

import (
	"encoding/json"
	"net/http"
	"tulahackTest/models"
)

func SendData(w http.ResponseWriter, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
		return err
	}

	_, err = w.Write(b)
	if err != nil {
		http.Error(w, "Unable to send data", http.StatusInternalServerError)
		return err
	}

	return nil
}

func ParseUser(r *http.Request) (models.User, error) {
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}