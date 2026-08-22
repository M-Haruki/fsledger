package stock

import (
	"errors"

	"github.com/google/uuid"
)

type stockSummary struct {
	id               uuid.UUID
	name             string
	hasAmount        bool
	currency         string
	currencyExponent int32
}

type stock struct {
	name             string
	hasAmount        bool
	currency         string
	currencyExponent int32
	description      string
	tags             []uuid.UUID
}

type tag struct {
	id   uuid.UUID
	name string
}

var ErrStockNameDuplicate = errors.New("stock name duplicate")
var ErrStockNotFound = errors.New("stock not found")
var ErrStockTagNameDuplicate = errors.New("stock tag name duplicate")
var ErrStockTagNotFound = errors.New("stock tag not found")
