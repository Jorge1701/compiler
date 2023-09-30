package utils

import (
	"fmt"
)

type Error struct {
	msg string
	pos *FilePosition
}

func NewError(msg string, pos *FilePosition) *Error {
	return &Error{msg, pos}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s at line %d and column %d", e.msg, e.pos.line, e.pos.column)
}
