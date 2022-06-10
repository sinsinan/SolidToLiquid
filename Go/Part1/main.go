package main

import (
	"fmt"
	"SolidToLiquid/parser"
)

func main() {
	parser := parser.ConstructParser()
	fmt.Println(parser.ParseExpression("2*3+-4"))
}