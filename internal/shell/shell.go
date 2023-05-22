package shell

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/fatih/color"
)

func Start() {
	handleExitSignals()

	reader := bufio.NewReader(os.Stdin)

	for {
		currentDir, err := getCurrentDir()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		fmt.Print("> ")
		color.New(color.FgRed, color.Bold).Print(currentDir + " ")

		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		if err = executeCommand(input); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func executeCommand(input string) error {
	input = strings.TrimSuffix(input, "\n")
	args := strings.Fields(input)

	if len(args) == 0 {
		return nil
	}

	return runShellCommands(args[0], args[1:])
}

func getCurrentDir() (string, error) {
	currentDir, err := os.Getwd()

	if err != nil {
		return currentDir, err
	}

	parts := strings.Split(currentDir, "/")
	return parts[len(parts)-1], nil
}

func handleExitSignals() {
	exitSignals := make(chan os.Signal, 1)
	signal.Notify(exitSignals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-exitSignals
		os.Exit(0)
	}()
}
