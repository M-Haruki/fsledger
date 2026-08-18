package common

import (
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
)

type Repository struct {
	queries *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{
		queries: q,
	}
}

// func (r *Repository) DBCheck(ctx context.Context) error {
// 	return r.queries.Check(ctx)
// }
