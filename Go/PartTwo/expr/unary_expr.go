package expr

import "SolidToLiquid/lexer"

type UnaryExpr struct{
	expr IExpr
	op lexer.Token
}

func ConstructUnaryExpr(expr IExpr, op lexer.Token) UnaryExpr {
	return UnaryExpr{expr: expr, op: op}
}

func (unaryExpr UnaryExpr) Evaluate() int {
	x := unaryExpr.expr.Evaluate()

	if unaryExpr.op == lexer.TOK_PLUS {
		return x
	} else if unaryExpr.op == lexer.TOK_SUB {
		return -x
	}  else {
		panic("Error invalid token for unary operation")
	}
}