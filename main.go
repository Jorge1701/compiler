package main

import (
	"bytes"
	"compiler/generator"
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
	err = t.GenerateTokens()
	if err != nil {
		panic(fmt.Sprintf("Tokenizer error: %s", err))
	}

	// Print tokens
	fmt.Println("=== Tokens === ")
	for _, t := range t.GetTokens() {
		fmt.Println(t.String())
	}

	p := parser.NewParser(t.GetTokens())
	nodeProg, err := p.GenerateNodes()
	if err != nil {
		fmt.Println(err)
	}

	// Print parse tree
	fmt.Println("=== Parse tree === ")
	parser.PrintNode(nodeProg)

	g := generator.NewGenerator(nodeProg)
	nasmBs := g.GenerateNASM()
	fmt.Println("=== NASM ===")
	fmt.Println(string(nasmBs))

	// Write asm file
	err = os.WriteFile("output.asm", nasmBs, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
