package tokenizer

import (
	"bytes"
	"fmt"
	"unicode"
)

type TokenType string

const (
	SEP  TokenType = "SEP"
	P_L  TokenType = "P_L"
	P_R  TokenType = "P_R"
	B_L  TokenType = "B_L"
	B_R  TokenType = "B_R"
	SB_L TokenType = "SB_L"
	SB_R TokenType = "SB_R"

	ADD TokenType = "ADD"
	SUB TokenType = "SUB"
	MUL TokenType = "MUL"
	DIV TokenType = "DIV"

	EQ TokenType = "EQ"

	INT  TokenType = "INT"
	EXIT TokenType = "EXIT"

	IDENTIFIER TokenType = "IDENTIFIER"
	LITERAL    TokenType = "LITERAL"
)

var singleRuneTokens = map[rune]TokenType{
	'\n': SEP,
	'(':  P_L,
	')':  P_R,
	'{':  B_L,
	'}':  B_R,
	'[':  SB_L,
	']':  SB_R,
	'+':  ADD,
	'-':  SUB,
	'*':  MUL,
	'/':  DIV,
	'=':  EQ,
}

var listOfKeywords = map[string]TokenType{
	"int":  INT,
	"exit": EXIT,
}

type Token struct {
	Type  TokenType
	Value string
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

// GenerateTokens analyzes the list of runes and generates a []Token
func (t *Tokenizer) GenerateTokens() (tokens []Token) {
	buff := bytes.NewBuffer([]byte{})

	for t.hasRune() {
		if tokenType, foundMatch := singleRuneTokens[t.peek()]; foundMatch {
			// Try to match a single rune token, tokens that match a single character

			tokens = append(tokens, Token{
				Type:  tokenType,
				Value: string(t.consume()),
			})
		} else if unicode.IsSpace(t.peek()) {
			// Other spaces can be consumed and ignored

			t.consume()
		} else if unicode.IsLetter(t.peek()) {
			// If we peek and its a letter then we need to read until next non-valid character
			// Then check if it's a keyword or identifier

			buff.Reset()

			// Write to buff while tokenizer has a rune that matches a valid keyword or identifier
			for unicode.IsLetter(t.peek()) || unicode.IsNumber(t.peek()) || '_' == t.peek() {
				buff.WriteRune(t.consume())
			}

			value := buff.String()

			tokenType, foundMatch := listOfKeywords[value]
			if foundMatch {
				// If value is in the list of keywords then we create a token of that type
				tokens = append(tokens, Token{
					Type:  tokenType,
					Value: value,
				})
			} else {
				// If it's not in the list of keywords then it's an identifier
				tokens = append(tokens, Token{
					Type:  IDENTIFIER,
					Value: value,
				})
			}
		} else if unicode.IsNumber(t.peek()) {
			// If we peek and its a number then we match all numbers for the literal

			buff.Reset()

			// Consume and add to buffer until there are no more numbers
			for unicode.IsNumber(t.peek()) {
				buff.WriteRune(t.consume())
			}

			// All numbers is a literal
			tokens = append(tokens, Token{
				Type:  LITERAL,
				Value: buff.String(),
			})
		}
	}

	return tokens
}

// hasRune return true if the are still runes left to tokenize
func (t *Tokenizer) hasRune() bool {
	return t.index < len(t.runes)
}

// peek returns the rune at current index without modifying it
func (t *Tokenizer) peek() rune {
	return t.runes[t.index]
}

// consume returns the current rune incrementing index so that
// the next peek or consume work with the next value
func (t *Tokenizer) consume() rune {
	r := t.peek()
	t.index++
	return r
}

// String returns a printable string that represents the token
func (t *Token) String() string {
	if t.Type == SEP {
		return fmt.Sprintf("(%s, \\n)", t.Type)
	} else {
		return fmt.Sprintf("(%s, %s)", t.Type, t.Value)
	}
}
