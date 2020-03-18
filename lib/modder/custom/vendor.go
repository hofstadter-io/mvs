package custom

import (
	"fmt"
	"os"
	"path"

	"github.com/hofstadter-io/mvs/lib/mod"
	"github.com/hofstadter-io/mvs/lib/modder/common"
	"github.com/hofstadter-io/mvs/lib/remote/git"
	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Vendor() error {
	m, err := mdr.LoadModule(".")
	if err != nil {
		return err
	}

	mdr.module = m
	fmt.Printf("Root Module: %v\n", m)
	mdr.depsMap = map[string]*mod.Module{}

	// TODO START par.Work here

	// TODO, build up errors
	// TODO merge require / replace
	// TODO compare to sum file
	for _, req := range mdr.module.Require {
		ref, refs, err := common.IndexGit(req.Path, req.Version)
		if err != nil {
			return err
		}

		fmt.Println(" ", req.Path, req.Version, ref)

		// TODO Later, after any real clone, during dep recursion or vendoring,
		// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
		m := &mod.Module{
			Module:  req.Path,
			Version: req.Version,
			Ref:     ref,
			Refs:    refs,
		}

		// save module to depsMap
		mdr.depsMap[req.Path] = m

		// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions

		// WORK SECTION ========================

		// clone the module and load
		clone, err := git.CloneRepoRef(req.Path, ref)
		m.Clone = clone
		if err != nil {
			return err
		}

		/*
			files, err := clone.FS.ReadDir("/")
			if err != nil { return err }

			for _, file := range files {
				fmt.Println("  -", file.Name())
			}
		*/

		// load dep module
		// dm, err := mdr.LoadModule("...")
		// if err != nil { return err }

		// WORK SECTION ========================
	}

	// XXX For Now...

	// Copyout based on config

	// XXX Actually want to...

	// TODO recurse here

	// XXX TEMP FINALLY, write out any vendor
	return mdr.writeVendor()
}

func (mdr *Modder) writeVendor() error {
	// TODO calc and update imported module "hashes"" here

	// make vendor dir if not present
	err := os.MkdirAll(mdr.ModsDir, 0755)
	if err != nil {
		return err
	}

	// write out each dep
	for _, m := range mdr.depsMap {
		baseDir := path.Join(mdr.ModsDir, m.Module)

		err = util.BillyCopyDir(baseDir, "/", m.Clone.FS)

	}

	return nil
}
