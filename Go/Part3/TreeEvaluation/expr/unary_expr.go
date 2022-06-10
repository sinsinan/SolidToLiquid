package expr

import "SolidToLiquid/lexer"

type UnaryExpr struct{
	Expr IExpr
	Op lexer.Token
}

func ConstructUnaryExpr(expr IExpr, op lexer.Token) UnaryExpr {
	return UnaryExpr{Expr: expr, Op: op}
}