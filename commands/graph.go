package commands

import (
	"github.com/spf13/cobra"
)

var graphLong = `print module requirement graph`

var GraphCmd = &cobra.Command{

	Use: "graph",

	Short: "print module requirement graph",

	Long: graphLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

	},
}
