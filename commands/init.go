package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/pkg"
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

		if RootLangPflag == "" {
			fmt.Println("language flag is required for this command")
			cmd.Usage()
			os.Exit(1)
		}
		err := pkg.Init(RootLangPflag, module)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}
