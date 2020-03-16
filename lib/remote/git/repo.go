package git

import (
	"gopkg.in/src-d/go-billy.v4"
	gogit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage"
)

type GitRepo struct {
	Store  storage.Storer
	FS     billy.Filesystem

	Repo   *gogit.Repository

	Remote *gogit.Remote

	FetchOptions *gogit.FetchOptions
	ListOptions *gogit.ListOptions
}

