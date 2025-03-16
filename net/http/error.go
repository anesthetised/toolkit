package http

import (
	"errors"
	"fmt"
)

// UnexpectedStatusCode represents status code as an error for convenience.
type UnexpectedStatusCode int

// Error returns a formatted error message.
func (code UnexpectedStatusCode) Error() string {
	return fmt.Sprintf("unexpected response status code: %d", code)
}

// StatusCodeIs matches an error to a status code.
func StatusCodeIs(err error, code int) bool {
	var errStatusCode UnexpectedStatusCode
	return errors.As(err, &errStatusCode) && int(errStatusCode) == code
}
