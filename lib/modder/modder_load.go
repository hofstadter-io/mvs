package modder

import (
	"github.com/go-git/go-billy/v5/osfs"
)

/* Reads the module files relative to the supplied dir from local FS
- ModFile
- SumFile
- MappingFile
*/

func (mdr *Modder) LoadMinimalFromFS(dir string) error {
	// Load the root module
	err := mdr.LoadRootFromFS(".")
	if err != nil {
		return err
	}

	return nil
}

func (mdr *Modder) LoadIndexDepsFromFS(dir string) error {
	// Load the root module
	err := mdr.LoadRootFromFS(".")
	if err != nil {
		return err
	}

	// Load the root module's deps
	err = mdr.LoadRootDeps()
	if err != nil {
		return err
	}
	// Recurse

	return nil
}

func (mdr *Modder) LoadRootFromFS(dir string) error {
	// Shortcut for no load modules, forget the reason for no load...
	if mdr.NoLoad {
		return nil
	}

	// Initialize filesystem
	mdr.FS = osfs.New(dir)

	// Initialzie Module related fields
	mdr.module = &Module{
		FS: mdr.FS,
	}
	mdr.depsMap = map[string]*Module{}

	// Load module files
	var err error
	err = mdr.LoadRootModFile()
	if err != nil {
		return err
	}

	err = mdr.LoadRootSumFile()
	if err != nil {
		return err
	}

	err = mdr.LoadRootMappingsFile()
	if err != nil {
		return err
	}

	return nil
}

// Loads the root modules mod file
func (mdr *Modder) LoadRootModFile() error {
	fn := mdr.ModFile
	m := mdr.module

	err := m.LoadModFile(fn, false /* Do load replace directives! */)
	if err != nil {
		return err
	}

	m.Module = m.ModFile.Module.Mod.Path

	return nil
}

// Loads the root modules sum file
func (mdr *Modder) LoadRootSumFile() error {
	fn := mdr.SumFile
	m := mdr.module

	err := m.LoadSumFile(fn)
	if err != nil {
		return err
	}

	return nil
}

// Loads the root modules mapping file
func (mdr *Modder) LoadRootMappingsFile() error {
	fn := mdr.MappingFile
	m := mdr.module

	err := m.LoadMappingFile(fn)
	if err != nil {
		return err
	}

	return nil
}
