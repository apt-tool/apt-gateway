package controller

import "errors"

var (
	errBadRequest   = errors.New("bad request")
	errValidRequest = errors.New("request is not valid")
	errDatabase     = errors.New("internal db error")

	errUserNotFound = errors.New("user not found")
	errToken        = errors.New("failed to create token")
)
