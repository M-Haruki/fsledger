package model

import (
	"github.com/google/uuid"
)

type Flow struct {
	From       uuid.UUID
	To         uuid.UUID
	FromAmount int64
	ToAmount   int64
	Tags       []uuid.UUID
}
