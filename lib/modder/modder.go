package modder

import (
  "fmt"
	"os"

	"gopkg.in/yaml.v3"
	"github.com/go-git/go-billy/v5"

	"github.com/hofstadter-io/mvs/lib/util"
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

	// filesystem
	FS billy.Filesystem `yaml:"-"`

	// root module
	module *Module `yaml:"-"`
	errors []error `yaml:"-"`

	// dependency modules (requires/replace)
	// dependencies shoule respect any .mvsconfig it finds along side the module files
	// module writers can then have local control over how their module is handeled during vendoring
	depsMap map[string]*Module `yaml:"-"`
}


func NewFromFile(lang, filepath string, FS billy.Filesystem) (*Modder, error) {

	bytes, err := util.BillyReadAll(filepath, FS)
	if err != nil {
		if _, ok := err.(*os.PathError); !ok && err.Error() != "file does not exist" {
			return nil, err
		}
		// The user has not setup a global $HOME/.mvs/mvsconfig file
		return nil, nil
	}

	var mdrMap map[string]*Modder
	err = yaml.Unmarshal(bytes, &mdrMap)
	if err != nil {
		return nil, err
	}

	mdr, ok := mdrMap[lang]
	if !ok {
	  return nil, fmt.Errorf("lang %q not found in %s", lang, filepath)
	}

	return mdr, nil
}
