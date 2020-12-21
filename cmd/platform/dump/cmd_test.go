package dump

import (
	"bytes"
	"github.com/onec-cli/cli/internal/test"
	"gotest.tools/v3/assert"
	"strings"
	"testing"
)

func TestNewDumpCommand(t *testing.T) {
	var b bytes.Buffer
	cli := test.NewFakeCli()
	cmd := NewDumpCommand(cli)
	cmd.SetOut(&b)
	err := cmd.Execute()
	assert.NilError(t, err)
	if !strings.HasPrefix(b.String(), "Dump...") {
		t.Error("Wrong description") //todo хрупкий тест
	}
}
