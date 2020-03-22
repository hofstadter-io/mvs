package langs

import (
	"github.com/hofstadter-io/mvs/lib/modder"
)

var (
	DefaultModders = make(map[string]*modder.Modder)

	DefaultModdersCue = map[string]string{
		"go":     GolangModder,
		"cue":    CuelangModder,
		"python": PythonModder,
	}

	GolangModder = `
go: {
	Name:          "go",
	Version:       "1.14",
	ModFile:       "go.mod",
	SumFile:       "go.sum",
	ModsDir:       "vendor",
	MappingFile:   "vendor/modules.txt",
	CommandInit:   [["go", "mod", "init"]],
	CommandGraph:  [["go", "mod", "graph"]],
	CommandTidy:   [["go", "mod", "tidy"]],
	CommandVendor: [["go", "mod", "vendor"]],
	CommandVerify: [["go", "mod", "verify"]],
}
`

	CuelangModder = `
cue: {
	Name:        "cue",
	Version:     "0.0.15",
	ModFile:     "cue.mods",
	SumFile:     "cue.sums",
	ModsDir:     "cue.mod/pkg",
	MappingFile: "cue.mod/modules.txt",
	InitTemplates: {
		"cue.mod/module.cue": """
			module: "{{ .Module }}"
		"""
	},
	VendorIncludeGlobs: [
		".mvsconfig.cue",
		"cue.mods",
		"cue.sums",
		"cue.mod/module.cue",
		"cue.mod/modules.txt",
		"**/*.cue"
	],
	VendorExcludeGlobs: ["cue.mod/pkg"],
	NoLoad: true,
}
`

	PythonModder = `
python: {
	Name:          "python",
	Version:       "3.8",
	ModFile:       "python.mod",
	SumFile:       "requirements.txt",
	ModsDir:       "vendor",
	MappingFile:   "vendor/modules.txt",
	CommandInit:   [["python", "-m", "venv", "venv"]],
	CommandVendor: [["bash", "-c", ". ./venv/bin/activate && pip install -r requirements.txt"]],
}
`
)
