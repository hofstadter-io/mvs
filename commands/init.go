package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initLong = `initialize a new module in the current directory, language flag is required`

var InitCmd = &cobra.Command{

	Use: "init -l <lang> <module>",

	Short: "initialize a new module in the current directory",

	Long: initLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Module'")
			cmd.Usage()
			os.Exit(1)
		}

		var module string

		if 0 < len(args) {

			module = args[0]

		}

		fmt.Println("Init", RootLangPflag, module)

	},
}
