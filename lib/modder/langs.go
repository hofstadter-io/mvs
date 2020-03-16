package modder

var (

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
		Version: "0.0",
		Copies: append(CommonCopies, []string{
			"module.cue",
			"pkg",
			"usr",
			"gen",
		}...),
	}

	ModderMap = map[string]Modder{
		"go":  GolangModder,
		"cue": CuelangModder,
	}
)
