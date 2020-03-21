package modder

import (
	"fmt"
	"os"
	"path"

	"github.com/hofstadter-io/mvs/lib/util"
)

var (
	// Common files to copy from modules, also includes the .md version of the filename
	definiteVendors = []string{
		"README",
		"LICENSE",
		"PATENTS",
		"CONTRIBUTORS",
		"SECURITY",
	}

	// cross product these endings
	endings = []string{"", ".md", ".txt"}
)

func (mdr *Modder) WriteVendor() error {
	// TODO calc and update imported module "hashes"" here

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
		for _, fn := range definiteVendors {
			for _, end := range endings {
				_, err := m.Clone.FS.Stat(fn + end)
				if err != nil {
					if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
						// some other error
						return err
					}
					// not found
					continue
				}

				// Found one!
				// TODO, these functions should just take 2 billy FS
				err = util.BillyCopyFile(baseDir, "/"+fn+end, m.Clone.FS)

			}
		}

		if len(mdr.VendorIncludeGlobs) > 0 || len(mdr.VendorExcludeGlobs) > 0 {
			// Just copy everything
			// TODO, these functions should just take 2 billy FS
			err = util.BillyGlobCopy(baseDir, "/", m.Clone.FS, mdr.VendorIncludeGlobs, mdr.VendorExcludeGlobs)
			if err != nil {
				return err
			}

		} else {
			// Just copy everything
			// TODO, these functions should just take 2 billy FS
			err = util.BillyCopyDir(baseDir, "/", m.Clone.FS)
			if err != nil {
				return err
			}

		}

	}

	return nil
}
