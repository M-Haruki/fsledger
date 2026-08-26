package repository

import "context"

func (r *Repository) DBCheck(ctx context.Context) error {
	return r.queries.Check(ctx)
}
