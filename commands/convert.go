package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var convertLong = `convert another package system to MVS, language flag is required`

var ConvertCmd = &cobra.Command{

	Use: "convert -l <lang> <file>",

	Short: "convert another package system to MVS.",

	Long: convertLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		if 0 >= len(args) {
			fmt.Println("missing required argument: 'File'")
			cmd.Usage()
			os.Exit(1)
		}

		var file string

		if 0 < len(args) {

			file = args[0]

		}

		fmt.Println("Convert", RootLangPflag, file)

	},
}
