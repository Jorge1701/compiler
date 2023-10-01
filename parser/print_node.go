package parser

import (
	"bytes"
	"compiler/tokenizer"
	"fmt"
)

const (
	lineMark     string = "│  "
	nodeMark     string = "├─ "
	lastNodeMark string = "└─ "
	space        string = "   "
)

func NodeToString(n interface{}, indent string, last bool, buff *bytes.Buffer) {
	buff.WriteString(fmt.Sprintf(indent))
	if last {
		buff.WriteString(fmt.Sprintf(lastNodeMark))
		indent += space
	} else {
		buff.WriteString(fmt.Sprintf(nodeMark))
		indent += lineMark
	}

	switch node := n.(type) {
	case *NodeProg:
		buff.WriteString(fmt.Sprintln("NodeProg"))

		for i, stmt := range *node.Stmts {
			NodeToString(&stmt, indent, i == len(*node.Stmts)-1, buff)
		}
	case *NodeStmt:
		buff.WriteString(fmt.Sprintln("NodeStmt"))

		switch node.T {
		case TypeNodeStmtInit:
			NodeToString(node.Init, indent, true, buff)
		case TypeNodeStmtReassign:
			NodeToString(node.Reassign, indent, true, buff)
		case TypeNodeStmtScope:
			NodeToString(node.Scope, indent, true, buff)
		case TypeNodeStmtExit:
			NodeToString(node.Exit, indent, true, buff)
		default:
			panicWith("statement")
		}
	case *NodeStmtInit:
		buff.WriteString(fmt.Sprintln("NodeTypeStmtInit"))

		NodeToString(node.Ident, indent, false, buff)
		NodeToString(node.Expr, indent, true, buff)
	case *NodeStmtReassign:
		buff.WriteString(fmt.Sprintln("NodeTypeStmtReassign"))

		NodeToString(node.Ident, indent, false, buff)
		NodeToString(node.Expr, indent, true, buff)
	case *NodeStmtScope:
		buff.WriteString(fmt.Sprintln("NodeTypeStmtScope"))

		for i, stmt := range *node.Scope.Stmts {
			NodeToString(&stmt, indent, i == len(*node.Scope.Stmts)-1, buff)
		}
	case *NodeStmtExit:
		buff.WriteString(fmt.Sprintln("NodeTypeStmtExit"))

		NodeToString(node.Expr, indent, true, buff)
	case *NodeExpr:
		buff.WriteString(fmt.Sprintln("NodeExpr"))

		switch node.T {
		case TypeNodeExprTerm:
			NodeToString(node.Term, indent, true, buff)
		case TypeNodeExprOper:
			NodeToString(node.Oper, indent, true, buff)
		default:
			panicWith("expresion")
		}
	case *NodeTerm:
		buff.WriteString(fmt.Sprintln("NodeTerm"))

		switch node.T {
		case TypeNodeTermLit:
			NodeToString(node.Lit, indent, true, buff)
		case TypeNodeTermIdent:
			NodeToString(node.Ident, indent, true, buff)
		default:
			panicWith("term")
		}
	case *NodeOper:
		buff.WriteString(fmt.Sprintln("NodeOper"))

		NodeToString(node.Oper, indent, false, buff)
		NodeToString(node.Lhs, indent, false, buff)
		NodeToString(node.Rhs, indent, true, buff)
	case *tokenizer.Token:
		buff.WriteString(fmt.Sprintf("Token %s\n", node.String()))
	default:
		panicWith("switch")
	}
}

func panicWith(caze string) {
	panic(fmt.Sprintf("PrintNode not implemented for case: %s", caze))
}
