package exec

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

type Modder struct {
	Name    string `yaml:"Name"`
	Version string `yaml:"Version"`

	// Module information
	ModFile  string              `yaml:"ModFile"`
	SumFile  string              `yaml:"SumFile"`
	ModsDir  string              `yaml:"ModsDir"`
	Checksum string              `yaml:"Checksum"`
	Commands map[string][]string `yaml:"Checksum"`
}

func (m *Modder) Init(module string) error {
	args := append(m.Commands["init"], module)
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *Modder) Graph() error {
	args := m.Commands["graph"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *Modder) Tidy() error {
	args := m.Commands["tidy"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *Modder) Vendor() error {
	args := m.Commands["vendor"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *Modder) Verify() error {
	args := m.Commands["verify"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *Modder) Load(dir string) error {
	panic("This function should never be called or implemented")
}
