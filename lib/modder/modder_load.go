package modder

import (
	"github.com/go-git/go-billy/v5/osfs"
)

/* Reads the module files relative to the supplied dir from local FS
- ModFile
- SumFile
- MappingFile
*/
func (mdr *Modder) LoadModuleFromFS(dir string) error {
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
	err = mdr.LoadModFile()
	if err != nil {
		return err
	}

	err = mdr.LoadSumFile()
	if err != nil {
		return err
	}

	return nil
}

// Loads the root modules mod file
func (mdr *Modder) LoadModFile() error {
	fn := mdr.ModFile
	m := mdr.module

	err := m.LoadModFile(fn)
	if err != nil {
		return err
	}

	return nil
}

// Loads the root modules sum file
func (mdr *Modder) LoadSumFile() error {
	fn := mdr.SumFile
	m := mdr.module

	err := m.LoadSumFile(fn)
	if err != nil {
		return err
	}

	return nil
}
