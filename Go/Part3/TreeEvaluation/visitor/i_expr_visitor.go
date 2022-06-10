package visitor

import "SolidToLiquid/expr"

type IExprVisitor interface {
	Visit(expression expr.IExpr) int
}