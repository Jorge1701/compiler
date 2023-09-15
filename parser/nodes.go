package parser

import (
	"compiler/tokenizer"
)

type NodeProg struct {
	Statements []interface{}
}

type NodeSalirLiteral struct {
	Literal tokenizer.Token
}

type NodeSalirIdentifier struct {
    Identifier tokenizer.Token
}

type NodeInitialize struct {
	Identifier tokenizer.Token
	Literal    tokenizer.Token
}
