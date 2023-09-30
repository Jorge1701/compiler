package parser

import (
	"compiler/tokenizer"
	"compiler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeTerm_WhenTypeLiteral(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.LITERAL, Value: "1"},
		{Type: tokenizer.IDENTIFIER, Value: "name"},
	})

	term, err := p.parseNodeTerm()

	assert.NotNil(t, term)
	assert.NoError(t, err)
	assert.Equal(t, TypeNodeTermLit, term.T)
	assert.Equal(t, "1", term.Lit.Value)
}

func TestParseNodeTerm_WhenTypeIdentifier(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.LITERAL, Value: "1"},
	})

	term, err := p.parseNodeTerm()

	assert.NotNil(t, term)
	assert.NoError(t, err)
	assert.Equal(t, TypeNodeTermIdent, term.T)
	assert.Equal(t, "name", term.Ident.Value)
}

func TestParseNodeTerm_WhenNoTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	term, err := p.parseNodeTerm()

	assert.Error(t, err)
	assert.Nil(t, term)
	assert.Equal(t, "Invalid term: No tokens left", err.Error())
}

func TestParseNodeTerm_WhenInvalidTerm(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.P_L, Value: "(", Pos: utils.NewPosition(2, 5)},
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.LITERAL, Value: "1"},
	})

	term, err := p.parseNodeTerm()

	assert.Nil(t, term)
	assert.Error(t, err)
	assert.Equal(t, "Invalid term: Unexpected token (P_L, '(') at line 2 and column 5", err.Error())
}
