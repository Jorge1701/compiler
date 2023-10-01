package tokenizer

// Represents all available tokens
var allTokenTypes = []TokenType{
	SEP,
	P_L,
	P_R,
	B_L,
	B_R,
	SB_L,
	SB_R,
	ADD,
	SUB,
	MUL,
	DIV,
	EQ,
	INT,
	BOOL,
	EXIT,
	IDENTIFIER,
	INT_LITERAL,
	BOOL_LITERAL,
}

// Represents all the token types for which the actual value matter when in Token form
var valueTokenTypes = []TokenType{
	IDENTIFIER,
	INT_LITERAL,
	BOOL_LITERAL,
}
