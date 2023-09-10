package main

import (
	"bytes"
	"compiler/tokenizer"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Missing argument: filename")
	}

	bs, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("Error reading file: ", err)
	}

	t := tokenizer.NewTokenizer(bytes.Runes(bs))

	tokens := t.Tokenize()
	for _, t := range tokens {
		fmt.Println(t.TokenType, t.Value)
	}
}
