package lib

import (
	"fmt"
	"os"

	"cuelang.org/go/cue/load"
	"github.com/hofstadter-io/mvs/lib/util"
)

func Hack(lang string, args []string) error {
	fmt.Println("Hack", args)

	bis := load.Instances([]string{}, nil)
	for _, bi := range bis {
		if bi.Err != nil {
			fmt.Println(bi.Err.Error())
			os.Exit(1)
		}
		i, err := util.CueRuntime.Build(bi)
		if err != nil {
			fmt.Println(bi.Err.Error())
			os.Exit(1)
		}
		fmt.Println(i.Lookup("a"))
	}

	return nil
}
