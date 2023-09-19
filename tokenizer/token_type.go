package tokenizer

type TokenType string

// Definition of all the available tokens in the language
const (
	SEP  TokenType = "SEP"
	P_L  TokenType = "P_L"
	P_R  TokenType = "P_R"
	B_L  TokenType = "B_L"
	B_R  TokenType = "B_R"
	SB_L TokenType = "SB_L"
	SB_R TokenType = "SB_R"

	ADD TokenType = "ADD"
	SUB TokenType = "SUB"
	MUL TokenType = "MUL"
	DIV TokenType = "DIV"

	EQ TokenType = "EQ"

	INT  TokenType = "INT"
	EXIT TokenType = "EXIT"

	IDENTIFIER TokenType = "IDENTIFIER"
	LITERAL    TokenType = "LITERAL"
)

// Defines a map of all the tokens that are a singe character (rune) for quick access
var singleRuneTokens = map[rune]TokenType{
	'\n': SEP,
	'(':  P_L,
	')':  P_R,
	'{':  B_L,
	'}':  B_R,
	'[':  SB_L,
	']':  SB_R,
	'+':  ADD,
	'-':  SUB,
	'*':  MUL,
	'/':  DIV,
	'=':  EQ,
}

// Definition of all the keywords in the language and their respective tokens
var listOfKeywords = map[string]TokenType{
	"int":  INT,
	"exit": EXIT,
}
