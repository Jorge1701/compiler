package utils

import (
	"fmt"
)

type Error struct {
	Msg    string
	Row    int
	Column int
}

func NewError(msg string, pos *FilePosition) *Error {
	return &Error{
		Msg:    msg,
		Row:    pos.Row,
		Column: pos.Column,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s at line %d and column %d", e.Msg, e.Row, e.Column)
}
