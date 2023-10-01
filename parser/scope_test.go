package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeScope(t *testing.T) {
	p := NewParser(generateTokensFor("{a=1}"))

	scope, err := p.parseNodeScope()

	assert.NoError(t, err)
	assert.NotNil(t, scope)
	assert.Equal(t, 5, p.index)
}
