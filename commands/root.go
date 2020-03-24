package commands

import (
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

var RootCmd = &cobra.Command{

	Use: "mvs",

	Short: "MVS is a polyglot dependency management tool based on go mods",

	Long: mvsLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// fmt.Println("PersistentPrerun", args)
		lib.InitLangs()

	},
}

func init() {
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
