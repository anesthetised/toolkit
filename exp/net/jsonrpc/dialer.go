package jsonrpc

import (
	"context"
	"crypto/tls"
	"errors"
	"io"
	"net"

	"golang.org/x/exp/jsonrpc2"
)

type DialerTLS struct {
	Dialer    jsonrpc2.Dialer
	TLSConfig *tls.Config
}

func NewDialerTLS(dialer jsonrpc2.Dialer, tlsConfig *tls.Config) jsonrpc2.Dialer {
	return &DialerTLS{
		Dialer:    dialer,
		TLSConfig: tlsConfig,
	}
}

func (d *DialerTLS) Dial(ctx context.Context) (io.ReadWriteCloser, error) {
	rwc, err := d.Dialer.Dial(ctx)
	if err != nil {
		return nil, err
	}

	conn, ok := rwc.(net.Conn)
	if !ok {
		return nil, errors.New("cast io.ReadWriteCloser to net.Conn")
	}

	tlsConn := tls.Client(conn, d.TLSConfig)
	if err = tlsConn.HandshakeContext(ctx); err != nil {
		_ = conn.Close()
		return nil, err
	}

	return tlsConn, nil
}

func DialTLS(ctx context.Context,
	dialer jsonrpc2.Dialer,
	binder jsonrpc2.Binder,
	tlsConfig *tls.Config,
) (*jsonrpc2.Connection, error) {
	d := NewDialerTLS(dialer, tlsConfig)
	return jsonrpc2.Dial(ctx, d, binder)
}

func Dial(ctx context.Context, network, address string, binder jsonrpc2.Binder) (*jsonrpc2.Connection, error) {
	return jsonrpc2.Dial(ctx, jsonrpc2.NetDialer(network, address, net.Dialer{}), binder)
}
