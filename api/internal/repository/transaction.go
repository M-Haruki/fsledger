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
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	// transaction
	transactionPgId, err := q.CreateTransaction(ctx, sqlc.CreateTransactionParams{
		Description: transaction.Description,
		Occurredat:  pgtype.Date{Time: time.Time(transaction.OccurredAt), InfinityModifier: pgtype.Finite, Valid: true},
	})
	if err != nil {
		return uuid.Nil, err
	}
	transactionId := uuid.UUID(transactionPgId.Bytes)
	// transaction tag
	err = setTags(ctx, q, transactionTag, transactionId, transaction.Tags)
	if err != nil {
		return uuid.Nil, err
	}
	// flows
	for _, flow := range transaction.Flows {
		flowId, err := q.CreateFlow(ctx, sqlc.CreateFlowParams{
			Transactionid: transactionPgId,
			Fromstockid:   pgtype.UUID{Bytes: flow.From, Valid: true},
			Tostockid:     pgtype.UUID{Bytes: flow.To, Valid: true},
			Amount:        flow.Amount,
		})
		if err != nil {
			return uuid.Nil, err
		}
		// flow tag
		err = setTags(ctx, q, flowTag, uuid.UUID(flowId.Bytes), flow.Tags)
		if err != nil {
			return uuid.Nil, err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}

	return transactionId, nil
}

func (r *Repository) GetTransaction(ctx context.Context, id uuid.UUID) (model.Transaction, error) {
	tx, err := r.db.Begin(ctx)
	if err != nil {
		return model.Transaction{}, err
	}
	defer tx.Rollback(ctx)
	q := sqlc.New(tx)

	res, err := q.GetTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
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

	tags, err := q.ListTagIDsByTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return model.Transaction{}, err
	}
	transaction.Tags = make([]uuid.UUID, len(tags))
	for i, tag := range tags {
		transaction.Tags[i] = uuid.UUID(tag.Bytes)
	}

	flows, err := q.ListFlowByTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return model.Transaction{}, nil
	}
	transaction.Flows = make([]model.Flow, len(flows))
	for i, flow := range flows {
		flowTags, err := q.ListTagIDsByFlow(ctx, pgtype.UUID{Bytes: flow.ID.Bytes, Valid: true})
		if err != nil {
			return model.Transaction{}, nil
		}
		transaction.Flows[i] = model.Flow{
			From:   uuid.UUID(flow.FromStockID.Bytes),
			To:     uuid.UUID(flow.ToStockID.Bytes),
			Amount: flow.Amount,
			Tags:   make([]uuid.UUID, len(flowTags)),
		}
		for j, tag := range flowTags {
			transaction.Flows[i].Tags[j] = uuid.UUID(tag.Bytes)
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return model.Transaction{}, err
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

	// transaction tag
	err = q.DeleteTransactionTagRelation(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	err = setTags(ctx, q, transactionTag, id, transaction.Tags)
	if err != nil {
		return err
	}
	// flows
	err = q.DeleteFlowByTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	for _, flow := range transaction.Flows {
		flowId, err := q.CreateFlow(ctx, sqlc.CreateFlowParams{
			Transactionid: pgtype.UUID{Bytes: id, Valid: true},
			Fromstockid:   pgtype.UUID{Bytes: flow.From, Valid: true},
			Tostockid:     pgtype.UUID{Bytes: flow.To, Valid: true},
			Amount:        flow.Amount,
		})
		if err != nil {
			return err
		}
		// flow tag
		err = setTags(ctx, q, flowTag, uuid.UUID(flowId.Bytes), flow.Tags)
		if err != nil {
			return err
		}
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
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

func (r *Repository) ListTransactionTags(ctx context.Context) ([]model.Tag, error) {
	res, err := r.queries.ListTransactionTags(ctx)
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

func (r *Repository) CreateTransactionTag(ctx context.Context, name string) (uuid.UUID, error) {
	id, err := r.queries.CreateTransactionTag(ctx, name)
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

func (r *Repository) GetTransactionTag(ctx context.Context, id uuid.UUID) (string, error) {
	name, err := r.queries.GetTransactionTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", model.ErrTagNotFound
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
		return model.ErrTagNotFound
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
		return model.ErrTagNotFound
	}
	return nil
}
