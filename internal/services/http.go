package services

import (
	"elma_hw_2/internal/actions"
	"elma_hw_2/internal/models"
	"elma_hw_2/pkg/responses"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
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
		expire := time.Now().Add(time.Minute * time.Duration(models.TokenLive))
		cookie := http.Cookie{Name: "access_tocken", Value: result.AccessToken, Path: "/", Expires: expire}
		http.SetCookie(w, &cookie)
		responses.Make(w, result, err)
	})

	r.Put("/user", func(w http.ResponseWriter, r *http.Request) {
		user := &models.User{}
		user.FromRequest(r)
		result, err := actions.UpdateByUser(user, r)
		responses.Make(w, result, err)
	})

	r.Get("/user/login", func(w http.ResponseWriter, r *http.Request) {
		result, data, err := actions.GetLogin()
		responses.DrawPage(w, result, data, err)
	})

	r.Post("/user/login", func(w http.ResponseWriter, r *http.Request) {
		result, err := actions.GetUserByLogin(r)
		if err != nil {
			responses.Make(w, models.Err{ErrNo: 1, Message: err.Error()}, err)
			return
		}

		expire := time.Now().Add(time.Minute * time.Duration(models.TokenLive))
		cookie := http.Cookie{Name: "access_tocken", Value: result.AccessToken, Path: "/", Expires: expire}
		http.SetCookie(w, &cookie)
		responses.Make(w, result, err)
	})

	r.Get("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		result, err := actions.GetHelloWorld()
		responses.Make(w, result, err)
	})

	return r
}
