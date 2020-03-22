package langs

import (
	"github.com/hofstadter-io/mvs/lib/modder"
)

var (
	DefaultModders = map[string]*modder.Modder{
		"go":  GolangModder,
		"cue": CuelangModder,
		"py":  PythonModder,
	}

	GolangModder = &modder.Modder{
		Name:             "go",
		Version:          "1.14",
		ModFile:          "go.mod",
		SumFile:          "go.sum",
		ModsDir:          "vendor",
		MappingFile:      "vendor/modules.txt",
		CommandInit:   []string{"go", "mod", "init"},
		CommandGraph:  []string{"go", "mod", "graph"},
		CommandTidy:   []string{"go", "mod", "tidy"},
		CommandVendor: []string{"go", "mod", "vendor"},
		CommandVerify: []string{"go", "mod", "verify"},
	}

	CuelangModder = &modder.Modder{
		Name:        "cue",
		Version:     "0.0.15",
		ModFile:     "cue.mods",
		SumFile:     "cue.sums",
		ModsDir:     "cue.mod/pkg",
		MappingFile: "cue.mod/modules.txt",
		InitTemplates: map[string]string{
			"cue.mod/module.cue": `module: "{{ .Module }}"
`,
		},
		VendorIncludeGlobs: []string{
			".mvsconfig.cue",
			"cue.mods",
			"cue.sums",
			"cue.mod/module.cue",
			"cue.mod/modules.txt",
			"**/*.cue",
		},
		VendorExcludeGlobs: []string{
			"cue.mod/pkg",
		},
	}

	PythonModder = &modder.Modder{
		Name:          "python",
		Version:       "3.8",
		ModFile:       "python.mod",
		SumFile:       "requirements.txt",
		ModsDir:       "vendor",
		MappingFile:   "vendor/modules.txt",
		CommandInit:   []string{"python", "-m", "venv", "venv"},
		CommandVendor: []string{"bash", "-c", ". ./venv/bin/activate && pip install -r requirements.txt"},
	}
)
