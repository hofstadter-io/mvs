package commands

import (
	"github.com/spf13/cobra"
)

var mvsLong = `MVS is a polyglot vendor management tool based on go mods.

mod file format:

  module = "<module path>"

  <name> = "version"

  require (
    ...
  )

  replace <module path> => <local path>
  ...`

var (
	RootLangPflag string
)

func init() {

	RootCmd.PersistentFlags().StringVarP(&RootLangPflag, "lang", "l", "", "The language or system prefix to process. The default is to discover and process all known.")

}

var RootCmd = &cobra.Command{

	Use: "mvs",

	Short: "MVS is a polyglot vendor management tool based on go mods",

	Long: mvsLong,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		// fmt.Println("PersistentPrerun", RootLangPflag, args)

	},
}

func init() {
	RootCmd.AddCommand(ConvertCmd)
	RootCmd.AddCommand(GraphCmd)
	RootCmd.AddCommand(InitCmd)
	RootCmd.AddCommand(TidyCmd)
	RootCmd.AddCommand(VendorCmd)
	RootCmd.AddCommand(VerifyCmd)
	RootCmd.AddCommand(HackCmd)
}
