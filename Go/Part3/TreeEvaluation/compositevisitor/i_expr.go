package compositevisitor

type IExpr interface {
	Accept(visitor IExprVisitor) int
}