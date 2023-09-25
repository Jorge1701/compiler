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

func PrintNode(n interface{}) {
	buff := bytes.NewBuffer([]byte{})
	NodeToString(n, "", true, buff)
	fmt.Println(buff.String())
}

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

		for i, stmt := range node.Stmts {
			NodeToString(stmt, indent, i == len(node.Stmts)-1, buff)
		}
	case NodeStmt:
		buff.WriteString(fmt.Sprintln("NodeStmt"))

		switch node.T {
		case TypeNodeStmtInit:
			NodeToString(node.Init, indent, true, buff)
		case TypeNodeStmtExit:
			NodeToString(node.Exit, indent, true, buff)
		default:
			caseNotImplemented("NodeStmt", indent, last, buff)
		}
	case *NodeStmtInit:
		buff.WriteString(fmt.Sprintln("NodeTypeStmtInit"))

		NodeToString(node.Ident, indent, false, buff)
		NodeToString(node.Expr, indent, true, buff)
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
			caseNotImplemented("NodeExpr", indent, last, buff)
		}
	case *NodeTerm:
		buff.WriteString(fmt.Sprintln("NodeTerm"))

		switch node.T {
		case TypeNodeTermLit:
			NodeToString(node.Lit, indent, true, buff)
		case TypeNodeTermIdent:
			NodeToString(node.Ident, indent, true, buff)
		default:
			caseNotImplemented("NodeTerm", indent, last, buff)
		}
	case *NodeOper:
		buff.WriteString(fmt.Sprintln("NodeOper"))

		NodeToString(node.Oper, indent, false, buff)
		NodeToString(node.Lhs, indent, false, buff)
		NodeToString(node.Rhs, indent, true, buff)
	case *tokenizer.Token:
		buff.WriteString(fmt.Sprintf("Token %s\n", node.String()))
	default:
		caseNotImplemented("switch", indent, last, buff)
	}
}

func caseNotImplemented(caze, indent string, last bool, buff *bytes.Buffer) {
	buff.WriteString(fmt.Sprintf("NodeToString case '%s' not implemented\n", caze))
}
