package parser

import (
	"compiler/tokenizer"
	"fmt"
)

const (
	TypeNodeStmtInit = iota
	TypeNodeStmtReassign
	TypeNodeStmtExit
)

type NodeStmt struct {
	T        byte
	Init     *NodeStmtInit
	Reassign *NodeStmtReassign
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

type NodeStmtExit struct {
	Expr *NodeExpr
}

func (p *Parser) parseNodeStmtInit() (NodeStmt, error) {
	if p.matchSeq(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ) {
		p.consume()
		ident := p.consume()
		p.consume()

		expr, err := p.parseNodeExpr(1)
		if err != nil {
			return NodeStmt{}, err
		}
		return NodeStmt{
			T: TypeNodeStmtInit,
			Init: &NodeStmtInit{
				Ident: ident,
				Expr:  &expr,
			},
		}, nil
	}
	return NodeStmt{}, fmt.Errorf("Error parsing initialization statement")
}

func (p *Parser) parseNodeStmtReassign() (NodeStmt, error) {
	if p.matchSeq(tokenizer.IDENTIFIER, tokenizer.EQ) {
		ident := p.consume()
		p.consume()

		expr, err := p.parseNodeExpr(1)
		if err != nil {
			return NodeStmt{}, err
		}
		return NodeStmt{
			T: TypeNodeStmtReassign,
			Reassign: &NodeStmtReassign{
				Ident: ident,
				Expr:  &expr,
			},
		}, nil
	}
	return NodeStmt{}, fmt.Errorf("Error parsing reassignment statement")
}

func (p *Parser) parseNodeStmtExit() (NodeStmt, error) {
	if p.peek().IsType(tokenizer.EXIT) {
		p.consume()
		expr, err := p.parseNodeExpr(1)
		if err != nil {
			return NodeStmt{}, err
		}
		return NodeStmt{
			T: TypeNodeStmtExit,
			Exit: &NodeStmtExit{
				Expr: &expr,
			},
		}, nil
	}
	return NodeStmt{}, fmt.Errorf("Error parsing exit statement")
}
