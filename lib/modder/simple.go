package modder

import (
	"fmt"
	"io/ioutil"
	"os"
)

type SimpleModder struct {
	Name    string
	Version string
	Copies  []string
}

func (m *SimpleModder) Init(module string) error {
	return initYaml(m.Name, module)
}

func initYaml(lang, module string) error {
	filename := fmt.Sprintf("%s.mod.yaml", lang)

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

	content := ""
	content += fmt.Sprintf("Language: %s\n", lang)
	content += fmt.Sprintf("LangVer: %s\n", "TBD")
	content += fmt.Sprintf("Module: %s\n", module)

	return ioutil.WriteFile(filename, []byte(content), 0644)
}

func (m *SimpleModder) Graph() error {
	return fmt.Errorf("%s SimpleModder - Graph not implemented", m.Name)
}

func (m *SimpleModder) Tidy() error {
	return fmt.Errorf("%s SimpleModder - Tidy not implemented", m.Name)
}

func (m *SimpleModder) Vendor() error {
	return fmt.Errorf("%s SimpleModder - Vendor not implemented", m.Name)
}

func (m *SimpleModder) Verify() error {
	return fmt.Errorf("%s SimpleModder - Verify not implemented", m.Name)
}
