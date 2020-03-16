package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

type ExecModder struct {
	Name     string
	Commands map[string][]string
}

func (m *ExecModder) Init(module string) error {
	args := append(m.Commands["init"], module)
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *ExecModder) Graph() error {
	args := m.Commands["graph"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *ExecModder) Tidy() error {
	args := m.Commands["tidy"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *ExecModder) Vendor() error {
	args := m.Commands["vendor"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *ExecModder) Verify() error {
	args := m.Commands["verify"]
	out, err := util.Exec(args)
	fmt.Println(out)
	return err
}

func (m *ExecModder) Load(dir string) error {
	panic("This function should never be called or implemented")
}
