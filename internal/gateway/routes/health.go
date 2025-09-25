package routes

import (
	"github.com/julienschmidt/httprouter"
	"github.com/saleh-ghazimoradi/CineQuery/internal/gateway/handlers"
	"net/http"
)

type HealthRoutes struct {
	healthHandler *handlers.HealthHandler
}

func (h *HealthRoutes) Health(router *httprouter.Router) {
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", h.healthHandler.Health)
}

func NewHealthRoutes(healthHandler *handlers.HealthHandler) *HealthRoutes {
	return &HealthRoutes{healthHandler: healthHandler}
}
