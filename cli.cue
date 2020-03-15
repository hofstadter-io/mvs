package cli

import (
  "github.com/hof-lang/cuemod--cli-golang:cli"
  "github.com/hof-lang/cuemod--cli-golang/schema"
)

Outdir: "./"

GEN : cli.Generator & {
  Cli: CLI
}

CLI : cli.Schema & {
  Name: "mvs"
  Package: "github.com/hofstadter-io/mvs"
  Short: "MVS is a vendor management tool based on go mods"
  OmitRun: true

  Commands: [
    schema.Command & {
      Name: "convert"
    },
    schema.Command & {
      Name: "download"
    },
    schema.Command & {
      Name: "edit"
    },
    schema.Command & {
      Name: "graph"
    },
    schema.Command & {
      Name: "init"
    },
    schema.Command & {
      Name: "tidy"
    },
    schema.Command & {
      Name: "vendor"
    },
    schema.Command & {
      Name: "verify"
    },
    schema.Command & {
      Name: "why"
    },
  ]
}


