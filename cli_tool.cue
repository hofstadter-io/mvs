package cli

import (
	"path"
	"tool/cli"
	"tool/exec"
	"tool/file"
)

command: gen: {
	task: step_1: exec.Run & {
		cmd: ["cue", "render"]
		stdout: string
	}
	task: step_2: exec.Run & {
		$after: task.step_1
		cmd: ["cue", "format"]
		stdout: string
	}
}

command: render: {

	var: {
		outdir: Outdir
	}

	for i, F in GEN.Out {

		task: "mkdir-\(i)": exec.Run & {
			cmd: ["mkdir", "-p", var.outdir + path.Dir(F.Filename)]
			stdout: string
		}

		task: "write-\(i)": file.Create & {
			deps: [ task["mkdir-\(i)"].stdout]

			filename: var.outdir + F.Filename
			contents: F.Out
			stdout:   string
		}

		task: "print-\(i)": cli.Print & {
			deps: [ task["write-\(i)"].stdout]
			text: task["write-\(i)"].filename
		}

	}

}

command: format: {
	var: {
		outdir: Outdir
	}

	task: shell: exec.Run & {
		cmd: ["bash", "-c", "cd \(var.outdir) && goimports -w -l ."]
		stdout: string
	}
}

command: init: {
	var: {
		outdir: Outdir
	}

	task: shell: exec.Run & {
		cmd: ["bash", "-c", "cd \(var.outdir) && go mod init \(CLI.Package)"]
		stdout: string
	}

	task: vendor: exec.Run & {
		dep: [ task.shell.stdout]
		cmd: ["bash", "-c", "cd \(var.outdir) && go mod vendor"]
		stdout: string
	}

}

command: vendor: {
	var: {
		outdir: Outdir
	}

	task: vendor: exec.Run & {
		cmd: ["bash", "-c", "cd \(var.outdir) && go mod vendor"]
		stdout: string
	}

}

command: build: {
	var: {
		outdir: Outdir
	}

	task: shell: exec.Run & {
		cmd: ["bash", "-c", "cd \(var.outdir) && go build -o \(CLI.Name) ."]
		stdout: string
	}

}