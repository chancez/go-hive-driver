package hive

import (
	"context"
	"net"
	"time"
)

var (
	_ Dialer = DefaultDialer{}
)

// Dialer is the dialer interface. It can be used to obtain more control over
// how Hive creates network connections.
type Dialer interface {
	Dial(network, address string) (net.Conn, error)
	DialTimeout(network, address string, timeout time.Duration) (net.Conn, error)
}

type DefaultDialer struct {
	d net.Dialer
}

func (d DefaultDialer) Dial(network, address string) (net.Conn, error) {
	return d.d.Dial(network, address)
}

func (d DefaultDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return d.DialContext(ctx, network, address)
}

func (d DefaultDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	return d.d.DialContext(ctx, network, address)
}
