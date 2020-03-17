package simple

import (
	"fmt"
	"strings"
	"github.com/go-git/go-git/v5/plumbing"

	"golang.org/x/mod/semver"

	"github.com/hofstadter-io/mvs/lib/remote/git"
)


func (m *Modder) Vendor() error {
	err := m.Load(".")
	if err != nil { return err }

	// TODO, build up errors
	for _, req := range m.module.Require {
		err = m.index(req.Path, req.Version)
		if err != nil { return err }
	}

	return nil
}

func (m *Modder) index(url, req string) error {
	fmt.Println("indexing:", url)

	if !semver.IsValid(req) {
		return fmt.Errorf("Invalid SemVer v2 %q", req)
	}

	repo, err := git.NewRemote(url)
	if err != nil { return err }

	refs, err := repo.RemoteRefs()
	if err != nil { return err }

	ref, hash := "", ""

	// handle v0.0.0
	if req == "v0.0.0" {
		ref, hash, err = findHead(refs)
		if err != nil { return err }
	} else {
		ref, hash, err = findMin(req, refs)
		if err != nil { return err }
	}

	if hash == "" {
		return fmt.Errorf("Did not find compatible version for %s @ %s", url, req)
	}

	fmt.Println("  found:", ref, hash)

	return nil
}

func findHead(refs []*plumbing.Reference) (string, string, error) {
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
			return ref.Hash().String(), ver, nil
		}
	}

	return "", "", nil
}

func findMin(req string, refs []*plumbing.Reference) (string, string, error) {

	min := ""
	hash := ""
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
						hash = fields[0]
					}
				}
			}
		}
	}

	return hash, min, nil
}
