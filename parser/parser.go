package parser

import (
	"compiler/tokenizer"
	"fmt"
	"log"
)

const (
	NodeTypeExprLit byte = iota
	NodeTypeExprIdent
	NodeTypeExprOper

	NodeTypeStmtInit
	NodeTypeStmtExit
)

type NodeProg struct {
	Stmts []NodeStmt
}

type NodeExpr struct {
	t     byte
	Lit   *tokenizer.Token
	Ident *tokenizer.Token
	Oper  *NodeOper
}

type NodeOper struct {
	Oper *tokenizer.Token
	Lhs  *NodeExpr
	Rhs  *NodeExpr
}

type NodeStmt struct {
	t    byte
	Init *NodeStmtInit
	Exit *NodeStmtExit
}

type NodeStmtInit struct {
	Ident *tokenizer.Token
	Expr  *NodeExpr
}

type NodeStmtExit struct {
	Expr *NodeExpr
}

type Parser struct {
	tokens []tokenizer.Token
	index  int
}

func (p *Parser) GenerateNodes() (*NodeProg, error) {
	nodeProg := &NodeProg{
		Stmts: []NodeStmt{},
	}

	for p.hasToken() {
		if p.match(tokenizer.INT) {
			node, err := p.parseNodeStmtInit()
			if err != nil {
				return &NodeProg{}, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.match(tokenizer.EXIT) {
			node, err := p.parseNodeStmtExit()
			if err != nil {
				return &NodeProg{}, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.match(tokenizer.SEP) {
			p.consume()
		} else {
			return &NodeProg{}, fmt.Errorf("Error in parsing, cannot parse token '%s'", p.peek().String())
		}
	}
	return nodeProg, nil
}

func (p *Parser) parseNodeStmtInit() (NodeStmt, error) {
	if p.match(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ) {
		p.consume()
		ident := p.consume()
		p.consume()

		expr, err := p.parseNodeExpr()
		if err != nil {
			return NodeStmt{}, err
		}
		return NodeStmt{
			t: NodeTypeStmtInit,
			Init: &NodeStmtInit{
				Ident: ident,
				Expr:  &expr,
			},
		}, nil
	}
	return NodeStmt{}, fmt.Errorf("Error parsing initialization statement")
}

func (p *Parser) parseNodeStmtExit() (NodeStmt, error) {
	if p.match(tokenizer.EXIT) {
		p.consume()
		expr, err := p.parseNodeExpr()
		if err != nil {
			return NodeStmt{}, err
		}
		return NodeStmt{
			t: NodeTypeStmtExit,
			Exit: &NodeStmtExit{
				Expr: &expr,
			},
		}, nil
	}
	return NodeStmt{}, fmt.Errorf("Error parsing exit statement")
}

func (p *Parser) parseNodeExpr() (NodeExpr, error) {
	var tokens []tokenizer.Token

	for !p.match(tokenizer.SEP) {
		if p.match(tokenizer.LITERAL, tokenizer.IDENTIFIER, tokenizer.ADD, tokenizer.SUB, tokenizer.MUL, tokenizer.DIV) {
			tokens = append(tokens, *p.consume())
		} else {
			return NodeExpr{}, fmt.Errorf("Error parsing expresion, invalid token '%s'", p.tokens[p.index].String())
		}
	}

	if len(tokens) == 1 {
		if tokens[0].Type == tokenizer.LITERAL {
			return NodeExpr{
				t:   NodeTypeExprLit,
				Lit: &tokens[0],
			}, nil
		} else if tokens[0].Type == tokenizer.IDENTIFIER {
			return NodeExpr{
				t:     NodeTypeExprIdent,
				Ident: &tokens[0],
			}, nil
		} else {
			return NodeExpr{}, fmt.Errorf("Error parsing expresion, invalid token '%s'", tokens[0].String())
		}
	} else if len(tokens) == 0 {
		return NodeExpr{}, fmt.Errorf("Error parsing expresion, no tokens found")
	}

	oper, err := parseNodeOper(tokens)
	if err == nil {
		return NodeExpr{
			t:    NodeTypeExprOper,
			Oper: &oper,
		}, nil
	}

	return NodeExpr{}, fmt.Errorf("Error parsing expresion")
}

func parseNodeOper(tokens []tokenizer.Token) (NodeOper, error) {
	// TODO
	return NodeOper{}, nil
}

func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

func (p *Parser) peek() *tokenizer.Token {
	return &p.tokens[p.index]
}

func (p *Parser) match(tokenTypes ...tokenizer.TokenType) bool {
	for i, tt := range tokenTypes {
		if i > len(p.tokens) || p.tokens[i].Type != tt {
			return false
		}
	}
	return true
}

func (p *Parser) consume() *tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return &t
}
