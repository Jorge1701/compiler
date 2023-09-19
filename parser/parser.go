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

func (p *Parser) GenerateNodes() (*NodeProg, error) {
	nodeProg := &NodeProg{
		Stmts: []NodeStmt{},
	}

	for p.hasToken() {
		if p.match(tokenizer.INT) {
			node, err := p.parseNodeStmtInit()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.match(tokenizer.EXIT) {
			node, err := p.parseNodeStmtExit()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.match(tokenizer.SEP) {
			p.consume()
		} else {
			return nodeProg, fmt.Errorf("Error in parsing, cannot parse token '%s'", p.peek().String())
		}
	}
	return nodeProg, nil
}

func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

func (p *Parser) hasTokens(amt int) bool {
	return p.index+amt < len(p.tokens)
}

func (p *Parser) peek() *tokenizer.Token {
	return &p.tokens[p.index]
}

func (p *Parser) peekAt(offSet int) *tokenizer.Token {
	return &p.tokens[p.index+offSet]
}

func (p *Parser) match(tokenTypes ...tokenizer.TokenType) bool {
	for i, tt := range tokenTypes {
		if p.index+i > len(p.tokens) || p.tokens[p.index+i].Type != tt {
			return false
		}
	}
	return true
}

func (p *Parser) matchAny(tokenTypes ...tokenizer.TokenType) bool {
	for _, tt := range tokenTypes {
		if p.tokens[p.index].Type == tt {
			return true
		}
	}
	return false
}

func (p *Parser) consume() *tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return &t
}
