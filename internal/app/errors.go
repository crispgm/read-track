package app

import "errors"

// Error codes
const (
	ErrCodeParams = iota + 100
	ErrCodeDBFailed
	ErrMarshalFailed
)

// Definition of errors
var (
	ErrInvalidExportType  = NewError(1001, errors.New("Invalid export type"))
	ErrInvalidAccessToken = NewError(1002, errors.New("Invalid access token"))
)

// Error .
type Error struct {
	code int
	err  error
}

// NewError .
func NewError(code int, err error) *Error {
	return &Error{
		code: code,
		err:  err,
	}
}

// Response returns error response
func (ae Error) Response() Response {
	return Response{
		Code:    ae.code,
		Message: ae.err.Error(),
	}
}
