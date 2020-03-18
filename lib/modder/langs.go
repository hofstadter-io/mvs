package modder

import (
	"github.com/hofstadter-io/mvs/lib/modder/exec"
	"github.com/hofstadter-io/mvs/lib/modder/simple"
)

var (
	// Default known modderr
	ModderMap = map[string]Modder{
		"go":  GolangModder,
		"cue": CuelangModder,
		"hof": HoflangModder,
	}
	// TODO, add custom Modders here (for simple) read from a ./.mvsconfig file

	// Common files to copy from modules, also includes the .md version of the filename
	CommonCopies = []string{
		"README",
		"README.md",
		"LICENSE",
		"LICENSE.md",
		"PATENTS",
		"PATENTS.md",
		"CONTRIBUTORS",
		"CONTRIBUTORS.md",
		"SECURITY",
		"SECURITY.md",
	}

	GolangModder = &exec.Modder{
		Name: "go",
		Commands: map[string][]string{
			"init":   []string{"go", "mod", "init"},
			"graph":  []string{"go", "mod", "graph"},
			"tidy":   []string{"go", "mod", "tidy"},
			"vendor": []string{"go", "mod", "vendor"},
			"verify": []string{"go", "mod", "verify"},
		},
	}

	CuelangModder = &simple.Modder{
		Name:    "cue",
		Version: "0.0.15",
		ModFile: "cue.mods",
		SumFile: "cue.sums",
		ModsDir: "cue.mod/pkg",
	}

	HoflangModder = &simple.Modder{
		Name:    "hof",
		Version: "0.0.0",
		ModFile: "hof.mods",
		SumFile: "hof.sums",
		ModsDir: "hof.mod/pkg",
	}
)
