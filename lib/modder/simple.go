package modder

import (
	"fmt"
)

type SimpleModder struct {
	Name    string
	Version string
	Copies  []string
}

func (m *SimpleModder) Init(module string) error {
	return fmt.Errorf("%s SimpleModder - Init not implemented", m.Name)
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
