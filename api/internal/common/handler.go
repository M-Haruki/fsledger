package common

import (
	"context"
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

type CommonHandler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(service *Service, log *slog.Logger) *CommonHandler {
	return &CommonHandler{
		service: service,
		log:     log,
	}
}

func (h *CommonHandler) HealthCheck(ctx context.Context, request openapi.HealthCheckRequestObject) (openapi.HealthCheckResponseObject, error) {
	err := h.service.HealthCheck(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "health check failed", "error", err)
		return openapi.HealthCheck503JSONResponse{Message: "Unhealthy"}, nil
	}
	return openapi.HealthCheck204Response{}, nil
}
