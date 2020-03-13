# mvs

A flexible MVS tool and library based on Go mods.

### Development

Make sure you have cue installed.

```shell
# Import the vendor (while this tool is not working)
mkdir -p cue.mod/pkg/github.com/hof-lang
pushd cue.mod/pkg/github.com/hof-lang
git clone https://github.com/hof-lang/cuemod--cli-golang
popd

# Generate code
cue gen        # Generate gocode

# Golang stuff
cue init       # Go mod init for the cue cli package
go mod vendor  # or just 'go mod'
go build

# MVS Help
./mvs help
```

