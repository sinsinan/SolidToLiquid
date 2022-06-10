package expr

type IExpr interface {
	Evaluate() int
}