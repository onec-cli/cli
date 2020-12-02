package cmd

import (
	"context"
	"github.com/briandowns/spinner"
	"github.com/onec-cli/cli/command/create"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v8errors "github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"log"
	"time"
)

var sp = spinner.New(spinner.CharSets[14], 100*time.Millisecond)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create CONNECTION_STRING...",
	Short: "Create new infobase",
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

func init() {
	rootCmd.AddCommand(createCmd)

	// Persistent Flags which will work for this command and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")
	createCmd.PersistentFlags().StringP("user", "u", "", "user")         //todo test
	createCmd.PersistentFlags().StringP("password", "p", "", "password") //todo test

	// Local flags which will only run when this command is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	//todo если нет shorthand то заменить на Flags().String
	createCmd.Flags().StringP("db-type", "", "PostgreSQL", "db server type")
	createCmd.Flags().StringP("db-server", "", "db", "db server") // todo заменить на localhost
	createCmd.Flags().StringP("db", "", "ib", "db name")
	createCmd.Flags().StringP("db-user", "", "postgres", "db user")
	createCmd.Flags().BoolP("db-create", "", true, "create db")

	createCmd.Flags().StringP("server", "", "", "server")

	// Viper bind
	viper.BindPFlag("usr", createCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("pwd", createCmd.PersistentFlags().Lookup("password"))

	viper.BindPFlag("dbms", createCmd.Flags().Lookup("db-type"))
	viper.BindPFlag("dbsrvr", createCmd.Flags().Lookup("db-server"))
	viper.BindPFlag("db", createCmd.Flags().Lookup("db"))
	viper.BindPFlag("dbuid", createCmd.Flags().Lookup("db-user"))
	viper.BindPFlag("crsqldb", createCmd.Flags().Lookup("db-create"))

	// Viper default

	//viper.SetDefault("user", "")
	//viper.SetDefault("password", "")
	//viper.SetDefault("db-type", "PostgreSQL")
}

func runCreate(args []string) {

	log.Println("Creation infobase started:")

	sp.Start()

	options, err := create.DefaultOptions(viper.AllSettings())
	if err != nil {
		return
	}

	infobases := create.NewInfobases(args, options...)

	for i, infobase := range infobases {

		log.Printf("infobase #%d\n", i+1)

		what, err := infobase.Command()
		if err != nil {
			log.Println("error: ", err)
			continue
		}

		platformRunner := runner.NewPlatformRunner(nil, what)
		sp.Stop()
		log.Printf("=> %v\n", platformRunner.Args())
		sp.Start()
		err = platformRunner.Run(context.Background())

		sp.Stop()
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
		sp.Start()
	}
	sp.Stop()
}
