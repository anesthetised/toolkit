package jsonrpc

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"golang.org/x/exp/jsonrpc2"
)

type writer struct {
	bw *bufio.Writer
	d  byte
}

func NewWriter(w io.Writer, delimiter byte) jsonrpc2.Writer {
	return &writer{bw: bufio.NewWriter(w), d: delimiter}
}

func (w *writer) Write(ctx context.Context, msg jsonrpc2.Message) (int64, error) {
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	default:
	}

	data, err := jsonrpc2.EncodeMessage(msg)
	if err != nil {
		return 0, fmt.Errorf("encode message: %w", err)
	}

	n, err := w.bw.Write(data)
	if err != nil {
		return 0, fmt.Errorf("write message: %w", err)
	}
	if err = w.bw.WriteByte(w.d); err != nil {
		return 0, fmt.Errorf("write delimiter: %w", err)
	}

	return int64(n), w.bw.Flush()
}
