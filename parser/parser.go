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
		if p.matchEqual(tokenizer.INT) {
			if p.matchSequence(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ, tokenizer.LITERAL, tokenizer.SEPARATOR) {
				p.consume() // Ignore INT
				id := p.consume()
				p.consume() // Ignore EQ
				lit := p.consume()

				node.Statements = append(node.Statements,
					&NodeInitialize{
						Identifier: id,
						Literal:    lit,
					},
				)

                p.consume() // Ignore SEPARATOR
			} else {
                log.Fatal("Parser: Unexpected token after 'int'")
			}
		} else if p.matchEqual(tokenizer.SALIR) {
			if p.matchSequence(tokenizer.SALIR, tokenizer.LITERAL, tokenizer.SEPARATOR) {
				p.consume() // Ignore SALIR

				node.Statements = append(node.Statements,
					&NodeSalirLiteral{
						Literal: p.consume(),
					},
				)
                
                p.consume() // Ignore SEPARATOR
			} else if p.matchSequence(tokenizer.SALIR, tokenizer.IDENTIFIER, tokenizer.SEPARATOR) {
				p.consume() // Ignore SALIR

				node.Statements = append(node.Statements,
					&NodeSalirIdentifier{
						Identifier: p.consume(),
					},
				)

                p.consume() // Ignore SEPARATOR
			} else {
                log.Fatal("Parser: Error expected literal after 'salir'")
			}
		} else {
            log.Fatal("Parser: Error unexpected token", p.tokens[p.index].TokenType, p.tokens[p.index].Value)
		}
	}

	return node
}

func (p *Parser) matchEqual(tokenType tokenizer.TokenType) bool {
	return p.tokens[p.index].TokenType == tokenType
}

func (p *Parser) matchAny(tokens ...tokenizer.TokenType) bool {
    for _, t := range tokens {
        if p.tokens[p.index].TokenType == t {
            return true
        }
    }
    return false
}

func (p *Parser) matchSequence(tokens ...tokenizer.TokenType) bool {
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

func (p *Parser) consume() tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return t
}
