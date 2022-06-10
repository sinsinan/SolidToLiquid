package parser

// BNF grammer

// expr -> term | term {+|-} expr
// term -> factor | factor {*|/} term
// factor -> Number | (expr) | {+|-} factor

import (
	"SolidToLiquid/lexer"
	"SolidToLiquid/stack"
)

type RDParser struct {
	lexer        lexer.Lexer
	currentToken lexer.Token
	valueStack   stack.CustomStack
}

func ConstructParser() RDParser {
	return RDParser{}
}

func (parser *RDParser) ParseExpression(exp string) interface{} {
	parser.lexer = lexer.ConstructLexer(exp)
	parser.valueStack = stack.ConstructCustomStack()

	parser.getToken()
	parser.expr()

	if parser.currentToken != lexer.TOK_EOF {
		panic("Error: parsing exited before end")
	}

	return parser.valueStack.Pop()
}

func (parser *RDParser) expr() {
	parser.term()
	if parser.currentToken == lexer.TOK_SUB || parser.currentToken == lexer.TOK_PLUS {
		op := parser.currentToken
		x := parser.valueStack.Pop()
		parser.getToken()
		parser.expr()
		y := parser.valueStack.Pop()
		parser.valueStack.Push(evaluateBinaryOp(x, y, op))
	}
}

func (parser *RDParser) term() {
	parser.factor()
	if parser.currentToken == lexer.TOK_MUL || parser.currentToken == lexer.TOK_DIV {
		op := parser.currentToken
		x := parser.valueStack.Pop()
		parser.getToken()
		parser.term()
		y := parser.valueStack.Pop()
		parser.valueStack.Push(evaluateBinaryOp(x, y, op))
	}
}

func (parser *RDParser) factor() {
	if parser.currentToken == lexer.TOK_NUM {
		parser.valueStack.Push(parser.lexer.GetNumber())
		parser.getToken()
	} else if parser.currentToken == lexer.TOK_PLUS || parser.currentToken == lexer.TOK_SUB {
		op := parser.currentToken
		parser.getToken()
		parser.factor()
		x := parser.valueStack.Pop()
		parser.valueStack.Push(evaluateUnaryOp(x, op))
	} else if parser.currentToken == lexer.TOK_OPAREN {
		parser.getToken()
		parser.expr()
		if parser.currentToken != lexer.TOK_CPAREN {
			panic("Closing for open parenthesis could not be found")
		}
		parser.getToken()
	} else {
		panic("Invalid token")
	}
}

func evaluateUnaryOp(x int, op lexer.Token) int {
	if op == lexer.TOK_SUB {
		return -x
	}
	return x
}

func evaluateBinaryOp(x, y int, op lexer.Token) int {
	if op == lexer.TOK_PLUS {
		return x + y
	} else if op == lexer.TOK_SUB {
		return x - y
	} else if op == lexer.TOK_MUL {
		return x * y
	} else {
		if y == 0 {
			panic("Error division by zero")
		}
		return x / y
	}
}

func (parser *RDParser) getToken() {
	parser.currentToken = parser.lexer.GetToken()
}
