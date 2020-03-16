package git

import (
	"fmt"
	"os"
	"strings"

	"gopkg.in/src-d/go-billy.v4/memfs"
	gogit "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/http"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

func NewRemote(srcUrl string) (*GitRepo, error) {

	rc := &config.RemoteConfig {
		Name: "origin",
		URLs: []string{
			"https://" + srcUrl,
		},
	}

	lo := &gogit.ListOptions{}

	if strings.Contains(srcUrl, "github.com") && os.Getenv("GITHUB_TOKEN") != "" {
		lo.Auth = &http.BasicAuth{
			Username: "github-token", // yes, this can be anything except an empty string
			Password: os.Getenv("GITHUB_TOKEN"),
		}
		// co.URL = "git@" + strings.Replace(srcUrl, "/", ":", 1)
	}

	fmt.Println("URL:", rc.URLs[0])

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	st := memory.NewStorage()
	remote := gogit.NewRemote(st, rc)

	return &GitRepo {
		Store: st,
		Remote: remote,
		ListOptions: lo,
	}, nil
}

func CloneRepo(srcUrl, srcVer string) (*GitRepo, error) {

	co := &gogit.CloneOptions{
		URL: "https://" + srcUrl,
	}

	if strings.Contains(srcUrl, "github.com") && os.Getenv("GITHUB_TOKEN") != "" {
		co.Auth = &http.BasicAuth{
			Username: "github-token", // yes, this can be anything except an empty string
			Password: os.Getenv("GITHUB_TOKEN"),
		}
		// co.URL = "git@" + strings.Replace(srcUrl, "/", ":", 1)
	}

	fmt.Println("URL:", co.URL)

	if srcVer != "" {
		co.ReferenceName = plumbing.ReferenceName(srcVer)
	}

	// Clones the repository into the worktree (fs) and storer all the .git
	// content into the storer
	st := memory.NewStorage()
 	fs := memfs.New()
	r, err := gogit.Clone(st, fs, co)
	if err != nil {
		return nil, err
	}

	return &GitRepo {
		Store: st,
		FS: fs,
		Repo: r,
	}, nil
}
