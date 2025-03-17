package jsonrpc

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/jsonrpc2"
)

func TestFramer(t *testing.T) {
	buf := &bytes.Buffer{}
	f := NewDelimitedFramer(Delimiter)

	ctx := context.Background()
	data := []byte(`{"jsonrpc":"2.0","id":1,"method":"test","params":["test123"]}`)

	t.Run("writer", func(t *testing.T) {
		w := f.Writer(buf)

		req, err := jsonrpc2.NewCall(jsonrpc2.Int64ID(1), "test", []any{"test123"})
		assert.NoError(t, err)

		_, err = w.Write(ctx, req)
		assert.NoError(t, err)
		assert.Equal(t, append(data, Delimiter), buf.Bytes())
	})

	t.Run("reader", func(t *testing.T) {
		r := f.Reader(buf)

		msg, n, err := r.Read(ctx)
		assert.NoError(t, err)
		assert.Equal(t, n, int64(len(data)))

		raw, err := jsonrpc2.EncodeMessage(msg)
		assert.NoError(t, err)
		assert.Equal(t, raw, data)
	})
}
