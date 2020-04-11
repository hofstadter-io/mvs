package github

import (
	"archive/zip"
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v30/github"
	"github.com/parnurzeal/gorequest"
)

func GetTagsSplit(client *github.Client, module string) ([]*github.RepositoryTag, error) {
	flds := strings.SplitN(module, "/", 1)
	domain, rest := flds[0], flds[1]

	if domain != "github.com" {
		return nil, fmt.Errorf("Github Tags Fetch called with non 'github.com' domain %q", module)
	}

	flds = strings.Split(rest, "/")
	owner, repo := flds[0], flds[1]
	tags, _, err := client.Repositories.ListTags(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func GetTags(client *github.Client, owner, repo string) ([]*github.RepositoryTag, error) {
	tags, _, err := client.Repositories.ListTags(context.Background(), owner, repo, nil)
	if err != nil {
		return nil, err
	}
	return tags, nil
}

func FetchTagZip(client *github.Client, tag *github.RepositoryTag) (*zip.Reader, error) {

	url := *tag.ZipballURL

	req := gorequest.New().Get(url)
	resp, data, errs := req.EndBytes()

	check := "http2: server sent GOAWAY and closed the connection"
	if len(errs) != 0 && !strings.Contains(errs[0].Error(), check) {
		fmt.Println("errs:", errs)
		fmt.Println("resp:", resp)
		fmt.Println("body:", len(data))
		return nil, errs[0]
	}

	if len(errs) != 0 || resp.StatusCode >= 500 {
		return nil, fmt.Errorf("Internal Error: " + string(resp.StatusCode))
	}
	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Bad Request: " + string(resp.StatusCode))
	}

	r := bytes.NewReader(data)

	zfile, err := zip.NewReader(r, int64(len(data)))

	return zfile, err
}
