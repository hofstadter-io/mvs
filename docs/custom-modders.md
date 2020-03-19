# Custom Modders

MVS gives you the ability to create
custom module systems, called Modders.
Modder is the struct name for the
internal code which controls how
modules and vendoring is handled.
You can configure as many of these as you like,
by providing global or local `.mvsconfig.yaml` files.

See the following two files for the configuration options
and defaults built into the tool:

- Definition: [lib/modder/modder.go](../lib/modder/modder.go)
- Defaults: [lib/langs.go](../lib/langs.go) and [.mvsconfig.yaml](../.mvsconfig.yaml)

You create your own modders by createing `.mvsconfig.yaml` files.
MVS will look for these in two places before any commands are run.

- A global `$HOME/.mvs/config.yaml`
- A local `./.mvsconfig.yaml`

_Dev note, fix issue #10 and migrate from yaml to cue_

### MVS config file format


```yaml
# These two need to be the same
<lang>:
  Name: "<lang>"
  # non-semver of the language
  Version: "#.#.#"

  # Common defaults, can be anything
  ModFile: "<lang>.mods"
  SumFile: "<lang>.sums"
  ModsDir: "<lang>.mod/pkg"
  Checksum: "<lang>.mod/checksum.txt"

  # Controls for modders that want to shell out
  # to common tools for certain commands
  NoLoad: false
  CommandInit: []
  CommandGraph: []
  CommandTidy: []
  CommandVendor: []
  CommandVerify: []

  # Runs on init for this language
  # filename/content key/pair values
  # uses the golang text/template library
  # inputs will be
  #   .Language
  #   .Module
  #   .Modder
  InitTemplates:
      <lang>.mod/module.<lang>: |
          module "{{ .Module }}"
  # Series of commands to be executed pre/post init
  InitPreCommands: [][]string
  InitPostCommands: [][]string

  # Same as the InitTemplates, but run during vendor command
  VendorTemplates:
      <lang>.mod/module.<lang>: |
          module "{{ .Module }}"

  VendorIncludeGlobs:
    - ".mvsconfig.yaml"
    - "<lang>.mods"
    - "<lang>.sums"
    - "<lang>.mod/module.<lang>"
    - "<lang>.mod/modules.txt"
    - "**/*.<lang>"
  VendorExcludeGlobs:
    - "<lang>.mod/pkg"
    
  # Series of commands to be executed pre/post vendoring
  VendorPreCommands: [][]string
  VendorPostCommands: [][]string
  
  # Use MVS to only manage the languages normal dependency file
  ManageFileOnly: false
  
  # Whether local replaces should use a symlink instead of copying files
  SymlinkLocalReplaces: false

  # Controls the code introspection for dependency determiniation
  IntrospectIncludeGlobs:
    - "**/*.<lang>"
  IntrospectExcludeGlobs:
    - "<lang>.mod/pkg"
  IntrospectExtractRegex:
    - "you will have to figure out a series of 'any match passes' regexps to pull out dependencies"
    
  # This field determines the prefix to place in front of
  # imports which have a single token or leverage package managers
  # This is currently futurology for building MVS for Python and JavaScript
  PackageManagerDefaultPrefix: "npm.js"
```


