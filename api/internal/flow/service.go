package flow

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

func (s *Service) ListFlowTags(ctx context.Context) ([]tag, error) {
	return s.repo.ListFlowTags(ctx)
}
func (s *Service) CreateFlowTag(ctx context.Context, name string) (uuid.UUID, error) {
	return s.repo.CreateFlowTag(ctx, name)
}
func (s *Service) GetFlowTag(ctx context.Context, id uuid.UUID) (string, error) {
	return s.repo.GetFlowTag(ctx, id)
}
func (s *Service) UpdateFlowTag(ctx context.Context, id uuid.UUID, name string) error {
	return s.repo.UpdateFlowTag(ctx, id, name)
}
func (s *Service) DeleteFlowTag(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteFlowTag(ctx, id)
}
