package main

import (
	"fmt"
	"p1/parser"
)

func main() {
	parser := parser.ConstructParser()
	fmt.Println(parser.ParseExpression("2*3+4"))
}