package generator

import (
	"bytes"
	"compiler/parser"
	"compiler/tokenizer"
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
	g.index++
	g.textBuff.WriteString(fmt.Sprintf("    mov rax, %s\n", value))
	g.textBuff.WriteString(fmt.Sprintf("    push rax ; index at %d\n", g.index))
}

func (g *Generator) pushReg(reg string) {
	g.index++
	g.textBuff.WriteString(fmt.Sprintf("    push %s ; index at %d\n", reg, g.index))
}

func (g *Generator) popTo(reg string) {
	g.index--
	g.textBuff.WriteString(fmt.Sprintf("    pop %s\n", reg))
}

func (g *Generator) popSize(size int) {
	g.index -= size
	g.textBuff.WriteString(fmt.Sprintf("    add rsp, %d ; Delete %d from stack\n", size*8, size))
}

func (g *Generator) generateTerm(term *parser.NodeTerm) {
	switch term.T {
	case parser.TypeNodeTermLit:
		g.textBuff.WriteString("    ; Term lit\n")
		g.push(term.Lit.Value)
	case parser.TypeNodeTermIdent:
		g.textBuff.WriteString(fmt.Sprintf("    ; Term ident '%s'\n", term.Ident.Value))
		v, exists := g.variables[term.Ident.Value]
		if !exists {
			fmt.Println("ERROR Variable is not initialized")
		}
		g.textBuff.WriteString(fmt.Sprintf("    mov rax, [rsp+%d]\n", (g.index-v.index)*8))
		g.pushReg("rax")
	default:
		fmt.Println("Not supported generateTerm")
	}
}

func (g *Generator) generateExpr(expr *parser.NodeExpr) {
	switch expr.T {
	case parser.TypeNodeExprTerm:
		g.textBuff.WriteString("    ; Expr term\n")
		g.generateTerm(expr.Term)
	case parser.TypeNodeExprOper:
		g.generateExpr(expr.Oper.Lhs)
		g.generateExpr(expr.Oper.Rhs)
		g.textBuff.WriteString("    ; Expr oper\n")
		g.textBuff.WriteString("    mov rax, [rsp+8]\n")
		g.textBuff.WriteString("    mov rbx, [rsp]\n")
		switch expr.Oper.Oper.Type {
		case tokenizer.ADD:
			g.textBuff.WriteString("    add rax, rbx\n")
		case tokenizer.SUB:
			g.textBuff.WriteString("    sub rax, rbx\n")
		case tokenizer.MUL:
			g.textBuff.WriteString("    mul rbx\n")
		case tokenizer.DIV:
			g.textBuff.WriteString("    div rbx\n")
		default:
			fmt.Println("Not supported operation generateExpr")
		}
		g.popSize(2)
		g.pushReg("rax")
	default:
		fmt.Println("Not supported generateExpr")
	}
}

func (g *Generator) generateStmt(stmt parser.NodeStmt) {
	switch stmt.T {
	case parser.TypeNodeStmtInit:
		g.generateExpr(stmt.Init.Expr)
		varName := stmt.Init.Ident.Value

		_, isUsed := g.variables[varName]
		if isUsed {
			fmt.Println("ERROR: Variable already initialized")
		}
		g.textBuff.WriteString(fmt.Sprintf("    ; Save variable '%s' at index %d\n", varName, g.index))
		g.variables[varName] = Variable{
			index: g.index,
		}
	case parser.TypeNodeStmtReassign:
		g.generateExpr(stmt.Reassign.Expr)
		varName := stmt.Reassign.Ident.Value
		v, existsVar := g.variables[varName]
		if !existsVar {
			fmt.Println("ERROR: Variable not initialized")
		}

		g.textBuff.WriteString(fmt.Sprintf("    ; Reassign variable '%s' at index %d\n", varName, g.index))
		g.popTo("rax")
		g.textBuff.WriteString(fmt.Sprintf("    mov [rsp+%d], rax\n", (g.index-v.index)*8))
	case parser.TypeNodeStmtExit:
		g.generateExpr(stmt.Exit.Expr)
		g.textBuff.WriteString("    ; Stmt Exit\n")
		g.textBuff.WriteString("    mov rdi, rax\n")
		g.textBuff.WriteString("    mov rax, 60\n")
		g.textBuff.WriteString("    syscall\n")
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
