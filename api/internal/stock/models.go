package stock

import (
	"github.com/google/uuid"
)

type stockOverview struct {
	id   uuid.UUID
	name string
}

type stock struct {
	name        string
	has_amount  bool
	currency    string
	description string
	tags        []uuid.UUID
}
