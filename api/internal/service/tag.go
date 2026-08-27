package service

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
)

func (s *Service) ListTags(ctx context.Context, tagT model.TagType) ([]model.Tag, error) {
	return s.repo.ListTags(ctx, tagT)
}

func (s *Service) CreateTag(ctx context.Context, tagT model.TagType, name string) (uuid.UUID, error) { // error: ErrTagNameDuplicate
	return s.repo.CreateTag(ctx, tagT, name)
}

func (s *Service) GetTag(ctx context.Context, tagT model.TagType, id uuid.UUID) (string, error) { // error: ErrTagNotFound
	return s.repo.GetTag(ctx, tagT, id)
}

func (s *Service) UpdateTag(ctx context.Context, tagT model.TagType, id uuid.UUID, name string) error { // error: ErrTagNotFound
	return s.repo.UpdateTag(ctx, tagT, id, name)
}

func (s *Service) DeleteTag(ctx context.Context, tagT model.TagType, id uuid.UUID) error { // error: ErrTagNotFound
	return s.repo.DeleteTag(ctx, tagT, id)
}
