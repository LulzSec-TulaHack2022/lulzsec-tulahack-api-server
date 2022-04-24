package application

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
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

			lat, _ := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
			long, _ := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)

			we, err := location.GetCurrentWeather(lat, long, app.Config().OWMApiKey)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to get weather", http.StatusInternalServerError)
				return
			}
		
			weather := models.Weather{
				City: we.Name,
				Temperature: we.Main.Temp,
				Humidity: we.Main.Humidity,
			}

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
			var dat map[string]string
			err := json.NewDecoder(r.Body).Decode(&dat)

			err = app.DB().DeleteFlower(dat["flower_id"])
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to delete flower", http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusOK)
		}

		if r.Method == http.MethodGet {
			ownerid := r.URL.Query().Get("owner_id")
			lat, err := strconv.ParseFloat(r.URL.Query().Get("latitude"), 64)
			long, err := strconv.ParseFloat(r.URL.Query().Get("longitude"), 64)

			flowers, err := app.DB().GetAllUserFlowers(ownerid)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to get list of flowers", http.StatusBadRequest)
				return
			}

			we, err := location.GetCurrentWeather(lat, long, app.Config().OWMApiKey)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to get forecast", http.StatusInternalServerError)
				return
			}

			weather := models.Weather{
				City: we.Name,
				Temperature: we.Main.Temp,
				Humidity: we.Main.Humidity,
				Illumination: rand.Float64() * 24,
			}

			G := 9.0

			for k, _ := range flowers {
				flowers[k].WaterPerMonth = int(G * weather.Temperature * weather.Illumination / float64(weather.Humidity) - float64(flowers[k].ID))
				if flowers[k].WaterPerMonth > 16 {
					flowers[k].WaterPerMonth = 16
				} else if flowers[k].WaterPerMonth < 2 {
					flowers[k].WaterPerMonth = 2
				}
			}

			//app.Info(v)
			//
			//if weather.Temperature > 24 {
			//	weather.WaterPerMonth = 8
			//} else if weather.Temperature > 18 {
			//	weather.WaterPerMonth = 6
			//} else {
			//	weather.WaterPerMonth = 3
			//}

			//weather.WaterPerMonth += int(float64(weather.WaterPerMonth - 1) * (math.Round(float64(weather.Humidity) / 100 - 0.2)))
			//

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

func Dead(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		if r.Method == http.MethodGet {
			flowerid := r.URL.Query().Get("flower_id")

			err := app.DB().Dead(flowerid)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to kill plant", http.StatusBadRequest)
				return
			}
		}
	}
}

func Watered(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		if r.Method == http.MethodGet {
			flowerid := r.URL.Query().Get("flower_id")

			err := app.DB().Watered(flowerid, false)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to water plant", http.StatusBadRequest)
				return
			}
		}

		if r.Method == http.MethodPut {
			flowerid := r.URL.Query().Get("flower_id")

			err := app.DB().Watered(flowerid, true)
			if err != nil {
				app.Error(err)
				http.Error(w, "Unable to water plant", http.StatusBadRequest)
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

