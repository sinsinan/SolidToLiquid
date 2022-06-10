package expr

import "SolidToLiquid/lexer"

type BinaryExpr struct{
	Expr1 IExpr
	Expr2 IExpr
	Op lexer.Token
}

func ConstructBinaryExpr(expr1 IExpr, expr2 IExpr, op lexer.Token) BinaryExpr {
	return BinaryExpr{Expr1: expr1, Expr2: expr2, Op: op}
}