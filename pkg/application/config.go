package application

// * Конфигурационная структура для инициализации объекта приложения

import (
	"fmt"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"net/http"
	"os"
)

type Config struct {
	Addr string
	Parameters []string
}

func NewConfig() *Config {
	return &Config{
		Addr: fmt.Sprintf(":%v", os.Getenv("PORT")),
		Parameters: []string{"low", "medium", "high", "any"},
	}
}

func (app *Application) configureRouter() {
	router := http.NewServeMux()

	// Маршруты API
	router.HandleFunc("/getcatalog", GetCatalog(app))
	router.HandleFunc("/getcurrentweather", GetCurrentWeather(app))
	router.HandleFunc("/adduser", AddUser(app))
	router.HandleFunc("/addflower", AddFlower(app))
	router.HandleFunc("/deleteflower", DeleteFlower(app))

	app.router = router
}

func (app *Application) configureLogger() {
	infologger := &logrus.Logger{
		Level: logrus.InfoLevel,
		Hooks: make(logrus.LevelHooks),
		Out: os.Stdout,
		Formatter: &prefixed.TextFormatter{
			DisableColors: false,
			TimestampFormat : "2006-01-02 15:04:05",
			FullTimestamp: true,
			ForceFormatting: true,
		},
	}

	errorlogger := &logrus.Logger{
		Level: logrus.InfoLevel,
		Hooks: make(logrus.LevelHooks),
		Out: os.Stdout,
		Formatter: &prefixed.TextFormatter{
			DisableColors: false,
			TimestampFormat : "2006-01-02 15:04:05",
			FullTimestamp: true,
			ForceFormatting: true,
		},
	}

	app.infoLog = infologger
	app.errorLog = errorlogger
}

