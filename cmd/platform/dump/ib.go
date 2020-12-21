package dump

import (
	"github.com/onec-cli/cli/cli"
	"github.com/spf13/cobra"
)

// NewDumpIBCommand creates a new cobra.Command for `onec platform dump ib`
func NewDumpIBCommand(_ cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "ib",
		Aliases: []string{"i"},
		Short:   "Dump IB...",
		Long: `Dump IB... A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			return
		},
	}
	return cmd
}
