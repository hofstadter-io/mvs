package mod

import (
	"github.com/go-git/go-git/v5/plumbing"

	"github.com/hofstadter-io/mvs/lang/sumfile"
	"github.com/hofstadter-io/mvs/lib/remote/git"
)

type ModSet map[string]*Module

type Module struct {
	// From mod/sum files
	Language string
	LangVer  string
	Module   string
	Version  string
	Require  []Require
	Replace  []Replace

	// If this module gets replaced
	ReplaceModule string
	ReplaceVersion string

	// nested sum file
	SumMod *sumfile.Sum
	// TODO modules.txt for checksums

	Errors []error
	Ref    *plumbing.Reference
	Refs   []*plumbing.Reference
	Clone  *git.GitRepo
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
