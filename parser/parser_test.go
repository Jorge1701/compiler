package parser

import (
	"compiler/tokenizer"
	"compiler/utils"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestGenerateNodes
func TestParser_GenerateNodes_WhenNoTokensLeft(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	err := p.GenerateNodes()

	assert.Error(t, err)
	assert.Equal(t, "Invalid program: No tokens left", err.Error())
}

func TestParser_GenerateNodes_WhenOk(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 2 + 3\n{a=6}\nexit a"))

	err := p.GenerateNodes()

	assert.NoError(t, err)
	assert.NotNil(t, p.GetNodes())
}

// TestNoTokensLeft
func TestParser_NoTokensLeft(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	err := p.noTokensLeft()

	assert.NotNil(t, err)
	assert.Equal(t, "No tokens left", err.Error())
}

// TestUnexpectedToken
func TestParser_UnexpectedToken(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.IDENTIFIER, Value: "name", Pos: utils.NewPosition(532, 23)},
	})

	err := p.unexpectedToken()

	assert.NotNil(t, err)
	assert.Equal(t, "Unexpected token (IDENTIFIER, 'name') at line 532 and column 23", err.Error())
}

func TestParser_UnexpectedToken_WhenNoTokensLeft(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	err := p.unexpectedToken()

	assert.NotNil(t, err)
	assert.Equal(t, "No tokens left", err.Error())
}

// TestUnexpectedTokenAt
func TestParser_UnexpectedTokenAt_WhenNoTokensLeft(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	err := p.unexpectedTokenAt(3)

	assert.NotNil(t, err)
	assert.Equal(t, "No tokens left", err.Error())
}

func TestParser_UnexpectedTokenAt_WhenLineBreak(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.SEP, Pos: utils.NewPosition(132, 43)},
	})

	err := p.unexpectedTokenAt(0)

	assert.NotNil(t, err)
	assert.Equal(t, "Unexpected line break at line 132 and column 43", err.Error())
}

func TestParser_UnexpectedTokenAt(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.LITERAL, Value: "123", Pos: utils.NewPosition(53, 2)},
	})

	err := p.unexpectedTokenAt(0)

	assert.NotNil(t, err)
	assert.Equal(t, "Unexpected token (LITERAL, '123') at line 53 and column 2", err.Error())
}

// TestConsume
func TestParser_Consume(t *testing.T) {
	tokens := generateTokensFor("int number = 123\n")
	p := NewParser(tokens)

	for i := range tokens {
		token := p.consume()
		assert.NotNil(t, token, fmt.Sprintf("Consumed token %d produced nil", i))
		assert.Equal(t, i+1, p.index)
	}

	assert.Nil(t, p.consume())
	assert.False(t, p.hasToken())
	assert.Equal(t, 5, p.index)
}

func TestParser_ConsumeEmpty(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	token := p.consume()

	assert.Nil(t, token)
}

// TestPeek
func TestParser_Peek(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1"))

	p.consume()
	token := p.peek()

	assert.NotNil(t, token)
	assert.Equal(t, tokenizer.IDENTIFIER, token.Type)
	assert.Equal(t, 1, p.index)
}

func TestParser_Peek_WhenNoTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	token := p.peek()

	assert.Nil(t, token)
}

// TestPeekAhead
func TestParser_PeekAhead(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1"))

	token := p.peekAhead(1)

	assert.NotNil(t, token)
	assert.Equal(t, tokenizer.IDENTIFIER, token.Type)
	assert.Equal(t, 0, p.index)
}

func TestParser_PeekAhead_WhenNoTokens(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1"))

	token := p.peekAhead(23)

	assert.Nil(t, token)
}

// TestHasToken
func TestParser_HasToken_WhenNoTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	assert.False(t, p.hasToken())
}

func TestParser_HasToken_WhenTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.INT, Value: "int"},
	})

	assert.True(t, p.hasToken())
}

// TestHasTokens
func TestParser_HasTokens_WhenNoTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	assert.False(t, p.hasTokens(3))
}

func TestParser_HasTokens_WhenTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.INT, Value: "int"},
		{Type: tokenizer.INT, Value: "int"},
	})

	assert.True(t, p.hasTokens(1))
	assert.True(t, p.hasTokens(2))
	assert.False(t, p.hasTokens(3))
}

// TestMatchSeq
func TestParser_MatchSeqTrue(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1"))

	match, iErr := p.matchSeq(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ, tokenizer.LITERAL)

	assert.True(t, match)
	assert.Equal(t, -1, iErr)
}

func TestParser_MatchSeqFalse(t *testing.T) {
	p := NewParser(generateTokensFor("int a = 1"))

	match, iErr := p.matchSeq(tokenizer.INT, tokenizer.EQ, tokenizer.LITERAL)

	assert.False(t, match)
	assert.Equal(t, 1, iErr)
}
func generateTokensFor(text string) []tokenizer.Token {
	t := tokenizer.NewTokenizer([]rune(text))
	t.GenerateTokens()
	return t.GetTokens()
}
