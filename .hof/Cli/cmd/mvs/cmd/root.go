package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/lib"

	"github.com/hofstadter-io/mvs/cmd/mvs/ga"
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

	PreRun: func(cmd *cobra.Command, args []string) {

		ga.SendGaEvent("root", strings.Join(args, "/"), 0)

	},
}

func init() {

	hf := RootCmd.HelpFunc()
	f := func(cmd *cobra.Command, args []string) {
		if RootCmd.Name() == cmd.Name() {
			as := strings.Join(args, "/")
			ga.SendGaEvent("root/help", as, 0)
		}
		hf(cmd, args)
	}
	RootCmd.SetHelpFunc(f)

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
