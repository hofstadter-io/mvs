package mod

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v3"
)

func Load(dir string) (ModSet, []error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		_, fn := path.Split(file.Name())
		if path.Ext(fn) == ".mod" {
			lang := path.Base(fn)
			fmt.Printf("Processing %s modules\n", lang)
			lmod, err := LoadRecurse(lang, dir)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			set[lang] = lmod
		}
	}
}

func LoadLocal(dir string) (ModSet, []error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return err
	}

	set := make(map[string]*Module)
	errs := make([]error, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		_, fn := path.Split(file.Name())
		if path.Ext(fn) == ".mod" {
			lang := path.Base(fn)
			fmt.Printf("Processing %s modules\n", lang)
			lmod, err := LoadModule(lang, dir)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			set[lang] = lmod
		}
	}

	return set, errs
}

func LoadRecurse(lang, dir string) (ModSet, []error) {

}

func LoadModule(lang, dir string) (*Module, error) {
	// XXX TEMP yaml this file
	modFn := lang + ".mod.yaml"
	sumFn := lang + ".sum.yaml"

	var modMod *Module
	modBytes, err := ioutil.ReadFile(path.Join(dir, modFn))
	if err != nil {
		return nil, nil, err
	} else {
		// TODO, replace this with a parser
		yerr := yaml.Unmarshal(modBytes, modMod)
		if yerr != nil {
			return yerr
		}
	}

	sumBytes, err := ioutil.ReadFile(path.Join(dir, modFn))
	if err != nil {
		if ok := err.(os.PathError); !ok {
			return err
		} else {
			sumBytes = []byte{}
		}
	} else {
		// TODO, replace this with a parser
		var sumMod *Module
		yerr := yaml.Unmarshal(sumBytes, sumMod)
		if yerr != nil {
			return yerr
		}
		modMod.SumMod = sumMod
	}

	return modMod, nil
}
