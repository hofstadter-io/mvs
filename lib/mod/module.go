package mod

type ModSet map[string]*Module

type Module struct {
	Language string
	LangVer  string
	Module   string
	Version  string
	Require  []Require
	Replace  []Replace
	SumMod   *Module
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
