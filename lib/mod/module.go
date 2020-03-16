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
	Path   string
	Source string
}
