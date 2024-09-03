package os

import (
	"context"
	"time"
)

func Sleep(ctx context.Context, delay time.Duration) error {
	t := time.NewTimer(delay)
	defer t.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-t.C:
		return nil
	}
}
