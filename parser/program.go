package parser

import (
	"compiler/tokenizer"
	"fmt"
)

type NodeProg struct {
	Stmts *[]NodeStmt
}

func (p *Parser) parseNodeProg() (*NodeProg, error) {
	if !p.hasToken() {
		return nil, fmt.Errorf("Invalid program: %s", p.noTokensLeft())
	}

	stmts := []NodeStmt{}

	for p.hasToken() {
		stmt, err := p.parseNodeStmt()
		if err != nil {
			return nil, err
		}

		stmts = append(stmts, *stmt)

		if p.hasToken() && p.peek().IsType(tokenizer.SEP) {
			p.consume()
		} else if p.hasToken() {
			return nil, p.unexpectedToken()
		}
	}

	return &NodeProg{
		Stmts: &stmts,
	}, nil
}
