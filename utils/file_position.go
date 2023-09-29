package utils

type FilePosition struct {
	Row    int
	Column int
}

// NewPosition returns a FilePosition with row and column
func NewPosition(row, column int) *FilePosition {
	return &FilePosition{
		Row:    row,
		Column: column,
	}
}
