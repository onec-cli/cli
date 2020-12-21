package dump

import (
	"github.com/onec-cli/cli/cli"
	"github.com/spf13/cobra"
)

// NewDumpCommand returns a cobra command for `dump` subcommands
func NewDumpCommand(_ cli.Cli) *cobra.Command {
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
	return cmd
}

//
//func runCreate(connectPaths []string) {
//
//	log.Println("Creation infobase started:")
//
//	Spinner.Start()
//
//	// todo DefaultOptions to def conn str
//	defaultOptions, err := platform.DefaultOptions(viper.AllSettings())
//	if err != nil {
//		return
//	}
//
//	infobases := platform.NewInfobases(connectPaths, defaultOptions...)
//
//	for i, infobase := range infobases {
//
//		log.Printf("infobase #%d\n", i+1)
//
//		what, err := infobase.Command()
//		if err != nil {
//			log.Println("error: ", err)
//			continue
//		}
//
//		opts := options()
//		platformRunner := runner.NewPlatformRunner(nil, what, opts...)
//
//		Spinner.Stop()
//		log.Printf("=> %v\n", platformRunner.Args())
//		Spinner.Start()
//		err = platformRunner.Run(context.Background())
//
//		Spinner.Stop()
//		if err != nil {
//			// todo много букв
//			errorContext := v8errors.GetErrorContext(err)
//			out, ok := errorContext["message"]
//			if ok {
//				err = v8errors.Internal.Wrap(err, out)
//			}
//			log.Println("error: ", err)
//		} else {
//			log.Println("infobase created")
//		}
//		Spinner.Start()
//	}
//	Spinner.Stop()
//}
//
//func options() []interface{} {
//
//	out := viper.GetString("out")
//	if out == "" {
//		return nil
//	}
//	t := viper.GetBool("out-trunc")
//
//	var o []interface{}
//	o = append(o, runner.WithOut(out, t))
//
//	return o
//}
