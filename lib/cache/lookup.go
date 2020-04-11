package cache

import (
	"fmt"
	"path/filepath"
	"os"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/osfs"
)

func Load(lang, mod, ver string) (FS billy.Filesystem, err error) {

	dir := filepath.Join(
		LocalCacheBaseDir,
		"mod",
		lang,
		mod,
		"@",
		ver,
	)

	fmt.Println("Cache Load:", dir)

	_, err = os.Lstat(dir)
	if err != nil {
		return nil, err
	}

	FS = osfs.New(dir)

	return FS, nil
}
