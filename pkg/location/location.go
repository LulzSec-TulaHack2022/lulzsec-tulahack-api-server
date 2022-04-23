package location

import (
	owm "github.com/briandowns/openweathermap"
)

func GetWeather(lat, long float64, key string) (*owm.CurrentWeatherData, error) {
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