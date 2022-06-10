package visitor

import (
	"SolidToLiquid/expr"
	"SolidToLiquid/lexer"
)

type ExprVisitor struct{}

func ConstructExprVisitor() ExprVisitor {
	return ExprVisitor{}
}

func Visit(expression expr.IExpr) int {
	switch typeToVisit := expression.(type) {
	case expr.BinaryExpr:
		return visitBinaryExpr(typeToVisit)
	case expr.NumericConstant:
		return visitNumericConstant(typeToVisit)
	case expr.UnaryExpr:
		return visitUnaryExpr(typeToVisit)
	default:
		panic("Invalid expression type")
	}
}

func visitUnaryExpr(unaryExpr expr.UnaryExpr) int {
		x := Visit(unaryExpr.Expr)

		if unaryExpr.Op == lexer.TOK_PLUS {
			return x
		} else if unaryExpr.Op == lexer.TOK_SUB {
			return -x
		}  else {
			panic("Error invalid token for unary operation")
		}
}

func visitNumericConstant(numericConstant expr.NumericConstant) int {
		return numericConstant.Number
}

func visitBinaryExpr(binaryExpr expr.BinaryExpr) int {
	x := Visit(binaryExpr.Expr1)
	y := Visit(binaryExpr.Expr2)

	if binaryExpr.Op == lexer.TOK_PLUS {
		return x + y
	} else if binaryExpr.Op == lexer.TOK_SUB {
		return x - y
	} else if binaryExpr.Op == lexer.TOK_MUL {
		return x * y
	} else if binaryExpr.Op == lexer.TOK_DIV {
		if y == 0 {
			panic("Error division by zero")
		}
		return x / y
	} else {
		panic("Error invalid token for binary operation")
	}
}
