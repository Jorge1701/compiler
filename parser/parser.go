package parser

import (
	"compiler/tokenizer"
	"compiler/utils"
	"fmt"
)

type Parser struct {
	tokens []tokenizer.Token
	index  int

	nodeProg *NodeProg
}

func NewParser(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

// GenerateNodes parses the list of tokens and creates a parse tree
func (p *Parser) GenerateNodes() error {
	prog, err := p.parseNodeProg()

	if err != nil {
		return err
	}

	p.nodeProg = prog

	return nil
}

// GetNodes returns generated nodes
func (p *Parser) GetNodes() *NodeProg {
	return p.nodeProg
}

// hasToken returns true if there is a next token
func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

// hasTokens returns true if there is still an 'amt' of tokens left
func (p *Parser) hasTokens(amt int) bool {
	return p.index+amt-1 < len(p.tokens)
}

// peek returns the current token without changing the index
func (p *Parser) peek() *tokenizer.Token {
	if p.index >= len(p.tokens) {
		return nil
	}

	return &p.tokens[p.index]
}

// peekAhead returns the token at a given position ahead of the current token without changing the index
func (p *Parser) peekAhead(offSet int) *tokenizer.Token {
	if p.index+offSet >= len(p.tokens) {
		return nil
	}

	return &p.tokens[p.index+offSet]
}

// matchSeq returns (true, -1) if the next n tokens match the given types in order
// returns (false, number) if some token did not match, the number refering to the index of the non matching token
func (p *Parser) matchSeq(tokenTypes ...tokenizer.TokenType) (bool, int) {
	for i, tt := range tokenTypes {
		if p.index+i > len(p.tokens) || p.tokens[p.index+i].Type != tt {
			return false, p.index + i
		}
	}
	return true, -1
}

// consume returns the current token and increases the index so that the next operations handle the next token
func (p *Parser) consume() *tokenizer.Token {
	if p.index >= len(p.tokens) {
		return nil
	}

	t := p.tokens[p.index]
	p.index++
	return &t
}

// noTokensLeft used when there should be more tokens to parse
func (p *Parser) noTokensLeft() *utils.Error {
	return utils.NewError("No tokens left", nil)
}

// unexpectedToken returns an error with current token and token position
func (p *Parser) unexpectedToken() *utils.Error {
	if p.index >= len(p.tokens) {
		return p.noTokensLeft()
	}
	return utils.NewError(fmt.Sprintf("Unexpected token %s", p.peek().String()), p.peek().Pos)
}

// unexpectedTokenAt returns an error referencing token at index
func (p *Parser) unexpectedTokenAt(index int) *utils.Error {
	if index >= len(p.tokens) {
		return utils.NewError("No tokens left", nil)
	}

	t := p.tokens[index]
	if t.IsType(tokenizer.SEP) {
		return utils.NewError("Unexpected line break", t.Pos)
	}

	return utils.NewError(fmt.Sprintf("Unexpected token %s", t.String()), t.Pos)
}
