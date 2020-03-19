package modder

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
	ModFile  string `yaml:"ModFile"`
	SumFile  string `yaml:"SumFile"`
	ModsDir  string `yaml:"ModsDir"`
	Checksum string `yaml:"Checksum"`

	// Commands override default, configuragble processing
	NoLoad        bool     `yaml:"NoLoad"` // for things like golang
	CommandInit   []string `yaml:"CommandInit"`
	CommandGraph  []string `yaml:"CommandGraph"`
	CommandTidy   []string `yaml:"CommandTidy"`
	CommandVendor []string `yaml:"CommandVendor"`
	CommandVerify []string `yaml:"CommandVerify"`

	// Init related fields
	// we need to create things like directories and files beyond the
	InitTemplates    map[string]string `yaml:"InitTemplates"`
	InitPreCommands  [][]string        `yaml:"InitPreCommands"`
	InitPostCommands [][]string        `yaml:"InitPostCommands"`

	// Vendor related fields
	// filesystem globs for discovering files we should copy over
	VendorIncludeGlobs []string `yaml:"VendorIncludeGlobs"`
	VendorExcludeGlobs []string `yaml:"VendorExcludeGlobs"`
	// Any files we need to generate
	VendorTemplates    map[string]string `yaml:"VendorTemplates"`
	VendorPreCommands  [][]string        `yaml:"VendorPreCommands"`
	VendorPostCommands [][]string        `yaml:"VendorPostCommands"`

	// Some more vendor controls
	ManageFileOnly       bool `yaml:"ManageFileOnly"`
	SymlinkLocalReplaces bool `yaml:"SymlinkLocalReplaces"`

	// Introspection Configuration(s)
	// filesystem globs for discovering files we should introspect
	// regexs for extracting package information
	IntrospectIncludeGlobs []string `yaml:"IntrospectIncludeGlobs"`
	IntrospectExcludeGlobs []string `yaml:"IntrospectExcludeGlobs"`
	IntrospectExtractRegex []string `yaml:"IntrospectExtractRegex"`

	PackageManagerDefaultPrefix string `yaml:"PackageManagerDefaultPrefix"`

	// root module
	module *mod.Module
	errors []error

	// dependency modules (requires/replace)
	// dependencies shoule respect any .mvsconfig it finds along side the module files
	// module writers can then have local control over how their module is handeled during vendoring
	depsMap map[string]*mod.Module
}
