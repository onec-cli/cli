package client_test

import (
	"github.com/onec-cli/cli/client"
	"gotest.tools/v3/assert"
	"testing"
)

func TestNewCliStreams(t *testing.T) {
	cli := client.NewClient()
	assert.Check(t, cli.In() != nil)
	assert.Check(t, cli.Out() != nil)
	assert.Check(t, cli.Err() != nil)
}
