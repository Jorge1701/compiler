package parser

import (
	"compiler/tokenizer"
	"fmt"
)

const (
	lineMark     string = "│  "
	nodeMark     string = "├─ "
	lastNodeMark string = "└─ "
	space        string = "   "
)

func PrintNode(n interface{}, indent string, last bool) {
	fmt.Printf(indent)
	if last {
		fmt.Printf(lastNodeMark)
		indent += space
	} else {
		fmt.Printf(nodeMark)
		indent += lineMark
	}

	switch node := n.(type) {
	case *NodeProg:
		fmt.Println("NodeProg")

		for i, stmt := range node.Stmts {
			PrintNode(stmt, indent, i == len(node.Stmts)-1)
		}
	case NodeStmt:
		fmt.Println("NodeStmt")

		switch node.t {
		case TypeNodeStmtInit:
			PrintNode(node.Init, indent, true)
		case TypeNodeStmtExit:
			PrintNode(node.Exit, indent, true)
		default:
			caseNotImplemented("NodeStmt", indent, last)
		}
	case *NodeStmtInit:
		fmt.Println("NodeTypeStmtInit")

		PrintNode(node.Ident, indent, false)
		PrintNode(node.Expr, indent, true)
	case *NodeStmtExit:
		fmt.Println("NodeTypeStmtExit")

		PrintNode(node.Expr, indent, true)
	case *NodeExpr:
		fmt.Println("NodeExpr")

		switch node.t {
		case TypeNodeExprTerm:
			PrintNode(node.Term, indent, true)
		case TypeNodeExprOper:
			PrintNode(node.Oper, indent, true)
		default:
			caseNotImplemented("NodeExpr", indent, last)
		}
	case *NodeTerm:
		fmt.Println("NodeTerm")

		switch node.t {
		case TypeNodeTermLit:
			PrintNode(node.Lit, indent, true)
		case TypeNodeTermIdent:
			PrintNode(node.Ident, indent, true)
		default:
			caseNotImplemented("NodeTerm", indent, last)
		}
	case *NodeOper:
		fmt.Println("NodeOper")

		PrintNode(node.Oper, indent, false)
		PrintNode(node.Lhs, indent, false)
		PrintNode(node.Rhs, indent, true)
	case *tokenizer.Token:
		fmt.Printf("Token %s\n", node.String())
	default:
		caseNotImplemented("switch", indent, last)
	}
}

func caseNotImplemented(caze, indent string, last bool) {
	fmt.Printf("PrintNode case '%s' not implemented\n", caze)
}
