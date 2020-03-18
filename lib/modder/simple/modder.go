package simple

import (
	"github.com/hofstadter-io/mvs/lib/mod"
)

type Modder struct {
	// Configuration
	Name      string
	Version   string
	ModFile   string
	SumFile   string
	ModsDir   string
	Copies    []string
	Templates []string

	// root module
	module *mod.Module
	errors []error

	// dependency modules (requires/replace)
	depsMap map[string]*mod.Module
}
