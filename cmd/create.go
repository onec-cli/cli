package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v8errors "github.com/v8platform/errors"
	"github.com/v8platform/runner"
	"log"
	"strings"
)

//type createOptions struct {
//	user string
//	password string
//}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create CONNECTION_STRING",
	Short: "Create new database",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Create infobase started")

		//viper.GetString("user"), viper.GetString("password")
		// todo https://github.com/v8platform/v8/issues/2

		where := new(infobase)
		what, err := CreateInfobase(args[0])

		if err != nil {
			log.Fatalln(err)
		}

		err = runner.Run(where, what)

		// todo чёт неудобно
		if err != nil {
			errorContext := v8errors.GetErrorContext(err)
			out, ok := errorContext["message"]
			if ok {
				err = v8errors.Internal.Wrap(err, out)
			}
			log.Fatalln(err)
		}

		//	infobase := conn.Infobase()

		//	log.Println("infobase created: " + infobase.ConnectionString())
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

	// Viper bind
	viper.BindPFlag("user", createCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", createCmd.PersistentFlags().Lookup("password"))

	// Viper default
	viper.SetDefault("user", "")
	viper.SetDefault("password", "")
}

type infobase struct{}

func (i *infobase) ConnectionString() string {
	return ""
}

type connectionString struct {
	connectionString string
	values           []string
}

func (c *connectionString) Command() string {
	return runner.CreateInfobase
}

func (c *connectionString) Check() error {
	return nil
}

func (c *connectionString) Values() []string {
	return c.values
}

func (c *connectionString) parse() error {

	switch {
	case strings.HasPrefix(strings.ToUpper(c.connectionString), "/F"):

		c.values = append(c.values, "File="+strings.TrimLeft(c.connectionString, "/F"))

	case strings.Contains(c.connectionString, "File=") ||
		strings.Contains(c.connectionString, "Srvr="):

		c.values = strings.Split(c.connectionString, ";")
		// TODO Надо почистить от пустых строк и артифактов
	default:
		return errors.New("invalid connection string format")
	}

	return nil
}

func CreateInfobase(c string) (runner.Command, error) {
	command := connectionString{connectionString: c}
	err := command.parse()
	if err != nil {
		return nil, err
	}
	return &command, nil
}
