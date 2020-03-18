package complex

import (
	"fmt"
)

func (m *Modder) Tidy() error {
	return fmt.Errorf("%s ComplexModder - Tidy not implemented", m.Name)
}
