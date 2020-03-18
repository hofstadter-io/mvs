package complex

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hofstadter-io/mvs/lang/modfile"
)

/* TODO
	- more configuration for intialization
*/

func (m *Modder) Init(module string) error {
	lang := m.Name
	filename := m.ModFile

	// make sure file does not exist
	_, err := ioutil.ReadFile(filename)
	// we read the file and it exists
	if err == nil {
		return fmt.Errorf("%s already exists", filename)
	}
	// error was not path error, so return
	if _, ok := err.(*os.PathError); !ok {
		return err
	}

	// Create empty modfile
	f, err := modfile.Parse(filename, nil, nil)
	if err != nil {
		return err
	}

	err = f.AddModuleStmt(module)
	if err != nil {
		return err
	}

	err = f.AddLanguageStmt(lang, m.Version)
	if err != nil {
		return err
	}

	bytes, err := f.Format()
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
