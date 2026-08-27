package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) CreateTransaction(ctx context.Context, transaction model.Transaction) (uuid.UUID, error) {
	transactionPgId, err := r.queries.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		Description: transaction.Description,
		Occurredat:  pgtype.Date{Time: time.Time(transaction.OccurredAt), InfinityModifier: pgtype.Finite, Valid: true},
	})
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.UUID(transactionPgId.Bytes), nil
}

func (r *Repository) GetTransaction(ctx context.Context, id uuid.UUID) (model.Transaction, error) {
	res, err := r.queries.GetTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return model.Transaction{}, model.ErrTransactionNotFound
		}
		return model.Transaction{}, err
	}
	transaction := model.Transaction{
		Description: res.Description,
		OccurredAt:  model.Date(res.OccurredAt.Time),
	}
	return transaction, nil
}

func (r *Repository) UpdateTransaction(ctx context.Context, id uuid.UUID, transaction model.Transaction) error {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	result, err := q.UpdateTransaction(ctx, sqlc.UpdateTransactionParams{
		Description: transaction.Description,
		Occurredat:  pgtype.Date{Time: time.Time(transaction.OccurredAt), InfinityModifier: pgtype.Finite, Valid: true},
		ID:          pgtype.UUID{Bytes: id, Valid: true},
	})
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return model.ErrTransactionNameDuplicate
			}
		}
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return model.ErrTransactionNotFound
	}
	return nil
}

func (r *Repository) DeleteTransaction(ctx context.Context, id uuid.UUID) error {
	result, err := r.queries.DeleteTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return model.ErrTransactionNotFound
	}
	return nil
}
