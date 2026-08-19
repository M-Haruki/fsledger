package server

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/db"
	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func newDB(ctx context.Context, cfg Config) (*pgxpool.Pool, *sqlc.Queries, error) {
	dbPool, err := db.New(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, nil, err
	}
	queries := sqlc.New(dbPool)
	return dbPool, queries, nil
}
