package structural

import "testing"

func Test_DecoratorPattern(t *testing.T) {

	DoubleWithLog := LogDecorate(Double)
	DoubleWithLog(5)
}

func Double(n int) int {
	return n * 2
}
