package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

/* Vendor reads in a module, determines dependencies, and writes out the vendor folder.

Will there be infinite recursion, or maybe just two levels?


This module & deps
	1) load this module
		a) check sum <-> mod files, need to determine if any updates here, for now w/o sum file
	2) process its sum/mod files
	3) for each of this mods deps dependency
	  a) fetch all refs
		b) find minimum
		c) add to depMap, when / if ... ? guessing how right now
		  - replaces are processed first
			- requires are processed second, so only add if not there, we shouldn't have duplicates in owr own mods files
			-
	  d) if added, clond the desired ref to memory

	4) Now loop over depMap to pull in secondary dependencies
	  - probably want to create a "newDeps" map here if we need to support wider recursion
		- basically follow the last block, but load idependently and merge after
		- do we need a separate modder when we process each dep?
		  - probably if we are going to enable each module to optionally specify local behavior
		- so first file we should read is the .mvsconfig, that maps <lang> to whatever

	F) Finally, write out the vendor directory
	  a) check <vendor-dir>/modules.txt and checksums
		b) write out if necessary
*/
func (mdr *Modder) Vendor() error {
	// TODO, run pre vendor commands here

	// Vendor Command Override
	if len(mdr.CommandVendor) > 0 {
		for _, cmd := range mdr.CommandGraph {
			out, err := util.Exec(cmd)
			fmt.Println(out)
			if err != nil {
				return err
			}
		}
	} else {
		// Otherwise, MVS venodiring
		err := mdr.VendorMVS()
		if err != nil {
			mdr.PrintErrors()
			return err
		}
	}

	// TODO, run post vendor commands here

	return nil
}

// The entrypoint to the MVS internal vendoring process
func (mdr *Modder) VendorMVS() error {
	var err error

	// Load minimal root module
	err = mdr.LoadMinimalFromFS(".")
	if err != nil {
		fmt.Println(err)
		// return err
	}

	err = mdr.CheckAndFetchRootDeps()
	if err != nil {
		fmt.Println(err)
		return err
	}

	for {
		deps, err := mdr.CheckAndFetchDepsDeps(mdr.depsMap)
		if err != nil {
			fmt.Println(err)
			return err
		}

		if len(deps) == 0 {
			break
		}
	}

	err = mdr.WriteVendor()
	if err != nil {
		return err
	}

	return nil
}
