package middleware

import (
	"fmt"
	"github.com/saleh-ghazimoradi/CineQuery/internal/helper"
	"net/http"
)

type Middleware struct {
	customErr *helper.CustomErr
}

func (m *Middleware) RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				m.customErr.ServerErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func NewMiddleware() *Middleware {
	return &Middleware{}
}
