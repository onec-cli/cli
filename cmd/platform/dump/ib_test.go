package dump

import (
	"github.com/onec-cli/cli/internal/test"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
	"testing"
)

func TestNewDumpIBCommand(t *testing.T) {
	cli := test.NewFakeCli()
	cmd := NewDumpIBCommand(cli)
	cmd.SetOut(cli.OutBuffer())
	err := cmd.Execute()
	assert.NilError(t, err)
	assert.Check(t, is.Contains(cli.OutBuffer().String(), `
Usage:
	onec platform dump ib`))

}
