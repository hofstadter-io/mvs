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

	fmt.Println("fetching:", url, "@", ver)
	repo, err := git.NewRemote(url)
	if err != nil {
		return err
	}

	refs, err := repo.RemoteRefs()
	if err != nil {
		return err
	}

	for _, ref := range refs {
		fmt.Println(ref)
	}

	fmt.Println("\ntotal:", len(refs))
	return nil
}
