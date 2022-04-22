package application

import (
	"encoding/json"
	"net/http"
	"tulahackTest/models"
)

func GetFlower(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		var flower models.Flower

		name := r.URL.Query().Get("name")

		flower, err := app.DB().GetFlowerInfo(name)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to find flower in catalog", http.StatusBadRequest)
		}

		data, err := json.Marshal(flower)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
		}

		_, err = w.Write(data)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to send data", http.StatusInternalServerError)
		}
	}
}