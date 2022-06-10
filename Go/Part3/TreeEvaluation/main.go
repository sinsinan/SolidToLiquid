package main

import (
	"fmt"
	"SolidToLiquid/parser"
)

func main() {
	parser := parser.ConstructParser()
	fmt.Println(parser.EvaluateExpression("2*3+4"))
}