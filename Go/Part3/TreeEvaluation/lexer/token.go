package lexer

type Token int

const (
	ILLEGAL_TOKEN Token = -1
	TOK_PLUS            = 1
	TOK_MUL             = 2
	TOK_DIV             = 3
	TOK_SUB             = 4
	TOK_OPAREN          = 5
	TOK_CPAREN          = 6
	TOK_NUM             = 7
	TOK_EOF             = 8
	TOK_DOT             = 9
)
