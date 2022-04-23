package application

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"tulahackTest/models"
	"tulahackTest/pkg/location"
)

func Catalog(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		if r.Method == http.MethodGet {
			var flowers []models.Flower

			flowers, err := app.DB().Catalog()
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to find flower in catalog", http.StatusBadRequest)
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
}

func CurrentWeather(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		if r.Method == http.MethodGet {
			var loc map[string]float64

			err := json.NewDecoder(r.Body).Decode(&loc)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to parse data", http.StatusBadRequest)
				return
			}

			we, err := location.GetWeather(loc["latitude"], loc["longitude"], app.Config().OWMApiKey)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to get weather", http.StatusInternalServerError)
				return
			}

			weather := models.Weather{
				City: we.Name,
				Temperature: we.Main.Temp,
				Humidity: we.Main.Humidity,
				Illumination: rand.Intn(41) + 60,
			}

			if weather.Temperature > 24 {
				weather.WaterPerMonth = 6
			} else if weather.Temperature > 18 {
				weather.WaterPerMonth = 4
			} else if weather.Temperature > 15 {
				weather.WaterPerMonth = 2
			} else {
				weather.WaterPerMonth = 1
			}

			//weather.WaterPerMonth -= int(float64(weather.WaterPerMonth - 1) * (math.Round(float64(weather.Humidity) / 100 - 0.5ф)))

			data, err := json.Marshal(weather)
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
}

func Flower(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)
		if r.Method == http.MethodPost {
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

		if r.Method == http.MethodDelete {
			flowerid := r.URL.Query().Get("flower_id")

			err := app.DB().DeleteFlower(flowerid)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to delete flower", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		}

		if r.Method == http.MethodGet {
			var dat map[string]string
			err := json.NewDecoder(r.Body).Decode(&dat)

			flowers, err := app.DB().GetAllUserFlowers(dat["owner_id"])
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
}

//func GetUserFlower(app *Application) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		CORS(w)
//
//		flowerid := r.URL.Query().Get("flower_id")
//
//		flower, err := app.DB().GetUserFlower(flowerid)
//		if err != nil {
//			app.Error(err)
//			http.Error(w, "Unable to get flower data", http.StatusBadRequest)
//			return
//		}
//
//		data, err := json.Marshal(flower)
//		if err != nil {
//			app.Error(err)
//			http.Error(w, "Unable to marshal data", http.StatusInternalServerError)
//			return
//		}
//
//		_, err = w.Write(data)
//		if err != nil {
//			app.Error(err)
//			http.Error(w, "Unable to send data", http.StatusInternalServerError)
//			return
//		}
//	}
//}

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

