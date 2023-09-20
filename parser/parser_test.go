package parser

import (
	"bytes"
	"compiler/tokenizer"
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	examplesDir string = "../examples"
	resultsDir  string = "../resources/tests/results"
)

var exampleFiles = []string{
	"exit_literal.tbd",
	"exit_variable.tbd",
	"operations.tbd",
}

func TestParseNodes(t *testing.T) {
	// Iterates over the example files
	for _, file := range exampleFiles {
		// Run a separate test for each example file
		t.Run(fmt.Sprintf("Test example file '%s'", file),
			func(t *testing.T) {
				testExampleFile(t, file)
			},
		)
	}
}

func testExampleFile(t *testing.T, exampleFile string) {
	// Generate tokens for example file
	nodeProg := generateNodes(t, fmt.Sprintf("%s/%s", examplesDir, exampleFile))

	// Read the expected results
	expectedNodes := readResultsFile(t, exampleFile)

	buff := bytes.NewBuffer([]byte{})
	NodeToString(nodeProg, "", true, buff)

	assert.Equal(t, expectedNodes, buff.String(), "Nodes do not match")
}

// generateNodes reads the example file, creates a tokenizer, parser
// and returns the generated nodes
func generateNodes(t *testing.T, file string) *NodeProg {
	bs, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Error opening example file '%s': %s", file, err)
	}

	tokenizer := tokenizer.NewTokenizer(bytes.Runes(bs))
	tokens := tokenizer.GenerateTokens()

	parser := NewParser(tokens)
	nodeProg, err := parser.GenerateNodes()
	if err != nil {
		t.Fatalf("Error generating nodes %s", err)
	}
	return nodeProg
}

// readResultsFile reads the contents of the '.nodes' file which contains
// the expected resulting nodes for the given 'file'
func readResultsFile(t *testing.T, file string) string {
	f, err := os.ReadFile(fmt.Sprintf("%s/%s.nodes", resultsDir, file))
	if err != nil {
		t.Fatalf("Error loading results file for '%s': %s", file, err)
	}
	return string(f)

}
