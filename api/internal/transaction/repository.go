package transaction

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

func (r *Repository) ListTransactionTags(ctx context.Context) ([]tag, error) {
	res, err := r.queries.ListTransactionTags(ctx)
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

func (r *Repository) CreateTransactionTag(ctx context.Context, name string) (uuid.UUID, error) {
	id, err := r.queries.CreateTransactionTag(ctx, name)
	if err != nil {
		if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
			if pgErr.Code == "23505" {
				// duplicate key
				return uuid.Nil, ErrTransactionTagNameDuplicate
			}
		}
		return uuid.Nil, err
	}
	return uuid.UUID(id.Bytes), nil
}

func (r *Repository) GetTransactionTag(ctx context.Context, id uuid.UUID) (string, error) {
	name, err := r.queries.GetTransactionTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", ErrTransactionTagNotFound
		}
		return "", err
	}
	return name, nil
}

func (r *Repository) UpdateTransactionTag(ctx context.Context, id uuid.UUID, name string) error {
	result, err := r.queries.UpdateTransactionTag(ctx, sqlc.UpdateTransactionTagParams{
		Name: name,
		ID:   pgtype.UUID{Bytes: id, Valid: true},
	})
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return ErrTransactionTagNotFound
	}
	return nil
}

func (r *Repository) DeleteTransactionTag(ctx context.Context, id uuid.UUID) error {
	result, err := r.queries.DeleteTransactionTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return ErrTransactionTagNotFound
	}
	return nil
}
