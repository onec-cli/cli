package cli_test

import (
	"github.com/onec-cli/cli/cli"
	"gotest.tools/v3/assert"
	"reflect"
	"testing"
)

func TestNewCliStreams(t *testing.T) {
	cli := cli.NewCli()
	assert.Check(t, cli.In() != nil)
	assert.Check(t, cli.Out() != nil)
	assert.Check(t, cli.Err() != nil)
}

func TestInfobaseError(t *testing.T) {
	cli := cli.NewCli()
	ib := cli.Infobase("foo")
	assert.Check(t, ib.Error() != nil)
}

func TestInfobaseSuccess(t *testing.T) {
	cli := cli.NewCli()
	ib1 := cli.Infobase("/F./foo")
	assert.Check(t, ib1 != nil)
	ib2 := cli.Infobase("/F./foo")
	assert.Check(t, ib1 == ib2)
	ib3 := cli.Infobase("/F./boo")
	assert.Check(t, !reflect.DeepEqual(ib1, ib3))
}
