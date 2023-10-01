package parser

import (
	"compiler/tokenizer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeScope_WhenNoTokensLeft(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	scope, err := p.parseNodeScope()

	assert.Error(t, err)
	assert.Nil(t, scope)
	assert.Equal(t, "Invalid scope: No tokens left", err.Error())
}

func TestParseNodeScope_WhenUnexpectedToken(t *testing.T) {
	p := NewParser(generateTokensFor("-"))

	scope, err := p.parseNodeScope()

	assert.Error(t, err)
	assert.Nil(t, scope)
	assert.Equal(t, "Invalid scope: Unexpected token (SUB, '-') at line 1 and column 1", err.Error())
}

func TestParseNodeScope_WhenInvalidStatement(t *testing.T) {
	p := NewParser(generateTokensFor("{int =}"))

	scope, err := p.parseNodeScope()

	assert.Error(t, err)
	assert.Nil(t, scope)
	assert.Equal(t, "Invalid statement: Unexpected token (EQ, '=') at line 1 and column 6", err.Error())
}

func TestParseNodeScope_WhenValidScopes(t *testing.T) {
	scopes := []string{
		"{a=1}",
		"{a = 1\n b = 3}",
		"{\na = 1\n b = 3\n}",
	}

	for _, scope := range scopes {
		t.Run(scope,
			func(t *testing.T) {
				tokens := generateTokensFor(scope)
				p := NewParser(tokens)

				scope, err := p.parseNodeScope()

				assert.NoError(t, err)
				assert.NotNil(t, scope)
				assert.Equal(t, len(tokens), p.index)
			},
		)
	}
}
