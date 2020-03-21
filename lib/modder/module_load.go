package modder

import (
	"fmt"
	"os"

	"github.com/hofstadter-io/mvs/lang/modfile"
	"github.com/hofstadter-io/mvs/lang/sumfile"
	"github.com/hofstadter-io/mvs/lib/util"
)

func (m *Module) LoadModFile(fn string) error {

	modBytes, err := util.BillyReadAll(fn, m.FS)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
			return err
		}
	} else {
		f, err := modfile.Parse(fn, modBytes, nil)
		if err != nil {
			return err
		}
		m.ModFile = f
		m.Language = f.Language.Name
		m.LangVer = f.Language.Version
		m.Module = f.Module.Mod.Path
		m.Version = f.Module.Mod.Version

		for _, req := range f.Require {
			r := Require{Path: req.Mod.Path, Version: req.Mod.Version}
			m.Require = append(m.Require, r)
		}
		for _, rep := range f.Replace {
			r := Replace{OldPath: rep.Old.Path, OldVersion: rep.Old.Version, NewPath: rep.New.Path, NewVersion: rep.New.Version}
			m.Replace = append(m.Replace, r)
		}

		err = m.MergeSelfDeps()
		if err != nil {
			return err
		}

	}

	return nil
}

func (m *Module) MergeSelfDeps() error {
	// Now merge self deps
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

func (m *Module) LoadSumFile(fn string) error {

	sumBytes, err := util.BillyReadAll(fn, m.FS)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
			return err
		}
	} else {
		sumMod, err := sumfile.ParseSum(sumBytes, fn)
		if err != nil {
			return err
		}
		m.SumFile = &sumMod
	}

	return nil
}
