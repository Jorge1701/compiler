package tokenizer

import (
	"bufio"
	"bytes"
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

func TestGenerateTokens(t *testing.T) {
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
	tokens := generateTokens(t, fmt.Sprintf("%s/%s", examplesDir, exampleFile))

	// Read the expected results
	expectedTokens := readResultsFile(t, exampleFile)

	assert.Equal(t, len(expectedTokens), len(tokens),
		"Amount of tokens is not the expected",
	)

	for i, expected := range expectedTokens {
		actual := tokens[i].String()

		assert.Equal(t, expected, actual,
			fmt.Sprintf("Non matching token at line %d", i+1),
		)
	}
}

// generateTokens reads the example file, creates a tokenizer
// and returns the generated tokens
func generateTokens(t *testing.T, file string) []Token {
	bs, err := os.ReadFile(file)
	if err != nil {
		t.Fatalf("Error opening example file '%s': %s", file, err)
	}

	tokenizer := NewTokenizer(bytes.Runes(bs))
	tokenizer.GenerateTokens()
	return tokenizer.GetTokens()
}

// readResultsFile reads the contents of the '.tokens' file which contains
// the expected tokens for the given 'file'
func readResultsFile(t *testing.T, file string) []string {
	f, err := os.Open(fmt.Sprintf("%s/%s.tokens", resultsDir, file))
	if err != nil {
		t.Fatalf("Error opening results file for '%s': %s", file, err)
	}

	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	f.Close()

	return lines
}
