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
	pgId, err := r.queries.CreateStock(ctx, sqlc.CreateStockParams{
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
	return id, nil
}

func (r *Repository) GetStock(ctx context.Context, id uuid.UUID) (model.Stock, error) {
	stock, err := r.queries.GetStock(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Stock{}, model.ErrStockNotFound
		}
		return model.Stock{}, err
	}
	return model.Stock{
		Name:             stock.Name,
		HasAmount:        stock.HasAmount,
		Currency:         stock.Currency,
		CurrencyExponent: stock.CurrencyExponent,
		Description:      stock.Description,
	}, nil
}

func (r *Repository) UpdateStock(ctx context.Context, id uuid.UUID, stock model.Stock) error {
	result, err := r.queries.UpdateStock(ctx, sqlc.UpdateStockParams{
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
