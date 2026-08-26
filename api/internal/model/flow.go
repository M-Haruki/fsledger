package model

import "errors"

var ErrFlowTagNameDuplicate = errors.New("flow tag name duplicate")
var ErrFlowTagNotFound = errors.New("flow tag not found")
