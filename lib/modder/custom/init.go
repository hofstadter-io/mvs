package custom

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"

	"github.com/hofstadter-io/mvs/lang/modfile"
)

/* TODO
- more configuration for intialization
*/

func (mdr *Modder) Init(module string) error {
	var err error

	err = mdr.initModFile(module)
	if err != nil {
		return err
	}

	// err = os.MkdirAll(mdr.ModsDir, 0755)
	if err != nil {
		return err
	}

	err = mdr.writeInitTemplates(module)
	if err != nil {
		return err
	}

	return nil
}

func (mdr *Modder) initModFile(module string) error {
	lang := mdr.Name
	filename := mdr.ModFile

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

	err = f.AddLanguageStmt(lang, mdr.Version)
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

func (mdr *Modder) writeInitTemplates(module string) error {

	for filename, templateStr := range mdr.InitTemplates {

		tmpl, err := template.New(filename).Parse(templateStr)
		if err != nil {
			return err
		}

		data := map[string]interface{}{
			"Language": mdr.Name,
			"Module":   module,
			"Modder":   mdr,
		}

		err = os.MkdirAll(path.Dir(filename), 0755)
		if err != nil {
			return err
		}

		file, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer file.Close()

		err = tmpl.Execute(file, data)
		if err != nil {
			return err
		}

	}

	return nil
}
