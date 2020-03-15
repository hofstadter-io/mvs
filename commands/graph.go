package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/pkg"
)

var graphLong = `print module requirement graph`

var GraphCmd = &cobra.Command{

	Use: "graph",

	Short: "print module requirement graph",

	Long: graphLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		err := pkg.Graph(RootLangPflag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}
