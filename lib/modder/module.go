package modder

import (
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/hofstadter-io/mvs/lang/sumfile"
	"github.com/hofstadter-io/mvs/lib/remote/git"
)

type Module struct {
	// From mod/sum files
	Language string
	LangVer  string
	Module   string
	Version  string
	Require  []Require
	Replace  []Replace

	// Merged version of local require / replace
	// Requires as replaces will not have the old fields set
	SelfDeps map[string]Replace

	// If this module gets replaced
	ReplaceModule  string
	ReplaceVersion string

	// nested sum file
	SumMod *sumfile.Sum
	// TODO modules.txt for mapping imports to vendors
	// TODO also a checksum?

	// TODO, is this modder a good idea for our nested
	//   .mvsconfig processing and vendoring
	Mdr *Modder

	Errors []error
	Ref    *plumbing.Reference
	Refs   []*plumbing.Reference
	Clone  *git.GitRepo
	FS     billy.Filesystem
}

type Require struct {
	Path    string
	Version string
}

type Replace struct {
	OldPath    string
	OldVersion string
	NewPath    string
	NewVersion string
}

// If no lang.sum, calc sum, degenerate of next
// if both, look for differences, calc sumc
// if diff, fetch and do normal thing
