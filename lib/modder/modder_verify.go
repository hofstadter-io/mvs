package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Verify() error {

	// Verify Command Override
	if len(mdr.CommandVerify) > 0 {
		out, err := util.Exec(mdr.CommandVerify)
		fmt.Println(out)
		if err != nil {
			return err
		}
	} else {
		// Otherwise, MVS venodiring
		err := mdr.VerifyMVS()
		if err != nil {
			mdr.PrintErrors()
			return err
		}
	}

	return nil
}

// The entrypoint to the MVS internal verify process
func (mdr *Modder) VerifyMVS() error {

	// Load minimal root module
	err := mdr.LoadMinimalFromFS(".")
	if err != nil {
		return err
	}

	// Load the root module's deps
	err = mdr.CompareModToSum()
	if err != nil {
		return err
	}

	return nil
}
