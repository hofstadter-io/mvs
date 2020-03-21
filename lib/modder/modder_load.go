package modder

import (
  "fmt"

	"github.com/go-git/go-billy/v5/osfs"
)

/* Reads the module files in
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

	// Now merge self deps
	m := mdr.module
	m.SelfDeps = map[string]Replace{}
	for _, req := range m.Require {
		if _, ok := m.SelfDeps[req.Path]; ok {
			return fmt.Errorf("Dependency %q required twice in %q", req.Path, m.Module)
		}
		m.SelfDeps[req.Path] = Replace {
			NewPath: req.Path,
			NewVersion: req.Version,
		}
	}

	dblReplace := map[string]Replace{}
	for _, rep := range m.Replace {
		// Check if replaced twice
		if _, ok := dblReplace[rep.OldPath]; ok {
			return fmt.Errorf("Dependency %q replaced twice in %q", rep.OldPath, m.Module)
		}
		dblReplace[rep.OldPath] = rep

		// Pull in require info if not in replace
		if req, ok := m.SelfDeps[rep.OldPath]; ok {
			if rep.OldVersion == "" {
				rep.OldVersion = req.NewVersion
			}
		}
		m.SelfDeps[rep.OldPath] = rep
	}


	return nil
}

func (mdr *Modder) LoadModFile() error {
	fn := mdr.ModFile
	m := mdr.module

	err := m.LoadModFile(fn)
	if err != nil {
		return err
	}

	return nil
}

func (mdr *Modder) LoadSumFile() error {
	fn := mdr.SumFile
	m := mdr.module

	err := m.LoadSumFile(fn)
	if err != nil {
		return err
	}

	return nil
}
