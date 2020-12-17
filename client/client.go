package client

import (
	"io"
	"os"
)

// Client represents the command line client.
type Client interface {
	In() io.ReadCloser
	Out() io.Writer
	Err() io.Writer
}

// client is an instance the command line client.
// Instances of the client can be returned from NewClient.
type client struct {
	in  io.ReadCloser
	out io.Writer
	err io.Writer
}

// NewClient returns a client instance with all operators applied on it.
// It applies by default the standard streams.
func NewClient() *client {
	c := &client{
		in:  os.Stdin,
		out: os.Stdout,
		err: os.Stderr,
	}
	return c
}

// In returns the reader used for stdin
func (c *client) In() io.ReadCloser {
	return c.in
}

// Out returns the writer used for stdout
func (c *client) Out() io.Writer {
	return c.out
}

// Err returns the writer used for stderr
func (c *client) Err() io.Writer {
	return c.err
}
