package main

import (
	"bytes"
	"compiler/parser"
	"compiler/tokenizer"
	"fmt"
	"log"
	"os"
)

func main() {
	// Read arguments
	if len(os.Args) != 2 {
		log.Fatal("Missing argument: filename")
	}

	// Read file
	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	// Tokenize the input
	t := tokenizer.NewTokenizer(bytes.Runes(bs))
	tokens := t.GenerateTokens()

	// Print tokens
	fmt.Println("=== Tokens === ")
	for _, t := range tokens {
		fmt.Println(t.String())
	}

	p := parser.NewParser(tokens)
	nodeProg, err := p.GenerateNodes()
	if err != nil {
		fmt.Println(err)
	}

	// Print parse tree
	fmt.Println("=== Parse tree === ")
	nodeProg.Print()
}
