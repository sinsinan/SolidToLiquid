package parser

// BNF grammer

// expr -> term | term {+|-} expr
// term -> factor | factor {*|/} term
// factor -> Number | (expr) | {+|-} factor

import (
	compositevisitor "SolidToLiquid/compositevisitor"
	"SolidToLiquid/lexer"
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

	return expr.Accept(compositevisitor.ConstructExprVisitor())
}

func (parser *RDParser) expr() compositevisitor.IExpr {
	expr := parser.term()
	if parser.currentToken == lexer.TOK_SUB || parser.currentToken == lexer.TOK_PLUS {
		op := parser.currentToken
		parser.getToken()
		expr2 := parser.expr()
		return compositevisitor.ConstructBinaryExpr(expr, expr2, op)
	}
	return expr
}

func (parser *RDParser) term() compositevisitor.IExpr {
	expr := parser.factor()
	if parser.currentToken == lexer.TOK_MUL || parser.currentToken == lexer.TOK_DIV {
		op := parser.currentToken
		parser.getToken()
		expr2 := parser.term()
		return compositevisitor.ConstructBinaryExpr(expr, expr2, op)
	}
	return expr
}

func (parser *RDParser) factor() compositevisitor.IExpr {
	if parser.currentToken == lexer.TOK_NUM {
		expr := compositevisitor.ConstructNumericConstant(parser.lexer.GetNumber())
		parser.getToken()
		return expr
	} else if parser.currentToken == lexer.TOK_PLUS || parser.currentToken == lexer.TOK_SUB {
		op := parser.currentToken
		parser.getToken()
		expr := parser.factor()
		return compositevisitor.ConstructUnaryExpr(expr, op)
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
