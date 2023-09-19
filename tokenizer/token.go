package tokenizer

import "fmt"

type Token struct {
	Type  TokenType
	Value string
}

// IsTerm returns true if the token can be a term in an operation
func (t *Token) IsTerm() bool {
	if t.MatchAny(LITERAL, IDENTIFIER) {
		return true
	}
	return false
}

// IsOperator returns true if the token is an operator
func (t *Token) IsOperator() bool {
	if t.MatchAny(ADD, SUB, MUL, DIV) {
		return true
	}
	return false
}

// GetPrec returns the precedence of a token, only makes sense for operators
func (t *Token) GetPrec() int {
	switch {
	case t.MatchAny(ADD, SUB):
		return 1
	case t.MatchAny(MUL, DIV):
		return 2
	default:
		return 0
	}
}

// MatchAny returns true if the type of the token matches any of the given types
func (t *Token) MatchAny(tokenTypes ...TokenType) bool {
	for _, tt := range tokenTypes {
		if t.Type == tt {
			return true
		}
	}
	return false
}

// String returns a printable string that represents the token
func (t *Token) String() string {
	if t.Type == SEP {
		return fmt.Sprintf("(%s, \\n)", t.Type)
	} else {
		return fmt.Sprintf("(%s, %s)", t.Type, t.Value)
	}
}
