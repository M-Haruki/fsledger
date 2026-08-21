package stock

import (
	"errors"

	"github.com/google/uuid"
)

type stockSummary struct {
	id   uuid.UUID
	name string
}

type stockRequest struct {
	name        string
	hasAmount   bool
	currency    string
	description string
	tags        []uuid.UUID
}

type stockResponse struct {
	name        string
	hasAmount   bool
	currency    string
	description string
	tags        []tag
}

type tag struct {
	id   uuid.UUID
	name string
}

var ErrStockNameDuplicate = errors.New("stock name duplicate")
var ErrStockNotFound = errors.New("stock not found")
var ErrStockTagNameDuplicate = errors.New("stock tag name duplicate")
var ErrStockTagNotFound = errors.New("stock tag not found")
