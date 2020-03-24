package modder

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/hofstadter-io/mvs/lib/util"
)

var (
	// Common files to copy from modules, also includes the .md version of the filename
	definiteVendors = []string{
		"README",
		"SECURITY",
		"AUTHORS",
		"CONTRIBUTORS",
		"COPYLEFT",
		"COPYING",
		"COPYRIGHT",
		"LEGAL",
		"LICENSE",
		"NOTICE",
		"PATENTS",
	}

)

func (mdr *Modder) WriteVendor() error {
	os.RemoveAll(mdr.ModsDir)

	// make vendor dir if not present
	err := os.MkdirAll(mdr.ModsDir, 0755)
	if err != nil {
		return err
	}

	// write out each dep
	for _, m := range mdr.depsMap {

		baseDir := path.Join(mdr.ModsDir, m.Module)
		// TODO make billy FS here

		fmt.Println("Copying", baseDir)

		// copy definite files always
		files, err := m.FS.ReadDir("/")
		if err != nil {
			return err
		}
		for _, file := range files {
			for _, fn := range definiteVendors {
				// Found one!
				if strings.HasPrefix(strings.ToUpper(file.Name()), fn) {
					// TODO, these functions should just take 2 billy FS
					err = util.BillyCopyFile(baseDir, "/"+file.Name(), m.FS)
					if err != nil {
						return err
					}
				}

			}
		}

		if len(mdr.VendorIncludeGlobs) > 0 || len(mdr.VendorExcludeGlobs) > 0 {
			// Just copy everything
			// TODO, these functions should just take 2 billy FS
			err = util.BillyGlobCopy(baseDir, "/", m.FS, mdr.VendorIncludeGlobs, mdr.VendorExcludeGlobs)
			if err != nil {
				return err
			}

		} else {
			// Just copy everything
			// TODO, these functions should just take 2 billy FS
			err = util.BillyCopyDir(baseDir, "/", m.FS)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
