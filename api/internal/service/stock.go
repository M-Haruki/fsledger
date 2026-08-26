package service

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
)

func (s *Service) ListStocks(ctx context.Context) ([]model.StockSummary, error) {
	return s.repo.ListStocks(ctx)
}
func (s *Service) CreateStock(ctx context.Context, stock model.Stock) (uuid.UUID, error) {
	return s.repo.CreateStock(ctx, stock)
}
func (s *Service) GetStock(ctx context.Context, id uuid.UUID) (model.Stock, error) {
	return s.repo.GetStock(ctx, id)
}
func (s *Service) UpdateStock(ctx context.Context, id uuid.UUID, stock model.Stock) error {
	return s.repo.UpdateStock(ctx, id, stock)
}
func (s *Service) DeleteStock(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteStock(ctx, id)
}
func (s *Service) ListStockTags(ctx context.Context) ([]model.Tag, error) {
	return s.repo.ListStockTags(ctx)
}
func (s *Service) CreateStockTag(ctx context.Context, name string) (uuid.UUID, error) {
	return s.repo.CreateStockTag(ctx, name)
}
func (s *Service) GetStockTag(ctx context.Context, id uuid.UUID) (string, error) {
	return s.repo.GetStockTag(ctx, id)
}
func (s *Service) UpdateStockTag(ctx context.Context, id uuid.UUID, name string) error {
	return s.repo.UpdateStockTag(ctx, id, name)
}
func (s *Service) DeleteStockTag(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteStockTag(ctx, id)
}
