package handler

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/openapi"
)

func (h *Handler) StockView(ctx context.Context, request openapi.StockViewRequestObject) (openapi.StockViewResponseObject, error) {
	return openapi.StockView500JSONResponse{Message: "Constructing"}, nil
}

func (h *Handler) TransactionView(ctx context.Context, request openapi.TransactionViewRequestObject) (openapi.TransactionViewResponseObject, error) {
	return openapi.TransactionView500JSONResponse{Message: "Constructing"}, nil
}

func (h *Handler) FlowView(ctx context.Context, request openapi.FlowViewRequestObject) (openapi.FlowViewResponseObject, error) {
	return openapi.FlowView500JSONResponse{Message: "Constructing"}, nil
}
