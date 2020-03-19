package custom

import (
	"github.com/hofstadter-io/mvs/lib/mod"
)

// This modder is for more complex, yet configurable module processing.
// You can have system wide and local custom configurations.
// The fields in this struct are alpha and are likely to change
type Modder struct {
	// MetaConfiguration
	Name    string `yaml:"Name"`
	Version string `yaml:"Version"`

	// Module information
	ModFile string `yaml:"ModFile"`
	SumFile string `yaml:"SumFile"`
	ModsDir string `yaml:"ModsDir"`

	// Init related fields
	// we need to create things like directories and files beyond the
	InitTemplates    []string `yaml:"InitTemplates"`
	InitPreCommands  []string `yaml:"InitPreCommands"`
	InitPostCommands []string `yaml:"InitPostCommands"`

	// Vendor related fields
	// filesystem globs for discovering files we should copy over
	VendorIncludeGlobs []string `yaml:"VendorIncludeGlobs"`
	VendorExcludeGlobs []string `yaml:"VendorExcludeGlobs"`
	// Any files we need to generate
	VendorTemplates    []string `yaml:"VendorTemplates"`
	VendorPreCommands  []string `yaml:"VendorPreCommands"`
	VendorPostCommands []string `yaml:"VendorPostCommands"`

	// Introspection Configuration(s)
	// filesystem globs for discovering files we should introspect
	// regexs for extracting package information
	IntrospectIncludeGlobs []string `yaml:"IntrospectIncludeGlobs"`
	IntrospectExcludeGlobs []string `yaml:"IntrospectExcludeGlobs"`
	IntrospectExtractRegex []string `yaml:"IntrospectExtractRegex"`

	// root module
	module *mod.Module
	errors []error

	// dependency modules (requires/replace)
	// dependencies shoule respect any .mvsconfig it finds along side the module files
	// module writers can then have local control over how their module is handeled during vendoring
	depsMap map[string]*mod.Module
}
