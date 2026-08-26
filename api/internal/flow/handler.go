package flow

import (
	"context"
	"errors"
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
	"github.com/google/uuid"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

type FlowHandler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(service *Service, log *slog.Logger) *FlowHandler {
	return &FlowHandler{
		service: service,
		log:     log,
	}
}

func (h *FlowHandler) ListFlowTags(ctx context.Context, request openapi.ListFlowTagsRequestObject) (openapi.ListFlowTagsResponseObject, error) {
	res, err := h.service.ListFlowTags(ctx)
	if err != nil {
		h.log.ErrorContext(ctx, "list flow tags failed", "error", err)
		return openapi.ListFlowTags500JSONResponse{Message: "internal server error"}, nil
	}
	response := make(openapi.Tags, len(res))
	for i, atag := range res {
		response[i].Id = openapi_types.UUID(atag.id)
		response[i].Name = atag.name
	}
	return openapi.ListFlowTags200JSONResponse(response), nil
}

func (h *FlowHandler) CreateFlowTags(ctx context.Context, request openapi.CreateFlowTagsRequestObject) (openapi.CreateFlowTagsResponseObject, error) {
	id, err := h.service.CreateFlowTag(ctx, request.Body.Name)
	if err != nil {
		if errors.Is(err, ErrFlowTagNameDuplicate) {
			return openapi.CreateFlowTags400JSONResponse{Message: "flow tag name duplicate"}, nil
		}
		h.log.ErrorContext(ctx, "create flow tag failed", "error", err)
		return openapi.CreateFlowTags500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.CreateFlowTags201JSONResponse{Id: openapi_types.UUID(id)}, nil
}

func (h *FlowHandler) GetFlowTag(ctx context.Context, request openapi.GetFlowTagRequestObject) (openapi.GetFlowTagResponseObject, error) {
	name, err := h.service.GetFlowTag(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, ErrFlowTagNotFound) {
			return openapi.GetFlowTag404JSONResponse{Message: "flow tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "get flow tag failed", "error", err)
		return openapi.GetFlowTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.GetFlowTag200JSONResponse{Name: name}, nil
}

func (h *FlowHandler) UpdateFlowTag(ctx context.Context, request openapi.UpdateFlowTagRequestObject) (openapi.UpdateFlowTagResponseObject, error) {
	err := h.service.UpdateFlowTag(ctx, uuid.UUID(request.Id), request.Body.Name)
	if err != nil {
		if errors.Is(err, ErrFlowTagNotFound) {
			return openapi.UpdateFlowTag404JSONResponse{Message: "flow tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "update flow tag failed", "error", err)
		return openapi.UpdateFlowTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.UpdateFlowTag204Response{}, nil
}

func (h *FlowHandler) DeleteFlowTag(ctx context.Context, request openapi.DeleteFlowTagRequestObject) (openapi.DeleteFlowTagResponseObject, error) {
	err := h.service.DeleteFlowTag(ctx, uuid.UUID(request.Id))
	if err != nil {
		if errors.Is(err, ErrFlowTagNotFound) {
			return openapi.DeleteFlowTag404JSONResponse{Message: "flow tag not found"}, nil
		}
		h.log.ErrorContext(ctx, "delete flow tag failed", "error", err)
		return openapi.DeleteFlowTag500JSONResponse{Message: "internal server error"}, nil
	}
	return openapi.DeleteFlowTag204Response{}, nil
}
