package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Tidy() error {

	// Tidy Command Override
	if len(mdr.CommandTidy) > 0 {
		for _, cmd := range mdr.CommandGraph {
			out, err := util.Exec(cmd)
			fmt.Println(out)
			if err != nil {
				return err
			}
		}
	} else {
		// Otherwise, MVS venodiring
		err := mdr.TidyMVS()
		if err != nil {
			mdr.PrintErrors()
			return err
		}
	}

	return nil
}

// The entrypoint to the MVS internal verify process
func (mdr *Modder) TidyMVS() error {

	// Load minimal root module
	err := mdr.LoadMinimalFromFS(".")
	if err != nil {
		return err
	}

	return nil
}
