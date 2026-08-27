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

func (r *Repository) DeleteTagRelations(ctx context.Context, tagT model.TagType, parent uuid.UUID) error {
	parentPgId := pgtype.UUID{Bytes: parent, Valid: true}
	var err error
	switch tagT {
	case model.StockTag:
		err = r.queries.DeleteStockTagRelation(ctx, parentPgId)
	case model.TransactionTag:
		err = r.queries.DeleteTransactionTagRelation(ctx, parentPgId)
	case model.FlowTag:
		err = r.queries.DeleteFlowTagRelation(ctx, parentPgId)
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) SetTags(ctx context.Context, tagT model.TagType, parent uuid.UUID, tags []uuid.UUID) error {
	parentPgId := pgtype.UUID{Bytes: parent, Valid: true}
	for _, tagId := range tags {
		var err error
		tagPgId := pgtype.UUID{Bytes: tagId, Valid: true}
		switch tagT {
		case model.StockTag:
			err = r.queries.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{
				StockID: parentPgId,
				TagID:   tagPgId,
			})
		case model.TransactionTag:
			err = r.queries.CreateTransactionTagRelation(ctx, sqlc.CreateTransactionTagRelationParams{
				TransactionID: parentPgId,
				TagID:         tagPgId,
			})
		case model.FlowTag:
			err = r.queries.CreateFlowTagRelation(ctx, sqlc.CreateFlowTagRelationParams{
				FlowID: parentPgId,
				TagID:  tagPgId,
			})
		}
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
	return nil
}

func setTagsOld(ctx context.Context, q *sqlc.Queries, tagT model.TagType, parent uuid.UUID, tags []uuid.UUID) error {
	for _, tagId := range tags {
		var err error
		switch tagT {
		case model.StockTag:
			err = q.CreateStockTagRelation(ctx, sqlc.CreateStockTagRelationParams{
				StockID: pgtype.UUID{Bytes: parent, Valid: true},
				TagID:   pgtype.UUID{Bytes: tagId, Valid: true},
			})
		case model.TransactionTag:
			err = q.CreateTransactionTagRelation(ctx, sqlc.CreateTransactionTagRelationParams{
				TransactionID: pgtype.UUID{Bytes: parent, Valid: true},
				TagID:         pgtype.UUID{Bytes: tagId, Valid: true},
			})
		case model.FlowTag:
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

func (r *Repository) ListTags(ctx context.Context, tagT model.TagType) ([]model.Tag, error) {
	var tags []model.Tag
	switch tagT {
	case model.StockTag:
		res, err := r.queries.ListStockTags(ctx)
		if err != nil {
			return nil, err
		}
		tags = make([]model.Tag, len(res))
		for i, tag := range res {
			tags[i] = model.Tag{
				Id:   uuid.UUID(tag.ID.Bytes),
				Name: tag.Name,
			}
		}
	case model.TransactionTag:
		res, err := r.queries.ListTransactionTags(ctx)
		if err != nil {
			return nil, err
		}
		tags = make([]model.Tag, len(res))
		for i, tag := range res {
			tags[i] = model.Tag{
				Id:   uuid.UUID(tag.ID.Bytes),
				Name: tag.Name,
			}
		}
	case model.FlowTag:
		res, err := r.queries.ListFlowTags(ctx)
		if err != nil {
			return nil, err
		}
		tags = make([]model.Tag, len(res))
		for i, tag := range res {
			tags[i] = model.Tag{
				Id:   uuid.UUID(tag.ID.Bytes),
				Name: tag.Name,
			}
		}
	}
	return tags, nil
}

func (r *Repository) CreateTag(ctx context.Context, tagT model.TagType, name string) (uuid.UUID, error) {
	var id pgtype.UUID
	var err error
	switch tagT {
	case model.StockTag:
		id, err = r.queries.CreateStockTag(ctx, name)
	case model.TransactionTag:
		id, err = r.queries.CreateTransactionTag(ctx, name)
	case model.FlowTag:
		id, err = r.queries.CreateFlowTag(ctx, name)
	}
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

func (r *Repository) GetTag(ctx context.Context, tagT model.TagType, id uuid.UUID) (string, error) {
	var name string
	var err error
	switch tagT {
	case model.StockTag:
		name, err = r.queries.GetStockTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	case model.TransactionTag:
		name, err = r.queries.GetTransactionTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	case model.FlowTag:
		name, err = r.queries.GetFlowTag(ctx, pgtype.UUID{Bytes: id, Valid: true})
	}
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", model.ErrTagNotFound
		}
		return "", err
	}
	return name, nil
}

func (r *Repository) UpdateTag(ctx context.Context, tagT model.TagType, id uuid.UUID, name string) error {
	pgId := pgtype.UUID{Bytes: id, Valid: true}
	var result pgconn.CommandTag
	var err error
	switch tagT {
	case model.StockTag:
		result, err = r.queries.UpdateStockTag(ctx, sqlc.UpdateStockTagParams{
			Name: name,
			ID:   pgId,
		})
	case model.TransactionTag:
		result, err = r.queries.UpdateTransactionTag(ctx, sqlc.UpdateTransactionTagParams{
			Name: name,
			ID:   pgId,
		})
	case model.FlowTag:
		result, err = r.queries.UpdateFlowTag(ctx, sqlc.UpdateFlowTagParams{
			Name: name,
			ID:   pgId,
		})
	}
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return model.ErrTagNotFound
	}
	return nil
}

func (r *Repository) DeleteTag(ctx context.Context, tagT model.TagType, id uuid.UUID) error {
	pgId := pgtype.UUID{Bytes: id, Valid: true}
	var result pgconn.CommandTag
	var err error
	switch tagT {
	case model.StockTag:
		result, err = r.queries.DeleteStockTag(ctx, pgId)
	case model.TransactionTag:
		result, err = r.queries.DeleteTransactionTag(ctx, pgId)
	case model.FlowTag:
		result, err = r.queries.DeleteFlowTag(ctx, pgId)
	}
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		// not found
		return model.ErrTagNotFound
	}
	return nil
}

func (r *Repository) ListTagsByParentID(ctx context.Context, tagT model.TagType, id uuid.UUID) ([]uuid.UUID, error) {
	pgId := pgtype.UUID{Bytes: id, Valid: true}
	var tags []pgtype.UUID
	var err error
	switch tagT {
	case model.StockTag:
		tags, err = r.queries.ListTagIDsByStock(ctx, pgId)
	case model.TransactionTag:
		tags, err = r.queries.ListTagIDsByTransaction(ctx, pgId)
	case model.FlowTag:
		tags, err = r.queries.ListTagIDsByFlow(ctx, pgId)
	}
	if err != nil {
		return nil, err
	}
	formattedTags := make([]uuid.UUID, len(tags))
	for i, tag := range tags {
		formattedTags[i] = uuid.UUID(tag.Bytes)
	}
	return formattedTags, nil
}
