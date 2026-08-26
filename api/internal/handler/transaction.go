package handler

import (
	"context"
	"errors"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func (h *Handler) CreateTransaction(ctx context.Context, request openapi.CreateTransactionRequestObject) (openapi.CreateTransactionResponseObject, error) {
	return openapi.CreateTransaction500JSONResponse{Message: "constructing"}, nil
}
func (h *Handler) GetTransaction(ctx context.Context, request openapi.GetTransactionRequestObject) (openapi.GetTransactionResponseObject, error) {
	return openapi.GetTransaction500JSONResponse{Message: "constructing"}, nil
}
func (h *Handler) UpdateTransaction(ctx context.Context, request openapi.UpdateTransactionRequestObject) (openapi.UpdateTransactionResponseObject, error) {
	return openapi.UpdateTransaction500JSONResponse{Message: "constructing"}, nil
}
func (h *Handler) DeleteTransaction(ctx context.Context, request openapi.DeleteTransactionRequestObject) (openapi.DeleteTransactionResponseObject, error) {
	return openapi.DeleteTransaction500JSONResponse{Message: "constructing"}, nil
}

func (h *Handler) ListTransactionTags(ctx context.Context, request openapi.ListTransactionTagsRequestObject) (openapi.ListTransactionTagsResponseObject, error) {
	res, err := h.service.ListTransactionTags(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "list transaction tags failed", "error", err)
		return openapi.ListTransactionTags500JSONResponse{Message: "internal server error"}, nil
	}
	response := make(openapi.Tags, len(res))
	for i, atag := range res {
		response[i].Id = openapi_types.UUID(atag.Id)
		response[i].Name = atag.Name
	}
	return openapi.ListTransactionTags200JSONResponse(response), nil
}

func (h *Handler) CreateTransactionTags(ctx context.Context, request openapi.CreateTransactionTagsRequestObject) (openapi.CreateTransactionTagsResponseObject, error) {
	id, err := h.service.CreateTransactionTag(ctx, request.Body.Name)
	if err != nil {
		if errors.Is(err, model.ErrTransactionTagNameDuplicate) {
			return openapi.CreateTransactionTags400JSONResponse{Message: "transaction tag name duplicate"}, nil
		}
		h.log.ErrorContext(ctx, "create transaction tag failed", "error", err)
		return openapi.CreateTransactionTags500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateTransactionTags201JSONResponse{Id: openapi_types.UUID(id)}, nil
}

func (h *Handler) GetTransactionTag(ctx context.Context, request openapi.GetTransactionTagRequestObject) (openapi.GetTransactionTagResponseObject, error) {
	name, err := h.service.GetTransactionTag(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, model.ErrTransactionTagNotFound) {
			return openapi.GetTransactionTag404JSONResponse{Message: "transaction tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "get transaction tag failed", "error", err)
		return openapi.GetTransactionTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.GetTransactionTag200JSONResponse{Name: name}, nil
}

func (h *Handler) UpdateTransactionTag(ctx context.Context, request openapi.UpdateTransactionTagRequestObject) (openapi.UpdateTransactionTagResponseObject, error) {
	err := h.service.UpdateTransactionTag(ctx, uuid.UUID(request.Id), request.Body.Name)
	if err != nil {
		if errors.Is(err, model.ErrTransactionTagNotFound) {
			return openapi.UpdateTransactionTag404JSONResponse{Message: "transaction tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "update transaction tag failed", "error", err)
		return openapi.UpdateTransactionTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.UpdateTransactionTag204Response{}, nil
}

func (h *Handler) DeleteTransactionTag(ctx context.Context, request openapi.DeleteTransactionTagRequestObject) (openapi.DeleteTransactionTagResponseObject, error) {
	err := h.service.DeleteTransactionTag(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, model.ErrTransactionTagNotFound) {
			return openapi.DeleteTransactionTag404JSONResponse{Message: "transaction tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "delete transaction tag failed", "error", err)
		return openapi.DeleteTransactionTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.DeleteTransactionTag204Response{}, nil
}
