cue: {
  Name: "cue"

  VendorIncludeGlobs: [
    "/cue.mods",
    "/cue.sums",
    "/cue.mod/module.cue",
    "/cue.mod/modules.txt",
    "**/*.cue"
  ]
	VendorExcludeGlobs: ["**/cue.mod/pkg/**"]
}
