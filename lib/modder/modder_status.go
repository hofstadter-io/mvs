package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Status() error {
	if len(mdr.CommandStatus) > 0 {
		out, err := util.Exec(mdr.CommandStatus)
		fmt.Println(out)
		return err
	}

	return fmt.Errorf("%s Modder - Status not implemented", mdr.Name)
}
