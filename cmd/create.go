package cmd

import (
	"context"
	"github.com/briandowns/spinner"
	"github.com/onec-cli/cli/api"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	v8errors "github.com/v8platform/errors"
	"github.com/v8platform/marshaler"
	"github.com/v8platform/runner"
	"log"
	"reflect"
	"time"
)

//type createOptions struct {
//	user string
//	password string
//}

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
	Run: func(cmd *cobra.Command, args []string) {

		log.Println("Create infobase started")

		//viper.GetString("user"), viper.GetString("password")
		//viper.GetString("dbms")

		//todo заполнить designer.CreateServerInfoBaseOptions (а точнее, создать свою структуру defOpt c v8
		//заполнить ее деф значениями viper
		opts := new(defaultOptions)
		opts.bindViper()
		marshal, err := marshaler.Marshal(opts)
		log.Println(marshal, err)
		//передать api.CreateInfobase
		//там после command.parse() добавить значения из деф после маршалинга получив []string ключ=значение
		//если подобного значения не было

		spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
		spinner.Start()

		infobases := api.CreateInfobase(args, marshal...)
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

//func spinner(delay time.Duration) {
//	for {
//		for _, r := range `-\|/` {
//			fmt.Printf("\r%c", r)
//			time.Sleep(delay)
//		}
//	}
//}

type defaultOptions struct {
	//тип используемого сервера баз данных:
	// MSSQLServer — Microsoft SQL Server;
	// PostgreSQL — PostgreSQL;
	// IBMDB2 — IBM DB2;
	// OracleDatabase — Oracle Database.
	DBMS string `v8:"DBMS, equal_sep" json:"dbms"`

	//имя сервера баз данных;
	DBSrvr string `v8:"DBSrvr, equal_sep" json:"db_srvr"`

	// имя базы данных в сервере баз данных;
	DB string `v8:"DB, equal_sep" json:"db_ref"`

	//имя пользователя сервера баз данных;
	DBUID string `v8:"DBUID, equal_sep" json:"db_user"`

	// создать базу данных в случае ее отсутствия ("Y"|"N".
	// "Y" — создавать базу данных в случае отсутствия,
	// "N" — не создавать. Значение по умолчанию — N).
	CrSQLDB bool `v8:"CrSQLDB, optional, equal_sep, bool_true=Y" json:"create_db"`
}

func (o *defaultOptions) bindViper() {

	st := reflect.TypeOf(*o)
	el := reflect.ValueOf(o).Elem()
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		v := viper.GetString(field.Name)
		f := el.FieldByName(field.Name)
		switch f.Interface().(type) {
		case string:
			f.SetString(v)
		case bool:
			f.SetBool(v == "true")
		}
	}
}
