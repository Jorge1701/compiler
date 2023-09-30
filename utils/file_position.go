package utils

type FilePosition struct {
	line   int
	column int
}

// NewPosition returns a FilePosition with line and column
func NewPosition(line, column int) *FilePosition {
	return &FilePosition{line, column}
}
