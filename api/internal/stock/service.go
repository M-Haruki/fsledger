package stock

import (
	"context"

	"github.com/google/uuid"
)

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) ListStocks(ctx context.Context) ([]stockSummary, error) {
	return s.repo.ListStocks(ctx)
}
func (s *Service) CreateStock(ctx context.Context, stock stockRequest) (uuid.UUID, error) {
	return s.repo.CreateStock(ctx, stock)
}
func (s *Service) GetStock(ctx context.Context, id uuid.UUID) (stockResponse, error) {
	return s.repo.GetStock(ctx, id)
}
func (s *Service) UpdateStock(ctx context.Context, id uuid.UUID, stock stockRequest) error {
	return s.repo.UpdateStock(ctx, id, stock)
}
func (s *Service) DeleteStock(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteStock(ctx, id)
}
