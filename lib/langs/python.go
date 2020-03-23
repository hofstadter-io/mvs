package langs

const PythonModder = `
python: {
	Name:          "python",
	Version:       "3.8",
	ModFile:       "python.mod",
	SumFile:       "requirements.txt",
	ModsDir:       "vendor",
	MappingFile:   "vendor/modules.txt",
	CommandInit:   [["python", "-m", "venv", "venv"]],
	CommandVendor: [["bash", "-c", ". ./venv/bin/activate && pip install -r requirements.txt"]],
}
`
