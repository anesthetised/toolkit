package http

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestListenAndServe(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	cancel()

	err := ListenAndServe(ctx, &http.Server{Addr: ":0"})
	assert.ErrorIs(t, err, context.Canceled)
}
