package server

import (
	"net/http"

	"api-test/movement/infrastructure"

	"github.com/gorilla/mux"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/movement", infrastructure.Post).Methods(http.MethodPost)
	r.HandleFunc("/movements", infrastructure.GetList).Methods(http.MethodGet)
	r.HandleFunc("/movement", infrastructure.Get).Methods(http.MethodGet)
	r.HandleFunc("/ping", infrastructure.Ping).Methods(http.MethodGet)
	/* Gracias a Gorilla podemos usar expresiones regulares para asegurarnos
	 de antemano que los parámetros pasados cumplen con la regla que queremos.
	r.HandleFunc("/gophers/I{D:[a-zA-Z0-9_]+}", a.fetchGopher).Methods(http.MethodGet)*/
	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
