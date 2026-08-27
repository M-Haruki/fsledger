package repository

import (
	"context"
	"errors"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type tagType int

const (
	stockTag tagType = iota
	transactionTag
	flowTag
)

func setTags(ctx context.Context, q *sqlc.Queries, tagT tagType, parent uuid.UUID, tags []uuid.UUID) error {
	for _, tagId := range tags {
		var err error
		switch tagT {
		case stockTag:
			err = q.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{
				StockID: pgtype.UUID{Bytes: parent, Valid: true},
				TagID:   pgtype.UUID{Bytes: tagId, Valid: true},
			})
		case transactionTag:
			err = q.CreateTransactionTagRelation(ctx, sqlc.CreateTransactionTagRelationParams{
				TransactionID: pgtype.UUID{Bytes: parent, Valid: true},
				TagID:         pgtype.UUID{Bytes: tagId, Valid: true},
			})
		case flowTag:
			err = q.CreateFlowTagRelation(ctx, sqlc.CreateFlowTagRelationParams{
				FlowID: pgtype.UUID{Bytes: parent, Valid: true},
				TagID:  pgtype.UUID{Bytes: tagId, Valid: true},
			})
		}
		if err != nil {
			println(err.Error())
			if pgErr, ok := errors.AsType[*pgconn.PgError](err); ok {
				if pgErr.Code == "23503" {
					// tag not found
					return model.ErrTagNotFound
				}
			}
			return err
		}
	}
	return nil
}
