package common

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

type CommonHandler struct {
	service *Service
}

func NewHandler(service *Service) *CommonHandler {
	return &CommonHandler{
		service: service,
	}
}

func (h *CommonHandler) HealthCheck(ctx context.Context, request openapi.HealthCheckRequestObject) (openapi.HealthCheckResponseObject, error) {
	err := h.service.HealthCheck(ctx)
	if err != nil {
		return openapi.HealthCheck500Response{}, nil
	}
	return openapi.HealthCheck204Response{}, nil
}
