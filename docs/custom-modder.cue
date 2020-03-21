// These two need to be the same
cue: {
	Name: "cue"
	// non-semver of the language
	Version: "#.#.#"

	// Common defaults, can be anything
	ModFile:  "<lang>.mods"
	SumFile:  "<lang>.sums"
	ModsDir:  "<lang>.mod/pkg"
	Checksum: "<lang>.mod/checksum.txt"

	// Controls for modders that want to shell out
	// to common tools for certain commands
	NoLoad: false
	CommandInit: [[string]]
	CommandGraph: [[string]]
	CommandTidy: [[string]]
	CommandVendor: [[string]]
	CommandVerify: [[string]]

	// Runs on init for this language
	// filename/content key/pair values
	// uses the golang text/template library
	// inputs will be
	//   .Language
	//   .Module
	//   .Modder
	InitTemplates: {
		"<lang>.mod/module.<lang>": """
          module "{{ .Module }}"
          """
	}
	// Series of commands to be executed pre/post init
	InitPreCommands: [[string]]
	InitPostCommands: [[string]]

	// Same as the InitTemplates, but run during vendor command
	VendorTemplates: {
		"<lang>.mod/module.<lang>": """
          module "{{ .Module }}"
          """
	}

	VendorIncludeGlobs: [
		".mvsconfig.cue",
		"<lang>.mods",
		"<lang>.sums",
		"<lang>.mod/module.<lang>",
		"<lang>.mod/modules.txt",
		"**/*.<lang>",
	]
	VendorExcludeGlobs: ["<lang>.mod/pkg"]

	// Series of commands to be executed pre/post vendoring
	VendorPreCommands: [[string]]
	VendorPostCommands: [[string]]

	// Use MVS to only manage the languages normal dependency file
	ManageFileOnly: false

	// Whether local replaces should use a symlink instead of copying files
	SymlinkLocalReplaces: false

	// Controls the code introspection for dependency determiniation
	IntrospectIncludeGlobs: ["**/*.<lang>"]
	IntrospectExcludeGlobs: ["<lang>.mod/pkg"]
	IntrospectExtractRegex: ["you will have to figure out a series of 'any match passes' regexps to pull out dependencies"]

	// This field determines the prefix to place in front of
	// imports which have a single token or leverage package managers
	// This is currently futurology for building MVS for Python and JavaScript
	PackageManagerDefaultPrefix: "npm.js"
}
