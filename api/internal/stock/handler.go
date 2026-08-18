package stock

import (
	"context"
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/google/uuid"
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
		h.log.ErrorContext(ctx, "list stock failed", "error", err)
		return openapi.ListStocks500JSONResponse{Message: "internal server error"}, nil
	}
	result := make(openapi.Stocks, len(raw))
	for i, data := range raw {
		result[i].Id = openapi_types.UUID(data.id)
		result[i].Name = data.name
	}
	return openapi.ListStocks200JSONResponse(result), nil
}

func (h *StockHandler) CreateStock(ctx context.Context, request openapi.CreateStockRequestObject) (openapi.CreateStockResponseObject, error) {
	req := stock{
		name:        request.Body.Name,
		has_amount:  request.Body.HasAmount,
		currency:    request.Body.Currency,
		description: request.Body.Description,
		tags:        make([]uuid.UUID, len(request.Body.Tags)),
	}
	for i, id := range request.Body.Tags {
		req.tags[i] = uuid.UUID(id)
	}
	id, err := h.service.repo.CreateStock(ctx, req)
	if err != nil {
		h.log.ErrorContext(ctx, "create stock failed", "error", err)
		return openapi.CreateStock500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateStock201JSONResponse{Id: openapi_types.UUID(id)}, nil
}
