package compositevisitor

type IExprVisitor interface {
	VisitUnaryExpr(UnaryExpr) int
	VisitNumericConstant(NumericConstant) int
	VisitBinaryExpr(BinaryExpr) int
}
