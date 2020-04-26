package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/lib"
)

var mvsLong = `MVS is a polyglot dependency management tool based on go mods.

mod file format:

  module = "<module path>"

  <name> = "version"

  require (
    ...
  )

  replace <module path> => <local path>
  ...`

func RootPersistentPreRun(args []string) (err error) {

	// fmt.Println("PersistentPrerun", args)
	lib.InitLangs()

	return err
}

var RootCmd = &cobra.Command{

	Use: "mvs",

	Short: "MVS is a polyglot dependency management tool based on go mods",

	Long: mvsLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = RootPersistentPreRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.AddCommand(InfoCmd)
	RootCmd.AddCommand(ConvertCmd)
	RootCmd.AddCommand(GraphCmd)
	RootCmd.AddCommand(StatusCmd)
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(TidyCmd)
	RootCmd.AddCommand(VendorCmd)
	RootCmd.AddCommand(VerifyCmd)
	RootCmd.AddCommand(HackCmd)
}

func initConfig() {

}
