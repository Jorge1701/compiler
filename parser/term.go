package parser

import (
	"compiler/tokenizer"
	"fmt"
)

const (
	TypeNodeTermLit = iota
	TypeNodeTermIdent
)

type NodeTerm struct {
	T     int
	Lit   *tokenizer.Token
	Ident *tokenizer.Token
}

func (p *Parser) parseNodeTerm() (*NodeTerm, error) {
	if p.hasToken() && p.peek().IsType(tokenizer.LITERAL) {
		return &NodeTerm{
			T:   TypeNodeTermLit,
			Lit: p.consume(),
		}, nil
	} else if p.hasToken() && p.peek().IsType(tokenizer.IDENTIFIER) {
		return &NodeTerm{
			T:     TypeNodeTermIdent,
			Ident: p.consume(),
		}, nil
	}

	return nil, fmt.Errorf("Invalid term: %s", p.unexpectedToken())
}
