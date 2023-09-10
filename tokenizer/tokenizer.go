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

	for t.index = 0; t.index < len(t.runes); t.index++ {
		if t.hasRune() && unicode.IsLetter(t.getRune()) {
			buf = append(buf, t.getRune())

			t.index++
			for t.hasRune() && unicode.IsLetter(t.getRune()) || unicode.IsNumber(t.getRune()) {
				buf = append(buf, t.getRune())
				t.index++
			}
			t.index--

			tokens = append(tokens, Token{TokenType: KEYWORD, Value: string(buf)})
			buf = []rune{}
		} else if t.hasRune() && unicode.IsSpace(t.getRune()) {
			continue
		} else if t.hasRune() && unicode.IsNumber(t.getRune()) {
			buf = append(buf, t.getRune())

			t.index++
			for t.hasRune() && unicode.IsNumber(t.getRune()) {
				buf = append(buf, t.getRune())
				t.index++
			}
			t.index--

			tokens = append(tokens, Token{TokenType: LITERAL, Value: string(buf)})
			buf = []rune{}
		} else if t.hasRune() {
			tokens = append(tokens, Token{TokenType: SEPARATOR, Value: string(t.getRune())})
		}
	}

	return tokens
}

func (t *Tokenizer) hasRune() bool {
	return t.index < len(t.runes)
}

func (t *Tokenizer) getRune() rune {
	return t.runes[t.index]
}
