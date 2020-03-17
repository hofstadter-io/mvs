package mod

import (
	"fmt"
	"io/ioutil"
	"path"
)

func Load(dir string) (map[string]ModSet, []error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, []error{err}
	}

	sets := make(map[string]ModSet)
	errs := make([]error, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		_, fn := path.Split(file.Name())
		if path.Ext(fn) == ".mod" {
			lang := path.Base(fn)
			fmt.Printf("Processing %s modules\n", lang)
			setL, errsL := LoadRecurse(lang, dir)
			if errsL != nil {
				errs = append(errs, errsL...)
				continue
			}

			sets[lang] = setL
		}
	}
	return sets, errs
}

func LoadRecurse(lang, dir string) (ModSet, []error) {

	return nil, nil
}

func LoadLocal(dir string) (map[string]ModSet, []error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, []error{err}
	}

	sets := make(map[string]ModSet)
	errs := make([]error, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		/*
		_, fn := path.Split(file.Name())
		if path.Ext(fn) == ".mod" {
			lang := path.Base(fn)
			fmt.Printf("Processing %s modules\n", lang)
			modL, err := LoadModule(lang, dir)
			if err != nil {
				errs = append(errs, err)
				continue
			}

			setL := ModSet{
				modL.Module: modL,
			}
			sets[lang] = setL
		}
		*/
	}

	return sets, errs
}

/*
func LoadModSet(lang, dir string) (ModSet, error) {
	// XXX TEMP yaml this file
	modFn := lang + ".mod.yaml"
	sumFn := lang + ".sum.yaml"

	var modMod *Module
	modBytes, err := ioutil.ReadFile(path.Join(dir, modFn))
	if err != nil {
		return nil, err
	} else {
		// TODO, replace this with a parser
		yerr := yaml.Unmarshal(modBytes, modMod)
		if yerr != nil {
			return nil, yerr
		}
	}

	sumBytes, err := ioutil.ReadFile(path.Join(dir, modFn))
	if err != nil {
		if ok := err.(*os.PathError); !ok {
			return nil, err
		} else {
			sumBytes = []byte{}
		}
	} else {
		// TODO, replace this with a parser
		var sumMod *Module
		yerr := yaml.Unmarshal(sumBytes, sumMod)
		if yerr != nil {
			return nil, yerr
		}
		modMod.SumMod = sumMod
	}

	return modMod, nil
}
*/

