package expr

import "SolidToLiquid/lexer"

type BinaryExpr struct{
	expr1 IExpr
	expr2 IExpr
	op lexer.Token
}

func ConstructBinaryExpr(expr1 IExpr, expr2 IExpr, op lexer.Token) BinaryExpr {
	return BinaryExpr{expr1: expr1, expr2: expr2, op: op}
}

func (binaryExpr BinaryExpr) Evaluate() int {
	x := binaryExpr.expr1.Evaluate()
	y:= binaryExpr.expr2.Evaluate()

	if binaryExpr.op == lexer.TOK_PLUS {
		return x + y
	} else if binaryExpr.op == lexer.TOK_SUB {
		return x - y
	} else if binaryExpr.op == lexer.TOK_MUL {
		return x * y
	} else if binaryExpr.op == lexer.TOK_DIV {
		if y == 0 {
			panic("Error division by zero")
		}
		return x / y
	} else {
		panic("Error invalid token for binary operation")
	}
}