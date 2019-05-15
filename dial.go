package hive

import (
	"context"
	"crypto/tls"
	"net"
	"time"
)

var (
	_ Dialer        = DialWrapper{}
	_ TimeoutDialer = DialWrapper{}
)

// Dialer is the dialer interface. It can be used to obtain more control over
// how Hive creates network connections.
type Dialer interface {
	Dial(network, address string) (net.Conn, error)
}

type TimeoutDialer interface {
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type DialWrapper struct {
	net.Dialer
}

func (d DialWrapper) Dial(network, address string) (net.Conn, error) {
	return d.Dialer.Dial(network, address)
}

func (d DialWrapper) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return d.DialContext(ctx, network, address)
}

func (d DialWrapper) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return d.Dialer.DialContext(ctx, network, address)
}

type TLSDialer struct {
	*tls.Config
}

func (d TLSDialer) Dial(network, address string) (net.Conn, error) {
	return tls.Dial(network, address, d.Config)
}
