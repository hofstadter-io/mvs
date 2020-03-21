package modder

import (
	"fmt"
	"strings"

	"github.com/hofstadter-io/mvs/lib/repos/git"
)


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

