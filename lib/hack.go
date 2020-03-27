package lib

import (
	"fmt"
	"os"

	"cuelang.org/go/cue/format"
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
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Println("trying to find a")
		fmt.Println(i.Lookup("a"))
		fmt.Println("trying to loop")
		s, err := i.Value().Struct()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		iter := s.Fields()
		for iter.Next() {
			fmt.Println("- iterator found")
			v := iter.Value()
			bytes, err := format.Node(v.Syntax())
			if err != nil {
				return err
			}
			fmt.Println(string(bytes))
		}
	}

	return nil
}
