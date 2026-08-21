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

func (r *Repository) CreateStock(ctx context.Context, stock stock) (uuid.UUID, error) {
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

func (r *Repository) GetStock(ctx context.Context, id uuid.UUID) (stock, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return stock{}, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	astock, err := q.GetStock(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return stock{}, ErrStockNotFound
		}
		return stock{}, err
	}
	tags, err := q.ListTagIDsByStock(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return stock{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return stock{}, err
	}

	res := stock{
		name:        astock.Name,
		hasAmount:   astock.HasAmount,
		currency:    astock.Currency,
		description: astock.Description,
		tags:        make([]uuid.UUID, len(tags)),
	}
	for i, atag := range tags {
		res.tags[i] = uuid.UUID(atag.Bytes)
	}

	return res, nil
}

func (r *Repository) UpdateStock(ctx context.Context, id uuid.UUID, stock stock) error {
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

func (r *Repository) ListStockTags(ctx context.Context) ([]tag, error) {
	res, err := r.queries.ListStockTags(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]tag, len(res))
	for i, atag := range res {
		result[i] = tag{
			id:   uuid.UUID(atag.ID.Bytes),
			name: atag.Name,
		}
	}
	return result, nil
}

func (r *Repository) CreateStockTag(ctx context.Context, name string) (uuid.UUID, error) {
	id, err := r.queries.CreateStockTag(ctx, name)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return uuid.Nil, ErrStockTagNameDuplicate
			}
		}
		return uuid.Nil, err
	}
	return uuid.UUID(id.Bytes), nil
}

func (r *Repository) GetStockTag(ctx context.Context, id uuid.UUID) (string, error) {
	name, err := r.queries.GetStockTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrStockTagNotFound
		}
		return "", err
	}
	return name, nil
}

func (r *Repository) UpdateStockTag(ctx context.Context, id uuid.UUID, name string) error {
	result, err := r.queries.UpdateStockTag(ctx, sqlc.UpdateStockTagParams{
		Name: name,
		ID:   pgtype.UUID{Bytes: id, Valid: true},
	})
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return ErrStockTagNotFound
	}
	return nil
}

func (r *Repository) DeleteStockTag(ctx context.Context, id uuid.UUID) error {
	result, err := r.queries.DeleteStockTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return ErrStockTagNotFound
	}
	return nil
}
