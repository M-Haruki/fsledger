package transaction

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

func (s *Service) ListTransactionTags(ctx context.Context) ([]tag, error) {
	return s.repo.ListTransactionTags(ctx)
}
func (s *Service) CreateTransactionTag(ctx context.Context, name string) (uuid.UUID, error) {
	return s.repo.CreateTransactionTag(ctx, name)
}
func (s *Service) GetTransactionTag(ctx context.Context, id uuid.UUID) (string, error) {
	return s.repo.GetTransactionTag(ctx, id)
}
func (s *Service) UpdateTransactionTag(ctx context.Context, id uuid.UUID, name string) error {
	return s.repo.UpdateTransactionTag(ctx, id, name)
}
func (s *Service) DeleteTransactionTag(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteTransactionTag(ctx, id)
}
