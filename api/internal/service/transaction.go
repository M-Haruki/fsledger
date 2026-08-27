package service

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
)

func (s *Service) CreateTransaction(ctx context.Context, transaction model.Transaction) (uuid.UUID, error) { // error: ErrTagNotFound
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)
	repo := s.repo.WithTx(tx)
	//
	id, err := repo.CreateTransaction(ctx, transaction)
	if err != nil {
		return uuid.Nil, err
	}
	err = repo.SetTags(ctx, model.TransactionTag, id, transaction.Tags)
	if err != nil {
		return uuid.Nil, err
	}
	for _, flow := range transaction.Flows {
		flowId, err := repo.CreateFlow(ctx, id, flow)
		if err != nil {
			return uuid.Nil, err
		}
		err = repo.SetTags(ctx, model.FlowTag, flowId, flow.Tags)
		if err != nil {
			return uuid.Nil, err
		}
	}
	//
	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	//
	return id, nil
}

func (s *Service) GetTransaction(ctx context.Context, id uuid.UUID) (model.Transaction, error) { // error: ErrTransactionNotFound
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return model.Transaction{}, err
	}
	defer tx.Rollback(ctx)
	repo := s.repo.WithTx(tx)
	//
	transaction, err := repo.GetTransaction(ctx, id)
	if err != nil {
		return model.Transaction{}, err
	}
	transactionTags, err := repo.ListTagsByParentID(ctx, model.TransactionTag, id)
	if err != nil {
		return model.Transaction{}, err
	}
	flows, flowIds, err := repo.ListFlowsByTransaction(ctx, id)
	if err != nil {
		return model.Transaction{}, err
	}
	transaction.Tags = transactionTags
	transaction.Flows = flows
	for i := range flows {
		flowTags, err := repo.ListTagsByParentID(ctx, model.FlowTag, flowIds[i])
		if err != nil {
			return model.Transaction{}, err
		}
		transaction.Flows[i].Tags = flowTags
	}
	//
	err = tx.Commit(ctx)
	if err != nil {
		return model.Transaction{}, err
	}
	//
	return transaction, nil
}

func (s *Service) UpdateTransaction(ctx context.Context, id uuid.UUID, transaction model.Transaction) error { // error: ErrTransactionNotFound, ErrTagNotFound
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	repo := s.repo.WithTx(tx)
	//
	err = repo.UpdateTransaction(ctx, id, transaction)
	if err != nil {
		return err
	}
	err = repo.DeleteTagRelations(ctx, model.TransactionTag, id)
	if err != nil {
		return err
	}
	err = repo.SetTags(ctx, model.TransactionTag, id, transaction.Tags)
	if err != nil {
		return err
	}
	err = repo.DeleteFlowsByTransaction(ctx, id)
	if err != nil {
		return err
	}
	for _, flow := range transaction.Flows {
		flowId, err := repo.CreateFlow(ctx, id, flow)
		if err != nil {
			return err
		}
		err = repo.SetTags(ctx, model.FlowTag, flowId, flow.Tags)
		if err != nil {
			return err
		}
	}
	//
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	//
	return nil
}

func (s *Service) DeleteTransaction(ctx context.Context, id uuid.UUID) error { // error: ErrTransactionNotFound
	return s.repo.DeleteTransaction(ctx, id)
}
