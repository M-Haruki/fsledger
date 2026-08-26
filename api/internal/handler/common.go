package handler

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

func (h *Handler) HealthCheck(ctx context.Context, request openapi.HealthCheckRequestObject) (openapi.HealthCheckResponseObject, error) {
	err := h.service.HealthCheck(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "health check failed", "error", err)
		return openapi.HealthCheck503JSONResponse{Message: "Unhealthy"}, nil
	}
	return openapi.HealthCheck204Response{}, nil
}
