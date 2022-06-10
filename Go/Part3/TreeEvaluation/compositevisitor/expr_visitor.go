package compositevisitor

import (
	"SolidToLiquid/lexer"
)

type ExprVisitor struct{}

func ConstructExprVisitor() ExprVisitor {
	return ExprVisitor{}
}


func (exprVisitor ExprVisitor) VisitUnaryExpr(unaryExpr UnaryExpr) int {
		x := unaryExpr.Expr.Accept(exprVisitor)

		if unaryExpr.Op == lexer.TOK_PLUS {
			return x
		} else if unaryExpr.Op == lexer.TOK_SUB {
			return -x
		}  else {
			panic("Error invalid token for unary operation")
		}
}

func (exprVisitor ExprVisitor) VisitNumericConstant(numericConstant NumericConstant) int {
		return numericConstant.Number
}

func (exprVisitor ExprVisitor) VisitBinaryExpr(binaryExpr BinaryExpr) int {
	x := binaryExpr.Expr1.Accept(exprVisitor)
	y := binaryExpr.Expr2.Accept(exprVisitor)

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
