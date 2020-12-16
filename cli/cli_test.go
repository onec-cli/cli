package cli_test

import (
	"github.com/onec-cli/cli/cli"
	"testing"
)

func TestNewCliStreams(t *testing.T) {
	got := cli.NewCli()
	assertStream(t, "input", func() bool {
		return got.In() != nil
	})
	assertStream(t, "output", func() bool {
		return got.Out() != nil
	})
	assertStream(t, "error", func() bool {
		return got.Err() != nil
	})
}

func assertStream(t *testing.T, name string, want func() bool) {
	t.Run(name, func(t *testing.T) {
		if !want() {
			t.Error(name, "<nil>")
		}
	})
}
