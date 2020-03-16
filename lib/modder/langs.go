package modder

var {

	// Common files to copy from modules, also includes the .md version of the filename
	CommonCopies := []string {
		"README",
		"LICENSE",
		"PATENTS",
		"CONTRIBUTORS"
		"SECURITY"
	}

	GolangModder := &ExecModder {
		Name: "go",
		Commands: map[string][]string{
			"init":    []string{"go", "init"},
			"graph":   []string{"go", "graph"},
			"tidy":    []string{"go", "tidy"},
			"vendor":  []string{"go", "vendor"},
			"verify":  []string{"go", "verify"},
		},
	}

	CuelangModder := &SimpleModder {
		Name: "cue"
		Version: "0.0"
		Copies: []string {
			CommonCopies...,
			"module.cue"
			"pkg",
			"usr",
			"gen",
		}
	}


}
