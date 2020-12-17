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
	"fmt"
	"github.com/onec-cli/cli/cli"
	"github.com/onec-cli/cli/cli/build"
	"github.com/onec-cli/cli/cmd/config"
	"github.com/onec-cli/cli/cmd/platform"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

func NewRootCommand(cli cli.Cli) *cobra.Command {

	cmd := &cobra.Command{
		Use:   build.APP_NAME,
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		SilenceErrors: true,
		Version:       fmt.Sprintf("%s, build %s, time %s", build.Version, build.GitCommit, build.Time),
	}

	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	//cmd.Flags().BoolP("version", "v", false, "Print version information and quit")
	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	cmd.SetOut(cli.Out())

	AddCommands(cli, cmd)

	return cmd
}

// AddCommands adds all the commands from ./cmd to the root command
func AddCommands(cli cli.Cli, cmd *cobra.Command) {
	cmd.AddCommand(platform.NewPlatformCommand(cli))
	cmd.AddCommand(config.NewConfigCommand(cli))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName("." + build.APP_NAME)
		viper.SetConfigType("json")

		//viper.WriteConfigAs(filepath.Join(home, ".cli.json"))//todo test
	}

	viper.SetEnvPrefix(build.APP_NAME)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
