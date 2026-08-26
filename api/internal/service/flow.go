package service

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
)

func (s *Service) ListFlowTags(ctx context.Context) ([]model.Tag, error) {
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
