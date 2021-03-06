// Package balancer defines APIs for load balancing in RSocket.
package balancer

import (
	"io"

	"github.com/rsocket/rsocket-go"
)

// Balancer manage input RSocket clients.
type Balancer interface {
	io.Closer
	// Put puts a new client.
	Put(client rsocket.Client)
	// PutLabel puts a new client with a label.
	PutLabel(label string, client rsocket.Client)
	// Next returns next balanced RSocket client.
	Next() rsocket.Client
	// OnLeave handle events when a client exit.
	OnLeave(fn func(label string))
}
