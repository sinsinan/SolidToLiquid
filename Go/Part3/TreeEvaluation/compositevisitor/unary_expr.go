package compositevisitor

import (
	"SolidToLiquid/lexer"
)

type UnaryExpr struct {
	Expr IExpr
	Op   lexer.Token
}

func ConstructUnaryExpr(expr IExpr, op lexer.Token) UnaryExpr {
	return UnaryExpr{Expr: expr, Op: op}
}

func (UnaryExpr UnaryExpr) Accept(visitor IExprVisitor) int {
	return visitor.VisitUnaryExpr(UnaryExpr)
}
