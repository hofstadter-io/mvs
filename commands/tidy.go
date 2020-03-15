package commands

import (
	"github.com/spf13/cobra"
)

var tidyLong = `add missinad and remove unused modules`

var TidyCmd = &cobra.Command{

	Use: "tidy",

	Short: "add missinad and remove unused modules",

	Long: tidyLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

	},
}
