package tag

import (
	"context"
	"log/slog"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

type TagHandler struct {
	service *Service
	log     *slog.Logger
}

func NewHandler(service *Service, log *slog.Logger) *TagHandler {
	return &TagHandler{
		service: service,
		log:     log,
	}
}

func (h TagHandler) ListTags(ctx context.Context, request openapi.ListTagsRequestObject) (openapi.ListTagsResponseObject, error) {
	return openapi.ListTags500JSONResponse{}, nil
}

func (h TagHandler) CreateTags(ctx context.Context, request openapi.CreateTagsRequestObject) (openapi.CreateTagsResponseObject, error) {
	return openapi.CreateTags201JSONResponse{}, nil
}

func (h TagHandler) DeleteTag(ctx context.Context, request openapi.DeleteTagRequestObject) (openapi.DeleteTagResponseObject, error) {
	return openapi.DeleteTag204Response{}, nil
}

func (h TagHandler) GetTag(ctx context.Context, request openapi.GetTagRequestObject) (openapi.GetTagResponseObject, error) {
	return openapi.GetTag200JSONResponse{}, nil
}

func (h TagHandler) UpdateTag(ctx context.Context, request openapi.UpdateTagRequestObject) (openapi.UpdateTagResponseObject, error) {
	return openapi.UpdateTag204Response{}, nil
}
