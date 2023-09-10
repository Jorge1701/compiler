package main

import (
	"bytes"
	"compiler/generator"
	"compiler/parser"
	"compiler/tokenizer"
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
	tokens := t.Tokenize()

	// Parse tokens
	p := parser.NewParser(tokens)
	node := p.Parse()

	// Generate output
	g := generator.NewGenerator(node)
	output := g.Generate()

	// Write asm file
	err = os.WriteFile("output.asm", output, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
