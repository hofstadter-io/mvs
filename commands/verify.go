package commands

import (
	"github.com/spf13/cobra"
)

var verifyLong = `verify dependencies have expected content`

var VerifyCmd = &cobra.Command{

	Use: "verify",

	Short: "verify dependencies have expected content",

	Long: verifyLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

	},
}
