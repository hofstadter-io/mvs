package common

import (
	"strings"

	"github.com/go-git/go-git/v5/plumbing"
	"golang.org/x/mod/semver"
)

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
