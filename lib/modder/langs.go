package modder

var (
	// TODO, add custom Modders here (for simple) read from a ./.mvsconfig file

	// Common files to copy from modules, also includes the .md version of the filename
	CommonCopies = []string{
		"README",
		"LICENSE",
		"PATENTS",
		"CONTRIBUTORS",
		"SECURITY",
	}

	GolangModder = &ExecModder{
		Name: "go",
		Commands: map[string][]string{
			"init":   []string{"go", "mod", "init"},
			"graph":  []string{"go", "mod", "graph"},
			"tidy":   []string{"go", "mod", "tidy"},
			"vendor": []string{"go", "mod", "vendor"},
			"verify": []string{"go", "mod", "verify"},
		},
	}

	CuelangModder = &SimpleModder{
		Name:    "cue",
		Version: "0.0.15",
		ModFile: "cue.mods",
		SumFile: "cue.sums",
		ModsDir: "cue.mod",
		Copies: append(CommonCopies, []string{
			"cue.mods",
			"cue.sums",
			"cue.mod/module.cue",
			"cue.mod/pkg/",
			"cue.mod/usr/",
			"cue.mod/gen/",
		}...),
	}

	HoflangModder = &SimpleModder{
		Name:    "hof",
		Version: "0.0.0",
		ModFile: "hof.mod",
		SumFile: "hof.sum",
		ModsDir: "hof.mods",
		Copies: append(CommonCopies, []string{
			"hof.mod",
			"hof.sum",
			"hof.mods/module.cue",
			"hof.mods/pkg/",
			"hof.mods/usr/",
			"hof.mods/gen/",
		}...),
	}

	ModderMap = map[string]Modder{
		"go":  GolangModder,
		"cue": CuelangModder,
		"hof": HoflangModder,
	}
)
