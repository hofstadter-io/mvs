package exec

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

type Modder struct {
	Name     string
	Commands map[string][]string
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
