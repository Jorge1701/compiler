package generator

import (
	"bytes"
	"compiler/parser"
	"fmt"
)

type Generator struct {
	nodeSalir *parser.NodeSalir
}

func NewGenerator(nodeSalir *parser.NodeSalir) *Generator {
	return &Generator{
		nodeSalir: nodeSalir,
	}
}

func (g *Generator) Generate() []byte {
	buff := bytes.NewBuffer([]byte{})

	buff.WriteString("global _start\n")
	buff.WriteString("_start:\n")

	buff.WriteString("    mov rax, 60\n")
	buff.WriteString(fmt.Sprintf("    mov rdi, %s\n", g.nodeSalir.NodeLiteral.Literal.Value))
	buff.WriteString("    syscall")

	return buff.Bytes()
}
