package application

import (
	"encoding/json"
	"net/http"
	"tulahackTest/models"
)

func GetFlower(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		var flowers []models.Flower

		flowers, err := app.DB().GetFlowers()
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to find flower in catalog", http.StatusBadRequest)
		}

		data, err := json.Marshal(flowers)
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