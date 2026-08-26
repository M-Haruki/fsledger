package transaction

import (
	"errors"

	"github.com/google/uuid"
)

type tag struct {
	id   uuid.UUID
	name string
}

var ErrTransactionTagNameDuplicate = errors.New("transaction tag name duplicate")
var ErrTransactionTagNotFound = errors.New("transaction tag not found")
