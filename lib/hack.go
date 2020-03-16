package lib

import (
	"fmt"
	"strings"

	"golang.org/x/mod/semver"

	"github.com/hofstadter-io/mvs/lib/remote/git"
)

func Hack(lang string, args []string) error {
	fmt.Println("Hack", args)

	url := args[0]
	req := ""
	if len(args) >= 2 {
		req = args[1]
	}

	fmt.Println("fetching:", url, "@", req)
	repo, err := git.NewRemote(url)
	if err != nil {
		return err
	}

	refs, err := repo.RemoteRefs()
	if err != nil {
		return err
	}

	vers := []string{}
	for _, ref := range refs {
		fields := strings.Fields(ref.String())
		v := fields[1]
		if strings.HasPrefix(v, "refs/tags/") {
			v = strings.TrimPrefix(v, "refs/tags/")
			if v[0:1] != "v" {
				v = "v" + v
			}

			if semver.IsValid(v) {
				vers = append(vers, v)
			}
		}
	}

	min := ""
	if len(vers) > 0 {
		for _, ver := range vers {
			// fmt.Println(ver)

			// find min
			if req != "" {
				fmt.Printf("- %s : %s ? %s  ~  %d %d\n", req, min, ver, semver.Compare(req, ver), semver.Compare(ver, min))
				// this version is >= to requested
				if semver.Compare(req, ver) <= 0 {
					// if this version is less than the current min, update
					if min == "" || semver.Compare(ver, min) < 0 {
						min = ver
					}
				}
			}
		}
	}

	fmt.Println("\ntotal:", len(vers), "\n")
	if req != "" {
		if min == "" {
			fmt.Println("no version found compatible with", req)
		} else {
			fmt.Printf("found MVS %s for %s\n\n", min, req)
		}
	}
	return nil
}
