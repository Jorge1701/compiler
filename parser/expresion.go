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
	T    int
	Term *NodeTerm
	Oper *NodeOper
}

func (p *Parser) parseNodeExpr(minPrec int) (*NodeExpr, error) {
	if p.hasTokens(3) && p.peekAhead(1).IsOperator() {
		oper, err := p.parseNodeExprOper(minPrec)
		if err != nil {
			return nil, err
		}

		return oper, nil
	} else if p.hasToken() && p.peek().MatchAny(tokenizer.LITERAL, tokenizer.IDENTIFIER) {
		term, _ := p.parseNodeTerm()

		return &NodeExpr{
			T:    TypeNodeExprTerm,
			Term: term,
		}, nil
	}

	return nil, fmt.Errorf("Invalid expresion: %s", p.unexpectedToken())
}
