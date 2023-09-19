package parser

import (
	"fmt"
	"strings"
)

func (n NodeProg) Print() {
	fmt.Println("+ NodeProg")
	for _, stmt := range n.Stmts {
		fmt.Printf("|")
		stmt.Print(0)
	}
}

func (stmt NodeStmt) Print(offset int) {
	switch stmt.t {
	case TypeNodeStmtInit:
		fmt.Printf("- NodeTypeStmtInit {Ident: '%s'}\n", stmt.Init.Ident.Value)
		stmt.Init.Expr.Print(offset + 1)
	case TypeNodeStmtExit:
		fmt.Printf("- NodeTypeStmtExit\n")
		stmt.Exit.Expr.Print(offset + 1)
	default:
		fmt.Println("NOT IMPLEMENTED NodeStmt.Print")
	}
}

func (expr NodeExpr) Print(offset int) {
	switch expr.t {
	case TypeNodeExprTerm:
		expr.Term.Print(offset + 1)
	case TypeNodeExprOper:
		expr.Oper.Print(offset + 1)
	default:
		fmt.Println("NOT IMPLEMENTED NodeExpr.Print")
	}
}

func (term NodeTerm) Print(offset int) {
	fmt.Printf("|%s- ", strings.Repeat(" ", offset))
	switch term.t {
	case TypeNodeTermLit:
		fmt.Printf("NodeTypeTermLit {Lit: '%s'}\n", term.Lit.Value)
	case TypeNodeTermIdent:
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
