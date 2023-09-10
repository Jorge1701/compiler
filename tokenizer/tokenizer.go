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

func Tokenize(runes []rune) []Token {
	var tokens []Token

	var buf []rune

	for i := 0; i < len(runes); i++ {
		if unicode.IsLetter(runes[i]) {
			buf = append(buf, runes[i])

			i++
			for unicode.IsLetter(runes[i]) || unicode.IsNumber(runes[i]) {
				buf = append(buf, runes[i])
				i++
			}
			i--

			tokens = append(tokens, Token{TokenType: KEYWORD, Value: string(buf)})
			buf = []rune{}
		} else if unicode.IsSpace(runes[i]) {
			continue
		} else if unicode.IsNumber(runes[i]) {
			buf = append(buf, runes[i])

			i++
			for unicode.IsNumber(runes[i]) {
				buf = append(buf, runes[i])
				i++
			}
			i--

			tokens = append(tokens, Token{TokenType: LITERAL, Value: string(buf)})
			buf = []rune{}
		} else {
			tokens = append(tokens, Token{TokenType: SEPARATOR, Value: string(runes[i])})
		}
	}

	return tokens
}
