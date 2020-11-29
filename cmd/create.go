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

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create CONNECTION_STRING...",
	Short: "Create new database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return runCreate(args)
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

	createCmd.Flags().String("claster-user", "", "claster user") // todo для теста дефолтных, эксперимент

	// Viper bind
	viper.BindPFlag("usr", createCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("pwd", createCmd.PersistentFlags().Lookup("password"))
	viper.BindPFlag("dbms", createCmd.Flags().Lookup("db-type"))
	viper.BindPFlag("dbsrvr", createCmd.Flags().Lookup("db-server"))
	viper.BindPFlag("db", createCmd.Flags().Lookup("db"))
	viper.BindPFlag("dbuid", createCmd.Flags().Lookup("db-user"))
	viper.BindPFlag("crsqldb", createCmd.Flags().Lookup("db-create"))

	viper.BindPFlag("tester", createCmd.Flags().Lookup("claster-user"))

	// Viper default

	//viper.SetDefault("user", "")
	//viper.SetDefault("password", "")
	//viper.SetDefault("db-type", "PostgreSQL")
}

func runCreate(args []string) error {

	log.Println("Create infobase started")

	//todo заполнить designer.CreateServerInfoBaseOptions (а точнее, создать свою структуру defOpt c v8
	//заполнить ее деф значениями viper
	options, err := create.GetDefaultOptions(viper.AllSettings())

	log.Println(options, err)
	//передать api.CreateInfobase
	//там после command.parse() добавить значения из деф после маршалинга получив []string ключ=значение
	//если подобного значения не было

	spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	spinner.Start()

	infobases := create.CreateInfobase(args, options...)
	for _, infobase := range infobases {
		what, err := infobase.Command()
		if err != nil {
			log.Println(err)
			continue
		}
		platformRunner := runner.NewPlatformRunner(nil, what)
		//go spinner(100 * time.Millisecond)
		err = platformRunner.Run(context.Background())
		// todo много букв
		if err != nil {
			errorContext := v8errors.GetErrorContext(err)
			out, ok := errorContext["message"]
			if ok {
				err = v8errors.Internal.Wrap(err, out)
			}
			log.Println(err)
		}
		spinner.Stop()
		log.Printf("New infobase created: %v", platformRunner.Args())
		spinner.Start()
	}
	spinner.Stop()

	return nil
}
