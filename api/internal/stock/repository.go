package stock

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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

func (r *Repository) CreateStock(ctx context.Context, stock stock) (uuid.UUID, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)
	id, err := q.CreateStock(ctx, sqlc.CreateStockParams{
		Name:        stock.name,
		Hasamount:   stock.has_amount,
		Currency:    stock.currency,
		Description: stock.description,
	})
	if err != nil {
		return uuid.Nil, err
	}
	for _, tagId := range stock.tags {
		err := q.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{StockID: id, TagID: pgtype.UUID{Bytes: tagId, Valid: true}})
		if err != nil {
			return uuid.Nil, err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.UUID(id.Bytes), nil
}
