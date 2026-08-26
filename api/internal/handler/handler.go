package handler

import (
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/service"
)

type Handler struct {
	service *service.Service
	log     *slog.Logger
}

func NewHandler(service *service.Service, log *slog.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}
