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
	T     byte
	Lit   *tokenizer.Token
	Ident *tokenizer.Token
}

func (p *Parser) parseNodeTerm() (NodeTerm, error) {
	if p.hasToken() && p.peek().IsType(tokenizer.LITERAL) {
		return NodeTerm{
			T:   TypeNodeTermLit,
			Lit: p.consume(),
		}, nil
	} else if p.hasToken() && p.peek().IsType(tokenizer.IDENTIFIER) {
		return NodeTerm{
			T:     TypeNodeTermIdent,
			Ident: p.consume(),
		}, nil
	}

	if p.hasToken() {
		return NodeTerm{}, fmt.Errorf("No tokens left")
	} else {
		return NodeTerm{}, fmt.Errorf("Non valid term '%s'", p.peek().String())
	}
}
