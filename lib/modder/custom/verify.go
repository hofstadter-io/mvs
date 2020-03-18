package custom

import (
	"fmt"
)

func (m *Modder) Verify() error {
	return fmt.Errorf("%s ComplexModder - Verify not implemented", m.Name)
}
