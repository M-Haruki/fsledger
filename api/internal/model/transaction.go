package model

import "errors"

var ErrTransactionTagNameDuplicate = errors.New("transaction tag name duplicate")
var ErrTransactionTagNotFound = errors.New("transaction tag not found")
