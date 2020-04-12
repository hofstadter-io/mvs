package langs

const CuelangModder = `
cue: {
	Name:        "cue"
	Version:     string | *"0.1.1"
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
	VendorIncludeGlobs: [...string] | *[]
	VendorExcludeGlobs: [...string] | *[
		"/.git/**",
		"**/cue.mod/pkg/**",
	]
}
`
