package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/pkg"
)

var verifyLong = `verify dependencies have expected content`

var VerifyCmd = &cobra.Command{

	Use: "verify",

	Short: "verify dependencies have expected content",

	Long: verifyLong,

	Run: func(cmd *cobra.Command, args []string) {

		// Argument Parsing

		err := pkg.Verify(RootLangPflag)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}
