package parser

import (
	"compiler/tokenizer"
	"fmt"
)

type NodeLiteral struct {
	Literal tokenizer.Token
}

type NodeSalir struct {
	NodeLiteral *NodeLiteral
}

func (p *Parser) parseLiteral() (*NodeLiteral, error) {
	if p.isType(tokenizer.LITERAL) {
		return &NodeLiteral{
			Literal: p.consume(),
		}, nil
	} else {
		return nil, fmt.Errorf("parseLiteral did not found a literal value")
	}
}
