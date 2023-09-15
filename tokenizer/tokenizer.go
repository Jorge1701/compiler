package tokenizer

import (
	"bytes"
	"log"
	"unicode"
)

type TokenType string

const (
	SALIR      TokenType = "SALIR"
	LITERAL    TokenType = "LITERAL"
	IDENTIFIER TokenType = "IDENTIFIER"
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
			if value == "salir" {
				tokens = append(tokens, Token{TokenType: SALIR, Value: value})
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
            log.Fatalf("Error unexpected token '%s'", string(t.consume()))
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
