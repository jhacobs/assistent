package command

import (
	"errors"
	"os"
)

var errNoPath = errors.New("path required")

func ChangeDir(args []string) error {
	if len(args) < 1 {
		return errNoPath
	}

	return os.Chdir(args[0])
}
