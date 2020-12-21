package platform

import (
	"github.com/onec-cli/cli/cli"
	"github.com/onec-cli/cli/cmd/platform/dump"
	"github.com/spf13/cobra"
)

// NewPlatformCommand returns a cobra command for `platform` subcommands
func NewPlatformCommand(cli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "platform",
		Aliases: []string{"p"},
		Short:   "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		//Run: func(cmd *cobra.Command, args []string) {
		//	fmt.Println("platform called")
		//},
	}
	cmd.AddCommand(
		NewCreateCommand(cli),
		NewRunCommand(cli),
		dump.NewDumpCommand(cli),
	)
	return cmd
}
