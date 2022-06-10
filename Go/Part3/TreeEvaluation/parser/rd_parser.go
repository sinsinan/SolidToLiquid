package parser

// BNF grammer

// expr -> term | term {+|-} expr
// term -> factor | factor {*|/} term
// factor -> Number | (expr) | {+|-} factor

import (
	expressions "SolidToLiquid/expr"
	"SolidToLiquid/lexer"
	"SolidToLiquid/visitor"
)

type RDParser struct {
	lexer        lexer.Lexer
	currentToken lexer.Token
}

func ConstructParser() RDParser {
	return RDParser{}
}

func (parser *RDParser) EvaluateExpression(exp string) int {
	parser.lexer = lexer.ConstructLexer(exp)

	parser.getToken()
	expr := parser.expr()

	if parser.currentToken != lexer.TOK_EOF {
		panic("Error: parsing exited before end")
	}

	return visitor.Visit(expr)
}

func (parser *RDParser) expr() expressions.IExpr {
	expr := parser.term()
	if parser.currentToken == lexer.TOK_SUB || parser.currentToken == lexer.TOK_PLUS {
		op := parser.currentToken
		parser.getToken()
		expr2 := parser.expr()
		return expressions.ConstructBinaryExpr(expr, expr2, op)
	}
	return expr
}

func (parser *RDParser) term() expressions.IExpr {
	expr := parser.factor()
	if parser.currentToken == lexer.TOK_MUL || parser.currentToken == lexer.TOK_DIV {
		op := parser.currentToken
		parser.getToken()
		expr2 := parser.term()
		return expressions.ConstructBinaryExpr(expr, expr2, op)
	}
	return expr
}

func (parser *RDParser) factor() expressions.IExpr {
	if parser.currentToken == lexer.TOK_NUM {
		expr := expressions.ConstructNumericConstant(parser.lexer.GetNumber())
		parser.getToken()
		return expr
	} else if parser.currentToken == lexer.TOK_PLUS || parser.currentToken == lexer.TOK_SUB {
		op := parser.currentToken
		parser.getToken()
		expr := parser.factor()
		return expressions.ConstructUnaryExpr(expr, op)
	} else if parser.currentToken == lexer.TOK_OPAREN {
		parser.getToken()
		expr := parser.expr()
		if parser.currentToken != lexer.TOK_CPAREN {
			panic("Closing for open parenthesis could not be found")
		}
		parser.getToken()
		return expr
	} else {
		panic("Invalid token")
	}
}

func (parser *RDParser) getToken() {
	parser.currentToken = parser.lexer.GetToken()
}
