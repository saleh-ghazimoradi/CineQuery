package handlers

import (
	"github.com/saleh-ghazimoradi/CineQuery/config"
	"github.com/saleh-ghazimoradi/CineQuery/internal/helper"
	"log/slog"
	"net/http"
)

type HealthHandler struct {
	config    *config.Config
	logger    *slog.Logger
	customErr *helper.CustomErr
}

func (h *HealthHandler) Health(w http.ResponseWriter, r *http.Request) {
	env := helper.Envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": h.config.Application.Environment,
			"version":     h.config.Application.Version,
		},
	}

	if err := helper.WriteJSON(w, http.StatusOK, env, nil); err != nil {
		h.customErr.ServerErrorResponse(w, r, err)
	}
}

func NewHealthHandler(logger *slog.Logger, config *config.Config, customErr *helper.CustomErr) *HealthHandler {
	return &HealthHandler{
		logger:    logger,
		config:    config,
		customErr: customErr,
	}
}
