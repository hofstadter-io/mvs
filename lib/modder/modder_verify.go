package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Verify() error {

	// Verify Command Override
	if len(mdr.CommandVerify) > 0 {
		for _, cmd := range mdr.CommandGraph {
			out, err := util.Exec(cmd)
			fmt.Println(out)
			if err != nil {
				return err
			}
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

	valid := true

	// Load minimal root module
	err := mdr.LoadMetaFromFS(".")
	if err != nil {
		return err
	}

	// Load the root module's deps
	present, missing, err := mdr.FindPresentMissingInSum()
	if err != nil {
		return err
	}

	// Invalid if there are missing deps
	if len(missing) > 0 {
		valid = false
	}

	for _, m := range missing {
		R := mdr.module.SelfDeps[m]
		fmt.Printf("Sumfile missing: %s@%s\n", R.NewPath, R.NewVersion)
		err := fmt.Errorf("Sumfile missing: %s@%s", R.NewPath, R.NewVersion)
		mdr.errors = append(mdr.errors, err)
	}

	fmt.Println("Present\n-----------------")
	for _, p := range present {
		R := mdr.module.SelfDeps[p]
		fmt.Println(R.NewPath, R.NewVersion)
	}
	fmt.Println("-----------------")


	if !valid {

		return fmt.Errorf("Vendoring is in an inconsistent state, please run 'mvs vendor %s' ", mdr.Name)
	}
	return nil
}
