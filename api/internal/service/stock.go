package service

import (
	"context"

	"github.com/M-Haruki/fsledger/api/internal/model"
	"github.com/google/uuid"
)

func (s *Service) ListStocks(ctx context.Context) ([]model.StockAbstract, error) {
	return s.repo.ListStocks(ctx)
}

func (s *Service) CreateStock(ctx context.Context, stock model.Stock) (uuid.UUID, error) { // error: ErrStockNameDuplicate, ErrTagNotFound
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	defer tx.Rollback(ctx)
	repo := s.repo.WithTx(tx)
	//
	id, err := repo.CreateStock(ctx, stock)
	if err != nil {
		return uuid.Nil, err
	}
	err = repo.SetTags(ctx, model.StockTag, id, stock.Tags)
	if err != nil {
		return uuid.Nil, err
	}
	//
	err = tx.Commit(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	//
	return id, nil
}

func (s *Service) GetStock(ctx context.Context, id uuid.UUID) (model.Stock, error) { // error: ErrStockNotFound
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return model.Stock{}, err
	}
	defer tx.Rollback(ctx)
	repo := s.repo.WithTx(tx)
	//
	stock, err := repo.GetStock(ctx, id)
	if err != nil {
		return model.Stock{}, err
	}
	tags, err := repo.ListTagsByParentID(ctx, model.StockTag, id)
	if err != nil {
		return model.Stock{}, err
	}
	//
	err = tx.Commit(ctx)
	if err != nil {
		return model.Stock{}, err
	}
	//
	stock.Tags = tags
	return stock, nil
}

func (s *Service) UpdateStock(ctx context.Context, id uuid.UUID, stock model.Stock) error { // error: ErrStockNotFound, ErrStockNameDuplicate, ErrTagNotFound
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	repo := s.repo.WithTx(tx)
	//
	err = repo.UpdateStock(ctx, id, stock)
	if err != nil {
		return err
	}
	err = repo.DeleteTagRelations(ctx, model.StockTag, id)
	if err != nil {
		return err
	}
	err = repo.SetTags(ctx, model.StockTag, id, stock.Tags)
	if err != nil {
		return err
	}
	//
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	//
	return nil
}

func (s *Service) DeleteStock(ctx context.Context, id uuid.UUID) error { // error: ErrStockNotFound
	return s.repo.DeleteStock(ctx, id)
}
