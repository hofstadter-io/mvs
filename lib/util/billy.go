package util

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/bmatcuk/doublestar"
	"github.com/go-git/go-billy/v5"
)

func BillyReadAllString(filename string, FS billy.Filesystem) (string, error) {
	bytes, err := BillyReadAll(filename, FS)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func BillyReadAll(filename string, FS billy.Filesystem) ([]byte, error) {
	f, err := FS.Open(filename)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(f)
}

// Copies dir in FS onto the os filesystem at baseDir
func BillyCopyDir(baseDir string, dir string, FS billy.Filesystem) error {
	// fmt.Println("DIR:  ", baseDir, dir)
	files, err := FS.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		longname := path.Join(dir, file.Name())
		// fmt.Println("DIR:  ", baseDir, dir, file.Name(), longname, outname)

		if file.IsDir() {
			err = BillyCopyDir(baseDir, longname, FS)
			if err != nil {
				return err
			}

		} else {
			err = BillyCopyFile(baseDir, longname, FS)
			if err != nil {
				return err
			}

		}
	}

	return nil
}

// Copies file in FS onto the os filesystem at baseDir
func BillyCopyFile(baseDir string, file string, FS billy.Filesystem) error {
	outName := path.Join(baseDir, file)

	err := os.MkdirAll(path.Dir(outName), 0755)
	if err != nil {
		return err
	}

	bf, err := FS.Open(file)
	if err != nil {
		return err
	}

	content, err := ioutil.ReadAll(bf)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outName, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Copies dir in FS onto the os filesystem at baseDir
//
func BillyGlobCopy(baseDir string, dir string, FS billy.Filesystem, includes, excludes []string) error {
	// fmt.Println("DIR:  ", baseDir, dir)
	files, err := FS.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		longname := path.Join(dir, file.Name())
		// fmt.Println("DIR:  ", baseDir, dir, file.Name(), longname, outname)
		// fmt.Println("GLOB?  ", longname)

		if file.IsDir() {
			err = BillyGlobCopy(baseDir, longname, FS, includes, excludes)
			if err != nil {
				return err
			}

		} else {

			include := false
			if len(includes) > 0 {
				for _, pattern := range includes {
					include, err = doublestar.PathMatch(pattern, longname)
					// fmt.Println("GLOB++  ", longname, pattern, include)
					if err != nil {
						return err
					}
					if include {
						break
					}
				}
			} else {
				include = true
			}

			exclude := false
			if len(excludes) > 0 {
				for _, pattern := range excludes {
					exclude, err = doublestar.PathMatch(pattern, longname)
					// fmt.Println("GLOB--  ", longname, pattern, exclude)
					if err != nil {
						return err
					}
					if exclude {
						break
					}
				}
			}

			// fmt.Println("COPY ==>", longname, include, exclude, include && !exclude)

			if include && !exclude {
				err = BillyCopyFile(baseDir, longname, FS)
				if err != nil {
					return err
				}
			}

		}
	}

	return nil
}

