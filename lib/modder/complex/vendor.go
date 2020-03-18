package complex

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"golang.org/x/mod/semver"

	"github.com/hofstadter-io/mvs/lib/mod"
	"github.com/hofstadter-io/mvs/lib/remote/git"
)

func (mdr *Modder) Vendor() error {
	m, err := mdr.LoadModule(".")
	if err != nil {
		return err
	}

	mdr.module = m
	fmt.Printf("Root Module: %v\n", m)
	mdr.depsMap = map[string]*mod.Module{}

	// TODO START par.Work here

	// TODO, build up errors
	// TODO merge require / replace
	// TODO compare to sum file
	for _, req := range mdr.module.Require {
		ref, refs, err := mdr.index(req.Path, req.Version)
		if err != nil {
			return err
		}

		fmt.Println(" ", req.Path, req.Version, ref)

		// TODO Later, after any real clone, during dep recursion or vendoring,
			// We should fill this in to respect modules' .mvsconfig, or portions of it depending on what we are doing
		m := &mod.Module{
			Module:  req.Path,
			Version: req.Version,
			Ref:     ref,
			Refs:    refs,
		}

		// save module to depsMap
		mdr.depsMap[req.Path] = m

		// TODO ADD par.Work here - clone and ilook for sum..., then do checks and actions

		// WORK SECTION ========================

		// clone the module and load
		clone, err := git.CloneRepoRef(req.Path, ref)
		m.Clone = clone
		if err != nil {
			return err
		}

		/*
			files, err := clone.FS.ReadDir("/")
			if err != nil { return err }

			for _, file := range files {
				fmt.Println("  -", file.Name())
			}
		*/

		// load dep module
		// dm, err := mdr.LoadModule("...")
		// if err != nil { return err }

		// WORK SECTION ========================
	}

	// XXX For Now...

	// Copyout based on config

	// XXX Actually want to...

	// TODO recurse here

	// XXX TEMP FINALLY, write out any vendor
	return mdr.writeVendor()
}

func (mdr *Modder) index(url, req string) (*plumbing.Reference, []*plumbing.Reference, error) {
	// fmt.Println("indexing:", url)

	if !semver.IsValid(req) {
		return nil, nil, fmt.Errorf("Invalid SemVer v2 %q", req)
	}

	repo, err := git.NewRemote(url)
	if err != nil {
		return nil, nil, err
	}

	refs, err := repo.RemoteRefs()
	if err != nil {
		return nil, nil, err
	}

	var minRef *plumbing.Reference
	verS := ""

	// handle v0.0.0
	if req == "v0.0.0" {
		minRef, verS, err = findHead(refs)
		if err != nil {
			return nil, refs, err
		}
	} else {
		minRef, verS, err = findMin(req, refs)
		if err != nil {
			return nil, refs, err
		}
	}

	if verS == "" {
		return nil, refs, fmt.Errorf("Did not find compatible version for %s @ %s", url, req)
	}

	// fmt.Println("  found:", ref, hash)

	return minRef, refs, nil
}

func findHead(refs []*plumbing.Reference) (*plumbing.Reference, string, error) {
	// find HEAD ref
	headRef := ""
	for _, ref := range refs {
		fields := strings.Fields(ref.String())

		// HEAD is the only line with 3 fields
		if len(fields) < 3 {
			continue
		}

		v := fields[2]
		if v == "HEAD" {
			headRef = fields[1]
			break
		}
	}

	// find hash for HEAD
	for _, ref := range refs {
		fields := strings.Fields(ref.String())
		ver := fields[1]

		if ver == headRef {
			return ref, ver, nil
		}
	}

	return nil, "", nil
}

func findMin(req string, refs []*plumbing.Reference) (*plumbing.Reference, string, error) {

	var minR *plumbing.Reference
	min := ""
	for _, ref := range refs {
		fields := strings.Fields(ref.String())
		ver := fields[1]
		if strings.HasPrefix(ver, "refs/tags/") {
			ver = strings.TrimPrefix(ver, "refs/tags/")
			if ver[0:1] != "v" {
				ver = "v" + ver
			}

			if semver.IsValid(ver) {
				if semver.Compare(req, ver) <= 0 {
					// if this version is less than the current min, update
					if min == "" || semver.Compare(ver, min) < 0 {
						min = ver
						minR = ref
					}
				}
			}
		}
	}

	return minR, min, nil
}

func (mdr *Modder) writeVendor() error {
	// TODO calc and update imported module "hashes"" here

	// make vendor dir if not present
	err := os.MkdirAll(mdr.ModsDir, 0755)
	if err != nil {
		return err
	}

	// write out each dep
	for _, m := range mdr.depsMap {
		baseDir := path.Join(mdr.ModsDir, m.Module)

		err = copyDir(baseDir, "/", m.Clone.FS)

	}

	return nil
}

func copyFile(baseDir string, file string, FS billy.Filesystem) error {
	outName := path.Join(baseDir, file)
	fmt.Println("CopyFile:", outName)

	bf, err := FS.Open(file)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(bf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outName, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func copyDir(baseDir string, dir string, FS billy.Filesystem) error {
	files, err := FS.ReadDir(dir)
	if err != nil {
		return err
	}

	fmt.Println("CopyDir: ", path.Join(baseDir, dir))

	for _, file := range files {

		if file.IsDir() {
			os.MkdirAll(path.Join(baseDir, dir, file.Name()), 0755)
			err = copyDir(baseDir, path.Join(dir, file.Name()), FS)
			if err != nil {
				return err
			}

		} else {
			err = copyFile(baseDir, path.Join(dir, file.Name()), FS)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
