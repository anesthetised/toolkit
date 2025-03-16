package http

import (
	"context"
	"fmt"
	"net/http"
)

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
