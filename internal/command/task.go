package command

import "errors"

var errCommandRequired = errors.New("Command required")
var errTaskNameRequired = errors.New("Task name required")
var errCommandNotFound = errors.New("Command not found")

func handleTaskCommands(args []string) error {
	if len(args) < 1 {
		return errCommandRequired
	}

	if len(args) != 2 {
		return errTaskNameRequired
	}

	switch args[0] {
	case "create":
		return createTask(args[1])
	case "delete":
		//
	case "complete":
		//
	}

	return errCommandNotFound
}

func createTask(taskName string) error {

}
