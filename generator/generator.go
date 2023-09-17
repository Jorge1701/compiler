package generator

import (
	"bytes"
	"compiler/parser"
	"fmt"
)

type Generator struct {
	nodeProg *parser.NodeProg
}

func NewGenerator(nodeProg *parser.NodeProg) *Generator {
	return &Generator{
		nodeProg: nodeProg,
	}
}

func (g *Generator) Generate() []byte {
	buff := bytes.NewBuffer([]byte{})

	dataBuff := bytes.NewBuffer([]byte{})
	dataBuff.WriteString("section .data\n")

	textBuff := bytes.NewBuffer([]byte{})
	textBuff.WriteString("section .text\n")
	textBuff.WriteString("global _start\n")
	textBuff.WriteString("_start:\n")

	for _, s := range g.nodeProg.Stmts {
		fmt.Println(s)
	}

	dataBuff.WriteString("\n")
	textBuff.WriteString("\n")

	// Exit without error at the end, if no other exit then this prevents segmentation fault
	textBuff.WriteString("    mov rax, 60\n")
	textBuff.WriteString("    mov rdi,0\n")
	textBuff.WriteString("    syscall\n")

	buff.WriteString(dataBuff.String())
	buff.WriteString(textBuff.String())

	return buff.Bytes()
}
