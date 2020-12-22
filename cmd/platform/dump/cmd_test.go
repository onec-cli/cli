package dump

import (
	"github.com/onec-cli/cli/internal/test"
	"gotest.tools/v3/assert"
	"io/ioutil"
	"testing"
)

//todo избыточный хрупкий тест? надо ли проверять подкоманды без действий?
func TestNewDumpCommand(t *testing.T) {
	cmd := NewDumpCommand(test.NewFakeCli(nil))
	cmd.SetOut(ioutil.Discard)
	err := cmd.Execute()
	assert.NilError(t, err)
}
