package tokenizer

import (
	"bytes"
	"compiler/keywords"
	"unicode"
)

type TokenType string

const (
	IDENTIFIER TokenType = "IDENTIFIER"
	KEYWORD    TokenType = "KEYWORD"
	LITERAL    TokenType = "LITERAL"
	SEPARATOR  TokenType = "SEPARATOR"
)

type Token struct {
	TokenType TokenType
	Value     string
}

type Tokenizer struct {
	runes []rune
	index int
}

func NewTokenizer(runes []rune) *Tokenizer {
	return &Tokenizer{
		runes: runes,
	}
}

func (t *Tokenizer) Tokenize() []Token {
	var tokens []Token

	buff := bytes.NewBuffer([]byte{})

	for t.hasRune() {
		if unicode.IsLetter(t.peek()) {
			buff.WriteRune(t.consume())

			for t.hasRune() && (unicode.IsLetter(t.peek()) || unicode.IsNumber(t.peek())) {
				buff.WriteRune(t.consume())
			}

			value := buff.String()
			if keywords.IsKeyWord(value) {
				tokens = append(tokens, Token{TokenType: KEYWORD, Value: value})
			} else {
				tokens = append(tokens, Token{TokenType: IDENTIFIER, Value: value})
			}
			buff.Reset()
		} else if unicode.IsSpace(t.peek()) {
			t.consume()
		} else if unicode.IsNumber(t.peek()) {
			buff.WriteRune(t.consume())

			for t.hasRune() && unicode.IsNumber(t.peek()) {
				buff.WriteRune(t.consume())
			}

			tokens = append(tokens, Token{TokenType: LITERAL, Value: buff.String()})
			buff.Reset()
		} else {
			tokens = append(tokens, Token{TokenType: SEPARATOR, Value: string(t.consume())})
		}
	}

	return tokens
}

func (t *Tokenizer) hasRune() bool {
	return t.index < len(t.runes)
}

func (t *Tokenizer) peek() rune {
	return t.runes[t.index]
}

func (t *Tokenizer) consume() rune {
	r := t.runes[t.index]
	t.index++
	return r
}
