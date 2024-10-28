package jsonrpc

import (
	"io"

	"golang.org/x/exp/jsonrpc2"
)

var Delimiter = byte('\n')

type delimitedFramer struct {
	delimeter byte
}

func NewDelimitedFramer(delimeter byte) jsonrpc2.Framer {
	return &delimitedFramer{delimeter: delimeter}
}

func (f *delimitedFramer) Reader(r io.Reader) jsonrpc2.Reader {
	return NewReader(r, f.delimeter)
}

func (f *delimitedFramer) Writer(w io.Writer) jsonrpc2.Writer {
	return NewWriter(w, f.delimeter)
}
