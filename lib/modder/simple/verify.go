package simple

import (
	"fmt"
)

func (m *Modder) Verify() error {
	return fmt.Errorf("%s SimpleModder - Verify not implemented", m.Name)
}

