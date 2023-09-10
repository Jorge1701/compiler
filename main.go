package main

import (
	"bytes"
	"compiler/tokenizer"
	"fmt"
	"log"
	"os"
)

func generateOutput(tokens []tokenizer.Token) []byte {
	buff := bytes.NewBuffer([]byte{})

	buff.WriteString("global _start\n")
	buff.WriteString("_start:\n")

	for i := 0; i < len(tokens); i++ {
		t := tokens[i]
		if t.TokenType == tokenizer.KEYWORD && tokenizer.KeyWord(t.Value) == tokenizer.SALIR {
			if i+1 < len(tokens) && tokens[i+1].TokenType == tokenizer.LITERAL {
				if i+2 < len(tokens) && tokens[i+2].TokenType == tokenizer.SEPARATOR && tokens[i+2].Value == ";" {
					buff.WriteString("    mov rax, 60\n")
					buff.WriteString(fmt.Sprintf("    mov rdi, %s\n", tokens[i+1].Value))
					buff.WriteString("    syscall")
				}
			}
		}
	}

	return buff.Bytes()
}

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

	// Generate output
	output := generateOutput(tokens)

	// Write asm file
	err = os.WriteFile("output.asm", output, 0755)
	if err != nil {
		log.Fatal(err)
	}
}
