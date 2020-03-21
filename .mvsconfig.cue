joo: {
  Name:    "foo"
  Version: "0.0.0"
  ModFile: "foo.mods"
  SumFile: "foo.sums"
  ModsDir: "foo.mod/"
}

moo: {
  Name:    "moo"
  Version: "0.0.0"
  ModFile: "MOO.mods"
  SumFile: "MOO.sums"
  ModsDir: "MOO.MOO/"
}

cue: {
  Name: "cue"
  Version: "0.0.16"
  ModFile: "cue.mods"
  SumFile: "cue.sums"
  ModsDir: "cue.mod/pkg"
  Checksum: "cue.mod/modules.txt"
  InitTemplates: {
    "cue.mod/module.cue": """
      module "{{ .Module }}"
      """
    }

  VendorIncludeGlobs: [
    ".mvsconfig.cue",
    "cue.mods",
    "cue.sums",
    "cue.mod/module.cue",
    "cue.mod/modules.txt",
    "**/*.cue"
  ]
  VendorExcludeGlobs: ["cue.mod/pkg"]

  IntrospectIncludeGlobs: ["**/*.cue"]
  IntrospectExcludeGlobs: ["cue.mod/pkg"]
  IntrospectExtractRegex: ["tbd... same as go import"]
}

go: {
  Name:          "go",
  Version:       "1.14",
  ModFile:       "go.mod",
  SumFile:       "go.sum",
  ModsDir:       "vendor",
  Checksum:      "vendor/modules.txt",
  CommandInit:   ["go", "mod", "init"],
  CommandGraph:  ["go", "mod", "graph"],
  CommandTidy:   ["go", "mod", "tidy"],
  CommandVendor: ["go", "mod", "vendor"],
  // CommandVerify: ["go", "mod", "verify"],
}
