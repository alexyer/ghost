package server

import (
	"errors"
	"fmt"
)

func GhostBaseError(msg string) error {
	return errors.New(fmt.Sprintf("ghost: %s", msg))
}

func GhostCmdError(cmdName, msg string) error {
	return GhostBaseError(fmt.Sprintf("%s: %s", cmdName, msg))
}
