package shell

import (
	"os"
	"os/exec"

	"github.com/jhacobs/jarvis/internal/command"
)

func runShellCommands(input string, args []string) error {
	switch input {
	case "exit":
		os.Exit(0)
	case "cd":
		return command.ChangeDir(args)
	}

	return runDefaultCommands(input, args)
}

func runDefaultCommands(input string, args []string) error {
	cmd := exec.Command(input, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
