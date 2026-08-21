package stock

import (
	"context"
	"errors"
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
	raw, err := h.service.ListStocks(ctx)
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
	req := stockRequest{
		name:        request.Body.Name,
		hasAmount:   request.Body.HasAmount,
		currency:    request.Body.Currency,
		description: request.Body.Description,
		tags:        make([]uuid.UUID, len(request.Body.Tags)),
	}
	for i, id := range request.Body.Tags {
		req.tags[i] = uuid.UUID(id)
	}
	id, err := h.service.CreateStock(ctx, req)
	if err != nil {
		if errors.Is(err, ErrStockNameDuplicate) {
			return openapi.CreateStock400JSONResponse{Message: "stock name duplicate"}, nil
		}
		if errors.Is(err, ErrStockTagNotFound) {
			return openapi.CreateStock400JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "create stock failed", "error", err)
		return openapi.CreateStock500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateStock201JSONResponse{Id: openapi_types.UUID(id)}, nil
}

func (h *StockHandler) GetStock(ctx context.Context, request openapi.GetStockRequestObject) (openapi.GetStockResponseObject, error) {
	id := uuid.UUID(request.Id)
	res, err := h.service.GetStock(ctx, id)
	if err != nil {
		if errors.Is(err, ErrStockNotFound) {
			return openapi.GetStock404JSONResponse{Message: "stock not found"}, nil
		}
		h.log.ErrorContext(ctx, "get stock failed", "error", err)
		return openapi.GetStock500JSONResponse{}, nil
	}
	response := openapi.StockGetData{
		Name:        res.name,
		HasAmount:   res.hasAmount,
		Currency:    res.currency,
		Description: res.description,
		Tags:        make(openapi.Tags, len(res.tags)),
	}
	for i, tag := range res.tags {
		response.Tags[i].Id = openapi_types.UUID(tag.id)
		response.Tags[i].Name = tag.name
	}
	return openapi.GetStock200JSONResponse(response), nil
}

func (h *StockHandler) UpdateStock(ctx context.Context, request openapi.UpdateStockRequestObject) (openapi.UpdateStockResponseObject, error) {
	req := stockRequest{
		name:        request.Body.Name,
		hasAmount:   request.Body.HasAmount,
		currency:    request.Body.Currency,
		description: request.Body.Description,
		tags:        make([]uuid.UUID, len(request.Body.Tags)),
	}
	for i, id := range request.Body.Tags {
		req.tags[i] = uuid.UUID(id)
	}
	err := h.service.UpdateStock(ctx, uuid.UUID(request.Id), req)
	if err != nil {
		if errors.Is(err, ErrStockNotFound) {
			return openapi.UpdateStock404JSONResponse{Message: "stock not found"}, nil
		}
		if errors.Is(err, ErrStockNameDuplicate) {
			return openapi.UpdateStock400JSONResponse{Message: "stock name duplicate"}, nil
		}
		if errors.Is(err, ErrStockTagNotFound) {
			return openapi.UpdateStock400JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "update stock failed", "error", err)
		return openapi.UpdateStock500JSONResponse{}, nil
	}
	return openapi.UpdateStock204Response{}, nil
}

func (h *StockHandler) DeleteStock(ctx context.Context, request openapi.DeleteStockRequestObject) (openapi.DeleteStockResponseObject, error) {
	err := h.service.DeleteStock(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, ErrStockNotFound) {
			return openapi.DeleteStock404JSONResponse{Message: "stock not found"}, nil
		}
		h.log.ErrorContext(ctx, "delete stock failed", "error", err)
		return openapi.DeleteStock500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.DeleteStock204Response{}, nil
}

func (h *StockHandler) ListStockTags(ctx context.Context, request openapi.ListStockTagsRequestObject) (openapi.ListStockTagsResponseObject, error) {
	res, err := h.service.ListStockTags(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "list stock tags failed", "error", err)
		return openapi.ListStockTags500JSONResponse{Message: "internal server error"}, nil
	}
	response := make(openapi.Tags, len(res))
	for i, atag := range res {
		response[i].Id = openapi_types.UUID(atag.id)
		response[i].Name = atag.name
	}
	return openapi.ListStockTags200JSONResponse(response), nil
}

func (h *StockHandler) CreateStockTags(ctx context.Context, request openapi.CreateStockTagsRequestObject) (openapi.CreateStockTagsResponseObject, error) {
	id, err := h.service.CreateStockTag(ctx, request.Body.Name)
	if err != nil {
		if errors.Is(err, ErrStockTagNameDuplicate) {
			return openapi.CreateStockTags400JSONResponse{Message: "stock tag name duplicate"}, nil
		}
		h.log.ErrorContext(ctx, "create stock tag failed", "error", err)
		return openapi.CreateStockTags500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateStockTags201JSONResponse{Id: openapi_types.UUID(id)}, nil
}

func (h *StockHandler) GetStockTag(ctx context.Context, request openapi.GetStockTagRequestObject) (openapi.GetStockTagResponseObject, error) {
	name, err := h.service.GetStockTag(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, ErrStockTagNotFound) {
			return openapi.GetStockTag404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "get stock tag failed", "error", err)
		return openapi.GetStockTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.GetStockTag200JSONResponse{Name: name}, nil
}

func (h *StockHandler) UpdateStockTag(ctx context.Context, request openapi.UpdateStockTagRequestObject) (openapi.UpdateStockTagResponseObject, error) {
	err := h.service.UpdateStockTag(ctx, uuid.UUID(request.Id), request.Body.Name)
	if err != nil {
		if errors.Is(err, ErrStockTagNotFound) {
			return openapi.UpdateStockTag404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "update stock tag failed", "error", err)
		return openapi.UpdateStockTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.UpdateStockTag204Response{}, nil
}

func (h *StockHandler) DeleteStockTag(ctx context.Context, request openapi.DeleteStockTagRequestObject) (openapi.DeleteStockTagResponseObject, error) {
	err := h.service.DeleteStockTag(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, ErrStockTagNotFound) {
			return openapi.DeleteStockTag404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "delete stock tag failed", "error", err)
		return openapi.DeleteStockTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.DeleteStockTag204Response{}, nil
}
