package common

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HealthCheck(ctx context.Context, request openapi.HealthCheckRequestObject) (openapi.HealthCheckResponseObject, error) {
	err := h.service.HealthCheck(ctx)
	if err != nil {
		return openapi.HealthCheck500Response{}, nil
	}
	return openapi.HealthCheck204Response{}, nil
}
