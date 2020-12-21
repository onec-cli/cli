package dump

import (
	"github.com/onec-cli/cli/cli"
	"github.com/spf13/cobra"
)

// NewDumpCommand returns a cobra command for `dump` subcommands
func NewDumpCommand(cli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "dump",
		Aliases: []string{"d"},
		Short:   "Dump...",
		Long: `Dump... A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	}
	cmd.AddCommand(NewDumpIBCommand(cli))
	return cmd
}
