package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Status() error {

	// Status Command Override
	if len(mdr.CommandStatus) > 0 {
		out, err := util.Exec(mdr.CommandStatus)
		fmt.Println(out)
		if err != nil {
			return err
		}
	} else {
		// Otherwise, MVS venodiring
		err := mdr.StatusMVS()
		if err != nil {
			mdr.PrintErrors()
			return err
		}
	}

	return nil
}

// The entrypoint to the MVS internal verify process
func (mdr *Modder) StatusMVS() error {
	var err error

	// Load minimal root module
	err = mdr.LoadMinimalFromFS(".")
	if err != nil {
		return err
	}


	fmt.Println("==================")

	mod := mdr.module
	sf := mod.SumFile

	err = mod.PrintSelfDeps()
	if err != nil {
		return err
	}

	fmt.Println("==================")

	if sf != nil {
		err = sf.Print()
		if err != nil {
			return err
		}

	} else {
		fmt.Printf("No sum file %q found for lang %q\n", mdr.SumFile, mdr.Name)
	}
	fmt.Println("==================")

	return nil
}
