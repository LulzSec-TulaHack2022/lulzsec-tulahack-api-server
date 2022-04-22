package application

import (
	"net/http"
	"tulahackTest/models"
)

func AddUser(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		user, err := ParseUser(r)
		if err != nil {
			app.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		err = app.DB("auth").Insert(user)
		if err != nil {
			app.Error(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		app.Info("Registered new user: ", user)
	}
}

func SendUser(app *Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		CORS(w)

		data, err := ParseUser(r)
		if err != nil {
			app.Error(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		user, err := app.DB("auth").Get(data.UserID)
		if err != nil {
			app.Error(err)
		}

		err = SendData(w, user.(models.User))
		if err != nil {
			app.Error(err)
		}
	}
}
