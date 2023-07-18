package project

import "errors"

var (
	ErrProjectNotFound    = errors.New("project not found")
	ErrFailedToGetProject = errors.New("failed to get project")
)
