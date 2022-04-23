package location

import (
	owm "github.com/briandowns/openweathermap"
)

func GetCurrentWeather(lat, long float64, key string) (*owm.CurrentWeatherData, error) {
	w, err := owm.NewCurrent("C", "RU", key)
	if err != nil {
		return nil, err
	}

	err = w.CurrentByCoordinates(&owm.Coordinates{
		Longitude: long,
		Latitude: lat,
	})
	if err != nil {
		return nil, err
	}

	return w, nil
}

func GetForecast(lat, long float64, key string) (*owm.ForecastWeatherData, error) {
	w, err := owm.NewForecast("16", "C", "RU", key)
	if err != nil {
		return nil, err
	}

	err = w.DailyByCoordinates(&owm.Coordinates{
		Longitude: long,
		Latitude: lat,
	}, 16)
	if err != nil {
		return nil, err
	}

	return w, nil

	//if weather.Temperature > 24 {
	//	weather.WaterPerMonth = 8
	//} else if weather.Temperature > 18 {
	//	weather.WaterPerMonth = 6
	//} else {
	//	weather.WaterPerMonth = 3
	//}
	//
	//weather.WaterPerMonth -= int(float64(weather.WaterPerMonth - 1) * (math.Round(float64(weather.Humidity) / 100 - 0.2)))

}