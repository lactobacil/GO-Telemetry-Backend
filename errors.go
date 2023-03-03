package errors

import "errors"

var ErrInternalServer = errors.New("internal server error")
var ErrInvalidUser = errors.New("invalid user")
var ErrInvalidPassword = errors.New("invalid password")
