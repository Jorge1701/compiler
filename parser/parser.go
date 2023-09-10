package parser

import (
	"compiler/keywords"
	"compiler/tokenizer"
	"log"
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

func (p *Parser) Parse() *NodeSalir {
	var node *NodeSalir

	for p.hasToken() {
		if p.isType(tokenizer.KEYWORD) && p.isValue(string(keywords.SALIR)) {
			p.consume()

			nodeLiteral, err := p.parseLiteral()
			if err != nil {
				log.Fatal("Invalid expression:", err)
			}

			node = &NodeSalir{
				NodeLiteral: nodeLiteral,
			}
		} else if p.isType(tokenizer.SEPARATOR) {
			p.consume()
		}
	}

	return node
}

func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

func (p *Parser) isType(tokenType tokenizer.TokenType) bool {
	return p.tokens[p.index].TokenType == tokenType
}

func (p *Parser) isValue(value string) bool {
	return p.tokens[p.index].Value == value
}

func (p *Parser) consume() tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return t
}
