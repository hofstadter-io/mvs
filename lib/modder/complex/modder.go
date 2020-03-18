package complex

import (
	"github.com/hofstadter-io/mvs/lib/mod"
)

// This modder is for more complex, yet configurable module processing.
// You can have system wide and local custom configurations.
	// The fields in this struct are alpha and are likely to change
type Modder struct {
	// MetaConfiguration
	Name      string
	Version   string

	// Module information
	ModFile   string
	SumFile   string
	ModsDir   string

	// Init related fields
	// we need to create things like directories and files beyond the
	InitTemplates []string
	InitPreCommands []string
	InitPostCommands []string

	// Vendor related fields
	// filesystem globs for discovering files we should copy over
	VendorIncludeGlobs   []string
	VendorExcludeGlobs   []string
	// Any files we need to generate
	VendorTemplates []string
	VendorPreCommands []string
	VendorPostCommands []string

	// Introspection Configuration(s)
	// filesystem globs for discovering files we should introspect
	// regexs for extracting package information
	IntrospectIncludeGlobs []string
	IntrospectExcludeGlobs []string
	IntrospectExtractRegex []string

	// root module
	module *mod.Module
	errors []error

	// dependency modules (requires/replace)
	// dependencies shoule respect any .mvsconfig it finds along side the module files
	// module writers can then have local control over how their module is handeled during vendoring
	depsMap map[string]*mod.Module
}
