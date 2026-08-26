package model

import (
	"errors"

	"github.com/google/uuid"
)

type StockSummary struct {
	Id               uuid.UUID
	Name             string
	HasAmount        bool
	Currency         string
	CurrencyExponent int32
}

type Stock struct {
	Name             string
	HasAmount        bool
	Currency         string
	CurrencyExponent int32
	Description      string
	Tags             []uuid.UUID
}

var ErrStockNameDuplicate = errors.New("stock name duplicate")
var ErrStockNotFound = errors.New("stock not found")
