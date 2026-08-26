package service

import "context"

func (s *Service) HealthCheck(ctx context.Context) error {
	return s.repo.DBCheck(ctx)
}
