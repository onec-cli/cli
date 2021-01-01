package dump

import (
	"fmt"
	"github.com/onec-cli/cli/cli"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
)

//todo кандидат на мув ап
func logIfError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

type dumpOptions struct {
	file         string
	ibConnection string
}

// NewDumpIBCommand creates a new cobra.Command for `onec platform dump ib`
func NewDumpIBCommand(cli cli.Cli) *cobra.Command {
	var opts dumpOptions
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
			opts.file = args[0]
			return runDumpIB(cli, opts)
		},
	}
	cmd.Flags().StringVar(&opts.ibConnection, "ibconnection", "", "ibconnection (required)")
	logIfError(cmd.MarkFlagRequired("ibconnection"))

	return cmd
}

func runDumpIB(cli cli.Cli, opts dumpOptions) error {
	if err := ValidateOutputPath(opts.file); err != nil {
		return errors.Wrap(err, "failed to export infobase")
	}

	err := cli.NewRunner(nil).DumpIB(opts.file)
	if err != nil {
		return err
	}

	//ib, err := v8.NewTempIB()
	//
	//what := v8.DumpIB(file)
	//platformRunner := runner.NewPlatformRunner(ib, what)
	//err = platformRunner.Run(nil)
	//if err != nil {
	//	return err
	//}

	return nil
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
