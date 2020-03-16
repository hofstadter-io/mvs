package lib

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/remote/git"
)

func Hack(lang string, args []string) error {
	fmt.Println("Hack", args)

	url := args[0]
	ver := ""
	if len(args) >= 2 {
		ver = args[1]
	}

	repo, err := git.FetchRepo(url, ver)
	if err != nil {
		return err
	}

	err = repo.GetRefs()
	if err != nil {
		return err
	}

	err = repo.GetTags()
	if err != nil {
		return err
	}

	return nil
}
