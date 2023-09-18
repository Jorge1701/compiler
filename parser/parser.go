package parser

import (
	"compiler/tokenizer"
	"fmt"
	"strings"
)

const (
	NodeTypeStmtInit byte = iota
	NodeTypeStmtExit

	NodeTypeExprTerm
	NodeTypeExprOper

	NodeTypeTermLit
	NodeTypeTermIdent
)

// [Program]
type NodeProg struct {
	Stmts []NodeStmt
}

// [Statement]
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

// [Expresion]
type NodeExpr struct {
	t    byte
	Term *NodeTerm
	Oper *NodeOper
}

// [Term]
type NodeTerm struct {
	t     byte
	Lit   *tokenizer.Token
	Ident *tokenizer.Token
}

// [Operation]
type NodeOper struct {
	Oper *tokenizer.Token
	Lhs  *NodeExpr
	Rhs  *NodeExpr
}

type Parser struct {
	tokens []tokenizer.Token
	index  int
}

func NewParser(tokens []tokenizer.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

func (p *Parser) GenerateNodes() (*NodeProg, error) {
	nodeProg := &NodeProg{
		Stmts: []NodeStmt{},
	}

	for p.hasToken() {
		if p.match(tokenizer.INT) {
			node, err := p.parseNodeStmtInit()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.match(tokenizer.EXIT) {
			node, err := p.parseNodeStmtExit()
			if err != nil {
				return nodeProg, err
			}
			nodeProg.Stmts = append(nodeProg.Stmts, node)
		} else if p.match(tokenizer.SEP) {
			p.consume()
		} else {
			return nodeProg, fmt.Errorf("Error in parsing, cannot parse token '%s'", p.peek().String())
		}
	}
	return nodeProg, nil
}

func (p *Parser) parseNodeStmtInit() (NodeStmt, error) {
	if p.match(tokenizer.INT, tokenizer.IDENTIFIER, tokenizer.EQ) {
		p.consume()
		ident := p.consume()
		p.consume()

		expr, err := p.parseNodeExpr(0)
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
		expr, err := p.parseNodeExpr(0)
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

func (p *Parser) parseNodeTerm() (NodeTerm, error) {
	if p.match(tokenizer.LITERAL) {
		return NodeTerm{
			t:   NodeTypeTermLit,
			Lit: p.consume(),
		}, nil
	} else if p.match(tokenizer.IDENTIFIER) {
		return NodeTerm{
			t:     NodeTypeTermIdent,
			Ident: p.consume(),
		}, nil
	}
	return NodeTerm{}, fmt.Errorf("Non valid term '%s'", p.peek().String())
}

func (p *Parser) parseNodeExpr(minPrec int) (NodeExpr, error) {
	if p.hasTokens(3) && p.peekAt(1).IsOperator() {
		oper, err := p.parseNodeOper(minPrec)
		if err != nil {
			return NodeExpr{}, err
		}
		return NodeExpr{
			t:    NodeTypeExprOper,
			Oper: &oper,
		}, nil
	} else if p.hasToken() && p.matchAny(tokenizer.LITERAL, tokenizer.IDENTIFIER) {
		term, err := p.parseNodeTerm()
		if err != nil {
			return NodeExpr{}, err
		}
		return NodeExpr{
			t:    NodeTypeExprTerm,
			Term: &term,
		}, nil
	}

	return NodeExpr{}, fmt.Errorf("Error parsing expresion")
}

func (p *Parser) parseNodeOper(minPrec int) (NodeOper, error) {
	term, err := p.parseNodeTerm()
	if err != nil {
		return NodeOper{}, err
	}

	lhs := NodeExpr{
		t:    NodeTypeExprTerm,
		Term: &term,
	}

	var rhs NodeExpr
	var oper *tokenizer.Token
	for {
		if !p.hasToken() || !p.peek().IsOperator() || p.peek().GetPrec() < minPrec {
			break
		}

		oper = p.consume()

		rhs, err = p.parseNodeExpr(oper.GetPrec() + 1)
		if err != nil {
			return NodeOper{}, err
		}
	}
	return NodeOper{
		Oper: oper,
		Lhs:  &lhs,
		Rhs:  &rhs,
	}, nil
}

func (p *Parser) hasToken() bool {
	return p.index < len(p.tokens)
}

func (p *Parser) hasTokens(amt int) bool {
	return p.index+amt < len(p.tokens)
}

func (p *Parser) peek() *tokenizer.Token {
	return &p.tokens[p.index]
}

func (p *Parser) peekAt(offSet int) *tokenizer.Token {
	return &p.tokens[p.index+offSet]
}

func (p *Parser) match(tokenTypes ...tokenizer.TokenType) bool {
	for i, tt := range tokenTypes {
		if p.index+i > len(p.tokens) || p.tokens[p.index+i].Type != tt {
			return false
		}
	}
	return true
}

func (p *Parser) matchAny(tokenTypes ...tokenizer.TokenType) bool {
	for _, tt := range tokenTypes {
		if p.tokens[p.index].Type == tt {
			return true
		}
	}
	return false
}

func (p *Parser) consume() *tokenizer.Token {
	t := p.tokens[p.index]
	p.index++
	return &t
}

func (n NodeProg) Print() {
	fmt.Println("+ NodeProg")
	for _, stmt := range n.Stmts {
		fmt.Printf("|")
		stmt.Print(0)
	}
}

func (stmt NodeStmt) Print(offset int) {
	switch stmt.t {
	case NodeTypeStmtInit:
		fmt.Printf("- NodeTypeStmtInit {Ident: '%s'}\n", stmt.Init.Ident.Value)
		stmt.Init.Expr.Print(offset + 1)
	case NodeTypeStmtExit:
		fmt.Printf("- NodeTypeStmtExit\n")
		stmt.Exit.Expr.Print(offset + 1)
	default:
		fmt.Println("NOT IMPLEMENTED NodeStmt.Print")
	}
}

func (expr NodeExpr) Print(offset int) {
	switch expr.t {
	case NodeTypeExprTerm:
		expr.Term.Print(offset + 1)
	case NodeTypeExprOper:
		expr.Oper.Print(offset + 1)
	default:
		fmt.Println("NOT IMPLEMENTED NodeExpr.Print")
	}
}

func (term NodeTerm) Print(offset int) {
	fmt.Printf("|%s- ", strings.Repeat(" ", offset))
	switch term.t {
	case NodeTypeTermLit:
		fmt.Printf("NodeTypeTermLit {Lit: '%s'}\n", term.Lit.Value)
	case NodeTypeTermIdent:
		fmt.Printf("NodeTypeTermIdent {Ident: '%s'}\n", term.Ident.Value)
	default:
		fmt.Println("NOT IMPLEMENTED NodeTerm.Print")
	}
}

func (oper NodeOper) Print(offset int) {
	fmt.Printf("|%s- ", strings.Repeat(" ", offset))
	fmt.Printf("NodeOper {Oper: '%s'}\n", oper.Oper.Value)
	oper.Lhs.Print(offset + 1)
	oper.Rhs.Print(offset + 1)
}
