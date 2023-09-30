package tokenizer

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateTokens_WhenEmptyRunes(t *testing.T) {
	tokenizer := NewTokenizer([]rune(""))

	err := tokenizer.GenerateTokens()

	assert.NoError(t, err)
	assert.Empty(t, tokenizer.GetTokens())
}

func TestGenerateTokens_UnexpectedSymbols(t *testing.T) {
	unexpectedSymbols := "|!\"#$%&?¡°¬@·~\\¿'´.,:;¨<>"

	for _, unexpectedSymbol := range unexpectedSymbols {
		t.Run(string(unexpectedSymbol),
			func(t *testing.T) {
				tokenizer := NewTokenizer([]rune{unexpectedSymbol})

				err := tokenizer.GenerateTokens()

				assert.Error(t, err)
				assert.Equal(t, fmt.Sprintf("Unexpected symbol '%c' at line 1 and column 1", unexpectedSymbol), err.Error())
			},
		)
	}
}

func TestGenerateTokens_UnexpectedSymbolAtLineAndColumn(t *testing.T) {
	text := "int a=1\na=!"

	tokenizer := NewTokenizer([]rune(text))

	err := tokenizer.GenerateTokens()

	assert.Error(t, err)
	assert.Equal(t, "Unexpected symbol '!' at line 2 and column 3", err.Error())
}

func TestGenerateTokens(t *testing.T) {
	text := "int vAriaBlE2Name23=1235\n{([exit])}\n+-*/"

	tokenizer := NewTokenizer([]rune(text))

	err := tokenizer.GenerateTokens()

	assert.NoError(t, err)
	assertTokens(t, tokenizer,
		"(INT, 'int')",
		"(IDENTIFIER, 'vAriaBlE2Name23')",
		"(EQ, '=')",
		"(LITERAL, '1235')",
		"(SEP, '\\n')",
		"(B_L, '{')",
		"(P_L, '(')",
		"(SB_L, '[')",
		"(EXIT, 'exit')",
		"(SB_R, ']')",
		"(P_R, ')')",
		"(B_R, '}')",
		"(SEP, '\\n')",
		"(ADD, '+')",
		"(SUB, '-')",
		"(MUL, '*')",
		"(DIV, '/')",
	)
}

// assertTokens asserts that every generated token should match the expected
func assertTokens(t *testing.T, tokenizer *Tokenizer, expectedTokens ...string) {
	equalLength := assert.Equal(t, len(expectedTokens), len(tokenizer.GetTokens()),
		"Length of generated tokens does not match",
	)

	if equalLength {
		for i, actualToken := range tokenizer.GetTokens() {
			tokenMatches := assert.Equal(t, expectedTokens[i], actualToken.String(),
				fmt.Sprintf("Token at position %d does not match", i),
			)

			// Returns after the first one that does not match get a shorter fail message
			if !tokenMatches {
				return
			}
		}
	}
}
