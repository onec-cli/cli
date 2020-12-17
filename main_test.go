package main_test

import (
	"bytes"
	"github.com/onec-cli/cli/cli/build"
	"io"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/onec-cli/cli/cmd"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

var discard = ioutil.NopCloser(bytes.NewBuffer(nil))

// todo вынести в internal/test
type FakeCli struct {
	in  io.Reader
	out io.Writer
	err io.Writer
}

// In returns the input stream the cli will use
func (c *FakeCli) In() io.Reader {
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

func NewFakeCli(opts ...func(*FakeCli)) *FakeCli {
	outBuffer := new(bytes.Buffer)
	errBuffer := new(bytes.Buffer)
	c := &FakeCli{
		out: outBuffer,
		err: errBuffer,
		in:  ioutil.NopCloser(strings.NewReader("")),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func runCliCommand(t *testing.T, r io.ReadCloser, w io.Writer, args ...string) error {
	t.Helper()
	if r == nil {
		r = discard
	}
	if w == nil {
		w = ioutil.Discard
	}
	in := func(cli *FakeCli) {
		cli.in = r
	}
	combined := func(cli *FakeCli) {
		cli.out = w
		cli.err = w

	}
	cli := NewFakeCli(in, combined)
	command := cmd.NewRootCommand(cli)
	command.SetArgs(args)
	return command.Execute()
}

func TestExitStatusForInvalidSubcommand(t *testing.T) {
	err := runCliCommand(t, nil, nil, "invalid")
	assert.Check(t, is.ErrorContains(err, "unknown command \"invalid\""))
}

func TestExitStatusForInvalidSubcommandWithHelpFlag(t *testing.T) {
	var b bytes.Buffer
	err := runCliCommand(t, nil, &b, "help", "invalid")
	assert.NilError(t, err)
	assert.Check(t, is.Contains(b.String(), "Unknown help topic [`invalid`]"))
}

func TestVersion(t *testing.T) {
	var b bytes.Buffer
	err := runCliCommand(t, nil, &b, "--version")
	assert.NilError(t, err)
	assert.Check(t, is.Contains(b.String(), build.APP_NAME+" version unknown, build unknown, time unknown"))
}
