package repository

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/db/sqlc"
	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (r *Repository) CreateFlow(ctx context.Context, transactionId uuid.UUID, flow model.Flow) (uuid.UUID, error) {
	flowId, err := r.queries.CreateFlow(ctx, sqlc.CreateFlowParams{
		Transactionid: pgtype.UUID{Bytes: transactionId, Valid: true},
		Fromstockid:   pgtype.UUID{Bytes: flow.From, Valid: true},
		Tostockid:     pgtype.UUID{Bytes: flow.To, Valid: true},
		Amount:        flow.Amount,
	})
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.UUID(flowId.Bytes), nil
}

func (r *Repository) ListFlowsByTransaction(ctx context.Context, id uuid.UUID) ([]model.Flow, []uuid.UUID, error) {
	flows, err := r.queries.ListFlowByTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return nil, nil, err
	}
	formattedFlows := make([]model.Flow, len(flows))
	flowIds := make([]uuid.UUID, len(flows))
	for i, flow := range flows {
		formattedFlows[i] = model.Flow{
			From:   uuid.UUID(flow.FromStockID.Bytes),
			To:     uuid.UUID(flow.ToStockID.Bytes),
			Amount: flow.Amount,
		}
		flowIds[i] = uuid.UUID(flow.ID.Bytes)
	}
	return formattedFlows, flowIds, nil
}

func (r *Repository) DeleteFlowsByTransaction(ctx context.Context, id uuid.UUID) error {
	err := r.queries.DeleteFlowByTransaction(ctx, pgtype.UUID{Bytes: id, Valid: true})
	if err != nil {
		return err
	}
	return nil
}
