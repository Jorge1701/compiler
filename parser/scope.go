package parser

import (
	"compiler/tokenizer"
)

type NodeScope struct {
	Stmts *[]NodeStmt
}

func (p *Parser) parseNodeScope() (*NodeScope, error) {
	if p.peek().IsType(tokenizer.B_L) {
		p.consume() // Ignore opening B_L
		if p.peek().IsType(tokenizer.SEP) {
			p.consume() // Ignore following SEP
		}

		stmts := []NodeStmt{}
		for p.hasToken() && !p.peek().IsType(tokenizer.B_R) {
			stmt, err := p.parseNodeStmt()
			if err != nil {
				return nil, err
			}

			stmts = append(stmts, *stmt)

			if p.peek().IsType(tokenizer.SEP) {
				p.consume() // Ignore SEP
			}
		}

		p.consume() // Ignore closing B_R

		return &NodeScope{
			Stmts: &stmts,
		}, nil
	} else {
		return nil, p.unexpectedToken()
	}
}
