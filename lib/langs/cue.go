package langs

const CuelangModder = `
cue: {
	Name:        "cue"
	Version:     string | *"0.0.15"
	ModFile:     string | * "cue.mods"
	SumFile:     string | * "cue.sums"
	ModsDir:     string | * "cue.mod/pkg"
	MappingFile: string | * "cue.mod/modules.txt"
	InitTemplates: {...} | *{
		"cue.mod/module.cue": """
			module: "{{ .Module }}"
		"""
		...
	}
	VendorIncludeGlobs: [...string] | *[
		"/cue.mods",
		"/cue.sums",
		"/cue.mod/module.cue",
		"/cue.mod/modules.txt",
		"**/*.cue"
	]
	VendorExcludeGlobs: [...string] | *["**/cue.mod/pkg/**"]
}
`
