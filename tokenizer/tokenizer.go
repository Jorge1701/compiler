package tokenizer

import (
	"unicode"
)

type TokenType string

const (
	KEYWORD   TokenType = "KEYWORD"
	LITERAL   TokenType = "LITERAL"
	SEPARATOR TokenType = "SEPARATOR"
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

	var buf []rune

	for t.hasRune() {
		if unicode.IsLetter(t.peek()) {
			buf = append(buf, t.consume())

			for t.hasRune() && unicode.IsLetter(t.peek()) || unicode.IsNumber(t.peek()) {
				buf = append(buf, t.consume())
			}

			tokens = append(tokens, Token{TokenType: KEYWORD, Value: string(buf)})
			buf = []rune{}
		} else if unicode.IsSpace(t.peek()) {
			t.consume()
		} else if unicode.IsNumber(t.peek()) {
			buf = append(buf, t.consume())

			for t.hasRune() && unicode.IsNumber(t.peek()) {
				buf = append(buf, t.consume())
			}

			tokens = append(tokens, Token{TokenType: LITERAL, Value: string(buf)})
			buf = []rune{}
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
