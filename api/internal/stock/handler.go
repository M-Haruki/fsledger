package stock

import (
	"context"
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type StockHandler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(service *Service, log *slog.Logger) *StockHandler {
	return &StockHandler{
		service: service,
		log:     log,
	}
}

func (h *StockHandler) ListStocks(ctx context.Context, request openapi.ListStocksRequestObject) (openapi.ListStocksResponseObject, error) {
	raw, err := h.service.repo.ListStocks(ctx)
	if err != nil {
		return openapi.ListStocks500JSONResponse{Error: "internal server error"}, nil
	}
	result := make(openapi.Stocks, len(raw))
	for i, data := range raw {
		result[i].Id = openapi_types.UUID(data.id)
		result[i].Name = data.name
	}
	return openapi.ListStocks200JSONResponse(result), nil
}

func (h *StockHandler) CreateStock(ctx context.Context, request openapi.CreateStockRequestObject) (openapi.CreateStockResponseObject, error) {
	return openapi.CreateStock201JSONResponse{}, nil
}
