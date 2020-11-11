/*
Copyright © 2020 Alexander Strizhachuk <a.strizhachuk@yandex.ru>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"github.com/spf13/viper"
	"log"

	"github.com/spf13/cobra"

	"github.com/v8platform/errors"
	"github.com/v8platform/v8"
)

var user string
var password string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create [connection string]",
	Short: "A brief description of your command",
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
		infobase := newInfobase(args[0])

		err := v8.Run(infobase, v8.CreateFileInfobase(""))

		// todo чёт неудобно
		if err != nil {
			errorContext := errors.GetErrorContext(err)
			out, ok := errorContext["message"]
			if ok {
				err = errors.Internal.Wrap(err, out)
			}
			log.Fatalln(err)
		}
		log.Println("Infobase created: " + infobase.Path())
	},
}

func newInfobase(connectionString string) v8.Infobase {

	path := connectionString
	v := v8.NewFileIB(path)

	return v
}

func infobaseFromString(connectionString string) v8.Infobase {

	return nil

}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	createCmd.PersistentFlags().StringP("user", "u", "", "user")         //todo test
	createCmd.PersistentFlags().StringP("password", "p", "", "password") //todo test

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	viper.BindPFlag("user", createCmd.PersistentFlags().Lookup("user"))
	viper.BindPFlag("password", createCmd.PersistentFlags().Lookup("password"))

	viper.SetDefault("user", "")
	viper.SetDefault("password", "")
}
