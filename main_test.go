package main_test

import (
	"bytes"
	"github.com/onec-cli/cli/cli/build"
	"github.com/onec-cli/cli/internal/test"
	"io"
	"io/ioutil"
	"testing"

	"github.com/onec-cli/cli/cmd"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

var discard = ioutil.NopCloser(bytes.NewBuffer(nil))

func runCliCommand(t *testing.T, r io.ReadCloser, w io.Writer, args ...string) error {
	t.Helper()
	if r == nil {
		r = discard
	}
	if w == nil {
		w = ioutil.Discard
	}
	in := func(cli *test.FakeCli) {
		cli.SetIn(r)
	}
	combined := func(cli *test.FakeCli) {
		cli.SetOut(w)
		cli.SetErr(w)

	}
	cli := test.NewFakeCli(in, combined)
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
	assert.Check(t, is.Contains(b.String(), build.AppName+" version unknown, build unknown, time unknown"))
}
