package compositevisitor

type NumericConstant struct{
	Number int
}

func ConstructNumericConstant(number int) NumericConstant {
	return NumericConstant{Number: number}
}

func (numericConstant NumericConstant) Accept(visitor IExprVisitor) int {
	return visitor.VisitNumericConstant(numericConstant)
}