package main

import (
	"net/http"
	"tulahackTest/pkg/application"
)

func main() {
	app := application.NewApplication()

	s := &http.Server{
		Addr: app.Config().Addr,
		Handler: app.Router(),
	}

	app.Info("Server is started on port ", app.Config().Addr)
	app.Error(s.ListenAndServe())
}
