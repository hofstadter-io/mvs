package modder

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v3"

	"github.com/hofstadter-io/mvs/lang/modfile"
	"github.com/hofstadter-io/mvs/lib/mod"
)

type SimpleModder struct {
	Name    string
	Version string
	ModFile string
	SumFile string
	ModsDir string
	Copies  []string
}

func (m *SimpleModder) Init(module string) error {
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

func (m *SimpleModder) Graph() error {
	return fmt.Errorf("%s SimpleModder - Graph not implemented", m.Name)
}

func (m *SimpleModder) Tidy() error {
	return fmt.Errorf("%s SimpleModder - Tidy not implemented", m.Name)
}

func (m *SimpleModder) Vendor() error {
	return m.Load(".")
	return fmt.Errorf("%s SimpleModder - Vendor not implemented", m.Name)
}

func (m *SimpleModder) Verify() error {
	return fmt.Errorf("%s SimpleModder - Verify not implemented", m.Name)
}

func (m *SimpleModder) Load(dir string) error {
	mdr, err := mod.LoadModule(m.Name, dir)
	if err != nil {
		return err
	}
	content, err := yaml.Marshal(mdr)
	if err != nil {
		return err
	}

	fmt.Printf("Module Contents:\n%s\n", string(content))
	return nil
}
