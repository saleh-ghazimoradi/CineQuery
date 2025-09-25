package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/CineQuery/internal/helper"
	"github.com/saleh-ghazimoradi/CineQuery/internal/middleware"
	"net/http"
)

type RegisterRoutes struct {
	CustomErr    *helper.CustomErr
	MiddleWares  *middleware.Middleware
	HealthRoutes *HealthRoutes
	MovieRoutes  *MovieRoutes
}

type Option func(*RegisterRoutes)

func WithCustomErr(customErr *helper.CustomErr) Option {
	return func(r *RegisterRoutes) {
		r.CustomErr = customErr
	}
}

func WithMiddleWares(middleWares *middleware.Middleware) Option {
	return func(r *RegisterRoutes) {
		r.MiddleWares = middleWares
	}
}

func WithHealthRoutes(healthRoutes *HealthRoutes) Option {
	return func(r *RegisterRoutes) {
		r.HealthRoutes = healthRoutes
	}
}

func WithMovieRoutes(movieRoutes *MovieRoutes) Option {
	return func(r *RegisterRoutes) {
		r.MovieRoutes = movieRoutes
	}
}

func (r *RegisterRoutes) Register() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(r.CustomErr.NotFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(r.CustomErr.MethodNotAllowedResponse)

	r.HealthRoutes.Health(router)
	r.MovieRoutes.Movie(router)

	return r.MiddleWares.RecoverPanic(router)
}

func NewRegisterRoutes(opts ...Option) *RegisterRoutes {
	r := &RegisterRoutes{}
	for _, opt := range opts {
		opt(r)
	}
	return r
}
