package parser

import (
	"compiler/tokenizer"
	"compiler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseNodeStmt
func TestParseNodeStmt_WhenEmptyTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	stmt, err := p.parseNodeStmt()

	assert.Error(t, err)
	assert.Nil(t, stmt)
	assert.Equal(t, "Invalid statement: No tokens left", err.Error())
}

func TestParseNodeStmt_WhentUnexpectedToken(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.SEP, Pos: utils.NewPosition(5, 3)},
	})

	stmt, err := p.parseNodeStmt()

	assert.Error(t, err)
	assert.Nil(t, stmt)
	assert.Equal(t, "Invalid statement: Unexpected token (SEP, '\\n') at line 5 and column 3", err.Error())
}

// TestParseNodeStmtInit
func TestParseNodeStmt_WhenStmtInitUnexpectedToken(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.INT},
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.LITERAL, Value: "1", Pos: utils.NewPosition(6, 2)},
	})

	stmt, err := p.parseNodeStmt()

	assert.Error(t, err)
	assert.Nil(t, stmt)
	assert.Equal(t, "Invalid statement: Unexpected token (LITERAL, '1') at line 6 and column 2", err.Error())
}

func TestParseNodeStmt_WhenStmtInit(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.INT},
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.EQ},
		{Type: tokenizer.LITERAL, Value: "1"},
	})

	stmt, err := p.parseNodeStmt()

	assert.NoError(t, err)
	assert.NotNil(t, stmt)
	assert.Equal(t, TypeNodeStmtInit, stmt.T)
	assert.Equal(t, "name", stmt.Init.Ident.Value)
	assert.NotNil(t, stmt.Init.Expr)
}

// TestParseNodeStmtReassign
func TestParseNodeStmt_WhenStmtReassignUnexpectedToken(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.LITERAL, Value: "1", Pos: utils.NewPosition(2, 74)},
	})

	stmt, err := p.parseNodeStmt()

	assert.Error(t, err)
	assert.Nil(t, stmt)
	assert.Equal(t, "Invalid statement: Unexpected token (LITERAL, '1') at line 2 and column 74", err.Error())
}

func TestParseNodeStmt_WhenStmtReassign(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.EQ},
		{Type: tokenizer.LITERAL, Value: "1"},
	})

	stmt, err := p.parseNodeStmt()

	assert.NoError(t, err)
	assert.NotNil(t, stmt)
	assert.Equal(t, TypeNodeStmtReassign, stmt.T)
	assert.Equal(t, "name", stmt.Reassign.Ident.Value)
	assert.NotNil(t, stmt.Reassign.Expr)
}

// TestParseNodeStmtReassign
func TestParseNodeStmt_WhenStmtExitUnexpectedToken(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.EXIT},
		{Type: tokenizer.P_L, Value: "(", Pos: utils.NewPosition(3, 7)},
	})

	stmt, err := p.parseNodeStmt()

	assert.Error(t, err)
	assert.Nil(t, stmt)
	assert.Equal(t, "Invalid statement: Invalid expresion: Unexpected token (P_L, '(') at line 3 and column 7", err.Error())
}

func TestParseNodeStmt_WhenStmtExit(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.EXIT},
		{Type: tokenizer.LITERAL, Value: "1"},
	})

	stmt, err := p.parseNodeStmt()

	assert.NoError(t, err)
	assert.NotNil(t, stmt)
	assert.Equal(t, TypeNodeStmtExit, stmt.T)
	assert.NotNil(t, stmt.Exit.Expr)
}

// TestParseNodeStmtScope
func TestParseNodeStmt_WhenStmtScope(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.B_L},
		{Type: tokenizer.INT},
		{Type: tokenizer.IDENTIFIER, Value: "name"},
		{Type: tokenizer.EQ},
		{Type: tokenizer.LITERAL, Value: "1"},
		{Type: tokenizer.B_R},
	})

	stmt, err := p.parseNodeStmt()

	assert.NoError(t, err)
	assert.NotNil(t, stmt)
}
