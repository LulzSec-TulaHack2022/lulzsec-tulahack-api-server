package application

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"tulahackTest/models"
)

func GetCatalog(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		var flowers []models.Flower

		flowers, err := app.DB().Catalog()
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

func GetCurrentWeather(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		weather := models.Weather{
			Temperature: app.Config().Parameters[rand.Intn(4)],
			Humidity: app.Config().Parameters[rand.Intn(4)],
			Illumination: rand.Intn(101),
		}

		data, err := json.Marshal(weather)
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

func AddUser(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to get data", http.StatusBadRequest)
			return
		}

		err = app.DB().AddUser(user)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to create new user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func AddFlower(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		var flower models.UserFlower

		err := json.NewDecoder(r.Body).Decode(&flower)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to parse flower data", http.StatusBadRequest)
			return
		}

		err = app.DB().AddFlower(flower)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to create new flower", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteFlower(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		flowerid := r.URL.Query().Get("flower_id")

		err := app.DB().DeleteFlower(flowerid)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to delete flower", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

func GetAllUserFlowers(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		userid := r.URL.Query().Get("owner_id")

		flowers, err := app.DB().GetAllUserFlowers(userid)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to get list of flowers", http.StatusBadRequest)
			return
		}

		data, err := json.Marshal(flowers)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to send data", http.StatusInternalServerError)
			return
		}
	}
}

func GetUserFlower(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		flowerid := r.URL.Query().Get("flower_id")

		flower, err := app.DB().GetUserFlower(flowerid)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to get flower data", http.StatusBadRequest)
			return
		}

		data, err := json.Marshal(flower)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
			return
		}

		_, err = w.Write(data)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to send data", http.StatusInternalServerError)
			return
		}
	}
}

func Dead(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		flowerid := r.URL.Query().Get("flower_id")

		err := app.DB().Dead(flowerid)
		if err != nil {
			app.Error(err)
			http.Error(w, "Unable to modify flower data", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

