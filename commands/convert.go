package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/pkg"
)

var convertLong = `convert another package system to MVS, language flag is required`

var ConvertCmd = &cobra.Command{

	Use: "convert -l <lang> <file>",

	Short: "convert another package system to MVS.",

	Long: convertLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'Filename'")
			cmd.Usage()
			os.Exit(1)
		}

		var filename string

		if 0 < len(args) {

			filename = args[0]

		}

		if RootLangPflag == "" {
			fmt.Println("language flag is required for this command")
			cmd.Usage()
			os.Exit(1)
		}
		err := pkg.Convert(RootLangPflag, filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}
