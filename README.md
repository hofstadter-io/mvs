# MVS

A flexible MVS tool and library based on Go mods.

Create module systems with [Minimum Version Selection](https://research.swtch.com/vgo-mvs) semantics
for any language, and generally any set of git repositories.
Manage polyglot and monorepo codebase dependencies
with more stability from a single tool.


### Features

- Based on go mods MVS system for better vendor stability.
- Custom module systems with custom file names and vendor directories.
- Recursive dependencies, version resolution, and code instrospection.
- Control configuration for naming, vendoring, and other behaviors.
- Polyglot support for multiple module systems and multiple languages within one tool.
- Works with any git system and supports the main features from go mods.
- Convert other vendor and module systems to MVS or just manage their files with MVS.

Language support:

- [Golang](https://golang.org) - exec's out to go tool
- [Cuelang](https://cuelang.org) - builtin in default using the custom module feature
- [Hoflang](https://hof-lang.org) - extends Cuelang with low-code capabilities (also a builtin custom)
- Custom - Create your own locally or globally with `.mvsconfig` files

Upcoming languages: Python and JavaScript so they
can have an MVS system and the benefits,
and `mvs` can start supporting package registries.
These language implementations may end up generation
the language specific files with version setup so that
the MVS resolution is realized with npm and pip.
We are also planning to fetch from common package managers.
Pull requests for new languages are welcome.

The main difference from go mods is that `mvs`, generally,
is not introspecting your code to determine dependencies.
It relies on user management of the `<lang>.mods` file.
Since golang is exec'd out to, introspection is supported,
and as more languages improve, we look to similarly
improve this situation in `mvs`.

_Note, we also default to the plural `<lang>.mods/sums` files and `<lang.mod>/` vendor directory.
This is 1) because cuelang reads from `cue.mod` directory, and 2) because it is less likely
to collide with existing directories.
You can also configure more behavior per language and module than go mods.
The go mods is shelled out to as a convience, and often languages impose restrictions._


### Install

```shell
go get github.com/hofstadter-io/mvs
```


### Usage

```shell
# Initialize this folder as a module
mvs init -l <lang> <module-path>

# Add your requirements
vim <lang>.mods  # go.mod like file

# Pull in dependencies
mvs vendor [-l <lang>]

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

require (
  # <module-path> <module-semver>
  github.com/hof-lang/cuemod--cli-golang v0.0.0      # This is latest on HEAD
  github.com/hof-lang/cuemod--cli-golang v0.1.5      # This is a tag v0.1.5 (can omit 'v' in tag, but not here)
)

replace (
  <module-path> <module-path|local-path>
  github.com/hof-lang/cuemod--cli-golang => github.com/hofstadter-io/hofmod-cli-golang
  github.com/hof-lang/cuemod--cli-golang => ../../cuelibs/cuemod--cli-golang
)
```


### Custom Module Systems

`.mvsconfig` allows you to define custom module systems.
With some simple configuration, you can create and control
and vendored module system based on `go mods`.

Coming soon


### Development

Want to help out?

Here are sommands if you want to develop `mvs`.

Make sure you have go 1.14 and [cue installed](https://cuelang.org/docs/install/).
We are mainly [developing with cuelang tip](https://github.com/cuelang/cue/blob/master/doc/contribute.md#overview-1) (just step-1 at the link)

```shell
# Fetch deps (go and cue)
go mod vendor  # or... mvs vendor -l go
mvs vendor -l cue

# Generate code
cue gen        # Generate gocode

# Build binary
go build

# Run local mvs
./mvs help
```

You can find us in the
[cucelang slack](https://join.slack.com/t/cuelang/shared_invite/enQtNzQwODc3NzYzNTA0LTAxNWQwZGU2YWFiOWFiOWQ4MjVjNGQ2ZTNlMmIxODc4MDVjMDg5YmIyOTMyMjQ2MTkzMTU5ZjA1OGE0OGE1NmE)
for now.


### Rational

- MVS has better semantics for vendoring and leads to more stable code
- Prototype for Cuelang module and vendor management
- JS and Python can have MVS while still using the remainder of the tool chains.

links about go mods

[Using go modules](https://blog.golang.org/using-go-modules)

[Go and Versioning](https://research.swtch.com/vgo)
