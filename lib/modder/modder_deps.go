package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/repos/git"
)

func (mdr *Modder) PrintRootDeps() error {
	fmt.Println("Root module self deps for", mdr.module.Module)
	err := mdr.module.PrintSelfDeps()
	if err != nil {
		return err
	}

	return nil
}

func (mdr *Modder) LoadRootDeps() error {
	fmt.Println("Loading self deps for", mdr.module.Module)

	err := mdr.module.LoadSelfDeps()
	if err != nil {
		return err
	}
	return nil
}

// This sets or overwrites the module
func (mdr *Modder) ReplaceDependency(m *Module) error {
	// Don't add the root module to the dependencies
	if mdr.module.Module == m.Module {
		return nil
	}
	// save module to depsMap, that's it? (yes)
	mdr.depsMap[m.Module] = m

	return nil
}

// If not set, justs adds. If set, takes the one with the greater version.
func (mdr *Modder) MvsMergeDependency(m *Module) error {
	// Don't add the root module to the dependencies
	if mdr.module.Module == m.Module {
		return nil
	}

	// TODO check for existing module, version comparison
	mdr.depsMap[m.Module] = m

	return nil
}

// TODO, break this function appart
func (mdr *Modder) addDependency(m *Module) error {
	// Don't add the root module to the dependencies
	if mdr.module.Module == m.Module {
		return nil
	}
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
