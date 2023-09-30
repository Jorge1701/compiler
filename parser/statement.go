package parser

import (
	"compiler/tokenizer"
	"fmt"
)

const (
	TypeNodeStmtInit = iota
	TypeNodeStmtReassign
	TypeNodeStmtScope
	TypeNodeStmtExit
)

type NodeStmt struct {
	T        byte
	Init     *NodeStmtInit
	Reassign *NodeStmtReassign
	Scope    *NodeStmtScope
	Exit     *NodeStmtExit
}

type NodeStmtInit struct {
	Ident *tokenizer.Token
	Expr  *NodeExpr
}

type NodeStmtReassign struct {
	Ident *tokenizer.Token
	Expr  *NodeExpr
}

type NodeStmtScope struct {
	Scope *NodeScope
}

type NodeStmtExit struct {
	Expr *NodeExpr
}

func (p *Parser) parseNodeStmt() (node *NodeStmt, err error) {
	switch p.peek().Type {
	case tokenizer.INT:
		node, err = p.parseNodeStmtInit()
	case tokenizer.IDENTIFIER:
		node, err = p.parseNodeStmtReassign()
	case tokenizer.B_L:
		node, err = p.parseNodeStmtScope()
	case tokenizer.EXIT:
		node, err = p.parseNodeStmtExit()
	default:
		err = p.unexpectedToken()
	}

	if err != nil {
		return nil, fmt.Errorf("Invalid statement: %s", err)
	} else {
		return node, nil
	}
}

func (p *Parser) parseNodeStmtInit() (*NodeStmt, error) {
	match, iErr := p.matchSeq(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ)
	if match {
		p.consume()          // Ignore INT
		ident := p.consume() // Save IDENTIFIER
		p.consume()          // Ignore ER

		expr, err := p.parseNodeExpr(1)
		if err != nil {
			return nil, err
		}

		return &NodeStmt{
			T: TypeNodeStmtInit,
			Init: &NodeStmtInit{
				Ident: ident,
				Expr:  &expr,
			},
		}, nil
	}
	return nil, p.unexpectedTokenAt(iErr)
}

func (p *Parser) parseNodeStmtReassign() (*NodeStmt, error) {
	match, iErr := p.matchSeq(tokenizer.IDENTIFIER, tokenizer.EQ)
	if match {
		ident := p.consume() // Save IDENTIFIER
		p.consume()          // Ignore EQ

		expr, err := p.parseNodeExpr(1)
		if err != nil {
			return nil, err
		}

		return &NodeStmt{
			T: TypeNodeStmtReassign,
			Reassign: &NodeStmtReassign{
				Ident: ident,
				Expr:  &expr,
			},
		}, nil
	}
	return nil, p.unexpectedTokenAt(iErr)
}

func (p *Parser) parseNodeStmtScope() (*NodeStmt, error) {
	scope, err := p.parseNodeScope()
	if err != nil {
		return nil, err
	}

	return &NodeStmt{
		T: TypeNodeStmtScope,
		Scope: &NodeStmtScope{
			Scope: &scope,
		},
	}, nil
}

func (p *Parser) parseNodeStmtExit() (*NodeStmt, error) {
	if p.peek().IsType(tokenizer.EXIT) {
		p.consume() // Ignore EXIT

		expr, err := p.parseNodeExpr(1)
		if err != nil {
			return nil, err
		}

		return &NodeStmt{
			T: TypeNodeStmtExit,
			Exit: &NodeStmtExit{
				Expr: &expr,
			},
		}, nil
	}
	return nil, p.unexpectedToken()
}
