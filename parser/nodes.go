package parser

import (
	"compiler/tokenizer"
	"fmt"
)

type NodeLiteral struct {
	literal tokenizer.Token
}

type NodeSalir struct {
	nodeLiteral *NodeLiteral
}

func (p *Parser) parseLiteral() (*NodeLiteral, error) {
	if p.isType(tokenizer.LITERAL) {
		return &NodeLiteral{
			literal: p.consume(),
		}, nil
	} else {
		return nil, fmt.Errorf("parseLiteral did not found a literal value")
	}
}
