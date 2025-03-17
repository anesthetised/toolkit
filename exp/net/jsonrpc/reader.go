package jsonrpc

import (
	"bufio"
	"context"
	"fmt"
	"io"

	"golang.org/x/exp/jsonrpc2"
)

type reader struct {
	br *bufio.Reader
	d  byte
}

func NewReader(r io.Reader, delimiter byte) jsonrpc2.Reader {
	return &reader{br: bufio.NewReader(r), d: delimiter}
}

func (r *reader) Read(ctx context.Context) (jsonrpc2.Message, int64, error) {
	select {
	case <-ctx.Done():
		return nil, 0, ctx.Err()
	default:
	}

	raw, err := r.br.ReadBytes(r.d)
	if err != nil {
		return nil, 0, err
	}

	msg, err := jsonrpc2.DecodeMessage(raw)
	if err != nil {
		return nil, 0, fmt.Errorf("decode message: %w", err)
	}

	return msg, int64(len(raw) - 1), nil
}
