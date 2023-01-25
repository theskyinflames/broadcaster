package app

import (
	"fmt"
)

// InvalidCommandError should be returned by the implementations of the interface when the handler does not receive the needed command.
type InvalidCommandError struct {
	expected string
	had      string
}

// NewInvalidCommandError is a constructor
func NewInvalidCommandError(expected string, had string) InvalidCommandError {
	return InvalidCommandError{expected: expected, had: had}
}

const errMsgInvalidCommand = "invalid command, expected '%s' but found '%s'"

func (e InvalidCommandError) Error() string {
	return fmt.Sprintf(errMsgInvalidCommand, e.expected, e.had)
}

// InvalidQueryError is self described
type InvalidQueryError struct {
	Expected string
	Had      string
}
