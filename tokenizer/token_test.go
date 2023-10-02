package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestMatchAny
func TestMatchAny_WhenMatch(t *testing.T) {
	token := newTokenType(IDENTIFIER)

	result := token.MatchAny(SEP, IDENTIFIER, B_L)

	assert.True(t, result)
}

func TestMatchAny_WhenDoesNotMatch(t *testing.T) {
	token := newTokenType(IDENTIFIER)

	result := token.MatchAny(SEP, INT_LITERAL, B_L)

	assert.False(t, result)
}

// TermIsType
func TestIsType_WhenMatch(t *testing.T) {
	token := newTokenType(IDENTIFIER)

	result := token.IsType(IDENTIFIER)

	assert.True(t, result)
}

func TestIsType_WhenDoesNotMatch(t *testing.T) {
	token := newTokenType(IDENTIFIER)

	result := token.IsType(P_L)

	assert.False(t, result)
}

// TestIsTerm
func TestIsTerm(t *testing.T) {
	term := []TokenType{IDENTIFIER, INT_LITERAL, BOOL_LITERAL}

	for _, tokenType := range allTokenTypes {
		t.Run(string(tokenType),
			func(t *testing.T) {
				token := newTokenType(tokenType)

				result := token.IsTerm()

				assert.Equal(t, token.MatchAny(term...), result)
			},
		)
	}
}

// TestIsOperator
func TestIsOperator(t *testing.T) {
	operators := []TokenType{ADD, SUB, DIV, MUL}

	for _, tokenType := range allTokenTypes {
		t.Run(string(tokenType),
			func(t *testing.T) {
				token := newTokenType(tokenType)

				result := token.IsOperator()

				assert.Equal(t, token.MatchAny(operators...), result)
			},
		)
	}
}

// TestGetPrec
func TestGetPrec(t *testing.T) {
	tokensWithPrecDefined := map[TokenType]int{
		ADD: 1,
		SUB: 1,
		MUL: 2,
		DIV: 2,
	}

	for _, tokenType := range allTokenTypes {
		t.Run(string(tokenType),
			func(t *testing.T) {
				token := newTokenType(tokenType)

				result := token.GetPrec()

				if prec, isDefined := tokensWithPrecDefined[tokenType]; isDefined {
					assert.Equal(t, prec, result)
				} else {
					assert.Equal(t, 0, result)
				}
			},
		)
	}
}

// TestString
func TestString(t *testing.T) {
	for _, tokenType := range allTokenTypes {
		t.Run(string(tokenType),
			func(t *testing.T) {
				token := newToken(tokenType, "expected")

				result := token.String()

				if tokenType == SEP {
					assert.Equal(t, "(SEP, '\\n')", result)
				} else {
					assert.Equal(t, fmt.Sprintf("(%s, 'expected')", tokenType), result)
				}
			},
		)
	}
}

func newTokenType(t TokenType) *Token {
	return &Token{
		Type: t,
	}
}

func newToken(t TokenType, v string) *Token {
	return &Token{
		Type:  t,
		Value: v,
	}
}
