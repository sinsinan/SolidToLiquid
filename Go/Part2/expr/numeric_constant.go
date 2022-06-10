package expr

type NumericConstant struct{
	number int
}

func ConstructNumericConstant(number int) NumericConstant {
	return NumericConstant{number: number}
}

func (numericConstant NumericConstant) Evaluate() int {
	return numericConstant.number
}