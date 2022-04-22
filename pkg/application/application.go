package application

// * Структура приложения, объединяет в себе все компоненты веб-приложения

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"tulahackTest/pkg/storage"
)

type Application struct {
	errorLog *logrus.Logger
	infoLog  *logrus.Logger
	router   *http.ServeMux
	storage  *storage.Storage
	config   *Config
}


func NewApplication() *Application {
	app := &Application{
		storage: storage.NewStorage(),
		config: NewConfig(),
	}

	app.configureRouter()
	app.configureLogger()

	return app
}

func (app *Application) Config() *Config {
	return app.config
}

func (app *Application) Router() *http.ServeMux {
	return app.router
}

func (app *Application) DB(repo string) storage.Repo {
	switch repo {
	case "auth":
		return app.storage.Authrepo()

	default:
		return nil
	}
}

func (app *Application) Info(content ...interface{}) {
	app.infoLog.Println(content...)
}

func (app *Application) Error(content ...interface{}) {
	app.errorLog.Error(content...)
}