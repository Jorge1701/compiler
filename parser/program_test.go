package parser

import (
	"compiler/tokenizer"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeProg_WhenNoTokensLeft(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	prog, err := p.parseNodeProg()

	assert.Error(t, err)
	assert.Nil(t, prog)
	assert.Equal(t, "Invalid program: No tokens left", err.Error())
}

func TestParseNodeProg_WhenInvalidStatement(t *testing.T) {
	p := NewParser(generateTokensFor("int ="))

	prog, err := p.parseNodeProg()

	assert.Error(t, err)
	assert.Nil(t, prog)
	assert.Equal(t, "Invalid statement: Unexpected token (EQ, '=') at line 1 and column 5", err.Error())
}

func TestParseNodeProg_WhenUnexpectedToken(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1\nexit a-"))

	prog, err := p.parseNodeProg()

	assert.Error(t, err)
	assert.Nil(t, prog)
	assert.Equal(t, "Unexpected token (SUB, '-') at line 2 and column 7", err.Error())
}

func TestParseNodeProg_WhenValidProgram(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1\nexit a"))

	prog, err := p.parseNodeProg()

	assert.NoError(t, err)
	assert.NotNil(t, prog)
	assert.Equal(t, 7, p.index)
}
