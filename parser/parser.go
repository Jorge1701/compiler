package parser

import (
	"compiler/tokenizer"
	"compiler/utils"
	"fmt"
)

type Parser struct {
	tokens []tokenizer.Token
	index  int
}

func NewParser(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

// GenerateNodes parses the list of tokens and returns a parse tree
func (p *Parser) GenerateNodes() (*NodeProg, error) {
	nodeProg := &NodeProg{
		Stmts: []NodeStmt{},
	}

	for p.hasToken() {
		stmt, err := p.parseNodeStmt()
		if err != nil {
			return nil, err
		}

		nodeProg.Stmts = append(nodeProg.Stmts, *stmt)

		if p.peek().IsType(tokenizer.SEP) {
			p.consume()
		} else {
			return nil, p.unexpectedToken()
		}
	}

	return nodeProg, nil
}

// hasToken returns true if there is a next token
func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

// hasTokens returns true if there is still an 'amt' of tokens left
func (p *Parser) hasTokens(amt int) bool {
	return p.index+amt < len(p.tokens)
}

// peek returns the current token without changing the index
func (p *Parser) peek() *tokenizer.Token {
	return &p.tokens[p.index]
}

// peekAhead returns the token at a given position ahead of the current token without changing the index
func (p *Parser) peekAhead(offSet int) *tokenizer.Token {
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
	t := p.tokens[p.index]
	p.index++
	return &t
}

// expected
func (p *Parser) expected(tokenType tokenizer.TokenType) *utils.Error {
	return utils.NewError(fmt.Sprintf("Expected token %s but got %s", tokenType, p.peek()), p.peek().Pos)
}

// unexpectedToken returns an error with current token and token position
func (p *Parser) unexpectedToken() *utils.Error {
	return utils.NewError(fmt.Sprintf("Unexpected token %s", p.peek().String()), p.peek().Pos)
}

// unexpectedTokenAt returns an error referencing token at index
func (p *Parser) unexpectedTokenAt(index int) *utils.Error {
	t := p.tokens[index]
	if t.IsType(tokenizer.SEP) {
		return utils.NewError("Unexpected line break", t.Pos)
	}
	return utils.NewError(fmt.Sprintf("Unexpected token %s", t.String()), t.Pos)
}
