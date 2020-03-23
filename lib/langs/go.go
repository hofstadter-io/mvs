package langs

const GolangModder = `
go: {
	Name:          "go",
	Version:       "1.14",
	ModFile:       "go.mod",
	SumFile:       "go.sum",
	ModsDir:       "vendor",
	MappingFile:   "vendor/modules.txt",
	CommandInit:   [["go", "mod", "init"]],
	CommandGraph:  [["go", "mod", "graph"]],
	CommandTidy:   [["go", "mod", "tidy"]],
	CommandVendor: [["go", "mod", "vendor"]],
	CommandVerify: [["go", "mod", "verify"]],
}
`
