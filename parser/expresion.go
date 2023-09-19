package parser

import (
	"compiler/tokenizer"
	"fmt"
)

const (
	TypeNodeExprTerm = iota
	TypeNodeExprOper
)

type NodeExpr struct {
	t    byte
	Term *NodeTerm
	Oper *NodeOper
}

func (p *Parser) parseNodeExpr(minPrec int) (NodeExpr, error) {
	if p.hasTokens(3) && p.peekAhead(1).IsOperator() {
		oper, err := p.parseNodeExprOper(minPrec)
		if err != nil {
			return NodeExpr{}, err
		}
		return oper, nil
	} else if p.hasToken() && p.matchAny(tokenizer.LITERAL, tokenizer.IDENTIFIER) {
		term, err := p.parseNodeTerm()
		if err != nil {
			return NodeExpr{}, err
		}
		return NodeExpr{
			t:    TypeNodeExprTerm,
			Term: &term,
		}, nil
	}

	return NodeExpr{}, fmt.Errorf("Error parsing expresion")
}
