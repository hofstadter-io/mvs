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

	fmt.Println("=====  Root  =====")

	for path, R := range mod.SelfDeps {

		// Missing sum file, need to fetch (cache -> internet)
		if sf == nil {
			// fmt.Printf("missing mod file, fetch %s %#+v\n", path, R)

			// Local REPLACE
			if strings.HasPrefix(R.NewPath, "./") || strings.HasPrefix(R.NewPath, "../") {
				fmt.Println("Local replace:", path)
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

				err = m.LoadMetaFiles(mdr.ModFile, mdr.SumFile, mdr.MappingFile, true /* ignoreReplace directives */)
				if err != nil {
					return err
				}

				err = mdr.MvsMergeDependency(m)
				if err != nil {
					return err
				}

				fmt.Println("  root loaded local module: %s", m.Module, m.Version, m.ReplaceModule, m.ReplaceVersion)

				continue
			}

			// TODO Lookup in cache

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

			err = m.LoadMetaFiles(mdr.ModFile, mdr.SumFile, mdr.MappingFile, true /* ignoreReplace directives */)
			if err != nil {
				return err
			}

			err = mdr.MvsMergeDependency(m)
			if err != nil {
				return err
			}

			fmt.Println("  root loaded remote module: %s", m.Module, m.Version, m.ReplaceModule, m.ReplaceVersion)
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

	fmt.Println("==== END ROOT ====")
	return nil
}

func (mdr *Modder) CheckAndFetchDepsDeps(deps map[string]*Module) (map[string]*Module, error) {
	// var err error
	// mod := mdr.module
	// sf := mod.SumFile

	fmt.Println("=====  Deps  =====")

	newDeps := map[string]*Module{}
	for path, M := range deps {

		fmt.Println("  dep module loading other:", mdr.module.Module, mdr.module.Version, mdr.module.ReplaceModule, mdr.module.ReplaceVersion)
		fmt.Println("     ", path, M.Module, M.Version, M.ReplaceModule, M.ReplaceVersion)

		for _, dep := range M.SelfDeps {
			fmt.Println("        ", dep.NewPath, dep.NewVersion)
			// Don't add the root module to the dependencies
			if mdr.module.Module == dep.NewPath || mdr.module.Module == dep.OldPath {
				continue
				// return nil
			}

			/*
				// TODO shortcut with sum and/or mapping files
				if sf == nil {
					fmt.Printf("missing mod file, fetch %s %#+v\n", path, R)
				}
			*/

			mod, ok := mdr.depsMap[dep.NewPath]
			// XXX Not sure why this doesn't have a version...
			// See change in module_load.go lines 28:29
			// Found an exiting dep, check versions
			if ok {
				fmt.Println("     found", mod.Module, mod.Version, mod.ReplaceModule, mod.ReplaceVersion)
				// Is the existing a replace already?
				if mod.ReplaceModule != "" {
					fmt.Println("     SKIPPING, replaced already by root module")
					continue
				}

				// is the current version equal or newer than the incoming version
				if semver.Compare(mod.Version, dep.NewVersion) >= 0 {
					fmt.Println("     FETCH UPDATE", mod.Version, dep.NewVersion)
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

			err = m.LoadMetaFiles(mdr.ModFile, mdr.SumFile, mdr.MappingFile, true /* ginoreReplace directives */)
			if err != nil {
				return newDeps, err
			}

			// we already checked the semver, but this should
			err = mdr.MvsMergeDependency(m)
			if err != nil {
				return newDeps, err
			}

			newDeps[m.Module] = m

		}

	}

	fmt.Println("==== End Deps ====")
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

func (mdr *Modder) FindPresentMissingInSum() ([]string, []string, error) {
	present := []string{}
	missing := []string{}

	mod := mdr.module
	sf := mod.SumFile
	if sf == nil {
		return nil, nil, fmt.Errorf("No sum file %q for %s, run 'mvs vendor [%s]' to fix.", mdr.SumFile, mdr.Name, mdr.Name)
	}

	for path, R := range mod.SelfDeps {
		ver := sumfile.Version{
			Path:    path,
			Version: R.NewVersion,
		}

		_, ok := sf.Mods[ver]
		if ok {
			present = append(present, path)
		} else {
			missing = append(missing, path)
		}
	}

	return present, missing, nil
}

func (mdr *Modder) CompareSumToVendor() error {

	return nil
}
