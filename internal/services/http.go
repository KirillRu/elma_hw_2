package services

import (
	"elma_hw_2/internal/actions"
	"elma_hw_2/internal/models"
	"elma_hw_2/pkg/responses"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

func New() ServerImplementation {
	return ServerImplementation{}
}

type ServerImplementation struct {
}

func (s ServerImplementation) BuildRoutes() http.Handler {
	// https://github.com/go-chi/chi/blob/master/_examples/limits/main.go
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		result, data, err := actions.GetMain()
		responses.DrawPage(w, result, data, err)
	})

	r.Get("/user", func(w http.ResponseWriter, r *http.Request) {
		result, data, err := actions.GetUser(r)
		responses.DrawPage(w, result, data, err)
	})

	r.Post("/user", func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		user.FromRequest(r)
		result, err := actions.RegByUser(user)
		responses.Make(w, result, err)
	})

	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		result, data, err := actions.GetLogin()
		responses.DrawPage(w, result, data, err)
	})

	r.Get("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		result, err := actions.GetHelloWorld()
		responses.Make(w, result, err)
	})

	return r
}
