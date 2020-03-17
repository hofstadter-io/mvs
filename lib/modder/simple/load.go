package simple

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/hofstadter-io/mvs/lib/mod"
)

func (m *Modder) Load(dir string) error {
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

