package di

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDI(t *testing.T) {
	testValue := &http.Server{}
	ctx := PutNamed(context.Background(), "", testValue)

	val, ok := GetNamed[*http.Server](ctx, "")
	assert.True(t, ok)
	assert.Equal(t, testValue, val)

	_, ok = GetNamed[*http.Server](ctx, "test")
	assert.False(t, ok)
}
