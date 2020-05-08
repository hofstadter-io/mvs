package cmd

import (
	"fmt"
	"os"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hofstadter-io/mvs/lib"

	"github.com/hofstadter-io/mvs/cmd/mvs/ga"
)

var vendorLong = `make a vendored copy of dependencies`

func VendorRun(args []string) (err error) {

	err = lib.ProcessLangs("vendor", args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return err
}

var VendorCmd = &cobra.Command{

	Use: "vendor [langs...]",

	Short: "make a vendored copy of dependencies",

	Long: vendorLong,

	PreRun: func(cmd *cobra.Command, args []string) {

		cs := strings.Fields(cmd.CommandPath())
		c := strings.Join(cs[1:], "/")
		ga.SendGaEvent(c, strings.Join(args, "/"), 0)

	},

	Run: func(cmd *cobra.Command, args []string) {
		var err error

		// Argument Parsing

		err = VendorRun(args)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}
