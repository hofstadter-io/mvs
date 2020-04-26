package main

import (
	"fmt"
	"os"

	"github.com/hofstadter-io/mvs/cmd/mvs/cmd"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
