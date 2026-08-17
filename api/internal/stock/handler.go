package stock

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

type StockHandler struct {
	// service *Service
}

func NewHandler(
// service *Service
) *StockHandler {
	return &StockHandler{
		// service: service,
	}
}

func (h *StockHandler) ListStocks(ctx context.Context, request openapi.ListStocksRequestObject) (openapi.ListStocksResponseObject, error) {
	return openapi.ListStocks200JSONResponse{}, nil
}
func (h *StockHandler) CreateStock(ctx context.Context, request openapi.CreateStockRequestObject) (openapi.CreateStockResponseObject, error) {
	return openapi.CreateStock201JSONResponse{}, nil
}
