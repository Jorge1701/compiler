package parser

import (
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

func (p *Parser) Parse() *NodeProg {
	node := &NodeProg{
		Statements: make([]interface{}, 0),
	}

	for p.hasToken() {
		if p.isType(tokenizer.INT) {
			if p.match(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ, tokenizer.LITERAL) {
				p.consume()
				id := p.consume()
				p.consume()
				lit := p.consume()

				node.Statements = append(node.Statements,
					&NodeInitialize{
						Identifier: id,
						Literal:    lit,
					},
				)
			} else {
				log.Fatal("Unexpected token after 'int'")
			}
		} else if p.isType(tokenizer.SALIR) {
			if p.match(tokenizer.SALIR, tokenizer.LITERAL) {
				p.consume()

				node.Statements = append(node.Statements,
					&NodeSalirLiteral{
						Literal: p.consume(),
					},
				)
			} else if p.match(tokenizer.SALIR, tokenizer.IDENTIFIER) {
				p.consume()

				node.Statements = append(node.Statements,
					&NodeSalirIdentifier{
						Identifier: p.consume(),
					},
				)
			} else {
				log.Fatal("Error expected literal after 'salir'")
			}
		} else {
			log.Fatal("Error unexpected token", p.tokens[p.index].TokenType, p.tokens[p.index].Value)
		}
	}

	return node
}

func (p *Parser) match(tokens ...tokenizer.TokenType) bool {
	for i, t := range tokens {
		if p.tokens[p.index+i].TokenType != t {
			return false
		}
	}
	return true
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
