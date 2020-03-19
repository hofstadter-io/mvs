package cli

import (
	"github.com/hof-lang/cuemod--cli-golang:cli"
	"github.com/hof-lang/cuemod--cli-golang/schema"
)

Outdir: "./"

GEN :: cli.Generator & {
	Cli: CLI
}

_CmdImports :: [
	schema.Import & {Path: CLI.Package + "/lib"},
]

CLI :: cli.Schema & {
	Name:    "mvs"
	Package: "github.com/hofstadter-io/mvs"

	Usage: "mvs"
	Short: "MVS is a polyglot vendor management tool based on go mods"
	Long: """
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
			Name:    "lang"
			Type:    "string"
			Default: ""
			Help:    "The language or system prefix to process. The default is to discover and process all known."
			Long:    "lang"
			Short:   "l"
		},
	]

	Imports: [
		schema.Import & {Path: CLI.Package + "/lib"},
	]

	PersistentPrerun: true
	PersistentPrerunBody: """
    // fmt.Println("PersistentPrerun", RootLangPflag, args)
    lib.InitLangs()
  """

	Commands: [
		schema.Command & {
			Name:  "langinfo"
			Usage: "langinfo [language]"
			Short: "print info about languages and modders known to mvs"
			Long: """
        print info about languages and modders known to mvs
          - no arg prints a list of known languages
          - an arg prints info about the language modder configuration that would be used
      """

			Args: [
				schema.Arg & {
					Name: "lang"
					Type: "string"
					Help: "name of the language to print info about"
				},
			]

			Imports: _CmdImports

			Body: """
      msg, err := lib.LangInfo(lang)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      fmt.Println(msg)
      """
		},
		schema.Command & {
			Name:  "convert"
			Usage: "convert <lang> <file>"
			Short: "convert another package system to MVS."
			Long:  Short

			Args: [
				schema.Arg & {
					Name:     "lang"
					Type:     "string"
					Required: true
					Help:     "name of the language to print info about"
				},
				schema.Arg & {
					Name:     "filename"
					Type:     "string"
					Required: true
					Help:     "existing package filename, depending on language"
				},
			]

			Imports: _CmdImports

			Body: """
      err := lib.Convert(lang, filename)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},
		schema.Command & {
			Name:  "graph"
			Usage: "graph"
			Short: "print module requirement graph"
			Long:  Short

			Imports: _CmdImports

			Body: """
      err := lib.Graph(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},
		schema.Command & {
			Name:  "init"
			Usage: "init <lang> <module>"
			Short: "initialize a new module in the current directory"
			Long:  Short

			Args: [
				schema.Arg & {
					Name:     "lang"
					Type:     "string"
					Required: true
					Help:     "name of the language to print info about"
				},
				schema.Arg & {
					Name:     "module"
					Type:     "string"
					Required: true
					Help:     "module name or path, depending on language"
				},
			]

			Imports: _CmdImports

			Body: """
      err := lib.Init(lang, module)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},
		schema.Command & {
			Name:  "tidy"
			Usage: "tidy"
			Short: "add missinad and remove unused modules"
			Long:  Short

			Imports: _CmdImports

			Body: """
      err := lib.Tidy(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},
		schema.Command & {
			Name:  "vendor"
			Usage: "vendor"
			Short: "make a vendored copy of dependencies"
			Long:  Short

			Imports: _CmdImports

			Body: """
      err := lib.Vendor(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},
		schema.Command & {
			Name:  "verify"
			Usage: "verify"
			Short: "verify dependencies have expected content"
			Long:  Short

			Imports: _CmdImports

			Body: """
      err := lib.Verify(RootLangPflag)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},
		schema.Command & {
			Name:   "hack"
			Usage:  "hack"
			Short:  "dev command"
			Long:   Short
			Hidden: true

			Imports: _CmdImports

			Body: """
      err := lib.Hack(RootLangPflag, args)
      if err != nil {
        fmt.Println(err)
        os.Exit(1)
      }
      """
		},

	]
}
