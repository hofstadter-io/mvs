package simple

import (
	"github.com/hofstadter-io/mvs/lib/mod"
)

type Modder struct {
	Name    string
	Version string
	ModFile string
	SumFile string
	ModsDir string
	Copies  []string

	module *mod.Module
}

