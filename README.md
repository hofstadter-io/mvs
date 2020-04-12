# MVS

`mvs` is a flexible tool and library based on Go mods.

Use and create module systems with [Minimum Version Selection](https://research.swtch.com/vgo-mvs) semantics
and manage dependencies `go mod` style.
Mix any set of language, code bases, git repositories, package managers, and subdirectories.
Manage polyglot and monorepo codebase dependencies with
[100% reproducible builds](https://github.com/golang/go/wiki/Modules#version-selection) from a single tool.


### Features

- Based on go mods MVS system, aiming for 100% reproducible builds.
- Recursive dependencies, version resolution, and code instrospection.
- Custom module systems with custom file names and vendor directories.
- Control configuration for naming, vendoring, and other behaviors.
- Polyglot support for multiple module systems and multiple languages within one tool.
- Works with any git system and supports the main features from go mods.
- Convert other vendor and module systems to MVS or just manage their files with MVS.

Language support:

- [golang](https://golang.org) - delegates to the go tool for the commands mirrored here
- [cuelang](https://cuelang.org) - builtin in default using the custom module feature
- [hofmods](https://hofstadter.io) - extends Cue with advanced code generation capabilities
- [custom](./docs/custom-modders.md) - Create your own locally or globally with `.mvsconfig` files

Upcoming languages: Python and JavaScript
so they can have an MVS system and the benefits,
and `mvs` can start supporting and fetching from package registries.
These language implementations will have flexibility to
manage with mvs and the common toolchain to varying degrees.
Pull requests for improved language support are welcome.

The main difference from go mods is that `mvs`, generally,
is not introspecting your code to determine dependencies.
It relies on user management of the `<lang>.mods` file.
Since golang is exec'd out to, introspection is supported,
and as more languages improve, we look to similarly
improve this situation in `mvs`.
A midstep to full custom implementation will be a
introspection custom module with some basic support
using file globs and regex lists.

_Note, we also default to the plural `<lang>.mods/sums` files and `<lang.mod>/` vendor directory.
This is 1) because cuelang reads from `cue.mod` directory, and 2) because it is less likely
to collide with existing directories.
You can also configure more behavior per language and module than go mods.
The go mods is shelled out to as a convience, and often languages impose restrictions._


### Install

[Releases for Linux, Mac, and Windows are available](https://github.com/hofstadter-io/mvs/releases).

Development setup (requires the [hof tool](https://github.com/hofstadter-io/hof)):

```shell
git clone https://github.com/hofstadter-io/mvs
cd mvs

go mod vendor
go install
mvs vendor cue

hof gen
git status # should be no difference
go install
```

### Usage

```shell
# Print known languages in the current directory
mvs info

# Initialize this folder as a module
mvs init <lang> <module-path>

# Add your requirements
vim <lang>.mods  # go.mod like file

# Pull in dependencies, no args discovers by *.mods and runs all
mvs vendor [langs...]

# See all of the commands
mvs help
```


### Module File

The module file holds the requirements for project.
It has the same format as a `go.mod` file.

```
# These are like golang import paths
#   i.e. github.com/hofstadter-io/mvs
module <module-path> 

# Information about the module type / version
#  some systems that take this into account
# go = 1.14
<lang> = <version>

# Required dependencies section
require (
  # <module-path> <module-semver>
  github.com/hof-lang/cuemod--cli-golang v0.0.0      # This is latest on HEAD
  github.com/hof-lang/cuemod--cli-golang v0.1.5      # This is a tag v0.1.5 (can omit 'v' in tag, but not here)
)

# replace <module-path> => <module-path|local-path> [version]
replace github.com/hof-lang/cuemod--cli-golang => github.com/hofstadter-io/hofmod-cli-golang v0.2.0
replace github.com/hof-lang/cuemod--cli-golang => ../../cuelibs/cuemod--cli-golang
```


### Custom Module Systems

`.mvsconfig.cue` allows you to define custom module systems.
With some simple configuration, you can create and control
and vendored module system based on `go mods`.
You can also define global configurations.

See the [custom-modder docs](./docs/custom-modders.md)
to learn more about writing
you own module systems.

This is the current Cue __modder__ configuration:

```cue
cue: {
  Name: "cue"
  Version: "0.1.1"
  ModFile: "cue.mods"
  SumFile: "cue.sums"
  ModsDir: "cue.mod/pkg"
  MappingFile: "cue.mod/modules.txt"
  InitTemplates: {
    "cue.mod/module.cue": """
      module "{{ .Module }}"
      """
    }

  VendorIncludeGlobs: []
  VendorExcludeGlobs: [
    "/.git/**",
    "**/cue.mod/pkg/**",
  ]
}
```


### Development

Want to help out?

Here are some commands if you want to develop `mvs`.

Make sure you have go 1.14 and [cue installed](https://cuelang.org/docs/install/).
We are mainly [developing with cuelang tip](https://github.com/cuelang/cue/blob/master/doc/contribute.md#overview-1) (just step-1 at the link)
Development setup requires the [hof tool](https://github.com/hofstadter-io/hof) as well.

```shell
# Fetch deps (go and cue)
mvs vendor

# Generate code
hof gen        # Generate gocode for the cmd implementation

# Validate cuecode that is input to hof
cue eval
cue export

# Build binary
go build

# Run local mvs
./mvs help
```

You may also like to look at the [hofmod-cli](https://github.com/hofstadter-io/hofmod-cli) project.
We use this to generate the CLI code and files for CI.
The `mvs` uses our [hof tool](https://github.com/hofstadter-io/hof) for code generation of itself
and is also imported into the hof tool as both cue and go modules,
embedding the `mvs` cli as the subcommand `hof mod`.

You can find us in the
[cuelang slack](https://join.slack.com/t/cuelang/shared_invite/enQtNzQwODc3NzYzNTA0LTAxNWQwZGU2YWFiOWFiOWQ4MjVjNGQ2ZTNlMmIxODc4MDVjMDg5YmIyOTMyMjQ2MTkzMTU5ZjA1OGE0OGE1NmE)
for now.


### Motivation

- MVS has better semantics for vendoring and gets us closer to 100% reproducible builds.
- JS and Python can have MVS while still using the remainder of the tool chains.
- Prototype for cuelang module and vendor management.
- We need a module system for our [hof-lang](https://hof-lang.org) project.

#### Links about go mods

[Using go modules](https://blog.golang.org/using-go-modules)

[Go and Versioning](https://research.swtch.com/vgo)

[More about version selection](https://github.com/golang/go/wiki/Modules#version-selection)



