package test

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"
)

type FakeClient struct {
	in  io.ReadCloser
	out io.Writer
	err io.Writer
}

// In returns the input stream the client will use
func (c *FakeClient) In() io.ReadCloser {
	return c.in
}

// Out returns the output stream (stdout) the client should write on
func (c *FakeClient) Out() io.Writer {
	return c.out
}

// Err returns the output stream (stderr) the client should write on
func (c *FakeClient) Err() io.Writer {
	return c.err
}

// SetIn sets the input of the client to the specified io.ReadCloser
func (c *FakeClient) SetIn(in io.ReadCloser) {
	c.in = in
}

// SetOut sets the stdout stream for the client to the specified io.Writer
func (c *FakeClient) SetOut(out io.Writer) {
	c.out = out
}

// SetErr sets the stderr stream for the client to the specified io.Writer
func (c *FakeClient) SetErr(err io.Writer) {
	c.err = err
}

// NewFakeCli returns a fake for the client.Cli interface
func NewFakeCli(opts ...func(*FakeClient)) *FakeClient {
	outBuffer := new(bytes.Buffer)
	errBuffer := new(bytes.Buffer)
	c := &FakeClient{
		out: outBuffer,
		err: errBuffer,
		in:  ioutil.NopCloser(strings.NewReader("")),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
