package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Tidy() error {
	if len(mdr.CommandTidy) > 0 {
		out, err := util.Exec(mdr.CommandTidy)
		fmt.Println(out)
		return err
	}

	return fmt.Errorf("%s ComplexModder - Tidy not implemented", mdr.Name)
}
