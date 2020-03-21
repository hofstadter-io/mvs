package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Verify() error {
	if len(mdr.CommandVerify) > 0 {
		out, err := util.Exec(mdr.CommandVerify)
		fmt.Println(out)
		return err
	}

	return fmt.Errorf("%s Modder - Verify not implemented", mdr.Name)
}
