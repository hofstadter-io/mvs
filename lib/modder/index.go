package modder

import (
	"fmt"

	"github.com/go-git/go-git/v5/plumbing"
	"golang.org/x/mod/semver"

	"github.com/hofstadter-io/mvs/lib/repos/git"
)

// given a git url, and a req (required) version
// returns the minReference, allRefferences, and error
func IndexGit(url, req string) (*plumbing.Reference, []*plumbing.Reference, error) {
	fmt.Println("indexing:", url, req)

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
