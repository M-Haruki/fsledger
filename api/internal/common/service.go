package common

import "context"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) HealthCheck(ctx context.Context) error {
	return s.repo.queries.Check(ctx)
}
