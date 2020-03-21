package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Graph() error {

	// Graph Command Override
	if len(mdr.CommandGraph) > 0 {
		out, err := util.Exec(mdr.CommandGraph)
		fmt.Println(out)
		if err != nil {
			return err
		}
	} else {
		// Otherwise, MVS venodiring
		err := mdr.GraphMVS()
		if err != nil {
			mdr.PrintErrors()
			return err
		}
	}

	return nil
}

// The entrypoint to the MVS internal verify process
func (mdr *Modder) GraphMVS() error {

	// Load minimal root module
	err := mdr.LoadMinimalFromFS(".")
	if err != nil {
		return err
	}

	return nil
}

