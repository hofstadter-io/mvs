package simple

import (
	"fmt"
)

func (m *Modder) Tidy() error {
	return fmt.Errorf("%s SimpleModder - Tidy not implemented", m.Name)
}
