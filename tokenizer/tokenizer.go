package tokenizer

import (
	"bytes"
	"compiler/utils"
	"fmt"
	"unicode"
)

type Tokenizer struct {
	runes []rune
	index int

	tokens []Token

	line            int
	lineIndexStart  int
	tokenIndexStart int
}

func NewTokenizer(runes []rune) *Tokenizer {
	return &Tokenizer{
		runes: runes,
	}
}

// GenerateTokens analyzes the list of runes and generates a list of tokens you can get with GetTokens(
func (t *Tokenizer) GenerateTokens() error {
	buff := bytes.NewBuffer([]byte{})

	for t.hasRune() {
		// Track index of start of token
		t.tokenIndexStart = t.index - t.lineIndexStart

		if tokenType, foundMatch := singleRuneTokens[t.peek()]; foundMatch {
			// Try to match a single rune token, tokens that match a single character

			value := t.consume()
			t.createToken(tokenType, string(value))

			if value == '\n' {
				// Track line in file
				t.line++
				// Track where the line starts to later figure out the token position
				t.lineIndexStart = t.index
			}
		} else if unicode.IsSpace(t.peek()) {
			// Other spaces can be consumed and ignored

			t.consume()
		} else if unicode.IsLetter(t.peek()) {
			// If we peek and its a letter then we need to read until next non-valid character
			// Then check if it's a keyword or identifier

			buff.Reset()

			// Write to buff while tokenizer has a rune that matches a valid keyword or identifier
			for t.hasRune() && (unicode.IsLetter(t.peek()) || unicode.IsNumber(t.peek()) || '_' == t.peek()) {
				buff.WriteRune(t.consume())
			}

			value := buff.String()

			tokenType, foundMatch := listOfKeywords[value]
			if foundMatch {
				// If value is in the list of keywords then we create a token of that type
				t.createToken(tokenType, value)
			} else if value == "true" || value == "false" {
				t.createToken(BOOL_LITERAL, value)
			} else {
				// If it's not any of the above it's an identifier
				t.createToken(IDENTIFIER, value)
			}
		} else if unicode.IsNumber(t.peek()) {
			// If we peek and its a number then we match all numbers for the literal

			buff.Reset()

			// Consume and add to buffer until there are no more numbers
			for t.hasRune() && unicode.IsNumber(t.peek()) {
				buff.WriteRune(t.consume())
			}

			// All numbers is a literal
			t.createToken(INT_LITERAL, buff.String())
		} else {
			// If there is an unknown symbol we just return an error
			return utils.NewError(
				fmt.Sprintf("Unexpected symbol '%c'", t.peek()),
				t.getTokenPosition(),
			)
		}
	}

	return nil
}

// GetTokens return list of generated tokens
func (t *Tokenizer) GetTokens() []Token {
	return t.tokens
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

// createToken creates a new token with the given info and tracked position in file
func (t *Tokenizer) createToken(tokenType TokenType, value string) {
	t.tokens = append(t.tokens, Token{
		Type:  tokenType,
		Value: value,
		Pos:   t.getTokenPosition(),
	})
}

func (t *Tokenizer) getTokenPosition() *utils.FilePosition {
	return utils.NewPosition(t.line+1, t.tokenIndexStart+1)
}
