package dump

import (
	"github.com/onec-cli/cli/internal/test"
	"gotest.tools/v3/assert"
	"io/ioutil"
	"testing"
)

func TestNewDumpCommand(t *testing.T) {
	cmd := NewDumpCommand(test.NewFakeCli())
	cmd.SetOut(ioutil.Discard)
	err := cmd.Execute()
	assert.NilError(t, err)
}
