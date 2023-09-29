package error_wrapper

import "fmt"

type Error struct {
	Msg    string
	Row    int
	Column int
}

func NewError(msg string, row, column int) *Error {
	return &Error{
		Msg:    msg,
		Row:    row,
		Column: column,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s at line %d and column %d", e.Msg, e.Row, e.Column)
}
