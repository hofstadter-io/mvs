package modder

import (
	"fmt"
	"strings"

	"github.com/go-git/go-billy/v5/osfs"
	"golang.org/x/mod/semver"

	"github.com/hofstadter-io/mvs/lib/parse/sumfile"
	"github.com/hofstadter-io/mvs/lib/repos/git"
)

func (mdr *Modder) CheckAndFetchRootDeps() error {
	// var err error
	mod := mdr.module
	sf := mod.SumFile

	// fmt.Println("=====  Root  =====")

	for path, R := range mod.SelfDeps {
		if sf == nil {
			// fmt.Printf("missing mod file, fetch %s %#+v\n", path, R)

			// Local REPLACE
			if strings.HasPrefix(R.NewPath, "./") || strings.HasPrefix(R.NewPath, "../") {
				// fmt.Println("Local replace:", R)
				m := &Module{
					// TODO Think about Replace syntax options and the existence of git
					// XXX  How does go mod handle this question
					Module:         R.OldPath,
					Version:        R.OldVersion,
					ReplaceModule:  R.NewPath,
					ReplaceVersion: R.NewVersion,
				}

				m.FS = osfs.New(R.NewPath)

				var err error

				err = m.LoadModFile(mdr.ModFile)
				if err != nil {
					return err
				}

				err = m.LoadSumFile(mdr.SumFile)
				if err != nil {
					// fmt.Println(err)
					// return err
				}

				err = m.LoadMappingFile(mdr.MappingFile)
				if err != nil {
					// fmt.Println(err)
					// return err
				}

				err = mdr.MvsMergeDependency(m)
				if err != nil {
					return err
				}

				// fmt.Printf("  module: %#+v\n", m)

				continue
			}

			// HANDLE remote and non-local replace the same way
			ref, refs, err := git.IndexGitRemote(R.NewPath, R.NewVersion)
			if err != nil {
				// Build up errors
				mod.Errors = append(mod.Errors, err)
				continue
			}

			// TODO Later, after any real clone, during dep recursion or vendoring,
			// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
			m := &Module{
				Module:  R.NewPath,
				Version: R.NewVersion,
				Ref:     ref,
				Refs:    refs,
			}
			// is this module already in the map
			//   a) from replaces
			//   b) duplicate entry
			//   c) if not replace, greater version required? (we eventually want the minimum download, given the maximum required)
			if _, ok := mdr.depsMap[m.Module]; ok {
				return fmt.Errorf("Dependency %q required twice", m.Module)
			}

			clone, err := git.CloneRepoRef(m.Module, m.Ref)
			m.Clone = clone
			if err != nil {
				return err
			}
			m.FS = m.Clone.FS

			// TODO load the modules .mvsconfig if present

			err = m.LoadModFile(mdr.ModFile)
			if err != nil {
				return err
			}

			err = m.LoadSumFile(mdr.SumFile)
			if err != nil {
				// fmt.Println(err)
				// return err
			}

			err = m.LoadMappingFile(mdr.MappingFile)
			if err != nil {
				// fmt.Println(err)
				// return err
			}

			err = mdr.MvsMergeDependency(m)
			if err != nil {
				return err
			}

			// fmt.Printf("  module: %#+v\n", m)

			continue
		}

		ver := sumfile.Version{
			Path:    path,
			Version: R.NewVersion,
		}

		_, ok := sf.Mods[ver]
		if !ok {
			// TODO fetch missing dep
			// fmt.Println("fetch missing mod->sum", ver, R, v)
		} else {
			// fmt.Println("match", ver, R, v)
			// TODO check mapping and vendor
			// verify contents are what is expected

			// TODO if not verify, fetch
		}

	}

	// fmt.Println("==================")
	return nil
}

func (mdr *Modder) CheckAndFetchDepsDeps(deps map[string]*Module) (map[string]*Module, error) {
	// var err error
	// mod := mdr.module
	// sf := mod.SumFile

	fmt.Println("=====  Deps  =====")

	newDeps := map[string]*Module{}
	for _, M := range deps {
		// fmt.Printf("  %s  %#+v\n", path, M)

		for _, dep := range M.SelfDeps {
			fmt.Println("    ", dep.NewPath, dep.NewVersion)
			/*
				// TODO shortcut with sum and/or mapping files
				if sf == nil {
					fmt.Printf("missing mod file, fetch %s %#+v\n", path, R)
				}
			*/

			mod, ok := mdr.depsMap[dep.NewPath]
			// XXX Not sure why this doesn't have a version...
			// See change in module_load.go lines 28:29
			// fmt.Printf("     found %#+v\n", mod)
			// Found an exiting dep, check versions
			if ok {
				// is the current version equal or newer than the incoming version
				if semver.Compare(mod.Version, dep.NewVersion) >= 0 {
					fmt.Println("     FETCH UPDATE", dep.NewVersion)
					continue
				}
			}

			// else fetch a new dependency
			// HANDLE remote and non-local replace the same way
			ref, refs, err := git.IndexGitRemote(dep.NewPath, dep.NewVersion)
			if err != nil {
				// Build up errors
				M.Errors = append(M.Errors, err)
				continue
			}

			// TODO Later, after any real clone, during dep recursion or vendoring,
			// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
			m := &Module{
				Module:  dep.NewPath,
				Version: dep.NewVersion,
				Ref:     ref,
				Refs:    refs,
			}

			clone, err := git.CloneRepoRef(m.Module, m.Ref)
			m.Clone = clone
			if err != nil {
				return newDeps, err
			}
			m.FS = m.Clone.FS

			// TODO load the modules .mvsconfig if present

			err = m.LoadModFile(mdr.ModFile)
			if err != nil {
				return newDeps, err
			}

			err = m.LoadSumFile(mdr.SumFile)
			if err != nil {
				// fmt.Println(err)
				// return err
			}

			err = m.LoadMappingFile(mdr.MappingFile)
			if err != nil {
				// fmt.Println(err)
				// return err
			}

			// we already checked the semver, but this should
			err = mdr.MvsMergeDependency(m)
			if err != nil {
				return newDeps, err
			}

			newDeps[m.Module] = m

		}

	}

	fmt.Println("==================")
	return newDeps, nil
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
		ver := sumfile.Version{
			Path:    path,
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
