package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

// ListenAndServe allows graceful shutdown via context.Context cancellation.
func ListenAndServe(ctx context.Context, s *http.Server) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.ListenAndServe()
		close(errCh)
	}()

	select {
	case <-ctx.Done():
		if err := s.Shutdown(context.Background()); err != nil {
			return fmt.Errorf("shutdown gracefully: %w", err)
		}
		return ctx.Err()
	case err := <-errCh:
		return fmt.Errorf("listen and serve: %w", err)
	}
}
