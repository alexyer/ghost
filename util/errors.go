package util

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

func GhostErrorf(formatString string, a ...interface{}) error {
	return GhostBaseError(fmt.Sprintf(formatString, a...))
}

var GhostEmptyMsg = GhostBaseError("got empty message")
