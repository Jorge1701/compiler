package parser

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

const result = `└─ NodeProg
   ├─ NodeStmt
   │  └─ NodeTypeStmtInit
   │     ├─ Token (IDENTIFIER, 'a')
   │     └─ NodeExpr
   │        └─ NodeOper
   │           ├─ Token (ADD, '+')
   │           ├─ NodeExpr
   │           │  └─ NodeTerm
   │           │     └─ Token (LITERAL, '1')
   │           └─ NodeExpr
   │              └─ NodeTerm
   │                 └─ Token (LITERAL, '4')
   ├─ NodeStmt
   │  └─ NodeTypeStmtInit
   │     ├─ Token (IDENTIFIER, 'b')
   │     └─ NodeExpr
   │        └─ NodeOper
   │           ├─ Token (SUB, '-')
   │           ├─ NodeExpr
   │           │  └─ NodeTerm
   │           │     └─ Token (LITERAL, '4')
   │           └─ NodeExpr
   │              └─ NodeTerm
   │                 └─ Token (LITERAL, '2')
   ├─ NodeStmt
   │  └─ NodeTypeStmtInit
   │     ├─ Token (IDENTIFIER, 'c')
   │     └─ NodeExpr
   │        └─ NodeOper
   │           ├─ Token (MUL, '*')
   │           ├─ NodeExpr
   │           │  └─ NodeTerm
   │           │     └─ Token (LITERAL, '2')
   │           └─ NodeExpr
   │              └─ NodeTerm
   │                 └─ Token (LITERAL, '2')
   ├─ NodeStmt
   │  └─ NodeTypeStmtInit
   │     ├─ Token (IDENTIFIER, 'd')
   │     └─ NodeExpr
   │        └─ NodeOper
   │           ├─ Token (DIV, '/')
   │           ├─ NodeExpr
   │           │  └─ NodeTerm
   │           │     └─ Token (LITERAL, '10')
   │           └─ NodeExpr
   │              └─ NodeTerm
   │                 └─ Token (LITERAL, '5')
   ├─ NodeStmt
   │  └─ NodeTypeStmtScope
   │     └─ NodeStmt
   │        └─ NodeTypeStmtReassign
   │           ├─ Token (IDENTIFIER, 'a')
   │           └─ NodeExpr
   │              └─ NodeOper
   │                 ├─ Token (ADD, '+')
   │                 ├─ NodeExpr
   │                 │  └─ NodeTerm
   │                 │     └─ Token (IDENTIFIER, 'b')
   │                 └─ NodeExpr
   │                    └─ NodeOper
   │                       ├─ Token (SUB, '-')
   │                       ├─ NodeExpr
   │                       │  └─ NodeTerm
   │                       │     └─ Token (IDENTIFIER, 'c')
   │                       └─ NodeExpr
   │                          └─ NodeTerm
   │                             └─ Token (IDENTIFIER, 'd')
   └─ NodeStmt
      └─ NodeTypeStmtExit
         └─ NodeExpr
            └─ NodeTerm
               └─ Token (IDENTIFIER, 'a')
`

func TestNodeToString(t *testing.T) {
	program := `int a = 1 + 4
    int b = 4 - 2
    int c = 2 * 2
    int d = 10 / 5
    {
        a = b + c - d
    }
    exit a
    `
	p := NewParser(generateTokensFor(program))

	err := p.GenerateNodes()
	assert.Nil(t, err)

	assert.NotPanics(t, func() {
		buff := bytes.NewBuffer([]byte{})
		NodeToString(p.GetNodes(), "", true, buff)

		assert.Equal(t, result, buff.String())
	})
}

func TestNodeToString_WhenCaseSwitchNotImplemented(t *testing.T) {
	assert.PanicsWithValue(t, "PrintNode not implemented for case: switch", func() {
		NodeToString("", "", true, bytes.NewBuffer([]byte{}))
	})
}

func TestNodeToString_WhenCaseStatementNotImplemented(t *testing.T) {
	stmt := &NodeStmt{
		T: 4,
	}

	assert.PanicsWithValue(t, "PrintNode not implemented for case: statement", func() {
		NodeToString(stmt, "", true, bytes.NewBuffer([]byte{}))
	})
}

func TestNodeToString_WhenCaseExprNotImplemented(t *testing.T) {
	expr := &NodeExpr{
		T: 2,
	}

	assert.PanicsWithValue(t, "PrintNode not implemented for case: expresion", func() {
		NodeToString(expr, "", true, bytes.NewBuffer([]byte{}))
	})
}

func TestNodeToString_WhenCaseTermNotImplemented(t *testing.T) {
	expr := &NodeTerm{
		T: 2,
	}

	assert.PanicsWithValue(t, "PrintNode not implemented for case: term", func() {
		NodeToString(expr, "", true, bytes.NewBuffer([]byte{}))
	})
}
