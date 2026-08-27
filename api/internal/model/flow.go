package model

import (
	"github.com/google/uuid"
)

type Flow struct {
	From   uuid.UUID
	To     uuid.UUID
	Amount int64
	Tags   []uuid.UUID
}
