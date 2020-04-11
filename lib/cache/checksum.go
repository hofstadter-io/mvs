package cache

import (
	"fmt"
	"path/filepath"
	"os"

	"golang.org/x/mod/sumdb/dirhash"
)

func Checksum(lang, mod, ver string) (string, error) {

	dir := filepath.Join(
		LocalCacheBaseDir,
		"mod",
		lang,
		mod,
		"@",
		ver,
	)

	fmt.Println("Cache Load:", dir)

	_, err := os.Lstat(dir)
	if err != nil {
		return "", err
	}

	h, err := dirhash.HashDir(dir, mod, dirhash.Hash1)

	return h, err
}

