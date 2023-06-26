package redis

import "errors"

var (
	ErrSaveSet     = errors.New("failed to save the set")
	ErrSetNotFound = errors.New("set not found")
)
