package parser

import (
	"compiler/tokenizer"
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
		if p.peek().IsType(tokenizer.INT) {
			node, err := p.parseNodeStmtInit()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.peek().IsType(tokenizer.IDENTIFIER) {
			node, err := p.parseNodeStmtReassign()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.peek().IsType(tokenizer.EXIT) {
			node, err := p.parseNodeStmtExit()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.peek().IsType(tokenizer.SEP) {
			p.consume()
		} else {
			return nodeProg, fmt.Errorf("Error in parsing, cannot parse token '%s'", p.peek().String())
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

// matchSeq returns true if the next n tokens match the given types in order
func (p *Parser) matchSeq(tokenTypes ...tokenizer.TokenType) bool {
	for i, tt := range tokenTypes {
		if p.index+i > len(p.tokens) || p.tokens[p.index+i].Type != tt {
			return false
		}
	}
	return true
}

// consume returns the current token and increases the index so that the next operations handle the next token
func (p *Parser) consume() *tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return &t
}
