package parser

import (
	"compiler/tokenizer"
	"compiler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeExpr_WhenUnexpectedToken(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.SB_L, Value: "[", Pos: utils.NewPosition(5, 3)},
	})

	expr, err := p.parseNodeExpr(1)

	assert.Error(t, err)
	assert.Nil(t, expr)
	assert.Equal(t, "Invalid expresion: Unexpected token (SB_L, '[') at line 5 and column 3", err.Error())
}

func TestParseNodeExpr_WhenEmptyTokens(t *testing.T) {
	p := NewParser([]tokenizer.Token{})

	expr, err := p.parseNodeExpr(1)

	assert.Error(t, err)
	assert.Nil(t, expr)
	assert.Equal(t, "Invalid expresion: No tokens left", err.Error())
}

func TestParseNodeExpr_WhenTerm(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.LITERAL, Value: "1"},
	})

	expr, err := p.parseNodeExpr(1)

	assert.NoError(t, err)
	assert.NotNil(t, expr)
	assert.Equal(t, TypeNodeExprTerm, expr.T)
	assert.NotNil(t, expr.Term)
	assert.Equal(t, 1, p.index)
}

func TestParseNodeExpr_WhenOperValid(t *testing.T) {
	operators := []tokenizer.Token{
		{Type: tokenizer.ADD, Value: "+"},
		{Type: tokenizer.SUB, Value: "-"},
		{Type: tokenizer.MUL, Value: "*"},
		{Type: tokenizer.DIV, Value: "/"},
	}

	for _, operator := range operators {
		t.Run(operator.String(),
			func(t *testing.T) {
				p := NewParser([]tokenizer.Token{
					{Type: tokenizer.LITERAL, Value: "1"},
					operator,
					{Type: tokenizer.LITERAL, Value: "2"},
				})

				expr, err := p.parseNodeExpr(1)

				assert.NoError(t, err)
				assert.NotNil(t, expr)
				assert.Equal(t, TypeNodeExprOper, expr.T)
				assert.NotNil(t, expr.Oper)
				assert.Equal(t, "1", expr.Oper.Lhs.Term.Lit.Value)
				assert.Equal(t, "2", expr.Oper.Rhs.Term.Lit.Value)
				assert.Equal(t, 3, p.index)
			},
		)
	}
}

func TestParseNodeExpr_WhenOperInvalidExpresion(t *testing.T) {
	p := NewParser([]tokenizer.Token{
		{Type: tokenizer.LITERAL, Value: "1"},
		{Type: tokenizer.ADD, Value: "+"},
		{Type: tokenizer.EQ, Value: "=", Pos: utils.NewPosition(3, 2)},
	})

	expr, err := p.parseNodeExpr(1)

	assert.Error(t, err)
	assert.Nil(t, expr)
	assert.Equal(t, "Invalid expresion: Unexpected token (EQ, '=') at line 3 and column 2", err.Error())
}
