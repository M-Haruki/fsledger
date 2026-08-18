package stock

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/google/uuid"
)

type Repository struct {
	queries *sqlc.Queries
}

func NewRepository(q *sqlc.Queries) *Repository {
	return &Repository{
		queries: q,
	}
}

func (r *Repository) ListStocks(ctx context.Context) ([]stockOverview, error) {
	raw, err := r.queries.ListStocks(ctx)
	result := make([]stockOverview, len(raw))
	for i, data := range raw {
		result[i] = stockOverview{
			id:   uuid.UUID(data.ID.Bytes),
			name: data.Name,
		}
	}
	return result, err
}
