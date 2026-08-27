package model

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Date time.Time

type Transaction struct {
	Description string
	OccurredAt  Date
	Tags        []uuid.UUID
	Flows       []Flow
}

var ErrTransactionNameDuplicate = errors.New("transaction name duplicate")
var ErrTransactionNotFound = errors.New("transaction not found")
