package tokenizer

import (
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

	result := token.MatchAny(SEP, LITERAL, B_L)

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
func TestIsTerm_WhenMatch(t *testing.T) {
	terms := []TokenType{LITERAL, IDENTIFIER}

	for _, term := range terms {
		t.Run(string(term),
			func(t *testing.T) {
				token := newTokenType(term)

				result := token.IsTerm()

				assert.True(t, result)
			},
		)
	}
}

func TestIsTerm_WhenDoesNotMatch(t *testing.T) {
	token := newTokenType(P_L)

	result := token.IsTerm()

	assert.False(t, result)
}

// TestIsOperator
func TestIsOperator_WhenMatch(t *testing.T) {
	operators := []TokenType{ADD, SUB, MUL, DIV}

	for _, operator := range operators {
		t.Run(string(operator),
			func(t *testing.T) {
				token := newTokenType(operator)

				result := token.IsOperator()

				assert.True(t, result)
			},
		)
	}
}

func TestIsOperator_WhenDoesNotMatch(t *testing.T) {
	token := newTokenType(EQ)

	result := token.IsOperator()

	assert.False(t, result)
}

// TestGetPrec
func TestGetPrec(t *testing.T) {
	cases := map[TokenType]int{
		ADD: 1,
		SUB: 1,
		MUL: 2,
		DIV: 2,
		SEP: 0,
		EQ:  0,
		P_L: 0,
		P_R: 0,
	}

	for tokenType, expected := range cases {
		t.Run(string(tokenType),
			func(t *testing.T) {
				token := newTokenType(tokenType)

				result := token.GetPrec()

				assert.Equal(t, expected, result)
			},
		)
	}
}

// TestString
func TestString_WhenTypeIsSep(t *testing.T) {
	token := newToken(SEP, "should-not-print")

	result := token.String()

	assert.Equal(t, "(SEP, '\\n')", result)
}

func TestString_WhenTypeIsAny(t *testing.T) {
	token := newToken(LITERAL, "value")

	result := token.String()

	assert.Equal(t, "(LITERAL, 'value')", result)
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
