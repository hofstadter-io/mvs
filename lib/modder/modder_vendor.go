package modder

import (
	"fmt"
	"strings"

	"github.com/hofstadter-io/mvs/lib/repos/git"
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
	// Vendor Command Override
	if len(mdr.CommandVendor) > 0 {
		out, err := util.Exec(mdr.CommandVendor)
		fmt.Println(out)
		return err
	}

	// Otherwise, MVS venodiring
	err := mdr.VendorMVS()
	if err != nil {
		mdr.PrintErrors()
		return err
	}

	return nil
}

// The entrypoint to the MVS internal vendoring process
func (mdr *Modder) VendorMVS() error {

	// Load the current module
	err := mdr.LoadModuleFromFS(".")
	if err != nil {
		return err
	}

	// XXX This is where we need to start changing behavior
	// right now, the below just loads and clones without much intelligence
	// Want to go dep by dep, performing the same checks

	// mdr.PrintSelfDeps()
	err = mdr.LoadSelfDeps()
	if err != nil {
		return err
	}

	// XXX OLD BELOW

	err = mdr.LoadRequires()
	if err != nil {
		return err
	}

	err = mdr.LoadReplaces()
	if err != nil {
		return err
	}

	return mdr.WriteVendor()
}

func (mdr *Modder) LoadRequires() error {
	// TODO Check if already good

	// TODO Check mvs system cache in $HOME/.mvs/cache

	// First process the require directives
	for _, req := range mdr.module.Require {
		ref, refs, err := git.IndexGitRemote(req.Path, req.Version)
		if err != nil {
			// Build up errors
			mdr.module.Errors = append(mdr.module.Errors, err)
			continue
		}

		// TODO Later, after any real clone, during dep recursion or vendoring,
		// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
		m := &Module{
			Module:  req.Path,
			Version: req.Version,
			Ref:     ref,
			Refs:    refs,
		}
		// is this module already in the map
		//   a) from replaces
		//   b) duplicate entry
		//   c) if not replace, greater version required? (we eventually want the minimum download, given the maximum required)
		if _, ok := mdr.depsMap[m.Module]; ok {
			return fmt.Errorf("Dependency %q required twice", m.Module)
		}

		// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions
		mdr.addDependency(m)

		// WORK SECTION ========================

	}

	return nil
}

func (mdr *Modder) LoadReplaces() error {
	// Now replace any dependencies, and possibly add new ones (user not required to specify both orig require and replace source)
	for _, rep := range mdr.module.Replace {

		// Handle local replaces
		if strings.HasPrefix(rep.NewPath, "./") || strings.HasPrefix(rep.NewPath, "../") {
			fmt.Println("Local replace:", rep)
			m := &Module{
				// TODO Think about Replace syntax options and the existence of git
				// XXX  How does go mod handle this question
				Module:         rep.OldPath,
				Version:        rep.OldVersion,
				ReplaceModule:  rep.NewPath,
				ReplaceVersion: rep.NewVersion,
			}

			// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions
			mdr.addDependency(m)
			continue
		}

		// Handle remote replaces

		// Update version if needed
		orig, ok := mdr.depsMap[rep.OldPath]
		// TODO Think about this some more with versions and the existence of git
		// XXX  How does go mod handle this question
		if ok {
			if rep.OldVersion == "" {
				rep.OldVersion = orig.Version
			}
		}

		// pretty normal module dep handling now
		ref, refs, err := git.IndexGitRemote(rep.NewPath, rep.NewVersion)
		if err != nil {
			// Build up errors
			mdr.module.Errors = append(mdr.module.Errors, err)
			continue
		}

		// TODO Later, after any real clone, during dep recursion or vendoring,
		// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
		m := &Module{
			Module:         rep.OldPath,
			Version:        rep.OldVersion,
			ReplaceModule:  rep.NewPath,
			ReplaceVersion: rep.NewVersion,
			Ref:            ref,
			Refs:           refs,
		}

		// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions
		mdr.addDependency(m)
	}

	return nil

	// now process the requirements, should skip any that exist already because they are replaces

}
