package stock

import (
	"context"
	"database/sql"
	"errors"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
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

func (r *Repository) ListStocks(ctx context.Context) ([]stockSummary, error) {
	raw, err := r.queries.ListStocks(ctx)
	result := make([]stockSummary, len(raw))
	for i, data := range raw {
		result[i] = stockSummary{
			id:   uuid.UUID(data.ID.Bytes),
			name: data.Name,
		}
	}
	return result, err
}

func (r *Repository) CreateStock(ctx context.Context, stock stockRequest) (uuid.UUID, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	id, err := q.CreateStock(ctx, sqlc.CreateStockParams{
		Name:        stock.name,
		Hasamount:   stock.hasAmount,
		Currency:    stock.currency,
		Description: stock.description,
	})
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return uuid.Nil, ErrStockNameDuplicate
			}
		}
		return uuid.Nil, err
	}
	for _, tagId := range stock.tags {
		err := q.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{StockID: id, TagID: pgtype.UUID{Bytes: tagId, Valid: true}})
		if err != nil {
			if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
				if pgErr.Code == "23503" {
					// tag not found
					return uuid.Nil, ErrStockTagNotFound
				}
			}
			return uuid.Nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return uuid.UUID(id.Bytes), nil
}

func (r *Repository) GetStock(ctx context.Context, id uuid.UUID) (stockResponse, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return stockResponse{}, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	stock, err := q.GetStock(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return stockResponse{}, ErrStockNotFound
		}
		return stockResponse{}, err
	}
	tags, err := q.ListTagsByStock(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return stockResponse{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return stockResponse{}, err
	}

	res := stockResponse{
		name:        stock.Name,
		hasAmount:   stock.HasAmount,
		currency:    stock.Currency,
		description: stock.Description,
		tags:        make([]tag, len(tags)),
	}
	for i, atag := range tags {
		res.tags[i] = tag{id: uuid.UUID(atag.ID.Bytes), name: atag.Name}
	}

	return res, nil
}

func (r *Repository) UpdateStock(ctx context.Context, id uuid.UUID, stock stockRequest) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	result, err := q.UpdateStock(ctx, sqlc.UpdateStockParams{
		ID:          pgtype.UUID{Bytes: id, Valid: true},
		Name:        stock.name,
		Hasamount:   stock.hasAmount,
		Currency:    stock.currency,
		Description: stock.description,
	})
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return ErrStockNameDuplicate
			}
		}
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return ErrStockNotFound
	}
	err = q.DeleteStockTagRelation(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	for _, tagId := range stock.tags {
		err := q.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{StockID: pgtype.UUID{Bytes: id, Valid: true}, TagID: pgtype.UUID{Bytes: tagId, Valid: true}})
		if err != nil {
			if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
				if pgErr.Code == "23503" {
					// tag not found
					return ErrStockTagNotFound
				}
			}
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) DeleteStock(ctx context.Context, id uuid.UUID) error {
	result, err := r.queries.DeleteStock(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return ErrStockNotFound
	}
	return nil
}
