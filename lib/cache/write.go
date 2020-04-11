package cache

import (
	"path/filepath"

	"github.com/go-git/go-billy/v5"

	"github.com/hofstadter-io/mvs/lib/util"
)

func Write(lang, remote, owner, repo, tag string, FS billy.Filesystem) error {
	outdir := filepath.Join(
		LocalCacheBaseDir,
		"mod",
		lang,
		remote,
		owner,
		repo + "@" + tag,
	)

	return util.BillyWriteDirToOS(outdir, "/", FS)
}
