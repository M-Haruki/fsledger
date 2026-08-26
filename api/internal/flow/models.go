package flow

import (
	"errors"

	"github.com/google/uuid"
)

type tag struct {
	id   uuid.UUID
	name string
}

var ErrFlowTagNameDuplicate = errors.New("flow tag name duplicate")
var ErrFlowTagNotFound = errors.New("flow tag not found")
