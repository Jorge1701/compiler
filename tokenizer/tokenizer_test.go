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

func TestGenerateTokens_AllCases(t *testing.T) {
	allTokens := "\n(){}[]+-*/=int bool exit a 1 true"

	tokekizer := NewTokenizer([]rune(allTokens))

	err := tokekizer.GenerateTokens()

	assert.NoError(t, err)
	matchedAll := assert.Equal(t, len(allTokenTypes), len(tokekizer.GetTokens()),
		"String allTokens does not provide an example of all the tokens",
	)

	// Check if there is an example for every posible token type
	if matchedAll {
		for i, token := range tokekizer.GetTokens() {
			assert.Equal(t, allTokenTypes[i], token.Type,
				fmt.Sprintf("Token at position %d does not match", i),
			)
		}
	}
}

func TestGenerateTokens_ValueCases(t *testing.T) {
	text := "vAriaBlE2Name23 true other_name ALL_CAPS a 1 32 false 436 12353"

	tokenizer := NewTokenizer([]rune(text))

	err := tokenizer.GenerateTokens()

	assert.NoError(t, err)
	assertTokens(t, tokenizer,
		"(IDENTIFIER, 'vAriaBlE2Name23')",
		"(BOOL_LITERAL, 'true')",
		"(IDENTIFIER, 'other_name')",
		"(IDENTIFIER, 'ALL_CAPS')",
		"(IDENTIFIER, 'a')",
		"(INT_LITERAL, '1')",
		"(INT_LITERAL, '32')",
		"(BOOL_LITERAL, 'false')",
		"(INT_LITERAL, '436')",
		"(INT_LITERAL, '12353')",
	)

	// Check if there is at least one example for every token type that expects a value
	typeHasExample := map[TokenType]bool{}
	for _, tokenType := range valueTokenTypes {
		typeHasExample[tokenType] = false
	}

	for _, token := range tokenizer.GetTokens() {
		typeHasExample[token.Type] = true
	}

	for tokenType, hasExample := range typeHasExample {
		if !hasExample {
			assert.Fail(t, fmt.Sprintf("An example for token type '%s' was not provided", tokenType))
		}
	}
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
