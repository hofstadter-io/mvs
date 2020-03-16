package git

import (
	"fmt"

	"gopkg.in/src-d/go-billy.v4"
	gogit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage"

)

type GitRepo struct {
	Store storage.Storer
	FS    billy.Filesystem
	Repo  *gogit.Repository
}

func (R *GitRepo) GetRefs() error {
	fmt.Println("Refs:")
	refs, err := R.Repo.References()
	if err != nil {
		return err
	}

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		fmt.Println(ref)
		if ref.Type() == plumbing.HashReference {
			fmt.Println(ref)
		}
		return nil
	})

	return err
}

func (R *GitRepo) GetTags() error {
	fmt.Println("\nTags:")
	iter, err := R.Repo.Tags()
	if err != nil {
		return err
	}

	err = iter.ForEach(func (ref *plumbing.Reference) error {
		obj, err := R.Repo.TagObject(ref.Hash())
		switch err {

		case nil:
			// Tag object present
			fmt.Println(*obj)

		case plumbing.ErrObjectNotFound:
			// Not a tag object

		default:
			// Some other error

		return err
		}
		return nil
	})

	return err
}
