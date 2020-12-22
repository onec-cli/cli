package dump_test

import (
	"errors"
	"github.com/onec-cli/cli/cmd/platform/dump"
	"github.com/onec-cli/cli/internal/test"
	"gotest.tools/v3/assert"
	"gotest.tools/v3/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNewDumpIBCommandErrors(t *testing.T) {
	testCases := []struct {
		name          string
		args          []string
		expectedError string
	}{
		{
			name:          "wrong-args",
			args:          []string{},
			expectedError: "accepts 1 arg(s), received 0",
		},
		{
			name: "wrong-path",
			args: func() []string {
				basedir, _ := ioutil.TempDir("", "dump-ib")
				defer os.RemoveAll(basedir)
				return []string{filepath.Join(basedir, "notexist")}
			}(),
			expectedError: "failed to export infobase: invalid output path",
		},
	}
	for _, tc := range testCases {
		cmd := dump.NewDumpIBCommand(test.NewFakeCli())
		cmd.SetOut(ioutil.Discard)
		cmd.SetErr(ioutil.Discard)
		cmd.SetArgs(tc.args)
		assert.ErrorContains(t, cmd.Execute(), tc.expectedError)
	}
}

func TestNewDumpIBToFile(t *testing.T) {
	dir := fs.NewDir(t, "dump-ib")
	defer dir.Remove()

	cmd := dump.NewDumpIBCommand(test.NewFakeCli())
	cmd.SetArgs([]string{dir.Join("foo.dt")})

	assert.NilError(t, cmd.Execute())
	expected := fs.Expected(t,
		fs.WithFile("foo.dt", "", fs.MatchAnyFileMode),
	)
	assert.Assert(t, fs.Equal(dir.Path(), expected))
}

//func TestNewDumpIBToFileIntegrate(t *testing.T) {
//	dir := fs.NewDir(t, "dump-ib")
//	defer dir.Remove()
//
//	cmd := dump.NewDumpIBCommand(test.NewFakeCli())
//	cmd.SetArgs([]string{dir.Join("foo.dt")})
//
//	assert.NilError(t, cmd.Execute())
//	expected := fs.Expected(t,
//		fs.WithFile("foo.dt", "", fs.MatchAnyFileMode),
//	)
//	assert.Assert(t, fs.Equal(dir.Path(), expected))
//}

func TestValidateOutputPath(t *testing.T) {
	basedir, err := ioutil.TempDir("", "dump-ib")
	assert.NilError(t, err)
	defer os.RemoveAll(basedir)
	dir := filepath.Join(basedir, "dir")
	notexist := filepath.Join(basedir, "notexist")
	err = os.MkdirAll(dir, 0755)
	assert.NilError(t, err)
	file := filepath.Join(dir, "file")
	err = ioutil.WriteFile(file, []byte("hi"), 0644)
	assert.NilError(t, err)
	var testcases = []struct {
		path string
		err  error
	}{
		{basedir, nil},
		{file, nil},
		{dir, nil},
		{dir + string(os.PathSeparator), nil},
		{notexist, nil},
		{notexist + string(os.PathSeparator), nil},
		{filepath.Join(notexist, "file"), errors.New("does not exist")},
	}

	for _, testcase := range testcases {
		t.Run(testcase.path, func(t *testing.T) {
			err := dump.ValidateOutputPath(testcase.path)
			if testcase.err == nil {
				assert.NilError(t, err)
			} else {
				assert.ErrorContains(t, err, testcase.err.Error())
			}
		})
	}
}