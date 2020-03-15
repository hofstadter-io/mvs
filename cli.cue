package cli

import (
  "github.com/hof-lang/cuemod--cli-golang:cli"
  "github.com/hof-lang/cuemod--cli-golang/schema"
)

Outdir: "./"

GEN : cli.Generator & {
  Cli: CLI
}

_CmdImports : [
  schema.Import & { Path: CLI.Package + "/pkg" }
]


CLI : cli.Schema & {
  Name: "mvs"
  Package: "github.com/hofstadter-io/mvs"

  Usage: "mvs"
  Short: "MVS is a polyglot vendor management tool based on go mods"
  Long:  """
  MVS is a polyglot vendor management tool based on go mods.

  mod file format:

    module = "<module path>"

    <name> = "version"

    require (
      ...
    )

    replace <module path> => <local path>
    ...
  """


  OmitRun: true

  Pflags: [
    schema.Flag & {
      Name: "lang"
      Type: "string"
      Default: ""
      Help: "The language or system prefix to process. The default is to discover and process all known."
      Long: "lang"
      Short: "l"
    },
    schema.Flag & {
      Name: "dryrun"
      Type: "string"
      Default: ""
      Help: "Print the command and do not execute."
      Long: "dry-run"
      Short: "d"
    }
  ]

  PersistentPrerun: true
  PersistentPrerunBody: """
    // fmt.Println("PersistentPrerun", RootLangPflag, args)
  """

  Commands: [
    schema.Command & {
      Name:   "convert"
      Usage:  "convert -l <lang> <file>"
      Short:  "convert another package system to MVS."
      Long:   "convert another package system to MVS, language flag is required"

      Args: [
        schema.Arg & {
          Name: "filename"
          Type: "string"
          Required: true
          Help: "existing package filename, depending on language"
        }
      ]

      Imports: _CmdImports

      Body: """
      if RootLangPflag == "" {
        fmt.Println("language flag is required for this command")
        cmd.Usage()
        os.Exit(1)
      }
      err := pkg.Convert(RootLangPflag, filename)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """

    },
    schema.Command & {
      Name:   "graph"
      Usage:  "graph"
      Short:  "print module requirement graph"
      Long:   Short

      Imports: _CmdImports

      Body: """
      err := pkg.Graph(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """

    },
    schema.Command & {
      Name:   "init"
      Usage:  "init -l <lang> <module>"
      Short:  "initialize a new module in the current directory"
      Long:   "initialize a new module in the current directory, language flag is required"

      Args: [
        schema.Arg & {
          Name: "module"
          Type: "string"
          Required: true
          Help: "module name or path, depending on language"
        }
      ]

      Imports: _CmdImports

      Body: """
      if RootLangPflag == "" {
        fmt.Println("language flag is required for this command")
        cmd.Usage()
        os.Exit(1)
      }
      err := pkg.Init(RootLangPflag, module)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """

    },
    schema.Command & {
      Name:   "tidy"
      Usage:  "tidy"
      Short:  "add missinad and remove unused modules"
      Long:   Short

      Imports: _CmdImports

      Body: """
      err := pkg.Tidy(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """

    },
    schema.Command & {
      Name:   "vendor"
      Usage:  "vendor"
      Short:  "make a vendored copy of dependencies"
      Long:   Short

      Imports: _CmdImports

      Body: """
      err := pkg.Vendor(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """

    },
    schema.Command & {
      Name:   "verify"
      Usage:  "verify"
      Short:  "verify dependencies have expected content"
      Long:   Short

      Imports: _CmdImports

      Body: """
      err := pkg.Verify(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """

    },
  ]
}


