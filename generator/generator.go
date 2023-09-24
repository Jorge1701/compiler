package generator

import (
	"bytes"
	"compiler/parser"
	"fmt"
)

type Generator struct {
	nodeProg *parser.NodeProg

	dataBuff bytes.Buffer
	textBuff bytes.Buffer

	variables map[string]Variable
	index     int
}

type Variable struct {
	index int
}

func NewGenerator(nodeProg *parser.NodeProg) *Generator {
	return &Generator{
		nodeProg:  nodeProg,
		dataBuff:  *bytes.NewBuffer([]byte{}),
		textBuff:  *bytes.NewBuffer([]byte{}),
		variables: map[string]Variable{},
		index:     0,
	}
}

func (g *Generator) push(value string) {
	g.textBuff.WriteString(fmt.Sprintf("    mov rax, %s\n", value))
	g.textBuff.WriteString("    push rax\n")
	g.index++
}

func (g *Generator) generateTerm(term *parser.NodeTerm) {
	switch term.T {
	case parser.TypeNodeTermLit:
		g.textBuff.WriteString("    ; Term lit\n")
		g.push(term.Lit.Value)
	case parser.TypeNodeTermIdent:
		g.textBuff.WriteString("    ; Term ident\n")
		v, exists := g.variables[term.Ident.Value]
		if !exists {
			fmt.Println("ERROR Variable is not initialized")
		}
		g.textBuff.WriteString(fmt.Sprintf("    mov rax, [rsp+%d]\n", v.index*8))
	default:
		fmt.Println("Not supported generateTerm")
	}
}

func (g *Generator) generateExpr(expr *parser.NodeExpr) {
	switch expr.T {
	case parser.TypeNodeExprTerm:
		g.generateTerm(expr.Term)
	default:
		fmt.Println("Not supported generateExpr")
	}
}

func (g *Generator) generateStmt(stmt parser.NodeStmt) {
	switch stmt.T {
	case parser.TypeNodeStmtExit:
		g.generateExpr(stmt.Exit.Expr)
		g.textBuff.WriteString("    ; Stmt Exit\n")
		g.textBuff.WriteString("    mov rdi, rax\n")
		g.textBuff.WriteString("    mov rax, 60\n")
		g.textBuff.WriteString("    syscall\n")
	case parser.TypeNodeStmtInit:
		g.generateExpr(stmt.Init.Expr)
		_, isUsed := g.variables[stmt.Init.Ident.Value]
		if isUsed {
			fmt.Println("ERROR: Variable already initialized")
		}
		g.variables[stmt.Init.Ident.Value] = Variable{
			index: g.index,
		}
	default:
		fmt.Println("Not supported generateStmt")
	}
}

func (g *Generator) GenerateNASM() []byte {
	buff := bytes.NewBuffer([]byte{})

	g.dataBuff.WriteString("section .data\n")

	g.textBuff.WriteString("section .text\n")
	g.textBuff.WriteString("global _start\n")
	g.textBuff.WriteString("_start:\n")

	for _, s := range g.nodeProg.Stmts {
		g.generateStmt(s)
	}

	g.dataBuff.WriteString("\n")
	g.textBuff.WriteString("\n")

	// Exit without error at the end, if no other exit then this prevents segmentation fault
	g.textBuff.WriteString("    mov rax, 60\n")
	g.textBuff.WriteString("    mov rdi,0\n")
	g.textBuff.WriteString("    syscall\n")

	buff.WriteString(g.dataBuff.String())
	buff.WriteString(g.textBuff.String())

	return buff.Bytes()
}
