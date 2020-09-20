package behavioral

type Operator interface {
	Apply(int, int) int
}

type Operation struct {
	Operator Operator
}

func (o *Operation) Operate(leftValue, rightValue int) int {
	return o.Operator.Apply(leftValue, rightValue)
}

type Addition struct{}

func (Addition) Apply(lVal, rVal int) int {
	return lVal + rVal
}

type Multiplication struct{}

func (Multiplication) Apply(lVal, rVal int) int {
	return lVal * rVal
}
