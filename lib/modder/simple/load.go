package simple

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v3"

	"github.com/hofstadter-io/mvs/lang/modfile"
	"github.com/hofstadter-io/mvs/lib/mod"
)

func (m *Modder) Load(dir string) error {
	mdr, err := m.LoadModule(dir)
	if err != nil {
		return err
	}
	m.module = mdr

	content, err := yaml.Marshal(mdr)
	if err != nil {
		return err
	}

	fmt.Printf("Module Contents:\n%s\n", string(content))

	return nil
}

func (m *Modder)LoadModule(dir string) (*mod.Module, error) {

	// XXX TEMP yaml this file
	modFn := m.ModFile
	sumFn := m.SumFile

	var modMod mod.Module
	modBytes, err := ioutil.ReadFile(path.Join(dir, modFn))
	if err != nil {
		return nil, err
	} else {
		f, err := modfile.Parse(modFn, modBytes, nil)
		if err != nil {
			return nil, err
		}
		modMod.Language = f.Language.Name
		modMod.LangVer = f.Language.Version
		modMod.Module = f.Module.Mod.Path
		modMod.Version = f.Module.Mod.Version
		for _, req := range f.Require {
			m := mod.Require{Path: req.Mod.Path, Version: req.Mod.Version}
			modMod.Require = append(modMod.Require, m)
		}
		for _, rep := range f.Replace {
			m := mod.Replace{OldPath: rep.Old.Path, OldVersion: rep.Old.Version, NewPath: rep.New.Path, NewVersion: rep.New.Version}
			modMod.Replace = append(modMod.Replace, m)
		}
	}

	sumBytes, err := ioutil.ReadFile(path.Join(dir, sumFn))
	if err != nil {
		if _, ok := err.(*os.PathError); !ok {
			return nil, err
		} else {
			sumBytes = []byte{}
		}
	} else {
		// TODO, replace this with a parser
		var sumMod mod.Module
		yerr := yaml.Unmarshal(sumBytes, &sumMod)
		if yerr != nil {
			return nil, yerr
		}
		modMod.SumMod = &sumMod
	}

	return &modMod, nil
}
