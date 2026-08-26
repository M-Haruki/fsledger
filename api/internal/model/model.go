package model

import (
	"errors"

	"github.com/google/uuid"
)

type Tag struct {
	Id   uuid.UUID
	Name string
}

var ErrTagNameDuplicate = errors.New("tag name duplicate")
var ErrTagNotFound = errors.New("tag not found")
