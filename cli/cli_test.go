package cli_test

import (
	"github.com/onec-cli/cli/cli"
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewCliStreams(t *testing.T) {
	cli := cli.NewCli()
	assert.Check(t, cli.In() != nil)
	assert.Check(t, cli.Out() != nil)
	assert.Check(t, cli.Err() != nil)
}
