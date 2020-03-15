package commands

import (
	"github.com/spf13/cobra"
)

var vendorLong = `make a vendored copy of dependencies`

var VendorCmd = &cobra.Command{

	Use: "vendor",

	Short: "make a vendored copy of dependencies",

	Long: vendorLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

	},
}
