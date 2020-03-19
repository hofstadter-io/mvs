package custom

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hofstadter-io/mvs/lib/mod"
	"github.com/hofstadter-io/mvs/lib/modder/common"
	"github.com/hofstadter-io/mvs/lib/remote/git"
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
	M, err := mdr.LoadModule(".")
	if err != nil {
		return err
	}

	mdr.module = M
	fmt.Printf("Root Module: %v\n", M)
	mdr.depsMap = map[string]*mod.Module{}

	// TODO START par.Work here

	// TODO compare to sum file

	// First process the require directives
	for _, req := range mdr.module.Require {
		ref, refs, err := common.IndexGit(req.Path, req.Version)
		if err != nil {
			// Build up errors
			mdr.module.Errors = append(mdr.module.Errors, err)
			continue
		}

		// TODO Later, after any real clone, during dep recursion or vendoring,
		// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
		m := &mod.Module{
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

	// Now replace any dependencies, and possibly add new ones (user not required to specify both orig require and replace source)
	for _, rep := range mdr.module.Replace {

		// Handle local replaces
		if strings.HasPrefix(rep.NewPath, "./") || strings.HasPrefix(rep.NewPath, "../") {
			fmt.Println("Local replace:", rep)
			m := &mod.Module{
				Module:  rep.OldPath,
				Version: rep.OldVersion,
				ReplaceModule:  rep.NewPath,
				ReplaceVersion: rep.NewVersion,
			}

			// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions
			mdr.addDependency(m)
			continue
		}

		// Handle remote replaces

		// Update version if needed
		orig, ok := mdr.depsMap[rep.OldPath];
		if ok {
			if rep.OldVersion == "" {
				rep.OldVersion = orig.Version
			}
		}

		// pretty normal module dep handling now
		ref, refs, err := common.IndexGit(rep.NewPath, rep.NewVersion)
		if err != nil {
			// Build up errors
			mdr.module.Errors = append(mdr.module.Errors, err)
			continue
		}

		// TODO Later, after any real clone, during dep recursion or vendoring,
		// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
		m := &mod.Module{
			Module:  rep.OldPath,
			Version: rep.OldVersion,
			ReplaceModule:  rep.NewPath,
			ReplaceVersion: rep.NewVersion,
			Ref:     ref,
			Refs:    refs,
		}

		// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions
		mdr.addDependency(m)
	}

	// now process the requirements, should skip any that exist already because they are replaces

	err = mdr.checkPrintErrors()
	if err != nil { return err }

	// XXX Actually want to recurse here
	// XXX for now, write out any vendor
	return mdr.writeVendor()
}

func (mdr *Modder) checkPrintErrors() error {
	var wasError error

	if len(mdr.module.Errors) > 0 {
		wasError = fmt.Errorf("Exiting due to errors during vendoring.")
		for _, err := range mdr.module.Errors {
			fmt.Println(err)
		}
	}

	for _, dep := range mdr.depsMap {
		if len(dep.Errors) > 0 {
			if wasError != nil {
				wasError = fmt.Errorf("Exiting due to errors during vendoring.")
			}
			for _, err := range dep.Errors {
				fmt.Println(err)
			}
		}
	}

	return wasError
}

func (mdr *Modder) addDependency(m *mod.Module) error {

		// save module to depsMap
		mdr.depsMap[m.Module] = m

		// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions

		// Should only happen with local replace right now
		if m.Ref == nil {
			clone, err := git.CloneLocalRepo(m.ReplaceModule)
			m.Clone = clone
			if err != nil {
				return err
			}
			return nil
		}

		// clone the module and load
		clone, err := git.CloneRepoRef(m.Module, m.Ref)
		m.Clone = clone
		if err != nil {
			return err
		}

		// Pushi into parallel worker here?
		// load dep module
		// dm, err := mdr.LoadModule("...")
		// if err != nil { return err }

	return nil
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
		if err != nil {
		  return err
		}

	}

	return nil
}
