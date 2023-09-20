package parser

import (
	"compiler/tokenizer"
	"fmt"
)

func (n NodeProg) Print() {
	fmt.Println("─ NodeProg")
	for i, stmt := range n.Stmts {
		var indent string
		last := i == len(n.Stmts)-1
		if last {
			indent = "  "
		} else {
			indent = "  "
		}
		stmt.Print(indent, last)
	}
}

func (stmt NodeStmt) Print(indent string, last bool) {
	indent = pIndent(indent, last)
	switch stmt.t {
	case TypeNodeStmtInit:
		fmt.Println("NodeTypeStmtInit")
		printToken(stmt.Init.Ident, indent, false)
		stmt.Init.Expr.Print(indent, true)
	case TypeNodeStmtExit:
		fmt.Println("NodeTypeStmtExit")
		stmt.Exit.Expr.Print(indent, true)
	default:
		fmt.Println("NOT IMPLEMENTED NodeStmt.Print")
	}
}

func (expr NodeExpr) Print(indent string, last bool) {
	switch expr.t {
	case TypeNodeExprTerm:
		expr.Term.Print(indent, last)
	case TypeNodeExprOper:
		expr.Oper.Print(indent, last)
	default:
		if last {
			fmt.Println("\\-")
		} else {
			fmt.Println("|-")
		}
		fmt.Println("NOT IMPLEMENTED NodeExpr.Print")
	}
}

func (term NodeTerm) Print(indent string, last bool) {
	indent = pIndent(indent, last)
	switch term.t {
	case TypeNodeTermLit:
		fmt.Println("NodeTypeTermLit")
		printToken(term.Lit, indent, true)
	case TypeNodeTermIdent:
		fmt.Println("NodeTypeTermIdent")
		printToken(term.Ident, indent, true)
	default:
		fmt.Println("NOT IMPLEMENTED NodeTerm.Print")
	}
}

func (oper NodeOper) Print(indent string, last bool) {
	indent = pIndent(indent, last)
	fmt.Println("NodeOper")
	printToken(oper.Oper, indent, false)
	oper.Lhs.Print(indent, false)
	oper.Rhs.Print(indent, true)
}

func pIndent(indent string, last bool) string {
	fmt.Printf(indent)
	if last {
		fmt.Printf("└─ ")
		indent += "   "
	} else {
		fmt.Printf("├─ ")
		indent += "│  "
	}
	return indent
}

func printToken(token *tokenizer.Token, indent string, last bool) {
	indent = pIndent(indent, last)
	fmt.Printf("Token %s\n", token.String())
}
