package platform

import (
	"context"
	"github.com/onec-cli/cli/cli"
	. "github.com/onec-cli/cli/cli/spinner"
	"github.com/onec-cli/cli/platform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v8errors "github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"log"
)

// NewCreateCommand creates a new cobra.Command for `cli platform create`
func NewCreateCommand(_ cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create CONNECTION_STRING...",
		Aliases: []string{"c"},
		Short:   "Create new infobase",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			runCreate(args)
		},
	}

	// Persistent Flags which will work for this command and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	cmd.PersistentFlags().StringP("user", "u", "", "user")         //todo test
	cmd.PersistentFlags().StringP("password", "p", "", "password") //todo test

	// Local flags which will only run when this command is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//todo если нет shorthand то заменить на Flags().String
	// db
	cmd.Flags().StringP("db-type", "", "PostgreSQL", "db server type")
	cmd.Flags().StringP("db-server", "", "db", "db server") // todo заменить на localhost
	cmd.Flags().StringP("db", "", "ib", "db name")
	cmd.Flags().StringP("db-user", "", "postgres", "db user")
	cmd.Flags().BoolP("db-create", "", true, "create db")

	// common
	cmd.Flags().StringP("out", "o", "", "out file")
	cmd.Flags().BoolP("out-trunc", "", false, "truncate out file")

	// Viper bind
	viper.BindPFlag("usr", cmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("pwd", cmd.PersistentFlags().Lookup("password"))

	viper.BindPFlag("dbms", cmd.Flags().Lookup("db-type"))
	viper.BindPFlag("dbsrvr", cmd.Flags().Lookup("db-server"))
	viper.BindPFlag("db", cmd.Flags().Lookup("db"))
	viper.BindPFlag("dbuid", cmd.Flags().Lookup("db-user"))
	viper.BindPFlag("crsqldb", cmd.Flags().Lookup("db-create"))

	viper.BindPFlag("out", cmd.Flags().Lookup("out"))
	viper.BindPFlag("out-trunc", cmd.Flags().Lookup("out-trunc"))

	// Viper default

	//viper.SetDefault("user", "")
	//viper.SetDefault("password", "")
	//viper.SetDefault("db-type", "PostgreSQL")

	return cmd
}

func runCreate(args []string) {

	log.Println("Creation infobase started:")

	Spinner.Start()

	// todo DefaultOptions to def conn str
	defaultOptions, err := platform.DefaultOptions(viper.AllSettings())
	if err != nil {
		return
	}

	infobases := platform.NewInfobases(args, defaultOptions...)

	for i, infobase := range infobases {

		log.Printf("infobase #%d\n", i+1)

		what, err := infobase.Command()
		if err != nil {
			log.Println("error: ", err)
			continue
		}

		opts := options()
		platformRunner := runner.NewPlatformRunner(nil, what, opts...)

		Spinner.Stop()
		log.Printf("=> %v\n", platformRunner.Args())
		Spinner.Start()
		err = platformRunner.Run(context.Background())

		Spinner.Stop()
		if err != nil {
			// todo много букв
			errorContext := v8errors.GetErrorContext(err)
			out, ok := errorContext["message"]
			if ok {
				err = v8errors.Internal.Wrap(err, out)
			}
			log.Println("error: ", err)
		} else {
			log.Println("infobase created")
		}
		Spinner.Start()
	}
	Spinner.Stop()
}

func options() []interface{} {

	out := viper.GetString("out")
	if out == "" {
		return nil
	}
	t := viper.GetBool("out-trunc")

	var o []interface{}
	o = append(o, runner.WithOut(out, t))

	return o
}
