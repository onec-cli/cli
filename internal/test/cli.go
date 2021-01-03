package test

import (
	"bytes"
	"github.com/onec-cli/cli/api"
	"github.com/onec-cli/cli/internal/platform"
	"io"
	"io/ioutil"
	"strings"
)

type FakeCli struct {
	in        io.ReadCloser
	out       io.Writer
	outBuffer *bytes.Buffer
	err       io.Writer
	ib        map[string]api.Infobase
}

// In returns the input stream the cli will use
func (c *FakeCli) In() io.ReadCloser {
	return c.in
}

// Out returns the output stream (stdout) the cli should write on
func (c *FakeCli) Out() io.Writer {
	return c.out
}

// Err returns the output stream (stderr) the cli should write on
func (c *FakeCli) Err() io.Writer {
	return c.err
}

// OutBuffer returns the stdout buffer
func (c *FakeCli) OutBuffer() *bytes.Buffer {
	return c.outBuffer
}

// NewInfobase returns the infobase the cli will use
func (c *FakeCli) Infobase(connPath string, opts ...string) api.Infobase {
	if ib, ok := c.ib["test"]; ok {
		return ib
	}
	c.ib["test"] = platform.NewInfobase(connPath, opts...)
	return c.ib["test"]
}

// SetIn sets the input of the cli to the specified io.ReadCloser
func (c *FakeCli) SetIn(in io.ReadCloser) {
	c.in = in
}

// SetOut sets the stdout stream for the cli to the specified io.Writer
func (c *FakeCli) SetOut(out io.Writer) {
	c.out = out
}

// SetErr sets the stderr stream for the cli to the specified io.Writer
func (c *FakeCli) SetErr(err io.Writer) {
	c.err = err
}

// NewFakeCli returns a fake for the cli.Cli interface
func NewFakeCli(ib api.Infobase, opts ...func(*FakeCli)) *FakeCli {
	outBuffer := new(bytes.Buffer)
	errBuffer := new(bytes.Buffer)
	c := &FakeCli{
		out:       outBuffer,
		err:       errBuffer,
		outBuffer: outBuffer,
		in:        ioutil.NopCloser(strings.NewReader("")),
		ib:        map[string]api.Infobase{"test": ib},
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}
