package modder

import (
	"fmt"

	"github.com/hofstadter-io/mvs/lib/util"
)

func (mdr *Modder) Graph() error {
	if len(mdr.CommandGraph) > 0 {
		out, err := util.Exec(mdr.CommandGraph)
		fmt.Println(out)
		return err
	}

	return fmt.Errorf("%s Modder - Graph not implemented", mdr.Name)
}
