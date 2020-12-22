package dump

import (
	"fmt"
	"github.com/onec-cli/cli/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	_ "github.com/v8platform/api"
	//"github.com/v8platform/runner"
	"os"
	"path/filepath"
)

// NewDumpIBCommand creates a new cobra.Command for `onec platform dump ib`
func NewDumpIBCommand(_ cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ib FILE",
		Aliases: []string{"i"},
		Short:   "Dump IB...",
		Long: `Dump IB... A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := ValidateOutputPath(args[0]); err != nil {
				return errors.Wrap(err, "failed to export infobase")
			}
			//ib, err := v8.NewTempIB()
			//
			//what := v8.DumpIB(args[0])
			//platformRunner := runner.NewPlatformRunner(ib, what)
			//err = platformRunner.Run(nil)
			//if err != nil {
			//	return err
			//}
			return nil
		},
	}
	return cmd
}

// todo кандидаты на вынос в отдельный модуль
// ValidateOutputPath validates the output paths of the `save` commands.
func ValidateOutputPath(path string) error {
	dir := filepath.Dir(filepath.Clean(path))
	if dir != "" && dir != "." {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			return errors.Errorf("invalid output path: directory %q does not exist", dir)
		}
	}
	// check whether `path` points to a regular file
	// (if the path exists and doesn't point to a directory)
	if fileInfo, err := os.Stat(path); !os.IsNotExist(err) {
		if err != nil {
			return err
		}

		if fileInfo.Mode().IsDir() || fileInfo.Mode().IsRegular() {
			return nil
		}

		if err := ValidateOutputPathFileMode(fileInfo.Mode()); err != nil {
			return errors.Wrapf(err, fmt.Sprintf("invalid output path: %q must be a directory or a regular file", path))
		}
	}
	return nil
}

// ValidateOutputPathFileMode validates the output paths of commands and serves as a
// helper to `ValidateOutputPath`
func ValidateOutputPathFileMode(fileMode os.FileMode) error {
	switch {
	case fileMode&os.ModeDevice != 0:
		return errors.New("got a device")
	case fileMode&os.ModeIrregular != 0:
		return errors.New("got an irregular file")
	}
	return nil
}
