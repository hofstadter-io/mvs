package modder

import (
	"fmt"
	"strings"

	"github.com/go-git/go-billy/v5/osfs"

	"github.com/hofstadter-io/mvs/lib/parse/sumfile"
)

func (mdr *Modder) CheckAndFetchRootDeps() error {
	// var err error
	mod := mdr.module
	sf := mod.SumFile

	fmt.Println("==================")

	for path, R := range mod.SelfDeps {
		if sf == nil {
			fmt.Printf("missing in mod file, fetch %s %#+v\n", path, R)

			if strings.HasPrefix(R.NewPath, "./") || strings.HasPrefix(R.NewPath, "../") {
				fmt.Println("Local replace:", R)
				m := &Module{
					// TODO Think about Replace syntax options and the existence of git
					// XXX  How does go mod handle this question
					Module:         R.OldPath,
					Version:        R.OldVersion,
					ReplaceModule:  R.NewPath,
					ReplaceVersion: R.NewVersion,
				}

				m.FS = osfs.New(R.NewPath)

				err := mdr.MvsMergeDependency(m)
				if err != nil {
					return err
				}

				fmt.Printf("  module: %#+v\n", m)
			}

			continue
		}

		ver := sumfile.Version {
			Path: path,
			Version: R.NewVersion,
		}

		v, ok := sf.Mods[ver]
		if !ok {
			// TODO fetch missing dep
			fmt.Println("fetch missing mod->sum", ver, R, v)
		} else {
			fmt.Println("match", ver, R, v)
			// TODO check mapping and vendor
			// verify contents are what is expected

			// TODO if not verify, fetch
		}

	}

	fmt.Println("==================")
	return nil
}

func (mdr *Modder) CompareModToSum() error {
	// var err error
	mod := mdr.module
	sf := mod.SumFile
	if sf == nil {
		return fmt.Errorf("No sum file %q for %s, run 'mvs vendor [%s]' to fix.", mdr.SumFile, mdr.Name, mdr.Name)
	}

	fmt.Println("==================")

	for path, R := range mod.SelfDeps {
		ver := sumfile.Version {
			Path: path,
			Version: R.NewVersion,
		}

		v, ok := sf.Mods[ver]
		if !ok {
			return fmt.Errorf("Mismatch between sum and mod files, run 'mvs vendor [%s]' to fix.", mdr.Name)
		} else {
			fmt.Println("match", ver, R, v)
			// TODO check mapping and vendor
			// verify contents are what is expected
		}

	}

	fmt.Println("==================")
	return nil
}

func (mdr *Modder) CompareSumToVendor() error {


	return nil
}
