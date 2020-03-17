# mvs

A flexible MVS tool and library based on Go mods.

Make module systems with "Minimum Version Selection" semantics
for any git repository. With MVS, you can create 

# TODO write about

- concept
- diffs from go
- mod/sum file format
- .mvsconfig format


### Install

```shell
go get github.com/hofstadter-io/mvs
```

### Usage

```shell
mvs help

mvs init -l <lang> <module-path>

vim <lang>.mods

mvs vendor
```

### Module File

The module file holds the requirements for project.

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

### Development

Make sure you have cue installed.

```shell
# Vendor deps (go and cue)
mvs vendor

# Generate code
cue gen        # Generate gocode

# Build binary
go build
```

