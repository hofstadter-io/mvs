package modder

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hofstadter-io/mvs/lib/repos/git"
	"github.com/hofstadter-io/mvs/lib/util"
)

var (
	// Common files to copy from modules, also includes the .md version of the filename
	definiteVendors = []string{
		"README",
		"LICENSE",
		"PATENTS",
		"CONTRIBUTORS",
		"SECURITY",
	}

	// cross product these endings
	endings = []string{"", ".md", ".txt"}
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

func (mdr *Modder) PrintSelfDeps() error {
	fmt.Println("Merged self deps for", mdr.module.Module)
	for path, R := range mdr.module.SelfDeps {
		fmt.Println("   ", path, "~", R.OldPath, R.OldVersion, "=>", R.NewPath, R.NewVersion)
	}

	return nil
}

func (mdr *Modder) LoadSelfDeps() error {
	fmt.Println("Loading self deps for", mdr.module.Module)
	for path, R := range mdr.module.SelfDeps {
		fmt.Println("   ", path, "~", R.OldPath, R.OldVersion, "=>", R.NewPath, R.NewVersion)

		// XXX is this the right place for this?
		// TODO Check if already good (i.e. ??? if in vendor and ok)
		// TODO Check mvs system cache in $HOME/.mvs/cache

		// We probably need to start module creating here

		// Handle local replaces
		if strings.HasPrefix(R.NewPath, "./") || strings.HasPrefix(R.NewPath, "../") {
			fmt.Println("Local Replace:", R.OldPath, R.OldVersion, "=>", R.NewPath, R.NewVersion)
			// is it git or not?

			return nil
		}

		// OTHERWISE... it's a remote repository

		// is it git or a package repository? TBD

	}

	return nil
}

func (mdr *Modder) LoadRequires() error {
	// TODO Check if already good

	// TODO Check mvs system cache in $HOME/.mvs/cache

	// First process the require directives
	for _, req := range mdr.module.Require {
		ref, refs, err := IndexGit(req.Path, req.Version)
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
		ref, refs, err := IndexGit(rep.NewPath, rep.NewVersion)
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

// This sets or overwrites the module
func (mdr *Modder) ReplaceDependency(m *Module) error {
	// save module to depsMap, that's it? (yes)
	mdr.depsMap[m.Module] = m

	return nil
}

// If not set, justs adds. If set, takes the one with the greater version.
func (mdr *Modder) MergeDependency(m *Module) error {

	// TODO check for existing module, version comparison
	mdr.depsMap[m.Module] = m

	return nil
}

// TODO, break this function appart
func (mdr *Modder) addDependency(m *Module) error {
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

func (mdr *Modder) WriteVendor() error {
	// TODO calc and update imported module "hashes"" here

	// make vendor dir if not present
	err := os.MkdirAll(mdr.ModsDir, 0755)
	if err != nil {
		return err
	}

	// write out each dep
	for _, m := range mdr.depsMap {

		baseDir := path.Join(mdr.ModsDir, m.Module)

		fmt.Println("Copying", baseDir)

		// copy special files
		for _, fn := range definiteVendors {
			for _, end := range endings {
				_, err := m.Clone.FS.Stat(fn + end)
				if err != nil {
					if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
						// some other error
						return err
					}
					// not found
					continue
				}

				// Found one!
				err = util.BillyCopyFile(baseDir, "/"+fn+end, m.Clone.FS)

			}
		}

		if len(mdr.VendorIncludeGlobs) > 0 || len(mdr.VendorExcludeGlobs) > 0 {
			// Just copy everything
			err = util.BillyGlobCopy(baseDir, "/", m.Clone.FS, mdr.VendorIncludeGlobs, mdr.VendorExcludeGlobs)
			if err != nil {
				return err
			}

		} else {
			// Just copy everything
			err = util.BillyCopyDir(baseDir, "/", m.Clone.FS)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
