package repository

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db      *pgxpool.Pool
	queries *sqlc.Queries
}

func NewRepository(d *pgxpool.Pool, q *sqlc.Queries) *Repository {
	return &Repository{
		db:      d,
		queries: q,
	}
}

func (r *Repository) BeginTx(ctx context.Context) (pgx.Tx, error) {
	return r.db.Begin(ctx)
}

func (r *Repository) WithTx(tx pgx.Tx) *Repository {
	return &Repository{
		db:      r.db,
		queries: r.queries.WithTx(tx),
	}
}
