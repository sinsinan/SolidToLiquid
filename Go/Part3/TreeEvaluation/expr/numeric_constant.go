package expr

type NumericConstant struct{
	Number int
}

func ConstructNumericConstant(number int) NumericConstant {
	return NumericConstant{Number: number}
}