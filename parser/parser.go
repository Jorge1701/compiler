package parser

import (
	"compiler/tokenizer"
)

type NodeExprLiteral struct {
	Literal *tokenizer.Token
}

type NodeExprIdentifier struct {
	Identifier *tokenizer.Token
}

type NodeExprOperation[nodeExpr NodeExpr] struct {
	LeftExpr  nodeExpr
	Operator  *tokenizer.Token
	RightExpr nodeExpr
}

type NodeExpr interface {
	NodeExprLiteral | NodeExprIdentifier
}

type NodeExit[T NodeExpr] struct {
	Expr T
}

type Parser struct {
	tokens []tokenizer.Token
	index  int
}

func NewParser(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) Parse() *NodeProg {
	return nil // Refactoring...
}

func (p *Parser) matchEqual(tokenType tokenizer.TokenType) bool {
	return p.tokens[p.index].Type == tokenType
}

func (p *Parser) matchAny(tokens ...tokenizer.TokenType) bool {
	for _, t := range tokens {
		if p.tokens[p.index].Type == t {
			return true
		}
	}
	return false
}

func (p *Parser) matchSequence(tokens ...tokenizer.TokenType) bool {
	for i, t := range tokens {
		if p.tokens[p.index+i].Type != t {
			return false
		}
	}
	return true
}

func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

func (p *Parser) consume() tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return t
}
