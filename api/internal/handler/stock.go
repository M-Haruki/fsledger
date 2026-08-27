package handler

import (
	"context"
	"errors"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) ListStocks(ctx context.Context, request openapi.ListStocksRequestObject) (openapi.ListStocksResponseObject, error) {
	raw, err := h.service.ListStocks(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "list stock failed", "error", err)
		return openapi.ListStocks500JSONResponse{Message: "internal server error"}, nil
	}
	result := make(openapi.Stocks, len(raw))
	for i, data := range raw {
		result[i].Id = openapi_types.UUID(data.Id)
		result[i].Name = data.Name
		result[i].HasAmount = data.HasAmount
		result[i].Currency = data.Currency
		result[i].CurrencyExponent = data.CurrencyExponent
	}
	return openapi.ListStocks200JSONResponse(result), nil
}

func (h *Handler) CreateStock(ctx context.Context, request openapi.CreateStockRequestObject) (openapi.CreateStockResponseObject, error) {
	req := model.Stock{
		Name:             request.Body.Name,
		HasAmount:        request.Body.HasAmount,
		Currency:         request.Body.Currency,
		CurrencyExponent: request.Body.CurrencyExponent,
		Description:      request.Body.Description,
		Tags:             make([]uuid.UUID, len(request.Body.Tags)),
	}
	for i, id := range request.Body.Tags {
		req.Tags[i] = uuid.UUID(id)
	}
	id, err := h.service.CreateStock(ctx, req)
	if err != nil {
		if errors.Is(err, model.ErrStockNameDuplicate) {
			return openapi.CreateStock409JSONResponse{Message: "stock name duplicate"}, nil
		}
		if errors.Is(err, model.ErrTagNotFound) {
			return openapi.CreateStock404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "create stock failed", "error", err)
		return openapi.CreateStock500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateStock201JSONResponse{Id: openapi_types.UUID(id)}, nil
}

func (h *Handler) GetStock(ctx context.Context, request openapi.GetStockRequestObject) (openapi.GetStockResponseObject, error) {
	id := uuid.UUID(request.Id)
	res, err := h.service.GetStock(ctx, id)
	if err != nil {
		if errors.Is(err, model.ErrStockNotFound) {
			return openapi.GetStock404JSONResponse{Message: "stock not found"}, nil
		}
		h.log.ErrorContext(ctx, "get stock failed", "error", err)
		return openapi.GetStock500JSONResponse{}, nil
	}
	response := openapi.StockData{
		Name:             res.Name,
		HasAmount:        res.HasAmount,
		Currency:         res.Currency,
		CurrencyExponent: res.CurrencyExponent,
		Description:      res.Description,
		Tags:             make([]openapi_types.UUID, len(res.Tags)),
	}
	for i, tag := range res.Tags {
		response.Tags[i] = openapi_types.UUID(tag)
	}
	return openapi.GetStock200JSONResponse(response), nil
}

func (h *Handler) UpdateStock(ctx context.Context, request openapi.UpdateStockRequestObject) (openapi.UpdateStockResponseObject, error) {
	req := model.Stock{
		Name:             request.Body.Name,
		HasAmount:        request.Body.HasAmount,
		Currency:         request.Body.Currency,
		CurrencyExponent: request.Body.CurrencyExponent,
		Description:      request.Body.Description,
		Tags:             make([]uuid.UUID, len(request.Body.Tags)),
	}
	for i, id := range request.Body.Tags {
		req.Tags[i] = uuid.UUID(id)
	}
	err := h.service.UpdateStock(ctx, uuid.UUID(request.Id), req)
	if err != nil {
		if errors.Is(err, model.ErrStockNotFound) {
			return openapi.UpdateStock404JSONResponse{Message: "stock not found"}, nil
		}
		if errors.Is(err, model.ErrStockNameDuplicate) {
			return openapi.UpdateStock409JSONResponse{Message: "stock name duplicate"}, nil
		}
		if errors.Is(err, model.ErrTagNotFound) {
			return openapi.UpdateStock404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "update stock failed", "error", err)
		return openapi.UpdateStock500JSONResponse{}, nil
	}
	return openapi.UpdateStock204Response{}, nil
}

func (h *Handler) DeleteStock(ctx context.Context, request openapi.DeleteStockRequestObject) (openapi.DeleteStockResponseObject, error) {
	err := h.service.DeleteStock(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, model.ErrStockNotFound) {
			return openapi.DeleteStock404JSONResponse{Message: "stock not found"}, nil
		} else if errors.Is(err, model.ErrStockCannotDelete) {
			return openapi.DeleteStock409JSONResponse{Message: "stock cannot delete"}, nil
		}
		h.log.ErrorContext(ctx, "delete stock failed", "error", err)
		return openapi.DeleteStock500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.DeleteStock204Response{}, nil
}

func (h *Handler) ListStockTags(ctx context.Context, request openapi.ListStockTagsRequestObject) (openapi.ListStockTagsResponseObject, error) {
	res, err := h.service.ListTags(ctx, model.StockTag)
	if err != nil {
		h.log.ErrorContext(ctx, "list stock tags failed", "error", err)
		return openapi.ListStockTags500JSONResponse{Message: "internal server error"}, nil
	}
	response := make(openapi.Tags, len(res))
	for i, atag := range res {
		response[i].Id = openapi_types.UUID(atag.Id)
		response[i].Name = atag.Name
	}
	return openapi.ListStockTags200JSONResponse(response), nil
}

func (h *Handler) CreateStockTags(ctx context.Context, request openapi.CreateStockTagsRequestObject) (openapi.CreateStockTagsResponseObject, error) {
	id, err := h.service.CreateTag(ctx, model.StockTag, request.Body.Name)
	if err != nil {
		if errors.Is(err, model.ErrTagNameDuplicate) {
			return openapi.CreateStockTags409JSONResponse{Message: "stock tag name duplicate"}, nil
		}
		h.log.ErrorContext(ctx, "create stock tag failed", "error", err)
		return openapi.CreateStockTags500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateStockTags201JSONResponse{Id: openapi_types.UUID(id)}, nil
}

func (h *Handler) GetStockTag(ctx context.Context, request openapi.GetStockTagRequestObject) (openapi.GetStockTagResponseObject, error) {
	name, err := h.service.GetTag(ctx, model.StockTag, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, model.ErrTagNotFound) {
			return openapi.GetStockTag404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "get stock tag failed", "error", err)
		return openapi.GetStockTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.GetStockTag200JSONResponse{Name: name}, nil
}

func (h *Handler) UpdateStockTag(ctx context.Context, request openapi.UpdateStockTagRequestObject) (openapi.UpdateStockTagResponseObject, error) {
	err := h.service.UpdateTag(ctx, model.StockTag, uuid.UUID(request.Id), request.Body.Name)
	if err != nil {
		if errors.Is(err, model.ErrTagNotFound) {
			return openapi.UpdateStockTag404JSONResponse{Message: "stock tag not found"}, nil
		} else if errors.Is(err, model.ErrTagNameDuplicate) {
			return openapi.UpdateStockTag409JSONResponse{Message: "stock tag name duplicate"}, nil
		}
		h.log.ErrorContext(ctx, "update stock tag failed", "error", err)
		return openapi.UpdateStockTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.UpdateStockTag204Response{}, nil
}

func (h *Handler) DeleteStockTag(ctx context.Context, request openapi.DeleteStockTagRequestObject) (openapi.DeleteStockTagResponseObject, error) {
	err := h.service.DeleteTag(ctx, model.StockTag, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, model.ErrTagNotFound) {
			return openapi.DeleteStockTag404JSONResponse{Message: "stock tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "delete stock tag failed", "error", err)
		return openapi.DeleteStockTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.DeleteStockTag204Response{}, nil
}
