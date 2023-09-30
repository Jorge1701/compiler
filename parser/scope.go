package parser

import (
	"compiler/tokenizer"
	"fmt"
)

type NodeScope struct {
	Stmts *[]NodeStmt
}

func (p *Parser) parseNodeScope() (NodeScope, error) {
	if p.peek().IsType(tokenizer.B_L) {
		p.consume() // Consume opening }
		p.consume() // Consume following SEP

		stmts := []NodeStmt{}
		for p.hasToken() && !p.peek().IsType(tokenizer.B_R) {
			stmt, err := p.parseNodeStmt()
			if err != nil {
				return NodeScope{}, fmt.Errorf("Error parsing statement inside scope")
			}

			stmts = append(stmts, *stmt)

			if p.peek().IsType(tokenizer.SEP) {
				p.consume()
			} else {
				return NodeScope{}, fmt.Errorf("Error expected SEP")
			}
		}

		p.consume() // Consume closing }
		return NodeScope{
			Stmts: &stmts,
		}, nil
	} else {
		return NodeScope{}, fmt.Errorf("Error parsing scope")
	}
}
