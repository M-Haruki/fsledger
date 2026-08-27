package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) ListStocks(ctx context.Context) ([]model.StockAbstract, error) {
	raw, err := r.queries.ListStocks(ctx)
	result := make([]model.StockAbstract, len(raw))
	for i, data := range raw {
		result[i] = model.StockAbstract{
			Id:               uuid.UUID(data.ID.Bytes),
			Name:             data.Name,
			HasAmount:        data.HasAmount,
			Currency:         data.Currency,
			CurrencyExponent: data.CurrencyExponent,
		}
	}
	return result, err
}

func (r *Repository) CreateStock(ctx context.Context, stock model.Stock) (uuid.UUID, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	pgId, err := q.CreateStock(ctx, sqlc.CreateStockParams{
		Name:             stock.Name,
		Hasamount:        stock.HasAmount,
		Currency:         stock.Currency,
		CurrencyExponent: stock.CurrencyExponent,
		Description:      stock.Description,
	})
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return uuid.Nil, model.ErrStockNameDuplicate
			}
		}
		return uuid.Nil, err
	}
	id := uuid.UUID(pgId.Bytes)

	err = setTags(ctx, q, stockTag, id, stock.Tags)
	if err != nil {
		return uuid.Nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *Repository) GetStock(ctx context.Context, id uuid.UUID) (model.Stock, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return model.Stock{}, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	astock, err := q.GetStock(ctx, pgtype.UUID{Bytes: id, Valid: true})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Stock{}, model.ErrStockNotFound
		}
		return model.Stock{}, err
	}
	tags, err := q.ListTagIDsByStock(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return model.Stock{}, err
	}
	err = tx.Commit(ctx)
	if err != nil {
		return model.Stock{}, err
	}

	res := model.Stock{
		Name:             astock.Name,
		HasAmount:        astock.HasAmount,
		Currency:         astock.Currency,
		CurrencyExponent: astock.CurrencyExponent,
		Description:      astock.Description,
		Tags:             make([]uuid.UUID, len(tags)),
	}
	for i, atag := range tags {
		res.Tags[i] = uuid.UUID(atag.Bytes)
	}

	return res, nil
}

func (r *Repository) UpdateStock(ctx context.Context, id uuid.UUID, stock model.Stock) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	result, err := q.UpdateStock(ctx, sqlc.UpdateStockParams{
		ID:               pgtype.UUID{Bytes: id, Valid: true},
		Name:             stock.Name,
		Hasamount:        stock.HasAmount,
		Currency:         stock.Currency,
		CurrencyExponent: stock.CurrencyExponent,
		Description:      stock.Description,
	})
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return model.ErrStockNameDuplicate
			}
		}
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return model.ErrStockNotFound
	}
	err = q.DeleteStockTagRelation(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	for _, tagId := range stock.Tags {
		err := q.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{StockID: pgtype.UUID{Bytes: id, Valid: true}, TagID: pgtype.UUID{Bytes: tagId, Valid: true}})
		if err != nil {
			if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
				if pgErr.Code == "23503" {
					// tag not found
					return model.ErrTagNotFound
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
		return model.ErrStockNotFound
	}
	return nil
}

func (r *Repository) ListStockTags(ctx context.Context) ([]model.Tag, error) {
	res, err := r.queries.ListStockTags(ctx)
	if err != nil {
		return nil, err
	}
	result := make([]model.Tag, len(res))
	for i, atag := range res {
		result[i] = model.Tag{
			Id:   uuid.UUID(atag.ID.Bytes),
			Name: atag.Name,
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
				return uuid.Nil, model.ErrTagNameDuplicate
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
			return "", model.ErrTagNotFound
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
		return model.ErrTagNotFound
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
		return model.ErrTagNotFound
	}
	return nil
}
