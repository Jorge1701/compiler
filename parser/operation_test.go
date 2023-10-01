package parser

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeExprOper_WhenInvalidTerm(t *testing.T) {
	p := NewParser(generateTokensFor("(+("))

	node, err := p.parseNodeExprOper(1)

	assert.Error(t, err)
	assert.Nil(t, node)
	assert.Equal(t, "Invalid term: Unexpected token (P_L, '(') at line 1 and column 1", err.Error())
}

func TestParseNodeExprOper_WhenInvalidExpresion(t *testing.T) {
	p := NewParser(generateTokensFor("2+("))

	node, err := p.parseNodeExprOper(1)

	assert.Error(t, err)
	assert.Nil(t, node)
	assert.Equal(t, "Invalid expresion: Unexpected token (P_L, '(') at line 1 and column 3", err.Error())
}

func TestParseNodeExprOper(t *testing.T) {
	operators := []rune{'+', '-', '/', '*'}
	simpleOperations := []string{}

	for _, operator := range operators {
		simpleOperations = append(simpleOperations, fmt.Sprintf("2%c2", operator))
		simpleOperations = append(simpleOperations, fmt.Sprintf("b %c2", operator))
		simpleOperations = append(simpleOperations, fmt.Sprintf("2 %c var", operator))
		simpleOperations = append(simpleOperations, fmt.Sprintf("name%cother", operator))
	}

	multiOperations := []string{}
	for _, operation := range simpleOperations {
		for _, operator := range operators {
			multiOperations = append(multiOperations,
				fmt.Sprintf("%s%c%s", operation, operator, operation),
			)
		}
	}

	for _, operation := range append(simpleOperations, multiOperations...) {
		t.Run(operation,
			func(t *testing.T) {
				tokens := generateTokensFor(operation)
				p := NewParser(tokens)

				node, err := p.parseNodeExprOper(1)

				assert.NoError(t, err)
				assert.NotNil(t, node)
				assert.Equal(t, len(tokens), p.index)
			},
		)
	}
}
