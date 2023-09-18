package api

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrTokenInactive      = errors.New("token is inactive")
	ErrNotAuthenticated   = errors.New("not authenticated")
)
